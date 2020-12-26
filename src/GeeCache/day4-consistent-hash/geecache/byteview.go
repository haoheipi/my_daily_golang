package geecache

type ByteView struct {
	value []byte
}

func (b ByteView) Len() int {
	return len(b.value)
}

func (b ByteView) String() string {
	return string(b.value)
}

func (b ByteView) ByteSlice() []byte {
	return cloneBytes(b.value)
}

func cloneBytes(value []byte) []byte {
	c := make([]byte, len(value))
	copy(c, value)
	return c
}
