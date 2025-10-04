// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func pushCommentToRight(line string, width int) string {
	// Trim the line to remove leading/trailing spaces
	trimmedLine := line // strings.TrimSpace(line)

	// Check if the line contains a comment
	if idx := strings.Index(trimmedLine, "//"); idx != -1 {
		// Split the line into code and comment
		codePart := strings.TrimRight(trimmedLine[:idx], "\t")
		codePart = strings.TrimRight(codePart, " ")
		commentPart := trimmedLine[idx:]

		lineLength := len(codePart)

		if strings.HasPrefix(codePart, "\t") {
			re := regexp.MustCompile("\t+")
			tabPrefixOccurance := len(re.FindString(codePart))
			lineLength = lineLength + (tabPrefixOccurance * 7)
			// codePart = fmt.Sprintf("%s%s", strings.Repeat(" ", tabPrefixOccurance*4), strings.TrimLeft(codePart, "\t"))
		}

		if strings.Contains(strings.TrimPrefix(codePart, "\t"), "\t") {
			re := regexp.MustCompile("\t+")
			for _, item := range re.FindAllString(strings.TrimPrefix(codePart, "\t"), -1) {
				tabPrefixOccurance := len(item)
				lineLength = lineLength + (tabPrefixOccurance * 5)
			}
			// codePart = fmt.Sprintf("%s%s", strings.Repeat(" ", tabPrefixOccurance*4), strings.TrimLeft(codePart, "\t"))
		}

		if strings.HasPrefix(codePart, " ") {
			re := regexp.MustCompile(`\s+`)
			spacePrefixOccurance := len(re.FindString(codePart))
			if spacePrefixOccurance%4 == 0 {
				numberOfTabs := int(spacePrefixOccurance / 4)
				codePart = fmt.Sprintf("%s%s", strings.Repeat("\t", numberOfTabs), strings.TrimLeft(codePart, " "))
				lineLength = lineLength - spacePrefixOccurance + (numberOfTabs * 8)
			}
		}

		if lineLength < width {
			spacesNeeded := width - lineLength
			codePart = fmt.Sprintf("%s%s", codePart, strings.Repeat(" ", spacesNeeded))
		}

		// Combine the code part and the comment part
		return fmt.Sprintf("%s %s", codePart, commentPart)
	} else {
		if strings.HasPrefix(line, " ") {
			re := regexp.MustCompile(`\s+`)
			spacePrefixOccurance := len(re.FindString(line))
			if spacePrefixOccurance%4 == 0 {
				numberOfTabs := int(spacePrefixOccurance / 4)
				line = fmt.Sprintf("%s%s", strings.Repeat("\t", numberOfTabs), strings.TrimLeft(line, " "))
			}
		}

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
			done = pushCommentToRight(code, 88)
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
	dirPath := "C:\\Users\\19081126D\\Downloads\\testing" //os.Args[1]

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
