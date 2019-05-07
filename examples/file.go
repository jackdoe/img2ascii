package main

import (
	"fmt"
	"os"

	"github.com/jackdoe/img2ascii"
)

func main() {
	ascii := string(img2ascii.MustFile2Ascii(os.Args[1], 80))
	fmt.Printf("%s", ascii)
}
