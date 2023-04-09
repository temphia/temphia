package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to SQLite database

	const memURL = "file::memory:?mode=memory&cache=shared"

	db, err := sql.Open("sqlite3", memURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
			BEFORE INSERT ON tablexyz
		BEGIN
		


			INSERT INTO tablexyz_log(row_id, mod_sig, type, row_version, data)
				VALUES (NEW.__id, '', 'INSERT', NEW.__version, NEW.data);
		END;

		CREATE TRIGGER tablexyz_trigger_update
			BEFORE UPDATE ON tablexyz
		BEGIN
			INSERT INTO tablexyz_log(row_id, mod_sig, type, row_version, data)
				VALUES (NEW.__id, '', 'UPDATE', NEW.__version, NEW.data);
		END;




	`)
	if err != nil {
		panic(err)
	}

	// Insert test data
	_, err = db.Exec(`
		INSERT INTO tablexyz (__id, __mod_sig, __version, data)
		VALUES (1, '', 0, 'testxyzyzyyzyzyzy')
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Test harness completed successfully")
}
