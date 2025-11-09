package command_test

import (
	"errors"
	"testing"

	"github.com/gloo-foo/testable/assertion"
	"github.com/gloo-foo/testable/run"
	command "github.com/yupsh/tr"
)

func TestTr_Basic(t *testing.T) {
	result := run.Command(command.Tr("a", "x")).
		WithStdinLines("abc").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"xbc"})
}

func TestTr_Delete(t *testing.T) {
	result := run.Command(command.Tr("a", "", command.Delete)).
		WithStdinLines("abc").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestTr_EmptyInput(t *testing.T) {
	result := run.Quick(command.Tr("a", "x"))
	assertion.NoError(t, result.Err)
	assertion.Empty(t, result.Stdout)
}

func TestTr_InputError(t *testing.T) {
	result := run.Command(command.Tr("a", "x")).
		WithStdinError(errors.New("read failed")).Run()
	assertion.ErrorContains(t, result.Err, "read failed")
}

