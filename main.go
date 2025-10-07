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
	// clangStylePath     string
)

func queryFunctionDeclarations() map[string]int {
	// [dirkarnez/libclang-experiments: Experiments with `libclang`](https://github.com/dirkarnez/libclang-experiments)
}

func insertMultiple() string {
	// [dirkarnez/go-insert-strings-to-file-offset](https://github.com/dirkarnez/go-insert-strings-to-file-offset/tree/main)
}

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
			codePart = fmt.Sprintf("%s%s", strings.Repeat(" ", tabPrefixOccurance*4), strings.TrimLeft(codePart, "\t"))
		}

		// if strings.Contains(strings.TrimPrefix(codePart, "\t"), "\t") {
		// 	re := regexp.MustCompile("\t+")
		// 	for _, item := range re.FindAllString(strings.TrimPrefix(codePart, "\t"), -1) {
		// 		tabPrefixOccurance := len(item)
		// 		lineLength = lineLength + (tabPrefixOccurance * 5)
		// 	}
		// 	// codePart = fmt.Sprintf("%s%s", strings.Repeat(" ", tabPrefixOccurance*4), strings.TrimLeft(codePart, "\t"))
		// }

		// if strings.HasPrefix(codePart, " ") {
		// 	re := regexp.MustCompile(`\s+`)
		// 	spacePrefixOccurance := len(re.FindString(codePart))
		// 	if spacePrefixOccurance%4 == 0 {
		// 		numberOfTabs := int(spacePrefixOccurance / 4)
		// 		codePart = fmt.Sprintf("%s%s", strings.Repeat("\t", numberOfTabs), strings.TrimLeft(codePart, " "))
		// 		lineLength = lineLength - spacePrefixOccurance + (numberOfTabs * 8)
		// 	}
		// }

		if lineLength < width {
			spacesNeeded := width - lineLength
			codePart = fmt.Sprintf("%s%s", codePart, strings.Repeat(" ", spacesNeeded))
		}

		// Combine the code part and the comment part
		return fmt.Sprintf("%s %s", codePart, commentPart)
	} else {
		// if strings.HasPrefix(line, " ") {
		// 	re := regexp.MustCompile(`\s+`)
		// 	spacePrefixOccurance := len(re.FindString(line))
		// 	if spacePrefixOccurance%4 == 0 {
		// 		numberOfTabs := int(spacePrefixOccurance / 4)
		// 		line = fmt.Sprintf("%s%s", strings.Repeat("\t", numberOfTabs), strings.TrimLeft(line, " "))
		// 	}
		// }

		if strings.HasPrefix(line, "\t") {
			re := regexp.MustCompile("\t+")
			tabPrefixOccurance := len(re.FindString(line))
			line = fmt.Sprintf("%s%s", strings.Repeat(" ", tabPrefixOccurance*4), strings.TrimLeft(line, "\t"))
		}
		return line
	}
}

