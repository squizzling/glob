package glob

type globEmptyInvert struct{}

func (gei *globEmptyInvert) Match(input string) bool {
	return input != ""
}

var _ = Matcher(&globEmptyInvert{})
