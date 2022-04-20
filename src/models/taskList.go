package models

import "time"

type taskList struct {
	TaskId          uint
	TaskName        string
	TaskDescription string
	TaskDeadline    time.Time
}
