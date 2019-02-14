package ascii

import "fmt"

// Oppgave 1b
// Implementer en funksjon som eksportere const ascii

// Funksjon tar en "string literal" med kun ASCII tegn og lager en utskrift på
// følgende format:
// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]
//
// Eksempel (utskriften kommer fra en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
// ...
func IterateOverASCIIStringLiteral(stringLiteral string) {
	// Kode for Oppgave 1a
	for i := 0; i < len(stringLiteral); i++ {
		// using %q here as several of the first characters are unprintable
		// including null and newlines etc.
		fmt.Printf("%X %q %b\n", stringLiteral[i], stringLiteral[i], stringLiteral[i])
	}
}

// Unix-like operating systems are known to use it as erase control character, i.e. to delete the previous character in the line mode. 

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen GreetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
// Gjelder oppgave 1c
func GreetingASCII() string {
	s := []byte{'"', 'H', 'e', 'l', 'l', 'o', ' ', ':', '-', ')', '"', '¥'}
	return string(s[:])
}
