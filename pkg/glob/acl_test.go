package glob

import (
	"testing"
)

// We avoid using complicated matchers, and stick with basic primitives as
// much as possible, because that logic is generally validated in glob_test.go

type aclTest struct {
	name  string
	allow []string
	deny  []string
	tests map[string]bool
}

// runACLTests will test all the provided inputs, and a dummy
// value, against both default = true and default = false
func runACLTests(t *testing.T, tests []aclTest) {
	const defaultTestValue = "default value"
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, def := range []bool{true, false} {
				a := NewACL(test.allow, test.deny, def)
				for input, expected := range test.tests {
					if a.Match(input) != expected {
						t.Errorf("w=%s, b=%s, input=%s, expected=%v, actual=%v", test.allow, test.deny, input, expected, !expected)
					}
				}
				if a.Match(defaultTestValue) != def {
					t.Errorf("w=%s, b=%s, input=zzz, expected=%v, actual=%v", test.allow, test.deny, def, !def)
				}
			}
		})
	}
}

func TestACLAllow(t *testing.T) {
	tests := []aclTest{{
		"basic-allow",
		[]string{"allow1", "allow2"},
		nil,
		map[string]bool{
			"allow1": true,
			"allow2": true,
		},
	}}
	runACLTests(t, tests)
}

func TestACLDeny(t *testing.T) {
	tests := []aclTest{{
		"basic-deny",
		nil,
		[]string{"deny1", "deny2"},
		map[string]bool{
			"deny1": false,
			"deny2": false,
		},
	}}
	runACLTests(t, tests)
}

func TestACLCombinedLists(t *testing.T) {
	tests := []aclTest{{
		"basic-combined-list",
		[]string{"allow1", "allow2"},
		[]string{"deny1", "deny2"},
		map[string]bool{
			"allow1": true,
			"allow2": true,
			"deny1":  false,
			"deny2":  false,
		},
	}, {
		"allow-general-deny-specific",
		[]string{"allow*"},
		[]string{"allow2"},
		map[string]bool{
			"allow1": true,
			"allow2": false,
			"allow3": true,
		},
	}, {
		"direct-conflict",
		[]string{"allow*"},
		[]string{"allow*"},
		map[string]bool{
			"allow1": false,
			"allow2": false,
			"allow3": false,
		},
	}}
	runACLTests(t, tests)
}
