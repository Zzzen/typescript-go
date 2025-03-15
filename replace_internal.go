package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk("internal", processFile)
	if err != nil {
		fmt.Printf("Error walking through files: %v\n", err)
	}
}

func processFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !strings.HasSuffix(path, ".go") {
		return nil
	}

	// Read the file
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", path, err)
	}

	// Create a reader for the content
	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)

	var newContent strings.Builder
	inImportBlock := false

	// Process the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if we're entering an import block
		if strings.HasPrefix(strings.TrimSpace(line), "import (") {
			inImportBlock = true
			newContent.WriteString(line + "\n")
			continue
		}

		// Check if we're leaving an import block
		if inImportBlock && strings.HasPrefix(strings.TrimSpace(line), ")") {
			inImportBlock = false
			newContent.WriteString(line + "\n")
			continue
		}

		// Handle single line imports
		if strings.HasPrefix(strings.TrimSpace(line), "import ") && strings.Contains(line, "\"github.com/microsoft/typescript-go/internal/") {
			line = strings.Replace(line, "/microsoft/typescript-go/internal/", "/Zzzen/typescript-go/use-at-your-own-risk/", 1)
		}

		// Replace "internal" with "use-at-your-own-risk" in import paths within import blocks
		if inImportBlock {
			line = strings.Replace(line, "/microsoft/typescript-go/internal/", "/Zzzen/typescript-go/use-at-your-own-risk/", 1)
		}

		newContent.WriteString(line + "\n")
	}

	// Create new path with "use-at-your-own-risk" instead of "internal"
	newPath := strings.Replace(path, "internal/", "use-at-your-own-risk/", 1)

	// Create directory structure if it doesn't exist
	err = os.MkdirAll(filepath.Dir(newPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating directories for %s: %v", newPath, err)
	}

	// Write the modified content to the new file
	err = os.WriteFile(newPath, []byte(newContent.String()), info.Mode())
	if err != nil {
		return fmt.Errorf("error writing file %s: %v", path, err)
	}

	// Remove the original file after successful write
	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("error removing original file %s: %v", path, err)
	}

	fmt.Printf("Processed: %s\n", path)
	return nil
}
