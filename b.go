package modb

import "fmt"
import "github.com/lcp8474140/modc"

var Version="v6.0.3"

func Print() {
	fmt.Println("modb", Version)
	modc.Print()
}
