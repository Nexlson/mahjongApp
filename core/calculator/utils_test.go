package calculator

import "testing"

func TestHello(t *testing.T) {
	str := Hello()

	expected := "Hello World"

	if expected != str {
		t.Errorf("Test output %s is not equal to expected output %s", str, expected)
	}
}