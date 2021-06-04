package main

import (
	"fmt"
	"strings"
)

func main()  {
	broken := "C# r#cks"
	replacer := strings.NewReplacer("#","o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
