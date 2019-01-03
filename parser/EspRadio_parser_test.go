package parser

import (
	"testing"
)

func TestNewEsperantoRadio(t *testing.T) {
	parser := NewEsperantoRadio()
	podcasts, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	_ = podcasts
	// for _, p := range podcasts {
	// 	fmt.Printf("%v\n", p.Podcast)
	// 	fmt.Printf("%v\n", p.Channel)
	// 	fmt.Println()
	// }

}
