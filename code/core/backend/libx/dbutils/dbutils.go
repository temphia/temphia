package dbutils

import (
	"database/sql"
	"fmt"

	"github.com/alecthomas/repr"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

func Execute(driver *sql.DB, qstr string) error {
	fmt.Printf(qstr)

	// fixme => check syntax using explain or sth
	_, err := driver.Exec(qstr)

	if err != nil {
		pp.Println("ERROR =>", err.Error())
	}

	return err
}

func Table(sess db.Session, name string) db.Collection {
	return sess.Collection(name)
}

func SelectScan(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	numColumns := len(columns)

	values := make([]interface{}, numColumns)
	for i := range values {
		values[i] = new(interface{})
	}

	var results []map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		dest := make(map[string]interface{}, numColumns)
		for i, column := range columns {
			dest[column] = *(values[i].(*interface{}))
		}
		results = append(results, dest)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func GetScan(rows *sql.Rows) (map[string]interface{}, error) {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	numColumns := len(columns)

	if !rows.Next() {
		repr.Println(rows.Err())
		return nil, sql.ErrNoRows
	}

	values := make([]interface{}, numColumns)
	for i := range values {
		values[i] = new(interface{})
	}

	if err := rows.Scan(values...); err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, numColumns)
	for i, column := range columns {
		result[column] = *(values[i].(*interface{}))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
