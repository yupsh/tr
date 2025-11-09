package command

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
)

type command gloo.Inputs[string, flags]

func Tr(parameters ...any) gloo.Command {
	return command(gloo.Initialize[string, flags](parameters...))
}

func (p command) Executor() gloo.CommandExecutor {
	// Get set1 and set2 from positional arguments
	set1 := ""
	set2 := ""
	if len(p.Positional) > 0 {
		set1 = p.Positional[0]
	}
	if len(p.Positional) > 1 {
		set2 = p.Positional[1]
	}

	return gloo.LineTransform(func(line string) (string, bool) {
		result := line

		// Delete mode
		if bool(p.Flags.Delete) {
			if bool(p.Flags.Complement) {
				// Delete characters NOT in set1
				var kept []rune
				for _, r := range result {
					if strings.ContainsRune(set1, r) {
						kept = append(kept, r)
					}
				}
				result = string(kept)
			} else {
				// Delete characters in set1
				for _, char := range set1 {
					result = strings.ReplaceAll(result, string(char), "")
				}
			}
		} else {
			// Translate mode
			if bool(p.Flags.Complement) {
				// Translate characters NOT in set1 to set2
				var translated []rune
				lastChar2 := rune(0)
				if len(set2) > 0 {
					lastChar2 = []rune(set2)[len([]rune(set2))-1]
				}

				for _, r := range result {
					if strings.ContainsRune(set1, r) {
						translated = append(translated, r)
					} else {
						translated = append(translated, lastChar2)
					}
				}
				result = string(translated)
			} else {
				// Normal translation
				runes1 := []rune(set1)
				runes2 := []rune(set2)

				var translated []rune
				for _, r := range result {
					found := false
					for i, c1 := range runes1 {
						if r == c1 {
							if i < len(runes2) {
								translated = append(translated, runes2[i])
							} else if len(runes2) > 0 {
								// Use last char of set2 if set1 is longer
								translated = append(translated, runes2[len(runes2)-1])
							}
							found = true
							break
						}
					}
					if !found {
						translated = append(translated, r)
					}
				}
				result = string(translated)
			}
		}

		// Squeeze mode
		if bool(p.Flags.Squeeze) {
			squeezed := []rune{}
			var lastRune rune
			for i, r := range result {
				if i == 0 || r != lastRune || !strings.ContainsRune(set2, r) {
					squeezed = append(squeezed, r)
				}
				lastRune = r
			}
			result = string(squeezed)
		}

		return result, true
	}).Executor()
}
