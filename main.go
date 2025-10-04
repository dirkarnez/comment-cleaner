// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func pushCommentToRight(code string, width int) string {
	// Split the code into lines
	lines := strings.Split(code, "\n")
	var result []string

	for _, line := range lines {
		// Trim the line to remove leading/trailing spaces
		trimmedLine := strings.TrimSpace(line)

		// Check if the line contains a comment
		if idx := strings.Index(trimmedLine, "//"); idx != -1 {
			// Split the line into code and comment
			codePart := trimmedLine[:idx]
			commentPart := trimmedLine[idx:]

			// Calculate the spaces needed to push the comment to the right
			lineLength := len(codePart)
			if lineLength < width {
				spacesNeeded := width - lineLength
				codePart = fmt.Sprintf("%s%s", codePart, strings.Repeat(" ", spacesNeeded))
			}

			// Combine the code part and the comment part
			result = append(result, fmt.Sprintf("%s %s", codePart, commentPart))
		} else {
			// If there's no comment, keep the line as is
			result = append(result, line)
		}
	}

	// Join the processed lines back together
	return strings.Join(result, "\n")
}

func main() {
	const commentWidth = 80

	// Process the code
	processedCode := pushCommentToRight("Hello, 世界 // 45", commentWidth)
	fmt.Println(processedCode)
}
