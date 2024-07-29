package main

import (
	"github.com/hunterjsb/gammoc/internal/config"
)

func main() {
	_, err := config.ReadDotenv(".env")
	if err != nil {
		panic(err)
	}
}
