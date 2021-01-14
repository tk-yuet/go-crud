package models

import (
	"encoding/json"
	"fmt"
	"log"

	. "github.com/tk-yuet/go-crud/internal/app/go-crud/database"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct {
	DbClient *MySqlClient
	User     *User
}

func NewUser(id int64, username string, pw string) *User {
	newInstance := User{
		Id:       id,
		Username: username,
		Password: pw,
	}
	return &newInstance
}

func NewUserModel(cli *MySqlClient, tm *User) *UserModel {
	newInstance := UserModel{
		DbClient: cli,
		User:     tm,
	}
	return &newInstance
}

func (um *UserModel) TableName() string {
	return "users"
}

func (u *User) ToJson() string {
	var jsonData []byte
	jsonData, err := json.Marshal(u)
	if err != nil {
		log.Println("err: ", err)
	}
	return string(jsonData)
}

func (um *UserModel) Insert() int64 {
	db := um.DbClient.Database
	sql := fmt.Sprintf("INSERT INTO %s(username, password) VALUES (?,?)", um.TableName())
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(um.User.Username, um.User.Password)
	lid, err := res.LastInsertId()
	um.User.Id = lid

	return lid
}

func (um *UserModel) Select(id int64) {
	db := um.DbClient.Database
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", um.TableName())
	res, err := db.Query(sql, id)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	if res.Next() {
		var user User
		err := res.Scan(&user.Id, &user.Username, &user.Password)

		if err != nil {
			log.Fatal(err)
		}

		um.User = &user
	} else {

		fmt.Println("No user found")
	}
}

func (um *UserModel) FindByUsername(username string) {
	db := um.DbClient.Database
	sql := fmt.Sprintf("SELECT * FROM %s WHERE username = ?", um.TableName())
	res, err := db.Query(sql, username)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	if res.Next() {
		var user User
		err := res.Scan(&user.Id, &user.Username, &user.Password)

		if err != nil {
			log.Fatal(err)
		}

		um.User = &user
	} else {

		fmt.Println("No user found")
	}
}

func (um *UserModel) Update(updateParams *User) {
	db := um.DbClient.Database
	if len(updateParams.Username) != 0 {
		um.User.Username = updateParams.Username
	}
	if len(updateParams.Password) != 0 {
		um.User.Password = updateParams.Password
	}
	sql := fmt.Sprintf("UPDATE %s SET username=?, password=? WHERE id=?", um.TableName())
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(um.User.Username, um.User.Password, um.User.Id)

}

func (um *UserModel) Delete(id int64) int64 {
	db := um.DbClient.Database
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?", um.TableName())

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
