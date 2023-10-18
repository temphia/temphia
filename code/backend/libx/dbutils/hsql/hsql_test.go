package hsql

import (
	"strings"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/rqlite/sql"
)

var testCases = []string{
	`SELECT authors.author_name, authors.author_email, author_submissions.article_title FROM authors
	INNER JOIN author_submissions ON authors.author_email=author_submissions.author_email;`,

	`SELECT a, b from (select a,b,c from tableb);`,
	`SELECT mpr, tolower(mno) FROM table_name where mpr=12 AND mno < '12' AND (pqr1 = 23 OR  dde IS NULL);`,
	`SELECT distinct(mno) FROM table_name;`,
	`SELECT count(mno) FROM table_name;`,
	`SELECT COUNT(mno) FROM table_name;`,
	`SELECT Customers.customer_id, Customers.first_name, Orders.amount
	FROM Customers
	INNER JOIN Orders
	ON Customers.customer_id = Orders.customer;`,

	`SELECT author_name AS "Author Name", author_email 
	AS "Author Emails" FROM authors;`,

	`SELECT authors.author_name, authors.author_email, author_submissions.article_title FROM authors
	LEFT JOIN author_submissions ON authors.author_email=author_submissions.author_email`,

	`SELECT authors.author_name, authors.author_email, author_submissions.article_title FROM authors
	RIGHT JOIN author_submissions ON authors.author_email=author_submissions.author_email`,

	`SELECT a1.author_name AS Author_A, a2.author_name 
	AS Author_B, a1.author_email FROM authors a1, authors a2;`,

	`SELECT author_name, author_pay, 
	CASE
	WHEN author_pay < 30000 THEN "New author"
	WHEN author_pay > 60000 THEN "Experienced Author"
	ELSE "Budding author"
	END AS "Author Experience"
	FROM authors`,

	`SELECT author_name, IFNULL(author_pay, 10000) 
	AS "Pay" FROM authors;`,

	`SELECT INSTR(author_name, author_email) 
	AS MatchName from authors;`,

	`SELECT TRIM(author_name) 
	AS "Trimmed Names" FROM authors;`,
}

func TestHsql(t *testing.T) {

	for _, tstr := range testCases {

		pp.Println("@processing", tstr)

		parser := sql.NewParser(strings.NewReader(tstr))

		stmt, err := parser.ParseStatement()

		if err != nil {
			t.Fatal(err)
		}

		v := &Visitor{
			tenantId:        "default0",
			group:           "test1",
			tns:             nil, //tns.New("shared"),
			inverseAliasMap: make(map[string]string),
		}

		err = (sql.Walk(v, stmt))
		if err != nil {
			t.Fatal(err)
		}

		pp.Println("@transformed", stmt.String())
		pp.Println("@alias", v.inverseAliasMap)

	}
}
