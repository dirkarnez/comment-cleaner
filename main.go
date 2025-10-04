// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func pushCommentToRight(line string, width int) string {
	// Trim the line to remove leading/trailing spaces
	trimmedLine := strings.TrimSpace(line)

	// Check if the line contains a comment
	if idx := strings.Index(trimmedLine, "//"); idx != -1 {
		// Split the line into code and comment
		codePart := strings.TrimRight(trimmedLine[:idx], "\t")
		commentPart := trimmedLine[idx:]

		// Calculate the spaces needed to push the comment to the right
		lineLength := len(codePart)
		if lineLength < width {
			spacesNeeded := width - lineLength
			codePart = fmt.Sprintf("%s%s", codePart, strings.Repeat(" ", spacesNeeded))
		}

		// Combine the code part and the comment part
		return fmt.Sprintf("%s %s", codePart, commentPart)
	} else {
		return line
	}
}

func Clean(folder, fileName string) error {
	fullPath := filepath.Join(folder, fileName)
	fmt.Println(fullPath)

	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}

	defer file.Close()

	outputFile, err := os.Create(filepath.Join(folder, fmt.Sprintf("%s.new", fileName)))
	if err != nil {
		return err
	}

	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	w := bufio.NewWriter(outputFile)

	commentBlockDetected := false
	commentBlockStartDetected := false
	commentBlockEndDetected := false

	var done string
	for scanner.Scan() {
		code := scanner.Text()
		if strings.HasPrefix(code, "//=") || strings.HasPrefix(code, "//*") || strings.HasPrefix(code, "/**") || strings.HasPrefix(code, "//-") {
			commentBlockDetected = true

			if commentBlockStartDetected {
				commentBlockEndDetected = true
			}

			if !commentBlockStartDetected && !commentBlockEndDetected {
				commentBlockStartDetected = true
			}
		} else {
			if commentBlockStartDetected && commentBlockEndDetected {
				commentBlockDetected = false
				commentBlockStartDetected = false
				commentBlockEndDetected = false
			}
		}

		if commentBlockDetected {
			done = code
		} else {
			done = pushCommentToRight(code, 80)
		}

		fmt.Println(done)
		fmt.Fprintln(w, done)
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return scanner.Err()
}

func main() {
	dirPath := os.Args[1]

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			Clean(dirPath, entry.Name())
		}
	}
}
