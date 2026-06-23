package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	envPath := ".env"

	bytes := make([]byte, 64)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Printf("Error generating random bytes: %v\n", err)
		return
	}
	newSecret := hex.EncodeToString(bytes)
	secretLine := fmt.Sprintf("JWT_SECRET=%s", newSecret)

	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		err := os.WriteFile(envPath, []byte(secretLine+"\n"), 0644)
		if err != nil {
			fmt.Printf("Error creating new .env file: %v\n", err)
			return
		}
		fmt.Println(".env file not found. Successfully created a new .env file with JWT_SECRET.")
		return
	}

	file, err := os.OpenFile(envPath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error opening .env file: %v\n", err)
		return
	}
	defer file.Close()

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading .env content: %v\n", err)
		return
	}
	content := string(contentBytes)

	var newLines []string
	lines := strings.Split(content, "\n")
	isUpdated := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "JWT_SECRET=") {
			newLines = append(newLines, secretLine)
			isUpdated = true
		} else {
			newLines = append(newLines, line)
		}
	}

	if !isUpdated {
		if len(lines) > 0 && lines[len(lines)-1] != "" {
			newLines = append(newLines, "")
		}
		newLines = append(newLines, secretLine)
	}

	if err := file.Truncate(0); err != nil {
		fmt.Printf("Error truncating file for rewrite: %v\n", err)
		return
	}
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Printf("Error resetting file pointer: %v\n", err)
		return
	}

	output := strings.Join(newLines, "\n")
	if _, err := file.WriteString(output); err != nil {
		fmt.Printf("Error writing new .env content: %v\n", err)
		return
	}

	if isUpdated {
		fmt.Println("JWT_SECRET in your .env file has been successfully updated with a new key.")
	} else {
		fmt.Println("JWT_SECRET has been successfully added to your .env file.")
	}
}
