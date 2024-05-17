package main

import (
	"fmt"
	"log"

	//"testing"
	"time"

	"restdb"
)

type User struct {
	ID        int
	Username  string
	Password  string
	LastLogin int64
	Admin     int
	Active    int
}

// PostgreSQL Connection details
var (
	Hostname = "localhost"
	Port     = 5433
	Username = "postgres"
	Password = "123"
	Database = "restapi"
)

// func TestDB(t *testing.T) {
// 	t.Run("test insertUser func", func(t *testing.T) {
// 		db := restdb.ConnectPostgres()
// 		defer db.Close()

// 		user := restdb.User{ID: 0, Username: "nick", Password: "admin", LastLogin: time.Now().Unix(), Admin: 1, Active: 1}
// 		if !restdb.InsertUser(user) {
// 			t.Errorf("got %v want %v", user, restdb.InsertUser(user))
// 		}
// 	})
// }

func main() {
	db := restdb.ConnectPostgres()
	fmt.Println(db)
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("Ping:", err)
		return
	}

	t := restdb.User{}
	fmt.Println(t)
	rows, err := db.Query(`SELECT "username" FROM "users"`)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(username)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Populating PostgreSQL")
	user := restdb.User{ID: 0, Username: "nick", Password: "admin", LastLogin: time.Now().Unix(), Admin: 1, Active: 1}
	if restdb.InsertUser(user) {
		fmt.Println("User inserted successfully.")
	} else {
		fmt.Println("Insert failed!")
	}

	nickUser := restdb.FindUserUsername(user.Username)
	fmt.Println("nick: ", nickUser)

	if restdb.DeleteUser(nickUser.ID) {
		fmt.Println("User Deleted.")
	} else {
		fmt.Println("User not Deleted.")
	}

	nickUser = restdb.FindUserUsername(user.Username)
	fmt.Println("nick: ", nickUser)

	if restdb.DeleteUser(nickUser.ID) {
		fmt.Println("User Deleted.")
	} else {
		fmt.Println("User not Deleted.")
	}

	if restdb.DeleteUser(nickUser.ID) {
		fmt.Println("User Deleted.")
	} else {
		fmt.Println("User not Deleted.")
	}
}
