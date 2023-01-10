package embedpg

import (
	"context"
	"database/sql"
	"fmt"
	"path"
	"time"

	epg "github.com/fergusstrange/embedded-postgres"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4/adapter/postgresql"
)

type EmbedPg struct {
	epg    *epg.EmbeddedPostgres
	conf   *postgresql.ConnectionURL
	folder string
	port   uint32
}

func New(folder string, port int) *EmbedPg {

	uconf, err := postgresql.ParseURL(fmt.Sprintf("postgres://demo:demo123@localhost:%d/demo?sslmode=disable", port))
	if err != nil {
		panic(err)
	}

	postgres := epg.NewDatabase(epg.DefaultConfig().
		Username(uconf.User).
		Password(uconf.Password).
		Database(uconf.Database).
		Version(epg.V12).
		RuntimePath(path.Join(folder, "./runtime")).
		DataPath(path.Join(folder, "./data")).
		Port(uint32(port)).
		StartTimeout(45 * time.Second))

	return &EmbedPg{
		epg:    postgres,
		port:   uint32(port),
		folder: folder,
		conf:   uconf,
	}
}

func (d *EmbedPg) GetPort() uint32 {
	return d.port
}

func (d *EmbedPg) Start() error {
	return d.epg.Start()
}

func (d *EmbedPg) Stop() error {
	return d.epg.Stop()
}

func (d *EmbedPg) CleanFolder() error {
	// "tmp/pgdata"

	return nil

}

func (d *EmbedPg) RunSchema(schema string) error {
	sess, err := postgresql.Open(d.conf)
	if err != nil {
		return err
	}

	exist, _ := sess.Collection("tenants").Exists()
	if exist {
		pp.Println("Schema looks fine, skipping")
		return nil
	}

	ctx, cFunc := context.WithTimeout(context.Background(), time.Minute*5)

	defer cFunc()

	pp.Println("before Creating schema")
	time.Sleep(time.Second)

	sdriver := sess.Driver().(*sql.DB)
	_, err = sdriver.ExecContext(ctx, schema)
	if err != nil {
		pp.Println("error occured while creating schema")
		return err
	}
	pp.Println("after Creating schema, looks good")

	return nil
}

// private
