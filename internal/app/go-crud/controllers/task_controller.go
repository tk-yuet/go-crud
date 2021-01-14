package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tk-yuet/go-crud/internal/app/go-crud/models"
)

type createResponse struct {
	Task string `json:"task"`
}

func (env *ControllerEnv) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask *models.Task
	var tm *models.TaskModel

	if r.Method == "POST" {
		name := r.FormValue("name")
		due := r.FormValue("due")
		completion := r.FormValue("completion")
		status := r.FormValue("status")
		newTask = models.NewTask(-1, name, due, completion, status)
		tm = models.NewTaskModel(env.DbClient, newTask)
		tm.Insert()
	}
	tm.Select(newTask.Id)
	fmt.Fprintf(w, "%s", tm.Task.ToJson())
}

func (env *ControllerEnv) ShowTask(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idstr := params.Get(":id")
	tm := models.NewTaskModel(env.DbClient, nil)
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	tm.Select(id)
	fmt.Fprintf(w, "%s", tm.Task.ToJson())
}

func (env *ControllerEnv) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateParams *models.Task
	params := r.URL.Query()
	idstr := params.Get(":id")
	err := json.NewDecoder(r.Body).Decode(&updateParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		log.Println(err.Error())
	}

	tm := models.NewTaskModel(env.DbClient, nil)
	tm.Select(id)
	tm.Update(updateParams)
	fmt.Fprintf(w, "%s", tm.Task.ToJson())
}

func (env *ControllerEnv) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idstr := params.Get(":id")
	tm := models.NewTaskModel(env.DbClient, nil)
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	tm.Delete(id)
	fmt.Fprintf(w, "Deleted %s", idstr)
}
