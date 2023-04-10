package sqlite

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
)

func temphiaDeleteRecord(conn *sqlite3.SQLiteConn) func(table, ctx, rowids string) error {

	// fixme add txn ?

	return func(table, ctx, rowids string) error {

		pp.Println("@temphia_delete_record", table, ctx, rowids)

		results := gjson.GetMany(ctx, "user_id", "user_sign", "init_sign")

		userId := results[0].String()
		userSign := results[1].String()
		initSign := results[2].String()

		activityTable := fmt.Sprintf("%s_activity", table)

		var buf bytes.Buffer

		buf.Write([]byte("DELETE FROM "))
		buf.WriteString(table)
		buf.Write([]byte(" WHERE __id IN ("))

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

		delquery := buf.String()

		pp.Println("@EXEC", delquery)

		_, err := conn.Exec(delquery, vids)
		if err != nil {
			pp.Println("@err_while_deleting", err)
			return err
		}

		errs := []error{}

		for _, vid := range vids {
			qstr := fmt.Sprintf("INSERT INTO %s( type, row_id, row_version, user_id, user_sign, init_sign) VALUES (?, ?, 0, ?, ?, ?);", activityTable)
			params := []driver.Value{"delete", vid, userId, userSign, initSign}
			_, err := conn.Exec(qstr, params)

			errs = append(errs, err)
		}

		pp.Println("@errs_@activity", errs)

		return errors.Join(errs...)
	}

}
