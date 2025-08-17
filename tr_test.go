package tr_test

import (
	"context"
	"os"
	"strings"

	"github.com/yupsh/tr"
	"github.com/yupsh/tr/opt"
)

func ExampleTr() {
	ctx := context.Background()
	input := strings.NewReader("hello world")

	cmd := tr.Tr("a-z", "A-Z") // Would need range expansion in real implementation
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output: hello world
}

func ExampleTr_delete() {
	ctx := context.Background()
	input := strings.NewReader("hello world")

	cmd := tr.Tr("l", opt.Delete)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output: heo word
}
