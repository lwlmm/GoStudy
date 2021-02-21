package unitTest

import "testing"

//testing.T 是 testing 包内置的测试用例
func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) should be 3")
	}
}