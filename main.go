package main

import (
	"fmt"

	"github.com/kecci/go-common-util/util"
)

func main() {
	person := []string{"Adi", "Budi", "Charlie"}
	
	// util.Contains
	isAdiExist := util.Contains(person, "Adi")
	fmt.Printf("is Adi exist: %v", isAdiExist)
}