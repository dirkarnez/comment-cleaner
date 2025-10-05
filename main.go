// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	clangFormatExePath string
	clangStylePath     string
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

func runClangFormat(folder, fileName string) error {
	cmd := exec.Command(clangFormatExePath, fmt.Sprintf("--style=file:%s", clangStylePath), fmt.Sprintf("--files=%s", filepath.Join(folder, fileName)))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
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

func IsCXXFile(fileName string) bool {
	// C language
	for _, extension := range []string{".c", ".h"} {
		if strings.HasSuffix(fileName, extension) {
			return true
		}
	}

	// C++ language
	for _, extension := range []string{".C", ".cc", ".cpp", ".cxx", ".c++", ".hh", ".hpp", ".hxx", ".h++"} {
		if strings.HasSuffix(fileName, extension) {
			return true
		}
	}

	return false
}

func GetClangFormatStylePath() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}

	clangStylePath := filepath.Join(filepath.Dir(path), ".clang-format")
	_, err = os.Stat(clangStylePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf(".clang-format cannot be found")
		} else {
			return "", fmt.Errorf(".clang-format cannot be accessed")
		}
	}

	return clangStylePath, nil
}

func GetClangFormatExecutablePath() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}

	clangFormatPath := filepath.Join(filepath.Dir(path), "clang-format.exe")
	_, err = os.Stat(clangFormatPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("clang-format.exe cannot be found")
		} else {
			return "", fmt.Errorf("clang-format.exe cannot be accessed")
		}
	}

	return clangFormatPath, nil
}

func main() {
	dirPath := "C:\\Users\\19081126D\\Downloads\\testing"
	// dirPath := os.Args[1]
	var err error = nil
	clangFormatExePath, err = GetClangFormatExecutablePath()
	if err != nil {
		log.Fatal(err)
	}

	clangStylePath, err = GetClangFormatStylePath()
	if err != nil {
		log.Fatal(err)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && IsCXXFile(entry.Name()) {
			Clean(dirPath, entry.Name())
		}
	}
}
