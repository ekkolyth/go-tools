package logging

import (
	"fmt"
	"log"
)

func Fatal(message string, err error) {
	log.Fatal(CLI_RED + "[FATAL]: " + CLI_RESET + message + err.Error() + "\n\n")
}

func Info(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(CLI_BLUE + "[INFO]: " + CLI_RESET + output)
}

func Warning(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(CLI_YELLOW + "[WARN]: " + CLI_RESET + output)
}

func Api(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(CLI_BRIGHT_CYAN + "[API]: " + CLI_RESET + output)
}

func Worker(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(CLI_BRIGHT_MAGENTA + "[WORKER]: " + CLI_RESET + output)
}
