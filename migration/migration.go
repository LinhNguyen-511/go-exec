package migration

import (
	"fmt"
	"go-exec/utils"
	"os"
	"strings"
	"time"
)

func Create() {
	currentPath, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
        return
    }

	// TODO: add some error handling mechanism
	migrationName := utils.Arguments[1]

	if(!isNameCorrectlyFormatted(migrationName)) {
		fmt.Println("Migration name should not contain any spaces.")
        return
	}

	currentDate := time.Now().Format(time.DateOnly)
	currentTime := time.Now().Format(time.TimeOnly)

	fileName := fmt.Sprintf("%s-%s_%s", currentDate, currentTime, migrationName)
	migrationDirectoryPath := fmt.Sprintf("%s/migrations", currentPath)
	migrationFilePath := fmt.Sprintf("%s/%s.go", migrationDirectoryPath, fileName)

	os.Mkdir(migrationDirectoryPath, 0750)

	file, err := os.Create(migrationFilePath)

	fmt.Println("created")
    if err != nil {
        panic(err)
    }
    defer file.Close()
}

func isNameCorrectlyFormatted(name string) bool {
	nameTrimmed := strings.Trim(name, "")

	return !strings.Contains(nameTrimmed, " ")
}