package glob

import (
	"testing"
)

type globTest struct {
	pattern  string
	input    string
	expected bool
}

func runGlobTests(t *testing.T, tests []globTest) {
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if matcher := NewMatcher(test.pattern); matcher.Match(test.input) != test.expected {
				t.Errorf("pattern=%s, input=%s, expected=%v, got=%v, matcher=%#v", test.pattern, test.input, test.expected, !test.expected, matcher)
			}
		})
		t.Run("invert "+test.input, func(t *testing.T) {
			if matcher := NewMatcher("!" + test.pattern); matcher.Match(test.input) == test.expected {
				t.Errorf("pattern=%s, input=%s, expected=%v, got=%v, matcher=%#v", test.pattern, test.input, !test.expected, test.expected, matcher)
			}
		})
	}
}

func TestGlobEmpty(t *testing.T) {
	tests := []globTest{
		{"", "foo", false},
		{"", "", true},
	}
	runGlobTests(t, tests)
}

func TestGlobExact(t *testing.T) {
	tests := []globTest{
		{"foo", "foo", true},
		{"foo", "bar", false},
		{"foo", "", false},
	}
	runGlobTests(t, tests)
}

func TestGlobMatcherWithPrefix(t *testing.T) {
	tests := []globTest{
		{"foo*", "foo", true},
		{"foo*", "foobar", true},
		{"foo*", "barfoo", false},
		{"foo*", "bar", false},
	}
	runGlobTests(t, tests)
}

func TestGlobMatcherWithPrefixMid(t *testing.T) {
	tests := []globTest{
		{"foo*baz*", "fooXbaz", true},
		{"foo*baz*", "fooXbazbar", true},
		{"foo*baz*", "barfooXbaz", false},
		{"foo*baz*", "bar", false},
	}
	runGlobTests(t, tests)
}

func TestGlobMatcherWithSuffix(t *testing.T) {
	tests := []globTest{
		{"*foo", "foo", true},
		{"*foo", "foobar", false},
		{"*foo", "barfoo", true},
		{"*foo", "bar", false},
	}
	runGlobTests(t, tests)
}

func TestGlobMatcherWithMidSuffix(t *testing.T) {
	tests := []globTest{
		{"*foo*baz", "foobaz", true},
		{"*foo*baz", "foobarbaz", true},
		{"*foo*baz", "barfoobaz", true},
		{"*foo*baz", "barbaz", false},
	}
	runGlobTests(t, tests)
}

func TestGlobMatcherWithPrefixSuffix(t *testing.T) {
	tests := []globTest{
		{"foo*bar", "foo", false},
		{"foo*bar", "bar", false},
		{"foo*bar", "baz", false},
		{"foo*bar", "foobar", true},
		{"foo*bar", "foobazbar", true},
	}
	runGlobTests(t, tests)
}

func TestGlobMatcherNaked(t *testing.T) {
	tests := []globTest{
		{"*foo*", "foo", true},
		{"*foo*", "bar", false},
		{"*foo*", "baz", false},
		{"*foo*", "foobar", true},
		{"*foo*", "barfoo", true},
	}
	runGlobTests(t, tests)
}

func TestGlobMatcherMultipleNaked(t *testing.T) {
	tests := []globTest{
		{"*foo*bar*", "foobar", true},
		{"*foo*bar*", "barbar", false},
		{"*foo*bar*", "bazbar", false},
		{"*foo*bar*", "foobarbar", true},
		{"*foo*bar*", "barfoobar", true},
	}
	runGlobTests(t, tests)
}
