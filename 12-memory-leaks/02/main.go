package main

import (
	"log"
	"os"
)

func main() {
	files := []string{"./a.txt", "./b.txt", "./c.txt"}

	for _, file := range files {
		func() {
			f, err := os.Open(file)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
		}()
	}
}
