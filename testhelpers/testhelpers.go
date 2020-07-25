package testhelpers

import "testing"

// CheckErrorNil ...
func CheckErrorNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error found " + err.Error())
	}
}

// CheckErrorNotNil ...
func CheckErrorNotNil(t *testing.T, err error) {
	if err != nil {
		return
	}
	t.Errorf("Error is nil")
}

func AssertEqual(t *testing.T, desired interface{}, actual interface{}) {
	if desired != actual {
		t.Errorf("Missmatch between desired `%v` and actual `%v`", desired, actual)
	}
}
