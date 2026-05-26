package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "HellO world",
			expected: []string{"hello", "world"},
		},
	}

for _, cs := range cases{
	actual := CleanInput(cs.input)
	if len(actual) != len(cs.expected){
		t.Errorf("the lengths are not equal %v actual and %v expected", len(actual), len(cs.expected))
		continue
	}
 
	for i, _ := range actual{
		if actual[i] != cs.expected[i] {
			t.Errorf("The word %v does not match %v", actual[i], cs.expected[i])
		}
	}

}
}