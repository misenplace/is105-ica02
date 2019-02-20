package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere
// Returnerer en slice av type []byte
//
func AllocateVar(b int) []byte {
	// Kode for Oppgave 5a
	return []byte{}
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
func AllocateMake(b int) []byte {
	// Kode for Oppgave 5a
	return make([]byte, b)
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	// kode her for 5b
	return slc[lidx:uidx]
}

// CopySlice ???
