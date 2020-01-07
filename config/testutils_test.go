package config

import (
	"fmt"
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestWithName(t *testing.T) {
	expected := "foo"
	o := createOption(withOptionName(expected))
	if expected != o.Name {
		t.Errorf(
			`expected name: "%s", actual: "%s"`,
			expected, o.Name,
		)
	}
}

func TestWithDependency(t *testing.T) {
	a := "foo"
	b := "bar"

	expectedA := fmt.Sprintf("${%s}", a)
	expectedB := fmt.Sprintf("${%s}", b)

	o := createOption(
		withOptionDependency(a),
		withOptionDependency(b),
	)

	actual, err := yaml.Marshal(o)
	if err != nil {
		t.Fatalf("unexpected error marshaling option: %s", err)
	}

	if !strings.Contains(string(actual), expectedA) {
		t.Errorf("option does not contain string: %s", expectedA)
	}

	if !strings.Contains(string(actual), expectedB) {
		t.Errorf("option does not contain string: %s", expectedB)
	}
}

func TestWithWhenDependency(t *testing.T) {
	a := "foo"
	b := "bar"

	foundA := false
	foundB := false

	o := createOption(
		withOptionWhenDependency(a),
		withOptionWhenDependency(b),
	)

	for _, value := range o.DefaultValues {
		for _, w := range value.When {
			for key := range w.Equal {
				switch key {
				case a:
					foundA = true
				case b:
					foundB = true
				default:
					t.Errorf("unexpected key: %s", key)
				}
			}
		}
	}

	if !foundA {
		t.Errorf("option does not contain when value: %s", a)
	}

	if !foundB {
		t.Errorf("option does not contain when value: %s", b)
	}
}
