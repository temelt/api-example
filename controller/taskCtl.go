package controller

import (
	"api-example/models"
	u "api-example/utils"
	"encoding/json"
	"net/http"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Task{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetTasksFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTasks(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
