package ugo

func Assert(b bool) {
	if !b {
		panic("Assertion failed")
	}
}

func NoError(err error) {
	if err != nil {
		panic("Unexpected error: " + err.Error())
	}
}
