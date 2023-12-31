package main

import (
	"fmt"
	"go-exec/migration"
	"go-exec/utils"
)

func main() {
	command := utils.GetCommand()

	switch command {
	case "migration":
		migration.Create()
	default: 
		fmt.Println("go-exec")
	}
}