package hexdump_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
)

func ExampleHexdump_basic() {
	// echo "Hello" | hexdump
	gloo.MustRun(
		Hexdump(strings.NewReader("Hello")),
	)
	// Output:
	// 00000000  48 65 6c 6c 6f                                    |Hello|
	// 00000005
}
