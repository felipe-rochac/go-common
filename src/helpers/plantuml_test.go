package helpers

import (
	"testing"
)

func Test_GeneratePlantUml(t *testing.T) {
	stackTrace := `goroutine 1 [running]:
runtime/debug.Stack(0xc0000a0000, 0x1000, 0x1000)
	/usr/local/go/src/runtime/debug/stack.go:24 +0x9d
main.main()
	/home/user/go/src/github.com/username/project/main.go:10 +0x20
goroutine 2 [running]:
net/http.(*conn).serve(0xc0000b0000, 0xc0000c0000)
	/usr/local/go/src/net/http/server.go:1825 +0x705
net/http.(*ServeMux).ServeHTTP(0xc0000d0000, 0xc0000e0000, 0xc0000f0000)
	/usr/local/go/src/net/http/server.go:2416 +0x1ad
net/http.serverHandler.ServeHTTP(0xc0000g0000, 0xc0000e0000, 0xc0000f0000)
	/usr/local/go/src/net/http/server.go:2843 +0xa3
net/http.(*conn).serve(0xc0000h0000, 0xc0000i0000)
	/usr/local/go/src/net/http/server.go:1925 +0x8ad
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1
goroutine 3 [running]:
os.(*File).Write(0xc0000j0000, 0xc0000k0000, 0x1000, 0x1000)
	/usr/local/go/src/os/file.go:176 +0x8d
fmt.Fprintf(0xc0000l0000, 0xc0000m0000, 0x1000, 0x1000)
	/usr/local/go/src/fmt/print.go:205 +0x9d
log.(*Logger).Output(0xc0000n0000, 0x2, 0xc0000o0000, 0x1000, 0x1000)
	/usr/local/go/src/log/log.go:179 +0x8d
log.Printf(0xc0000p0000, 0xc0000q0000, 0x1000, 0x1000)
	/usr/local/go/src/log/log.go:320 +0x9d
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1
`
	expected := `@startuml
left to right direction
skinparam packageStyle rect
"runtime/debug.Stack(0xc0000a0000, 0x1000, 0x1000)" --> "main.main()"
"main.main()" --> "net/http.(*conn).serve(0xc0000b0000, 0xc0000c0000)"
"net/http.(*conn).serve(0xc0000b0000, 0xc0000c0000)" --> "net/http.(*ServeMux).ServeHTTP(0xc0000d0000, 0xc0000e0000, 0xc0000f0000)"
"net/http.(*ServeMux).ServeHTTP(0xc0000d0000, 0xc0000e0000, 0xc0000f0000)" --> "net/http.serverHandler.ServeHTTP(0xc0000g0000, 0xc0000e0000, 0xc0000f0000)"
"net/http.serverHandler.ServeHTTP(0xc0000g0000, 0xc0000e0000, 0xc0000f0000)" --> "net/http.(*conn).serve(0xc0000h0000, 0xc0000i0000)"
"net/http.(*conn).serve(0xc0000h0000, 0xc0000i0000)" --> "os.(*File).Write(0xc0000j0000, 0xc0000k0000, 0x1000, 0x1000)"
"os.(*File).Write(0xc0000j0000, 0xc0000k0000, 0x1000, 0x1000)" --> "fmt.Fprintf(0xc0000l0000, 0xc0000m0000, 0x1000, 0x1000)"
"fmt.Fprintf(0xc0000l0000, 0xc0000m0000, 0x1000, 0x1000)" --> "log.(*Logger).Output(0xc0000n0000, 0x2, 0xc0000o0000, 0x1000, 0x1000)"
"log.(*Logger).Output(0xc0000n0000, 0x2, 0xc0000o0000, 0x1000, 0x1000)" --> "log.Printf(0xc0000p0000, 0xc0000q0000, 0x1000, 0x1000)"
@enduml
`

	actual := GeneratePlantUML(stackTrace)
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}

	// err := os.WriteFile("output.puml", []byte(expected), 0644)
	// if err != nil {
	// 	t.Fatalf("failed to write to file: %v", err)
	// }
}
