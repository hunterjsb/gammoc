package main

import (
	"fmt"

	"github.com/hunterjsb/gammoc/src/config"
)

func main() {
	s, err := config.ReadDotenv(".env")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
