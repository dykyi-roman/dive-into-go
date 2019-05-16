package Infrastructure

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	Domain "github.com/dykyi-roman/dive-into-go/server/Domain/Model"
	_ "github.com/go-sql-driver/mysql"
)

type MySLQClient struct{}

func (obj MySLQClient) dbConnect() (db *sql.DB) {

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_PASSWORD"),
	))

	if err != nil {
		panic(err.Error())
	}

	return db
}

func (obj MySLQClient) Get() []Domain.User {
	db := obj.dbConnect()

	stmt, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	model := Domain.User{}
	res := []Domain.User{}
	for stmt.Next() {
		var id int
		var name string
		var age int
		err = stmt.Scan(&id, &name, &age)
		if err != nil {
			panic(err.Error())
		}
		model.Id = id
		model.Name = name
		model.Age = age
		res = append(res, model)
	}

	defer db.Close()

	return res
}

func (obj MySLQClient) Insert(r *http.Request) bool {
	db := obj.dbConnect()
	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		stmt, err := db.Prepare("INSERT INTO user(name, age) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		stmt.Exec(name, age)
	}
	defer db.Close()

	return true
}

func (obj MySLQClient) Delete(r *http.Request) bool {
	db := obj.dbConnect()

	id := r.URL.Query().Get("id")
	stmt, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(id)
	defer db.Close()

	return true
}

func (obj MySLQClient) Update(r *http.Request) bool {
	db := obj.dbConnect()

	//TODO:: To be continue...

	defer db.Close()
	return true
}
