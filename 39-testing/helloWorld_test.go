package hello

import "testing"

func TestHelloWorld(t *testing.T) {
  t.Errorf("HelloWorld (%q) == (%q), want %q ", "hello", "world",  "hello")
}


/*
ZhiwendeMacBook-Pro:39-testing brianguo$ go test
--- FAIL: TestHelloWorld (0.00s)
    helloWorld_test.go:6: HelloWorld ("hello") == ("world"), want "hello" 
FAIL
exit status 1
FAIL	hello	0.687s
ZhiwendeMacBook-Pro:39-testing brianguo$

*/
