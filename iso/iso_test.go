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

// checks if string is writable using standard ASCII and Extended ASCII
// func TestGreetingExtendedASCII(t *testing.T) {
// 	val := greetingExtendedASCII()
// 	for i := 0; i < len(val); i++ {
// 		if val[i] > unicode.MaxLatin1 {
// 			t.Errorf("Return value contains non-ascii value: %q", val[i])
// 		}
// 	}
// }
