package self

func (b *Binding) selfGetFile(file string) ([]byte, error) {
	return b.pacman.BprintGetBlob(b.handle.Namespace, b.handle.BprintId, file)
}

func (b *Binding) selfAddFile(file string, data []byte) error {
	return b.pacman.BprintNewBlob(b.handle.Namespace, b.handle.BprintId, file, data)
}

func (b *Binding) selfUpdateFile(file string, data []byte) error {
	return b.pacman.BprintUpdateBlob(b.handle.Namespace, b.handle.BprintId, file, data)
}
