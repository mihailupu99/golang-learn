package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	grandiose := "Da### Dan##"
	replacer := strings.NewReplacer("#", "o")
	replacer2 := strings.NewReplacer("#", "daaaan")
	fixed := replacer.Replace(broken)
	fixed2 := replacer2.Replace(grandiose)
	fmt.Println(fixed)
	fmt.Println(fixed2)
}
