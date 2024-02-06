package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func execute(body func(db *sql.DB)) error {
	connection := "user=postgres password=2458173671 dbname=modding_server sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return err
	}

	body(db)

	db.Close()
	return err
}

func processSingle(statement string, queryArgs []any, scanArgs []any) error {
	var err error = nil
	body := func(db *sql.DB) {
		row := db.QueryRow(statement, queryArgs...)
		err = row.Scan(scanArgs...)
	}

	execute(body)

	return err
}

func Create(table string, columns string, values string, args []any) error {
	return processSingle(fmt.Sprintf("insert into %s (%s) values (%s) returning *", table, columns, values), args[1:], args)
}

func Read(table string, args []any) error {
	return processSingle(fmt.Sprintf("select * from %s where id = $1", table), args[0:1], args)
}

func Update(table string, columns string, args []any) error {
	return processSingle(fmt.Sprintf("update %s set %s where id = $1 returning *", table, columns), args, args)
}

func Delete(table string, args []any) error {
	return processSingle(fmt.Sprintf("delete from %s where id = $1 returning *", table), args[0:1], args)
}

func GetList(table string, target string, key string, owner_id int) ([]int, error) {
	var err error = nil
	var ids []int
	body := func(db *sql.DB) {
		var rows *sql.Rows
		rows, err = db.Query(fmt.Sprintf("select %s from %s where %s = $1", target, table, key), owner_id)
		if err != nil {
			return
		}

		for rows.Next() {
			var id int
			err = rows.Scan(&id)
			if err != nil {
				return
			}

			ids = append(ids, id)
		}
	}

	execute(body)

	return ids, err
}
