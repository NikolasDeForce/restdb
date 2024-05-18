package main

import (
	"testing"

	"github.com/NikolasDeForce/restdb"
)

// PostgreSQL Connection details
var (
	Hostname = "localhost"
	Port     = 5433
	Username = "postgres"
	Password = "123"
	Database = "restapi"
)

func TestDB(t *testing.T) {
	var user = restdb.User{ID: 2, Username: "nick", Password: "admin", LastLogin: 1716000604, Admin: 1, Active: 1}

	t.Run("check InsetUser, should return user already exist", func(t *testing.T) {
		db := restdb.ConnectPostgres()
		defer db.Close()

		got := restdb.InsertUser(user)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("find user by Username", func(t *testing.T) {
		db := restdb.ConnectPostgres()
		defer db.Close()

		got := restdb.FindUserUsername(user.Username)

		if got != user {
			t.Errorf("got %v want %v", got, user)
		}
	})

	t.Run("find user by ID", func(t *testing.T) {
		db := restdb.ConnectPostgres()
		defer db.Close()

		got := restdb.FindUserID(user.ID)

		if got != user {
			t.Errorf("got %v want %v", got, user)
		}
	})

	t.Run("check user admin or not", func(t *testing.T) {
		db := restdb.ConnectPostgres()
		defer db.Close()

		if !restdb.IsUserAdmin(user) {
			t.Errorf("got %v want %v", restdb.IsUserAdmin(user), true)
		}
	})

	t.Run("check valid user or not", func(t *testing.T) {
		db := restdb.ConnectPostgres()
		defer db.Close()

		if !restdb.IsUserValid(user) {
			t.Errorf("got %v want %v", restdb.IsUserValid(user), true)
		}
	})
}
