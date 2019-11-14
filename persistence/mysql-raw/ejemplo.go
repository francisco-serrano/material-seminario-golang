package mysql_raw

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

func CloseStatement(stmt *sql.Stmt) {
	if err := stmt.Close(); err != nil {
		panic(err)
	}
}

type Car struct {
	Id    int
	Brand string
	Model string
}

func Run() {
	db, err := sql.Open("mysql", "root:root@/sample_database")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	queryCreate := "create table if not exists `cars` (" +
		"id integer not null auto_increment," +
		"brand varchar(255)," +
		"model varchar(255)," +
		"primary key (id))"

	stmtCreate, _ := db.Prepare(queryCreate)
	if _, err = stmtCreate.Exec(); err != nil {
		panic(err)
	}
	defer CloseStatement(stmtCreate)

	queryInsert := "insert into `cars` (brand, model) values (?, ?);"

	stmtInsert, _ := db.Prepare(queryInsert)
	if _, err = stmtInsert.Exec("vw", "golf"); err != nil {
		panic(err)
	}
	defer CloseStatement(stmtInsert)

	querySelect := "select * from `cars` where id = ?"

	stmtSelect, _ := db.Prepare(querySelect)
	var aux Car
	if err = stmtSelect.QueryRow(1).Scan(&aux.Id, &aux.Brand, &aux.Model); err != nil {
		panic(err)
	}

	fmt.Println("aux", aux)

	defer CloseStatement(stmtSelect)
}

func RunSelect() {
	db, err := sql.Open("mysql", "root:root@/sample_database")
	if err != nil {
		panic(err) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	selectQuery := "select * from `cars`"

	stmtSelect, _ := db.Prepare(selectQuery)
	defer stmtSelect.Close()

	rows, _ := stmtSelect.Query()
	for rows.Next() {
		var id, model, brand string

		if err = rows.Scan(&id, &model, &brand); err != nil {
			panic(err)
		}

		fmt.Println(id, model, brand)
	}

}
