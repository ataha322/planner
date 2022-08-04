package controllers

import (
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"planner.xyi/src/database"
	"planner.xyi/src/middlewares"
	"planner.xyi/src/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	if !ValidateEmail(data["email"]) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email is not valid",
		})
	}

	user := models.User{
		Firstname: data["first_name"],
		Username:  data["username"],
		Email:     data["email"],
	}
	user.SetPassword(data["password"])
	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	// Implementation of both username and email login
	// ? probably possible to refactor
	if ValidateEmail(data["email"]) {
		database.DB.Where("email = ?", data["email"]).First(&user)

		if user.Id == 0 {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "Invalid Credentials!",
			})
		}

		if err := user.ComparePassword(data["password"]); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "Wrong password",
			})
		}

	} else {
		database.DB.Where("username = ?", data["username"]).First(&user)

		if user.Id == 0 {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "Invalid Credentials!",
			})
		}

		if err := user.ComparePassword(data["password"]); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "Wrong password",
			})
		}
	}

	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials!",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Message": "success",
	})
}

func User(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	id, _ := middlewares.GetUserId(c)

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
	}
	user.Id = id

	database.DB.Model(&user).Updates(&user)
	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{}
	user.Id = id

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)
	return c.JSON(user)
}

// Checks if email follows RFC 5322 structure and if
// email domain exists in MX records
func ValidateEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !emailRegex.MatchString(email) {
		return false
	}
	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
