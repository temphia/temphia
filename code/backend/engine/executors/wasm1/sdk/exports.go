package tasmsdk

func AllocBytes(size int32) int32 {
	return allocBytes(size)
}

func FreeBytes(addr int32) {
	freeBytes(addr)
}
