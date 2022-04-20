package models

import "time"

type taskList struct {
	TaskId          uint       `json:"task_id"`
	TaskName        string     `json:"task_name"`
	TaskDescription string     `json: "task_description`
	TaskDeadline    time.Time  `json: "task_deadline"`
	SubtaskList     []taskList `json:"SubtaskList"`
}
