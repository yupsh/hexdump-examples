package hexdump_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/hexdump"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleHexdump_fromFile_basic() {
	// hexdump testdata/binary.txt
	gloo.MustRun(
		Hexdump(gloo.File("testdata/binary.txt")),
	)
	// Output:
	// 00000000  48 65 6c 6c 6f                                    |Hello|
	// 00000005
}
