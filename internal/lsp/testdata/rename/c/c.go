package c

import "github.com/kdy1/tools/internal/lsp/rename/b"

func _() {
	b.Hello() //@rename("Hello", "Goodbye")
}
