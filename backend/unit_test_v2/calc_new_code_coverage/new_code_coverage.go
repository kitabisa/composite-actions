package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <diff-file> <coverage-file>")
		return
	}

	diffFile := os.Args[1]
	coverageFile := os.Args[2]

	// Parse the diff file to get changed files
	changedFiles, err := parseNewFiles(diffFile)
	if err != nil {
		fmt.Printf("Error processing diff file: %v\n", err)
		return
	}

	// Get coverage data using go tool cover
	fileCoverages, errArr := getCoverageData(coverageFile, changedFiles)
	// Check if changed files exist in coverage data and print their coverage
	var fileCoverageResult string
	var fileCoverageBelowThreshold string
	fileCoverageResult += "\n"
	fileCoverageBelowThreshold += "\n"
	for file, coverage := range fileCoverages {
		if coverage < 80 {
			fileCoverageBelowThreshold += fmt.Sprintf("File: %s, Coverage: %.2f%%\n", file, coverage)
		} else {
			fileCoverageResult += fmt.Sprintf("File: %s, Coverage: %.2f%%\n", file, coverage)
		}
	}

	fmt.Printf("Coverage accepted %v\n", fileCoverageResult)

	if len(errArr) > 0 {
		fmt.Printf("Coverage below threshold: %v\n", fileCoverageBelowThreshold)
	}
}

func parseNewFiles(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	newFiles := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "+++ b/") {
			filePath := strings.TrimPrefix(line, "+++ b/")
			newFiles[filePath] = true
		} else if strings.HasPrefix(line, "--- a/") {
			filePath := strings.TrimPrefix(line, "--- a/")
			delete(newFiles, filePath) // Remove if it's not a new file
		}
	}

	return newFiles, scanner.Err()
}

func getCoverageData(coverageFile string, changedFiles map[string]bool) (fileCoverages map[string]float64, errArr []error) {
	cmd := exec.Command("go", "tool", "cover", "-func", coverageFile)
	output, err := cmd.Output()
	if err != nil {
		errArr = append(errArr, err)
		return
	}

	fileCoverages = make(map[string]float64)
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		for file := range changedFiles {
			cleanFile := strings.TrimSpace(file)
			if strings.Contains(line, cleanFile) {
				// Compile a regular expression to match one or more consecutive tabs
				re := regexp.MustCompile(`\t+`)

				// Replace all occurrences of multiple tabs with a single tab
				strArr := re.ReplaceAllString(line, "\t")
				strArrResult := strings.Split(strArr, "\t")
				coveragePercent := strings.TrimSuffix(strArrResult[2], "%")

				// Check if the file path contains "internal/" and coverage is 0
				if strings.Contains(strArrResult[0], "internal/") {
					floatVal, _ := strconv.ParseFloat(coveragePercent, 64)
					fileCoverages[fmt.Sprintf("%s %s", file, strArrResult[1])] = floatVal
					if floatVal < 80 {
						errArr = append(errArr, fmt.Errorf("file %s with function %s is in 'internal/' directory and has %s coverage", cleanFile, strArrResult[1], coveragePercent))
					}
				}

			}
		}
	}

	return
}
