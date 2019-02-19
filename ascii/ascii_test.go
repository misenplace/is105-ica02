package ascii

import (
	"testing"
	"unicode"
)

var expected = "\"Hello :-)\""

func TestGreetingASCII(t *testing.T) {
	if val := greetingASCII(); val != expected {
		t.Errorf("greetingASCII() returned %s, expected %v", val, expected)
	}
}

func TestGreetingASCIIIsAscii(t *testing.T) {
	val := greetingASCII()
	for i := 0; i < len(val); i++ {
		if val[i] > unicode.MaxASCII {
			t.Errorf("Return value contains non-ascii value: %q", val[i])
		}
	}
}
