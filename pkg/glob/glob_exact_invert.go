package glob

type globExactInvert struct {
	part string
}

func (gei *globExactInvert) Match(input string) bool {
	return input != gei.part
}

var _ = Matcher(&globExactInvert{})
