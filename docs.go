package vdocs

import (
	"fmt"
	"strings"
)

const (
	opAnd = iota
	opOr
)

func getOpName(op int) string {
	if op == opAnd {
		return "and"
	} else if op == opOr {
		return "or"
	}

	panic("unknown operator")
}

type Validation struct {
	Descriptions []string
	Required     bool
	Format       string
}

type Documentor interface {
	GetRuleDescription(value string) string
}

type RuleWithoutValue struct {
	Documentor
	DescTemplate string
}

func (r RuleWithoutValue) GetRuleDescription(value string) string {
	return r.DescTemplate
}

type RuleWithValue struct {
	Documentor
	DescTemplate string
}

func (r RuleWithValue) GetRuleDescription(value string) string {
	return fmt.Sprintf(r.DescTemplate, value)
}

type RuleWithField struct {
	Documentor
	DescTemplate string
	Op           int
}

func (r RuleWithField) GetRuleDescription(value string) string {
	spaceSeparated := splitValue(value)

	var fields []string
	for _, field := range spaceSeparated {
		fields = append(fields, field)
	}

	var v string
	if len(fields) > 1 {
		x, fields := fields[len(fields)-1], fields[:len(fields)-1] // pop stack
		v = strings.Join(fields, ", ")
		if r.Op == opAnd {
			v = fmt.Sprintf("%s and %s", v, string(x))
		} else {
			v = fmt.Sprintf("%s or %s", v, string(x))
		}
	} else {
		v = strings.Join(fields, ", ")
	}

	return fmt.Sprintf(r.DescTemplate, v)
}

type RuleWithFieldValue struct {
	Documentor
	DescTemplate string
	Op           int
}

func (r RuleWithFieldValue) GetRuleDescription(value string) string {
	spaceSeparated := splitValue(value)

	v := spaceSeparated[0]
	if len(spaceSeparated) > 1 {
		var values []string
		for _, chunk := range chunkSlice(spaceSeparated, 2) {
			values = append(values, strings.Join(chunk, "="))
		}
		if len(values) > 1 {
			x, values := values[len(values)-1], values[:len(values)-1] // pop stack
			v = strings.Join(values, ", ")
			if r.Op == opAnd {
				v = fmt.Sprintf("%s and %s", v, string(x))
			} else {
				v = fmt.Sprintf("%s or %s", v, string(x))
			}
		} else {
			v = strings.Join(values, ", ")
		}
	}

	return fmt.Sprintf(r.DescTemplate, v)
}
func splitValue(value string) []string {
	quoted := false
	return strings.FieldsFunc(value, func(r rune) bool {
		if r == '"' || r == '\'' {
			quoted = !quoted
		}
		return !quoted && r == ' '
	})
}

// GetFieldDescription returns the description for a fields validation
func GetFieldDescription(rule string, value string) string {
	if documentor, ok := bakedIn[rule]; ok {
		return documentor.GetRuleDescription(value)
	}

	return ""
}

// GetFieldDocs returns docs given validation rules for given string of rules (validation tag)
func GetFieldDocs(rules string) Validation {
	if rules == "" {
		return Validation{}
	}

	separatedRules := strings.Split(rules, ",")

	var descriptions []string
	var required bool
	var format string
	for _, rule := range separatedRules {
		separatedRule := strings.Split(rule, "=")

		name := separatedRule[0]

		if r, ok := bakedIn[name]; ok {

			if len(separatedRule) > 1 {
				descriptions = append(descriptions, r.GetRuleDescription(separatedRule[1]))
				continue
			}

			descriptions = append(descriptions, r.GetRuleDescription(""))

			if name == "required" {
				required = true
			}

			if _, ok := r.(RuleWithoutValue); ok {
				if name != "required" && name != "isdefault" {
					format = name
				}
			}
		}
	}

	return Validation{
		Descriptions: descriptions,
		Required:     required,
		Format:       format,
	}
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
