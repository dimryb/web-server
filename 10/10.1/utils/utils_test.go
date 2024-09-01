package utils_test

import (
	"example/utils"
	"testing"
)

func TestFoo(t *testing.T) {
	if utils.Foo() != 42 {
		t.Logf("expected %d\n", 42)
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	defer func() {
		t.Log("defer")
	}()

	// t.Log("1")
	// t.Fail()
	// t.Log("2")
	// t.FailNow()
	// t.Log("3")

	//t.Fatal
}
