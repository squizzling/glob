package glob

// ACL implements the Matcher interface to apply standard ACL
// style criteria to the provided input.
type ACL struct {
	Allow   []Matcher
	Deny    []Matcher
	Default bool
}

// Match will evaluate the lists, with the following rules:
// - If anything in the Deny list matches, the value does not match
// - If anything in the Allow list matches, the value matches
// - If nothing matches, the Default is returned
//
// Thread safe under the following conditions:
// - All Matchers must be thread safe (the ones provided by this package are)
// - The ACL struct itself is not being modified by another thread
func (acl *ACL) Match(value string) bool {
	for _, m := range acl.Deny {
		if m.Match(value) {
			return false
		}
	}

	for _, m := range acl.Allow {
		if m.Match(value) {
			return true
		}
	}

	return acl.Default
}

// NewACL is a helper function to create an ACL style Matcher, using
// the provided strings as glob style Matchers.  Additional Matchers
// can be added to the lists later if desired.
func NewACL(allow []string, deny []string, def bool) *ACL {
	acl := &ACL{
		Default: def,
	}
	for _, pattern := range allow {
		acl.Allow = append(acl.Allow, NewMatcher(pattern))
	}
	for _, pattern := range deny {
		acl.Deny = append(acl.Deny, NewMatcher(pattern))
	}
	return acl
}

var _ = Matcher(&ACL{})
