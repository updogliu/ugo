package ugo

func Assert(b bool) {
	if !b {
		panic("Assertion failed")
	}
}
