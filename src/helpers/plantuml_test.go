package helpers

import "testing"

func Test_GeneratePlantUml(t *testing.T) {
	stackTrace := `goroutine 1 [running]:)
runtime/debug.Stack(0xc0000a0000, 0x1000, 0x1000)
	/usr/local/go/src/runtime/debug/stack.go:24 +0x9d
main.main()
	/home/user/go/src/github.com/username/project/main.go:10 +0x20
`
	expected := `@startuml
[*] --> main
main --> stack.go:24
stack.go:24 --> main.go:10
@enduml
`

	actual := GeneratePlantUML(stackTrace)
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
