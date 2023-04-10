package sqlite

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
)

func temphiaDeleteRecord(conn *sqlite3.SQLiteConn) func(table, ctx, rowids string) error {

	// fixme add txn ?

	return func(table, ctx, rowids string) error {

		results := gjson.GetMany(ctx, "user_id", "user_sign", "init_sign")

		userId := results[0].String()
		userSign := results[1].String()
		initSign := results[2].String()

		activityTable := fmt.Sprintf("xd_%s_activity", table)

		var buf bytes.Buffer

		buf.Write([]byte("DELETE FROM "))
		buf.WriteString(table)
		buf.Write([]byte("WHERE __id IN ("))

		rowStrs := strings.Split(rowids, ",")
		vids := make([]driver.Value, 0, len(rowStrs))

		for idx, rstr := range rowStrs {
			if idx != 0 {
				buf.Write([]byte(`, `))
			}
			buf.WriteByte('?')

			id, _ := strconv.ParseInt(rstr, 10, 64)
			vids = append(vids, id)
		}

		buf.Write([]byte(`);`))

		_, err := conn.Exec(buf.String(), vids)
		if err != nil {
			return err
		}

		errs := []error{}

		for _, vid := range vids {
			_, err := conn.Exec(
				fmt.Sprintf("INSERT INTO %s( type, row_id, row_version, user_id, user_sign, init_sign) VALUES (?, ?, ?, ?, ?, ?);", activityTable),
				[]driver.Value{"delete", vid, 0, userId, userSign, initSign},
			)
			errs = append(errs, err)
		}

		return errors.Join(errs...)
	}

}
