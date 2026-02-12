package logging

import (
	"fmt"
	"log"

	"github.com/ekkolyth/go-tools/config"
)

func Fatal(message string, err error) {
	log.Fatal(config.CLI_RED + "[FATAL]: " + config.CLI_RESET + message + err.Error() + "\n\n")
}

func Info(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(config.CLI_BLUE + "[INFO]: " + config.CLI_RESET + output)
}

func Warning(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(config.CLI_YELLOW + "[WARN]: " + config.CLI_RESET + output)
}

func Api(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(config.CLI_BRIGHT_CYAN + "[API]: " + config.CLI_RESET + output)
}

func Worker(message string, args ...any) {
	output := fmt.Sprintf(message, args...)
	log.Println(config.CLI_BRIGHT_MAGENTA + "[WORKER]: " + config.CLI_RESET + output)
}
