package util

import "testing"

func TestSlicesEqual(t *testing.T) {
	a := []string{"abc", "def"}
	b := []string{"abc", "def"}

	if !SlicesEqual(a, b) {
		t.Errorf("%v and %v should be equal", a, b)
	}
}

func TestSlicesEqual2(t *testing.T) {
	a := []string{"abc", "def"}
	b := []string{"abc", "def", "ghi"}

	if SlicesEqual(a, b) {
		t.Errorf("%v and %v should not be equal", a, b)
	}
}

func TestSlicesEqual3(t *testing.T) {
	a := []string{"abc", "def"}
	b := []string{"abc", "defg"}

	if SlicesEqual(a, b) {
		t.Errorf("%v and %v should not be equal", a, b)
	}
}

func TestReadWordList(t *testing.T) {
	filePath := "../../testdata/voc.txt"
	res, err := ReadWordList(filePath)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if len(res) == 0 {
		t.Errorf("wordlist should not be empty")
		t.Fail()
	}
}

func TestReadWordListErr(t *testing.T) {
	filePath := "../../testdata/does-not-exist"
	_, err := ReadWordList(filePath)
	if err == nil {
		t.Errorf("should throw error, if file does not exist")
		t.Fail()
	}
}
