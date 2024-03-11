package helper

import (
	"fmt"
	"log"
	"strings"
)

func LoggingError(name string, err error) {
	fmt.Printf("\n")
	log.Default()
	
	fmt.Printf("\n%s\n", name)
	fmt.Println(strings.Repeat("=", 20))
	fmt.Printf("%v\n", err)
}