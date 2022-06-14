package models

import (
	"time"
)

// TODO: fix pointers
//what's the problem?
type Task struct {
	Model
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_description"`
	TaskDeadline    time.Time `json:"task_deadline"`
	TaskLists       []Task    `json:"SubtaskList"`
}

func (task *Task) GetTimeLeft() time.Duration {
	return task.TaskDeadline.Sub(time.Now())
}

func (task *Task) SetDescription(newDescription string) {
	task.TaskDescription = newDescription
}
