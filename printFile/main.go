package main

import (
	"fmt"
	"os"
)

func main() {
	fl:=os.Args[0]
	fmt.Println(os.Open(fl))
}
