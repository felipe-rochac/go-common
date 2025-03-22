package helpers

import (
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
	lines := strings.Split(stackTrace, "\n")
	var plantUML strings.Builder

	plantUML.WriteString("@startuml\n")
	plantUML.WriteString("[*] --> main\n") // Start from main

	callStack := []string{"main"} // Keep track of calls to prevent duplicates for recursive or looped function calls, and to form the correct sequence.

	for _, line := range lines {
		if strings.Contains(line, ".go:") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				funcName := parts[0]
				funcName = strings.TrimSuffix(funcName, "(") // remove the ( if present
				funcName = strings.TrimSuffix(funcName, ")") // remove the ) if present
				funcNameParts := strings.Split(funcName, "/")
				funcName = funcNameParts[len(funcNameParts)-1] // grab only the function name, not the whole path

				if len(callStack) > 0 {
					lastFunc := callStack[len(callStack)-1]
					plantUML.WriteString(fmt.Sprintf("%s --> %s\n", lastFunc, funcName))
				}
				callStack = append(callStack, funcName)
			}
		}
	}

	plantUML.WriteString("@enduml\n")
	return plantUML.String()
}
