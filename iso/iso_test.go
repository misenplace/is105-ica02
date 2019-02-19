package iso

import (
	"testing"
	"unicode"
)

// checks if string is ONLY extended ascii
// this will usually fail, as it requires the string be ONLY extended ascii
// and no normal chars
func TestGreetingExtendedASCIIONly(t *testing.T) {
	val := GreetingExtendedASCII()
	for i := 0; i < len(val); i++ {
		// if val[i] is not between hex 80 and hex FF throw error
		if !(val[i] >= unicode.MaxASCII && val[i] <= unicode.MaxLatin1 ) {
			t.Errorf("greetingExtendedASCII() returns non-extended ASCII value: %q", val[i])
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
