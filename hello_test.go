package hello

/*
Naming convention used for a test file is filename_test.go
*/
import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, Test!"
	got := Say("Test")
	if want != got {
		t.Errorf("wanted %s got %s", want, got)
	}
}

/*
Command to run test file in current directory:
⚡➜ src go test ./...
?       hello/main      [no test files]
ok      hello   0.402s
⚡➜ src go test ./...
?       hello/main      [no test files]
ok      hello   (cached)
⚡➜ src go test ./...
?       hello/main      [no test files]
--- FAIL: TestSayHello (0.00s)
    hello_test.go:12: wanted Hello, Harsh! got Hello, Test!
FAIL
FAIL    hello   0.277s
FAIL
⚡➜ src
*/
