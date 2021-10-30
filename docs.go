package docs

import (
	"fmt"
	"strings"
)

const (
	_ = iota
	matchStrategyValuesOnly
	matchStrategyFieldsOnly
	matchStrateguFieldWithValues
)

// GetFieldDescription returns the description for a fields validation
func GetFieldDescription(validation string, value string) string {
	if desc, ok := bakedIn[validation]; ok {
		if value == "" {
			return desc
		}

		return fmt.Sprintf(desc, value)
	}

	return ""
}

// GetFieldDocs returns docs given a string of validation rules
func GetFieldDocs(rules string) (ret []string) {
	if rules == "" {
		return nil
	}

	separatedRules := strings.Split(rules, ",")

	for _, rule := range separatedRules {
		separatedRule := strings.Split(rule, "=")
		r := separatedRule[0]
		fmt.Println(r)
		var v string
		if len(separatedRule) > 1 {
			v = separatedRule[1]

			quoted := false
			spaceSeparated := strings.FieldsFunc(v, func(r rune) bool {
				if r == '"' || r == '\'' {
					quoted = !quoted
				}
				return !quoted && r == ' '
			})

			// If 2 or more values, it must be "FieldN valueN ..." e.g. "required_if=Field value"
			if len(spaceSeparated) > 1 {
				var values []string
				for _, chunk := range chunkSlice(spaceSeparated, 2) {
					values = append(values, strings.Join(chunk, "="))
				}
				v = strings.Join(values, ", ")
			}
		}

		ret = append(ret, GetFieldDescription(r, v))
	}

	return ret
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
