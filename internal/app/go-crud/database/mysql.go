package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/tk-yuet/go-crud/internal/app/go-crud/config"
)

type MySqlClient struct {
	Database *sql.DB
	Config   *MySqlConfig
}

func NewMySqlClient(conf *MySqlConfig) *MySqlClient {
	newInstance := MySqlClient{
		Database: nil,
		Config:   conf,
	}
	return &newInstance
}

func (cli *MySqlClient) Connect() {
	db, err := sql.Open(cli.Config.DbDriver, cli.Config.ConnectionString())
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	cli.Database = db
}

func (cli *MySqlClient) Disconnect() {
	cli.Database.Close()
}

func (cli *MySqlClient) UseTestDbAndTable() {
	db := cli.Database

	sql := "CREATE DATABASE IF NOT EXISTS db;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Database db created successfully..")
	}

	_, err = db.Exec("USE db")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("DB selected successfully..")
	}

	sql = "DROP TABLE IF EXISTS `tasks`;"
	stmt, err = db.Prepare(sql)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Drop tasks created successfully..")
	}

	sql = "DROP TABLE IF EXISTS `users`;"
	stmt, err = db.Prepare(sql)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Drop users created successfully..")
	}

	sql = "CREATE TABLE `tasks`(" +
		"`id` int unsigned NOT NULL AUTO_INCREMENT," +
		"`name` varchar(30) NOT NULL," +
		"`due` date NOT NULL," +
		"`completion`  date NOT NULL," +
		"`status`  ENUM('yet-started', 'wip', 'completed')," +
		"PRIMARY KEY (id)" +
		") ENGINE=`InnoDB` AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;"

	stmt, err = db.Prepare(sql)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Table tasks created successfully..")
	}

	sql = "CREATE TABLE `users`(" +
		"`id` int unsigned NOT NULL AUTO_INCREMENT," +
		"`username` varchar(256) NOT NULL," +
		"`password` varchar(256) NOT NULL," +
		"PRIMARY KEY (id)," +
		"UNIQUE (username)" +
		") ENGINE=`InnoDB` AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;"

	stmt, err = db.Prepare(sql)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Table users created successfully..")
	}

}
