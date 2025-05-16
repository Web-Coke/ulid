//go:build armbe || arm64be || m68k || mips || mips64 || mips64p32 || ppc || ppc64 || s390 || s390x || shbe || sparc || sparc64

package ulid

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (id ULID) Compare(other ULID) int {
	ih := uint64(id[0x00]) | uint64(id[0x01])<<8 | uint64(id[0x02])<<16 | uint64(id[0x03])<<24 | uint64(id[0x04])<<32 | uint64(id[0x05])<<40 | uint64(id[0x06])<<48 | uint64(id[0x07])<<56
	oh := uint64(other[0x00]) | uint64(other[0x01])<<8 | uint64(other[0x02])<<16 | uint64(other[0x03])<<24 | uint64(other[0x04])<<32 | uint64(other[0x05])<<40 | uint64(other[0x06])<<48 | uint64(other[0x07])<<56
	if ih > oh {
		return 1
	}
	if ih < oh {
		return -1
	}
	il := uint64(id[0x08]) | uint64(id[0x09])<<8 | uint64(id[0x0A])<<16 | uint64(id[0x0B])<<24 | uint64(id[0x0C])<<32 | uint64(id[0x0D])<<40 | uint64(id[0x0E])<<48 | uint64(id[0x0F])<<56
	ol := uint64(other[0x08]) | uint64(other[0x09])<<8 | uint64(other[0x0A])<<16 | uint64(other[0x0B])<<24 | uint64(other[0x0C])<<32 | uint64(other[0x0D])<<40 | uint64(other[0x0E])<<48 | uint64(other[0x0F])<<56
	if il > ol {
		return 1
	}
	if il < ol {
		return -1
	}
	return 0
}
