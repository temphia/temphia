package bprint

// SelfAddFile(file string, data []byte) error
// SelfUpdateFile(file string, data []byte) error

// SelfAddDataFile(file string, data []byte) error
// SelfUpdateDataFile(file string, data []byte) error
// SelfGetDataFile(file string) ([]byte, error)
// SelfListDataFiles() (map[string]string, error)
// SelfDeleteDataFile(file string) error

/*


func (b *SelfBindings) selfGetFile(file string) ([]byte, error) {
	return b.pacman.BprintGetBlob(b.handle.Namespace, b.handle.BprintId, file)
}

func (b *SelfBindings) selfAddFile(file string, data []byte) error {
	return b.pacman.BprintNewBlob(b.handle.Namespace, b.handle.BprintId, file, data, true)
}

func (b *SelfBindings) selfUpdateFile(file string, data []byte) error {
	return b.pacman.BprintUpdateBlob(b.handle.Namespace, b.handle.BprintId, file, data)
}

func (b *SelfBindings) selfAddDataFile(file string, data []byte) error {

	return b.pacman.BprintNewBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file), data, true)

}

func (b *SelfBindings) selfUpdateDataFile(file string, data []byte) error {
	return b.pacman.BprintUpdateBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file), data)
}

func (b *SelfBindings) selfGetDataFile(file string) ([]byte, error) {
	return b.pacman.BprintGetBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file))
}

func (b *SelfBindings) selfListDataFiles() (map[string]string, error) {

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

func (b *SelfBindings) selfDeleteDataFile(file string) error {
	return b.pacman.BprintDeleteBlob(b.handle.Namespace, b.handle.BprintId, dataFile(file))
}

func dataFile(file string) string {
	return fmt.Sprintf("data_%s", file)
}


*/
