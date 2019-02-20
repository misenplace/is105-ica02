package treasure


import (
	"bytes"
	"fmt"
)

// PrintTreasureUTF8 tar en "string literal" som INN-data
// og skriver ut en korrekt text på med UTF-8 koding
// Koden er for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) {
	treasureBytes := []byte(treasureString)

	// Replace returns a copy with the bytes replaced
	// therefore, it's needed to set treasureBytes to the new copy.
	treasureBytes = bytes.Replace(treasureBytes, []byte("\xe5"), []byte("\xc3\xa5"), -1)
	treasureBytes = bytes.Replace(treasureBytes, []byte("\xe6"), []byte("\xc3\xa6"), -1)
	treasureBytes = bytes.Replace(treasureBytes, []byte("\xf8"), []byte("\xc3\xb8"), -1)
	fmt.Printf("%q", treasureBytes)
}
