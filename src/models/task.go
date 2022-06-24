package models

import (
	"time"
)

type Task struct {
	Model
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_description"`
	TaskDeadline    time.Time `json:"task_deadline"`
	//TaskLists       []Task    `json:"SubtaskList" gorm:"many2many:subtask_list"`
	//IsSubTask       bool      `json:"is_subtask"`
	//UserId          uint      `json:"user_id" gorm:"foreignKey: UserId"`
	//
	//       this is disabled temporarily until minimum functionality is achieved
}

func (task *Task) GetTimeLeft() time.Duration {
	return task.TaskDeadline.Sub(time.Now())
}
