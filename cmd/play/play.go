package main

import (
	"database/sql"
	"fmt"

	"github.com/k0kubun/pp"

	_ "github.com/mattn/go-sqlite3"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
)

func main() {
	// Connect to SQLite database

	const memURL = "file::memory:?mode=memory&cache=shared"

	db, err := sql.Open("sqlite3", memURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query(`SELECT json('{"JSON": true}');`)
	if err != nil {
		panic(err)
	}

	pp.Println(dbutils.SelectScan(result))

	// Create table for testing
	_, err = db.Exec(`
		CREATE TABLE tablexyz (
			__id INTEGER,
			__mod_ctx TEXT,
			__version INTEGER,
			data TEXT
		);

		CREATE TABLE tablexyz_log (
			id INTEGER,
			row_id INTEGER,
			mod_sig TEXT,
			payload TEXT,
			user_id TEXT,
			type TEXT,
			row_version INTEGER,
			data TEXT
		)


	`)
	if err != nil {
		panic(err)
	}

	// Create trigger
	_, err = db.Exec(`

		CREATE TRIGGER tablexyz_trigger_insert
			AFTER INSERT ON tablexyz
		BEGIN
			INSERT INTO tablexyz_log(row_id, payload, type, row_version, data)
				VALUES (NEW.__id, json_object('data', NEW.data) , 'INSERT', NEW.__version, NEW.data);
		END;

		CREATE TRIGGER tablexyz_trigger_update
			AFTER UPDATE ON tablexyz
		BEGIN
			INSERT INTO tablexyz_log(row_id, payload, type, row_version, data)
				VALUES (NEW.__id, json_object('data', NEW.data), 'UPDATE', NEW.__version, NEW.data);
		END;




	`)
	if err != nil {
		panic(err)
	}

	// Insert test data
	_, err = db.Exec(`
		INSERT INTO tablexyz (__id, __mod_ctx, __version, data)
		VALUES (1, '', 0, 'testxyzyzyyzyzyzy')
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Test harness completed successfully")
}
