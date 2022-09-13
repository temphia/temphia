package self

import (
	"fmt"
	"strings"
)

func (b *Binding) selfGetFile(file string) ([]byte, error) {
	return b.pacman.BprintGetBlob(b.handle.Namespace, b.handle.BprintId, file)
}

func (b *Binding) selfAddFile(file string, data []byte) error {
	return b.pacman.BprintNewBlob(b.handle.Namespace, b.handle.BprintId, file, data)
}

func (b *Binding) selfUpdateFile(file string, data []byte) error {
	return b.pacman.BprintUpdateBlob(b.handle.Namespace, b.handle.BprintId, file, data)
}

func (b *Binding) selfAddDataFile(file string, data []byte) error {

	return b.pacman.BprintNewBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file), data)

}

func (b *Binding) selfUpdateDataFile(file string, data []byte) error {
	return b.pacman.BprintUpdateBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file), data)
}

func (b *Binding) selfGetDataFile(file string) ([]byte, error) {
	return b.pacman.BprintGetBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file))
}

func (b *Binding) selfListDataFile() (map[string]string, error) {

	files, err := b.pacman.BprintListBlobs(b.handle.Namespace, b.handle.BprintId)
	if err != nil {
		return nil, err
	}

	fmap := make(map[string]string)
	for f, v := range files {
		if !strings.HasPrefix(f, "data_") {
			continue
		}
		fmap[f] = v
	}

	return fmap, nil
}

func (b *Binding) selfDeleteDataFile(file string) error {
	return b.pacman.BprintDeleteBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file))
}

func dataFile(file string) string {
	return fmt.Sprintf("data_%s", file)
}
