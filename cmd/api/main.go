package main

import (
	"fmt"
	"log"

	"github.com/m00n3r-dev/forgecom-api/internal/config"
)

func main() {

	cnf, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config")
	}

	fmt.Println(cnf.DBHost)

}
