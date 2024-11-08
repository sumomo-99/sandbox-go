package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/marcboeker/go-duckdb"
	"log"
)

func db() error {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		return err
	}
	defer db.Close()
	defer fmt.Println("closing db")

	_, err = db.Exec("CREATE TABLE test (a INTEGER, b VARCHAR)")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO test VALUES (1, 'foo'), (2, 'bar')")
	if err != nil {
		return err
	}

	var (
		a int
		b string
	)

	row := db.QueryRow("SELECT a, b FROM test")
	err = row.Scan(&a, &b)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("no rows")
	} else if err != nil {
		return err
	}

	fmt.Println(a, b)

	return nil
}

func main() {
	err := db()
	if err != nil {
		log.Fatal(err)
	}
}

