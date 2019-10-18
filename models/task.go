package models

import (
	u "api-example/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title  string `json:"name"`
	Detail string `json:"phone"`
	UserId uint   `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (task *Task) Validate() (map[string]interface{}, bool) {

	if task.Title == "" {
		return u.Message(false, "Task title should be on the payload"), false
	}

	if task.Detail == "" {
		return u.Message(false, "Task detail should be on the payload"), false
	}

	if task.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (task *Task) Create() map[string]interface{} {

	if resp, ok := task.Validate(); !ok {
		return resp
	}

	GetDB().Create(task)

	resp := u.Message(true, "success")
	resp["task"] = task
	return resp
}

func GetTask(id uint) *Task {

	task := &Task{}
	err := GetDB().Table("tasks").Where("id = ?", id).First(task).Error
	if err != nil {
		return nil
	}
	return task
}

func GetTasks(user uint) []*Task {

	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", user).Find(&tasks).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tasks
}