func RunClangFormat(folder, fileName string) error {
	// args := []string{fmt.Sprintf("--style=file:%s", clangStylePath), "-i", filepath.Join(folder, fileName)}
	args := []string{`--style={ BasedOnStyle: LLVM,  AlignAfterOpenBracket: Align, AlignArrayOfStructures: None, AlignConsecutiveAssignments: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCompound: false, AlignFunctionDeclarations: false, AlignFunctionPointers: false, PadOperators: true }, AlignConsecutiveBitFields: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCompound: false, AlignFunctionDeclarations: false, AlignFunctionPointers: false, PadOperators: false }, AlignConsecutiveDeclarations: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCompound: false, AlignFunctionDeclarations: true, AlignFunctionPointers: false, PadOperators: false }, AlignConsecutiveMacros: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCompound: false, AlignFunctionDeclarations: false, AlignFunctionPointers: false, PadOperators: false }, AlignConsecutiveShortCaseStatements: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCaseArrows: false, AlignCaseColons: false }, AlignConsecutiveTableGenBreakingDAGArgColons: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCompound: false, AlignFunctionDeclarations: false, AlignFunctionPointers: false, PadOperators: false }, AlignConsecutiveTableGenCondOperatorColons: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCompound: false, AlignFunctionDeclarations: false, AlignFunctionPointers: false, PadOperators: false }, AlignConsecutiveTableGenDefinitionColons: { Enabled: false, AcrossEmptyLines: false, AcrossComments: false, AlignCompound: false, AlignFunctionDeclarations: false, AlignFunctionPointers: false, PadOperators: false }, AlignEscapedNewlines: Right, AlignOperands: Align, AlignTrailingComments: { Kind: Always, OverEmptyLines: 0 }, AllowAllArgumentsOnNextLine: true, AllowAllParametersOfDeclarationOnNextLine: true, AllowBreakBeforeNoexceptSpecifier: Never, AllowShortBlocksOnASingleLine: Never, AllowShortCaseExpressionOnASingleLine: true, AllowShortCaseLabelsOnASingleLine: false, AllowShortCompoundRequirementOnASingleLine: true, AllowShortEnumsOnASingleLine: true, AllowShortFunctionsOnASingleLine: All, AllowShortIfStatementsOnASingleLine: Never, AllowShortLambdasOnASingleLine: All, AllowShortLoopsOnASingleLine: false, AllowShortNamespacesOnASingleLine: false, AlwaysBreakAfterDefinitionReturnType: None, AlwaysBreakBeforeMultilineStrings: false, UseTab: Never, IndentWidth: 4, TabWidth: 4, BreakBeforeBraces: Allman, AllowShortIfStatementsOnASingleLine: false, IndentCaseLabels: false, ColumnLimit: 0, AccessModifierOffset: -4, NamespaceIndentation: All, FixNamespaceComments: false, IndentPPDirectives: BeforeHash }`, "-i", filepath.Join(folder, fileName)}
	cmd := exec.Command(clangFormatExePath, args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ReadFileAsLines(folder, fileName string) ([]string, error) {
	fullPath := filepath.Join(folder, fileName)

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, err
	}

	return lines, nil
}
func Clean(folder, fileName string) error {
	fullPath := filepath.Join(folder, fileName)
	// fmt.Println(fullPath)

	RunClangFormat(folder, fileName)

	lines, err := ReadFileAsLines(folder, fileName)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	commentBlockDetected := false
	commentBlockStartDetected := false
	commentBlockEndDetected := false

	var done string
	for _, code := range lines {
		codeTrimmedSpaces := strings.TrimSpace(code)
		if strings.HasPrefix(codeTrimmedSpaces, "//=") || strings.HasPrefix(codeTrimmedSpaces, "//*") || strings.HasPrefix(codeTrimmedSpaces, "/**") || strings.HasPrefix(codeTrimmedSpaces, "/*=") || strings.HasPrefix(codeTrimmedSpaces, "//-") {
			commentBlockDetected = true

			if commentBlockStartDetected {
				commentBlockEndDetected = true
			}

			if !commentBlockStartDetected && !commentBlockEndDetected {
				commentBlockStartDetected = true
			}
		}

		if commentBlockDetected {
			if strings.HasPrefix(codeTrimmedSpaces, "//") {
				done = code
				if commentBlockStartDetected && commentBlockEndDetected {
					commentBlockDetected = false
					commentBlockStartDetected = false
					commentBlockEndDetected = false
				}
			} else {
				commentBlockDetected = false
				commentBlockStartDetected = false
				commentBlockEndDetected = false
			}
		}

		done = pushCommentToRight(code, 88)
		fmt.Fprintln(w, done)
	}

	return w.Flush()

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

// func GetClangFormatStylePath() (string, error) {
// 	path, err := os.Executable()
// 	if err != nil {
// 		return "", err
// 	}

// 	clangStylePath := filepath.Join(filepath.Dir(path), ".clang-format")
// 	_, err = os.Stat(clangStylePath)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return "", fmt.Errorf(".clang-format cannot be found")
// 		} else {
// 			return "", fmt.Errorf(".clang-format cannot be accessed")
// 		}
// 	}

// 	return clangStylePath, nil
// }

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

	// clangStylePath, err = GetClangFormatStylePath()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && IsCXXFile(entry.Name()) {
			Clean(dirPath, entry.Name())
			Clean(dirPath, entry.Name())
		}
	}
}
