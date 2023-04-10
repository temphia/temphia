package main

import (
	"database/sql"
	"fmt"

	"github.com/k0kubun/pp"

	_ "github.com/mattn/go-sqlite3"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
)

func main() {

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

	_, err = db.Exec(`
		CREATE TABLE tablexyz (
			__id integer primary key autoincrement not null,
			__mod_ctx text,
			__version integer,
			title text not null,
			count integer not null
		);

		CREATE TABLE tablexyz_log (
			id integer primary key autoincrement not null,
			type text not null,
			row_id integer not null,
			row_version integer not null,
			user_id text not null DEFAULT '',
			user_sign text not null DEFAULT '',
			init_sign text not null DEFAULT '',
			payload text not null DEFAULT '',
			message text not null DEFAULT '',
			created_at text not null default current_timestamp
		)


	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`

			CREATE TRIGGER tablexyz_trigger_insert
				AFTER INSERT ON tablexyz
			BEGIN
				INSERT INTO tablexyz_log(
					type,
					row_id,
					row_version,
					user_id,
					user_sign,
					init_sign,
					payload
				)
				VALUES (
					'insert', 
					NEW.__id, 
					NEW.__version, 
					COALESCE(json_extract(NEW.__mod_ctx, '$.user_id' ), ''), 
					COALESCE(json_extract(NEW.__mod_ctx, '$.user_sign'), ''), 
					COALESCE(json_extract(NEW.__mod_ctx, '$.init_sign'), ''),
					json_object(
						'title', NEW.title
						)
					);
			END;


			CREATE TRIGGER tablexyz_trigger_update
				AFTER UPDATE ON tablexyz
			BEGIN
				INSERT INTO tablexyz_log(
					type,
					row_id,
					row_version,
					user_id,
					user_sign,
					init_sign,
					payload
				)
				VALUES (
					'update', 
					NEW.__id, 
					NEW.__version, 
					COALESCE(json_extract(NEW.__mod_ctx, '$.user_id' ), ''), 
					COALESCE(json_extract(NEW.__mod_ctx, '$.user_sign'), ''), 
					COALESCE(json_extract(NEW.__mod_ctx, '$.init_sign'), ''),
					json_object(
						'title', NEW.title
					)
				);
			END

		`)
	if err != nil {
		panic(err)
	}

	// user_sign

	_, err = db.Exec(`
			INSERT INTO tablexyz ( __mod_ctx, __version, title, count)
			VALUES ('{"user_id": "user12"}', 0, 'ping pong', 12)
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
			UPDATE tablexyz SET  __mod_ctx = '{"user_id": "user13"}', title='moded', __version = __version + 1   WHERE __id = 1;
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Test harness completed successfully")
}
