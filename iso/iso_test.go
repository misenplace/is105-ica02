package iso

import (
	"testing"
)

func TestGreetingExtendedASCIIONly(t *testing.T) {
	val := GreetingExtendedASCII()
	for _, i := range val {
		if !(i > 0x80 && i < 0xff) {
			t.Errorf("greetingExtendedASCII() returns non-extended ASCII value: %q - %v", i, i)
		}
	}
}