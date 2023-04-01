package main

import (
	"archive/zip"
	"io"
	"os"
	"time"

	"github.com/k0kubun/pp"
)

func IoMain() {

	file, err := os.CreateTemp(os.TempDir(), "mnop.zip")
	if err != nil {
		panic(err)
	}

	writer := zip.NewWriter(file)
	wr, err := writer.Create("xyz")
	if err != nil {
		panic(err)
	}

	pp.Println(wr.Write([]byte(`Test 123`)))
	pp.Println(writer.Close())

	file.Seek(0, 0)

	doSth(file)

	time.Sleep(time.Hour)

}

func doSth(reader io.ReadCloser) {
	file, err := os.CreateTemp(os.TempDir(), "psqr.zip")
	if err != nil {
		panic(err)
	}

	pp.Println(io.Copy(file, reader))
}
