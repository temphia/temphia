package dyncore

import "github.com/upper/db/v4"

func GroupTable(sess db.Session) db.Collection {
	return sess.Collection("data_table_groups")
}

func Table(sess db.Session) db.Collection {
	return sess.Collection("data_tables")
}

func TableColumn(sess db.Session) db.Collection {
	return sess.Collection("data_table_columns")
}
