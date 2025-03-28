package helpers

import (
	"bufio"
	"fmt"
	"strings"
)

// GeneratePlantUML generates a PlantUML sequence diagram from a stack trace.
//
// example:
// buf := make([]byte, 1024)
//
// runtime.Stack(buf, false)
//
// # Get the stack trace
//
// stackTrace := string(buf)
//
// # Parse and generate PlantUML
//
// plantUML := generatePlantUML(stackTrace)
//
// fmt.Println(plantUML)
func GeneratePlantUML(stackTrace string) string {
	scanner := bufio.NewScanner(strings.NewReader(stackTrace))
	var calls []string
	var lastFunction string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "\t") { // Function location line
			continue
		}
		if strings.HasPrefix(line, "goroutine") || strings.HasPrefix(line, "runtime.goexit") {
			continue
		}

		funcName := fmt.Sprintf("\"%s\"", strings.TrimSpace(line)) // Wrap function names in quotes
		if lastFunction != "" {
			calls = append(calls, fmt.Sprintf(`%s --> %s`, lastFunction, funcName))
		}
		lastFunction = funcName
	}

	var sb strings.Builder
	sb.WriteString("@startuml\n")
	// sb.WriteString("scale 2\n")
	// sb.WriteString("left to right direction\n")
	sb.WriteString("top to bottom direction\n")
	sb.WriteString("skinparam packageStyle rect\n")
	for _, call := range calls {
		sb.WriteString(call + "\n")
	}
	sb.WriteString("@enduml\n")

	return sb.String()
}
