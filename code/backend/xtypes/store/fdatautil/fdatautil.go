package fdatautil

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func ReadAndClose(data store.FData) ([]byte, error) {
	defer data.Close()

	out, err := data.AsBytes()
	return out, err
}

type Peekable interface {
	InnerFile() string
}

func ExtractZipAndClose(data store.FData, target string) error {

	pdata, ok := data.(Peekable)
	afile := ""

	defer func() {
		data.Close()
		os.Remove(afile)

	}()

	if ok {
		afile = pdata.InnerFile()
	} else {

		zfile, err := os.CreateTemp("", "*temphia_cab_folder.zip")
		if err != nil {
			return err
		}

		reader, err := data.AsReader()
		if err != nil {
			return err
		}

		_, err = io.Copy(zfile, reader)
		if err != nil {
			return err
		}

		err = zfile.Sync()
		if err != nil {
			return err
		}

		afile = zfile.Name()
		zfile.Close()

	}

	return xutils.ExtractZip(pdata.InnerFile(), target)

}

func WriteAndClose(rw http.ResponseWriter, file string, data store.FData) {
	defer data.Close()

	out, err := data.AsBytes()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ffiles := strings.Split(file, ".")

	ctype := ""
	switch ffiles[1] {
	case "js":
		ctype = httpx.CtypeJS
	case "css":
		ctype = httpx.CtypeCSS
	default:
		ctype = http.DetectContentType(out)
	}

	rw.Header().Set("Context-Type", ctype)
	rw.Write(out)

}
