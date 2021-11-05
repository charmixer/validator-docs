package vdocs

import (
	"fmt"
	"testing"
)

type ruleTest struct {
	rule string
	want string
}

func TestRuleWithoutValue(t *testing.T) {
	var tests []ruleTest
	for i, d := range bakedIn {
		var rule RuleWithoutValue
		switch d.(type) {
		case RuleWithoutValue:
			rule = d.(RuleWithoutValue)
			break
		default:
			continue
		}

		test := ruleTest{
			rule: i,
			want: rule.DescTemplate,
		}
		tests = append(tests, test)
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.rule)
		t.Run(testname, func(t *testing.T) {
			got := GetFieldDocs(tt.rule)
			if tt.want != got.Descriptions[0] {
				t.Errorf(`got "%s", want "%s"`, got.Descriptions[0], tt.want)
			}
		})
	}
}

func TestRuleWithField(t *testing.T) {
	var tests []ruleTest
	for i, d := range bakedIn {
		var rule RuleWithField
		switch d.(type) {
		case RuleWithField:
			rule = d.(RuleWithField)
			break
		default:
			continue
		}

		tests = append(tests, []ruleTest{
			{
				rule: fmt.Sprintf("%s=%s", i, "Field"),
				want: fmt.Sprintf(rule.DescTemplate, "Field"),
			},
			{
				rule: fmt.Sprintf("%s=%s", i, "Field1 Field2"),
				want: fmt.Sprintf(rule.DescTemplate, fmt.Sprintf("Field1 %s Field2", getOpName(rule.Op))),
			},
			{
				rule: fmt.Sprintf("%s=%s", i, "Field1 Field2 Field3"),
				want: fmt.Sprintf(rule.DescTemplate, fmt.Sprintf("Field1, Field2 %s Field3", getOpName(rule.Op))),
			},
		}...)
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.rule)
		t.Run(testname, func(t *testing.T) {
			got := GetFieldDocs(tt.rule)
			if tt.want != got.Descriptions[0] {
				t.Errorf(`got "%s", want "%s"`, got.Descriptions[0], tt.want)
			}
		})
	}
}

func TestRuleWithFieldValue(t *testing.T) {
	var tests []ruleTest
	for i, d := range bakedIn {
		var rule RuleWithFieldValue
		switch d.(type) {
		case RuleWithFieldValue:
			rule = d.(RuleWithFieldValue)
			break
		default:
			continue
		}

		tests = append(tests, []ruleTest{
			{
				rule: fmt.Sprintf("%s=%s", i, "Field value"),
				want: fmt.Sprintf(rule.DescTemplate, "Field=value"),
			},
			{
				rule: fmt.Sprintf("%s=%s", i, "Field value Field2 value2"),
				want: fmt.Sprintf(rule.DescTemplate, fmt.Sprintf("Field=value %s Field2=value2", getOpName(rule.Op))),
			},
			{
				rule: fmt.Sprintf("%s=%s", i, "Field value Field2 value2 Field3 value3"),
				want: fmt.Sprintf(rule.DescTemplate, fmt.Sprintf("Field=value, Field2=value2 %s Field3=value3", getOpName(rule.Op))),
			},
			{
				rule: fmt.Sprintf("%s=%s", i, "Field 'value with spaces'"),
				want: fmt.Sprintf(rule.DescTemplate, "Field='value with spaces'"),
			},
			{
				rule: fmt.Sprintf("%s=%s", i, "Field 'value with spaces' OtherField value"),
				want: fmt.Sprintf(rule.DescTemplate, fmt.Sprintf("Field='value with spaces' %s OtherField=value", getOpName(rule.Op))),
			},
			{
				rule: fmt.Sprintf("%s=%s", i, "Field value Field2 value2 Field3 value3"),
				want: fmt.Sprintf(rule.DescTemplate, fmt.Sprintf("Field=value, Field2=value2 %s Field3=value3", getOpName(rule.Op))),
			},
		}...)
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.rule)
		t.Run(testname, func(t *testing.T) {
			got := GetFieldDocs(tt.rule)
			if tt.want != got.Descriptions[0] {
				t.Errorf(`got "%s", want "%s"`, got.Descriptions[0], tt.want)
			}
		})
	}
}

func TestFieldWithValuesRules(t *testing.T) {
	var tests []ruleTest
	for i, d := range bakedIn {
		var rule RuleWithValue
		switch d.(type) {
		case RuleWithValue:
			rule = d.(RuleWithValue)
			break
		default:
			continue
		}

		test := ruleTest{
			rule: fmt.Sprintf("%s=%s", i, "x"),
			want: fmt.Sprintf(rule.DescTemplate, "x"),
		}
		tests = append(tests, test)
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.rule)
		t.Run(testname, func(t *testing.T) {
			got := GetFieldDocs(tt.rule)
			if tt.want != got.Descriptions[0] {
				t.Errorf(`got "%s", want "%s"`, got.Descriptions[0], tt.want)
			}
		})
	}

}
