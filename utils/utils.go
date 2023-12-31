package utils

import "os"

var Arguments []string = os.Args[1:]

func GetCommand() string {
	command := Arguments[0]

	return command
}