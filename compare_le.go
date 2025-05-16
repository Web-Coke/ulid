//go:build 386 || amd64 || amd64p32 || alpha || arm || arm64 || loong64 || mipsle || mips64le || mips64p32le || nios2 || ppc64le || riscv || riscv64 || sh || wasm

package ulid

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (id ULID) Compare(other ULID) int {
	ih := uint64(id[0x07]) | uint64(id[0x06])<<8 | uint64(id[0x05])<<16 | uint64(id[0x04])<<24 | uint64(id[0x03])<<32 | uint64(id[0x02])<<40 | uint64(id[0x01])<<48 | uint64(id[0x00])<<56
	oh := uint64(other[0x07]) | uint64(other[0x06])<<8 | uint64(other[0x05])<<16 | uint64(other[0x04])<<24 | uint64(other[0x03])<<32 | uint64(other[0x02])<<40 | uint64(other[0x01])<<48 | uint64(other[0x00])<<56
	if ih > oh {
		return 1
	}
	if ih < oh {
		return -1
	}
	il := uint64(id[0x0F]) | uint64(id[0x0E])<<8 | uint64(id[0x0D])<<16 | uint64(id[0x0C])<<24 | uint64(id[0x0B])<<32 | uint64(id[0x0A])<<40 | uint64(id[0x09])<<48 | uint64(id[0x08])<<56
	ol := uint64(other[0x0F]) | uint64(other[0x0E])<<8 | uint64(other[0x0D])<<16 | uint64(other[0x0C])<<24 | uint64(other[0x0B])<<32 | uint64(other[0x0A])<<40 | uint64(other[0x09])<<48 | uint64(other[0x08])<<56
	if il > ol {
		return 1
	}
	if il < ol {
		return -1
	}
	return 0
}
