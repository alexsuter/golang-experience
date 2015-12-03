package main

import "fmt"

func main() {
	// Slice Literal
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	// make-Funktion
	var s []byte
	s = make([]byte, 5)
	// s == []byte{0, 0, 0, 0, 0}
	
	s = make([]byte, 5)
	// len: 5 cap: 5
	s = append(s, 3)
	// len: 6 cap: 16 <- automatisch erweitert
	
	s = make([]byte, 10)
	for i := 0; i < len(s); i++ {
		s[i] = byte(i)
	}

	// before [0 1 2 3 4 5 6 7 8 9]
	AddOneToEachElement(s)
	//after [1 2 3 4 5 6 7 8 9 10]
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
