package models

import (
	"time"
)

// TODO: fix pointers
type taskList struct {
	ID              uint      `json: "task_id"`
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json: "task_description`
	TaskDeadline    time.Time `json: "task_deadline"`
	// TaskList        []taskList `json:"SubtaskList"`
}
