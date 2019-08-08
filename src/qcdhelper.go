package main

import (
	"log"
	"os"

	"github.com/eykrehbein/quickmove/src/utils"
)

func main() {
	err := utils.Commander(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
