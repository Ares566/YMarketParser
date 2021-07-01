package clickhouse

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go"
)

type InstanceClickhouse struct {
	DB *sql.DB // connection
}

func NewClickHouseConnection() InstanceClickhouse {
	connect, err := sql.Open("localhost", "tcp://127.0.0.1:9000?username=default&password=123456&database=test")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		// TODO обработка ошибки
		log.Fatal(err)
	}

	instance := InstanceClickhouse{DB: connect}
	InitClickhouseTables(&instance)

	fmt.Println("Successfully connected clickhouse")

	return instance

	// var (
	// 	tx, _   = connect.Begin()
	// 	stmt, _ = tx.Prepare("INSERT INTO example (country_code, os_id, browser_id, categories, action_day, action_time) VALUES (?, ?, ?, ?, ?, ?)")
	// )
	// defer stmt.Close()

	// for i := 0; i < 100; i++ {
	// 	if _, err := stmt.Exec(
	// 		"RU",
	// 		10+i,
	// 		100+i,
	// 		clickhouse.Array([]int16{1, 2, 3}),
	// 		time.Now(),
	// 		time.Now(),
	// 	); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// if err := tx.Commit(); err != nil {
	// 	log.Fatal(err)
	// }

	// rows, err := connect.Query("SELECT country_code, os_id, browser_id, categories, action_day, action_time FROM example")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var (
	// 		country               string
	// 		os, browser           uint8
	// 		categories            []int16
	// 		actionDay, actionTime time.Time
	// 	)
	// 	if err := rows.Scan(&country, &os, &browser, &categories, &actionDay, &actionTime); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s", country, os, browser, categories, actionDay, actionTime)
	// }

	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := connect.Exec("DROP TABLE example"); err != nil {
	// 	log.Fatal(err)
	// }

}

func InitClickhouseTables(clickhouse *InstanceClickhouse) {

	_, err := clickhouse.DB.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id			UUID
		url 		String,
		categories  Array(String),
		name 		String,
		label 		String,
		mark 		String,
		rpl 		String,
		oldPrice 	String,
		discount 	String,
		curPrice 	String,
		action_day   Date
	) engine=Memory
	`)

	if err != nil {
		log.Fatal(err)
	}
}
