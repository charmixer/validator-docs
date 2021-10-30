package docs

import (
	"fmt"
	"strings"
	"testing"
)

type ruleTest struct {
	rule string
	want string
}

func TestSimpleRules(t *testing.T) {
	var tests []ruleTest
	for i, d := range bakedIn {
		if strings.Contains(d, "%") {
			continue
		}

		test := ruleTest{
			rule: i,
			want: d,
		}
		tests = append(tests, test)
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.rule)
		t.Run(testname, func(t *testing.T) {
			got := GetFieldDocs(tt.rule)
			if tt.want != got[0] {
				t.Errorf(`got "%s", want "%s"`, got, tt.want)
			}
		})
	}

}

func TestFieldRules(t *testing.T) {
	tests := []ruleTest{
		{
			rule: "required_if=Field value",
			want: fmt.Sprintf(bakedIn["required_if"], "Field=value"),
		},
		{
			rule: "required_if=Field value Field2 value2",
			want: fmt.Sprintf(bakedIn["required_if"], "Field=value, Field2=value2"),
		},
		{
			rule: "required_if=Field 'value with spaces'",
			want: fmt.Sprintf(bakedIn["required_if"], "Field='value with spaces'"),
		},
		{
			rule: "required_if=Field 'value with spaces' OtherField value",
			want: fmt.Sprintf(bakedIn["required_if"], "Field='value with spaces', OtherField=value"),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.rule)
		t.Run(testname, func(t *testing.T) {
			got := GetFieldDocs(tt.rule)
			if tt.want != got[0] {
				t.Errorf(`got "%s", want "%s"`, got[0], tt.want)
			}
		})
	}
}
