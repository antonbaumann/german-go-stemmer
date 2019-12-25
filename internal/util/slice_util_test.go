package util

import (
	"testing"
)

func TestHasSuffix(t *testing.T) {
	arr := []rune{'a', 'b', 'c', 'd'}
	suff := []rune{'c', 'd'}

	if !HasSuffix(arr, suff) {
		t.Errorf("%v should be suffix of %v", suff, arr)
	}
}

func TestHasSuffix2(t *testing.T) {
	arr := []rune{'a', 'b', 'c', 'd'}
	suff := []rune{'a', 'b', 'c', 'd'}

	if !HasSuffix(arr, suff) {
		t.Errorf("%v should be suffix of %v", suff, arr)
	}
}

func TestHasSuffix3(t *testing.T) {
	var arr []rune
	var suff []rune

	if !HasSuffix(arr, suff) {
		t.Errorf("%v should be suffix of %v", suff, arr)
	}
}

func TestHasSuffix4(t *testing.T) {
	arr := []rune{'a', 'b', 'c', 'd'}
	suff := []rune{'a', 'b', 'c'}

	if HasSuffix(arr, suff) {
		t.Errorf("%v should not be suffix of %v", suff, arr)
	}
}

func TestHasSuffix5(t *testing.T) {
	arr := []rune{'a', 'b', 'c', 'd'}
	suff := []rune{'z', 'a', 'b', 'c', 'd'}

	if HasSuffix(arr, suff) {
		t.Errorf("%v should not be suffix of %v", suff, arr)
	}
}
