package tr

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/tr/opt"
)


// Flags represents the configuration options for the tr command
type Flags = localopt.Flags
// Command implementation
type command opt.Inputs[string, Flags]

// Tr creates a new tr command with the given parameters
func Tr(parameters ...any) yup.Command {
	return command(opt.Args[string, Flags](parameters...))
}

func (c command) Execute(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
	if len(c.Positional) < 1 {
		fmt.Fprintln(stderr, "tr: missing operand")
		return fmt.Errorf("missing operand")
	}

	set1 := c.Positional[0]
	var set2 string
	if len(c.Positional) > 1 {
		set2 = c.Positional[1]
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		result := c.processLine(line, set1, set2)
		fmt.Fprintln(output, result)
	}

	return scanner.Err()
}

func (c command) processLine(line, set1, set2 string) string {
	if bool(c.Flags.Delete) {
		return c.deleteChars(line, set1)
	}

	if bool(c.Flags.Squeeze) {
		return c.squeezeChars(line, set1)
	}

	return c.translateChars(line, set1, set2)
}

func (c command) deleteChars(line, set1 string) string {
	deleteSet := make(map[rune]bool)
	for _, r := range set1 {
		deleteSet[r] = true
	}

	var result strings.Builder
	for _, r := range line {
		if !deleteSet[r] {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func (c command) squeezeChars(line, set1 string) string {
	squeezeSet := make(map[rune]bool)
	for _, r := range set1 {
		squeezeSet[r] = true
	}

	var result strings.Builder
	var lastChar rune
	var inSqueeze bool

	for _, r := range line {
		if squeezeSet[r] {
			if !inSqueeze || r != lastChar {
				result.WriteRune(r)
				inSqueeze = true
				lastChar = r
			}
		} else {
			result.WriteRune(r)
			inSqueeze = false
		}
	}

	return result.String()
}

func (c command) translateChars(line, set1, set2 string) string {
	if set2 == "" {
		return line
	}

	// Build translation map
	transMap := make(map[rune]rune)
	runes1 := []rune(set1)
	runes2 := []rune(set2)

	for i, r1 := range runes1 {
		if i < len(runes2) {
			transMap[r1] = runes2[i]
		} else {
			// If set2 is shorter, use last character
			transMap[r1] = runes2[len(runes2)-1]
		}
	}

	var result strings.Builder
	for _, r := range line {
		if newR, exists := transMap[r]; exists {
			result.WriteRune(newR)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
