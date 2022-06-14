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
	TaskLists       []Task    `json:"SubtaskList" gorm: "many2many:subtask_list"`
	IsSubTask       bool      `json: "is_subtask"`
	UserId          uint      `json: "user_id" gorm: "foreignKey: UserId"`
}

func (task *Task) GetTimeLeft() time.Duration {
	return task.TaskDeadline.Sub(time.Now())
}
