package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const Username, Pass, DB_Name = "root", "strongRoots@911", "learning"

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

type Data struct {
	id   int
	name string
}

func main() {

	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", Username, Pass, DB_Name)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	result, err := db.Exec("INSERT INTO data VALUES(4, \"Pineapple\")")
	checkError(err)
	lastInsertedId, err := result.LastInsertId()
	checkError(err)
	fmt.Println("last inserted id: ", lastInsertedId)
	rowsAffected, err := result.RowsAffected()
	checkError(err)
	fmt.Println("rows affected: ", rowsAffected)

	//select and display all rows in database
	rows, err := db.Query("SELECT * FROM data")
	checkError(err)
	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}
}
