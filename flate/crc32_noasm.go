//+build !amd64 noasm appengine

// Copyright 2015, Klaus Post, see LICENSE for details.

package flate

func init() {
	useSSE42 = false
}

func crc32sse(a []byte) hash {
	panic("no assembler")
}

func crc32sseAll(a []byte, dst []hash) {
	panic("no assembler")
}
