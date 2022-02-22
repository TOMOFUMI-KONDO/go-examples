package sql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id   *int
	name *string
)

func Record() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fetch(db)
	prepareFetch(db)
	fetchSingle(db)
}

func fetch(db *sql.DB) {
	rows, err := db.Query(
		"select e.id from experiments e join networks n on e.network_id = n.id where n.name = ?",
		"tcp",
	)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("failed to scan row: %s", err)
		}
		fmt.Printf("id: %d\n", *id)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
}

func prepareFetch(db *sql.DB) {
	stmt, err := db.Prepare("select e.id from experiments e join networks n on e.network_id = n.id where n.name = ?")
	if err != nil {
		log.Fatalf("failed to prepare: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query("tcp")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	fmt.Println("case tcp")
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("failed to scan row: %s", err)
		}
		fmt.Printf("id: %d\n", *id)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("case quic")
	rows, err = stmt.Query("quic")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("failed to scan row: %s", err)
		}
		fmt.Printf("id: %d\n", *id)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
}

func fetchSingle(db *sql.DB) {
	if err := db.QueryRow(`select e.id, n.name from experiments e
join networks n on e.network_id = n.id
where n.name = ?
order by e.created_at desc
limit 1`, "tcp",
	).Scan(&id, &name); err != nil {
		log.Fatalf("failed to query row: %v", err)
	}
	fmt.Printf("id: %d, name: %s\n", *id, *name)
}

