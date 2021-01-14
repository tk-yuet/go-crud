package models

import (
	"encoding/json"
	"fmt"
	"log"

	. "github.com/tk-yuet/go-crud/internal/app/go-crud/database"
)

type Task struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Due        string `json:"due"`
	Completion string `json:"completion"`
	Status     string `json:"status"`
}

type TaskModel struct {
	DbClient *MySqlClient
	Task     *Task
}

func NewTask(id int64, name string, due string, completion string, status string) *Task {
	newInstance := Task{
		Id:         id,
		Name:       name,
		Due:        due,
		Completion: completion,
		Status:     status,
	}
	return &newInstance
}

func NewTaskModel(cli *MySqlClient, tm *Task) *TaskModel {
	newInstance := TaskModel{
		DbClient: cli,
		Task:     tm,
	}
	return &newInstance
}

func (tm *TaskModel) TableName() string {
	return "tasks"
}

func (tm *Task) ToJson() string {
	var jsonData []byte
	jsonData, err := json.Marshal(tm)
	if err != nil {
		log.Println("err: ", err)
	}
	return string(jsonData)
}

func (tm *TaskModel) Insert() int64 {
	db := tm.DbClient.Database
	sql := fmt.Sprintf("INSERT INTO %s(name, due, completion, status) VALUES (?,?,?,?)", tm.TableName())
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(tm.Task.Name, tm.Task.Due, tm.Task.Completion, tm.Task.Status)
	lid, err := res.LastInsertId()
	tm.Task.Id = lid

	return lid
}

func (tm *TaskModel) Select(id int64) {
	db := tm.DbClient.Database
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tm.TableName())
	res, err := db.Query(sql, id)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	if res.Next() {
		var task Task
		err := res.Scan(&task.Id, &task.Name, &task.Due, &task.Completion, &task.Status)

		if err != nil {
			log.Fatal(err)
		}
		tm.Task = &task
	} else {

		fmt.Println("No task found")
	}
}

func (tm *TaskModel) Update(updateParams *Task) {
	db := tm.DbClient.Database
	if len(updateParams.Name) != 0 {
		tm.Task.Name = updateParams.Name
	}
	if len(updateParams.Due) != 0 {
		tm.Task.Due = updateParams.Due
	}
	if len(updateParams.Completion) != 0 {
		tm.Task.Completion = updateParams.Completion
	}
	if len(updateParams.Status) != 0 {
		tm.Task.Status = updateParams.Status
	}
	sql := fmt.Sprintf("UPDATE %s SET name=?, due=?, completion=?, status = ? WHERE id=?", tm.TableName())
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(tm.Task.Name, tm.Task.Due, tm.Task.Completion, tm.Task.Status, tm.Task.Id)

}

func (tm *TaskModel) Delete(id int64) int64 {
	db := tm.DbClient.Database
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tm.TableName())

	res, err := db.Exec(sql, id)

	if err != nil {
		panic(err.Error())
	}

	affectedRows, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return affectedRows
}
