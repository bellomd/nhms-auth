package authrepo

import (
	"log"

	"github.com/jmoiron/sqlx"

	// The import will the used by other files and packages.
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE t_user (
	id SERIAL PRIMARY KEY,
	first_name varchar(60),
	last_name varchar(60),
	email varchar(250),
	password varchar(60),
	phone_number varchar(60),
	address text,
	created_by varchar(255),
	last_modified_by varchar(255),
	created_date date,
	last_modified_date date
);`

// Db holds database connection.
var Db *sqlx.DB

const (
	// CreatedBy hold the default created by user.
	CreatedBy = "SYSTEM"
)

func init() {
	var err error
	Db, err = sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=@postgres!")
	if err != nil {
		log.Fatalln(err)
	}
	//Db.MustExec(schema)
}
