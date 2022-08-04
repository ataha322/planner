package models

import (
	"time"
)

type Task struct {
	Model
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_description"`
	TaskDeadline    time.Time `json:"task_deadline"`
	UserId          uint      `json:"user_id"`
}

/**
* returns deadline datetime of a task
 */

func (task *Task) GetTimeLeft() time.Duration {
	return task.TaskDeadline.Sub(time.Now()) // ? S1024 - find out what's the difference between time.Until and t.Sub
}
