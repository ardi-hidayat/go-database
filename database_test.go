package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('budi', 'Budi')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
	}
}

func TestQueryCust(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("==================================")
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date : ", birthDate.Time)
		}
		fmt.Println("Married : ", married)
		fmt.Println("Created At : ", createdAt)

	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "admin"

	// script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login : ", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "eko';DROP TABLE user; #"
	password := "eko"

	script := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "ard@gmail.com"
	comment := "test komentar"

	script := "INSERT INTO comments(email, comment) VALUES(?,  ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment", insertId)
}

func TestPrepareStmt(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "ard" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Comment Id : ", lastInsertId)
	}

}
