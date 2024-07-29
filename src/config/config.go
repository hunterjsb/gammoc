package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadDotenv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf := make([]byte, 32*1024) // define your buffer size here.

	for {
		n, err := file.Read(buf)

		if n > 0 {
			fmt.Print(buf[:n]) // your read buffer.
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}

}
