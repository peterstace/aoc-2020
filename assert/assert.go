package assert

import "testing"

func IntEq(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got=%v want=%v", got, want)
	}
}
