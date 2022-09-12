package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const OUTPUT = ".env.example"

func Execute(location, delimiter string) error {
	functionName := "Exenv"

	log.Println(location)

	file, err := os.Open(location)
	if err != nil {
		log.Printf("[%s]: unable to open file - %s", functionName, err)
		return err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("[%s]: unable to read file - %s", functionName, err)
		return err
	}

	lines := strings.Split(string(content), "\n")
	keys := []string{}

	for _, line := range lines {
		if len(line) == 0 || !strings.Contains(line, delimiter) {
			continue
		}
		key := strings.Split(line, delimiter)[0]
		keys = append(keys, key)
	}

	if len(keys) == 0 {
		err := fmt.Errorf("unable to generate file from %s", location)
		log.Printf("[%s]: %s", functionName, err)
		return err
	}

	result := strings.Join(keys, "=\n")

	file, err = os.Create(OUTPUT)
	if err != nil {
		log.Printf("[%s]: unable to create file - %s", functionName, err)
		return err
	}

	if _, err := io.WriteString(file, result); err != nil {
		log.Printf("[%s]: unable to read file - %s", functionName, err)
		return err
	}

	log.Printf("Generate: %s", OUTPUT)
	return nil
}
