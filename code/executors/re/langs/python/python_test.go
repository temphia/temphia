package python

import (
	"testing"
)

func TestPyRunner(t *testing.T) {

}

/*



		var buf bytes.Buffer

		z := zip.NewWriter(&buf)
		defer z.Close()

		for _, fk := range []string{"main.py", "start.sh"} {

			rfile, err := os.Open(path.Join("testdata", fk))
			if err != nil {
				return nil, err
			}

			defer rfile.Close()

			wfile, err := z.Create(fk)
			if err != nil {
				return nil, err
			}

			if _, err := io.Copy(wfile, rfile); err != nil {
				return nil, err
			}
		}

		return nil, nil
	},

*/
