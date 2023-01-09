package embedpg

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"time"

	epg "github.com/fergusstrange/embedded-postgres"
	"github.com/upper/db/v4/adapter/postgresql"
)

type EmbedPg struct {
	epg    *epg.EmbeddedPostgres
	conf   *postgresql.ConnectionURL
	folder string
	port   uint32
}

func New(folder string) *EmbedPg {
	port, err := getFreePort()
	if err != nil {
		panic(err)
	}

	uconf, err := postgresql.ParseURL(fmt.Sprintf("postgres://demo:demo123@localhost:%d/demo?sslmode=disable", port))
	if err != nil {
		panic(err)
	}

	postgres := epg.NewDatabase(epg.DefaultConfig().
		Username(uconf.User).
		Password(uconf.Password).
		Database(uconf.Database).
		Version(epg.V12).
		RuntimePath(folder).
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

	exist, _ := sess.Collection("users").Exists()

	if exist {
		return nil
	}

	ctx, cFunc := context.WithTimeout(context.Background(), time.Minute*5)

	defer cFunc()

	sdriver := sess.Driver().(*sql.DB)
	_, err = sdriver.ExecContext(ctx, schema)
	if err != nil {
		return err
	}

	return nil
}

// private

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
