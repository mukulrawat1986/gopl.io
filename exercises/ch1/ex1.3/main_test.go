package main

import (
	"bytes"
	"testing"
)

func TestEcho2(t *testing.T) {
	buff := bytes.Buffer{}
	inp := []string{"Hello", "World", "this", "is", "dog"}

	echo2(inp, &buff)

	got := buff.String()
	want := "Hello World this is dog\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEcho3(t *testing.T) {
	buff := bytes.Buffer{}
	inp := []string{"Hello", "World", "this", "is", "dog"}

	echo3(inp, &buff)

	got := buff.String()
	want := "Hello World this is dog\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkEcho2(b *testing.B) {
	var buff bytes.Buffer
	for i := 0; i < b.N; i++ {
		echo2([]string{"Hello", "World", "this", "is", "Dog"}, &buff)
	}
}

func BenchmarkEcho3(b *testing.B) {
	var buff bytes.Buffer
	for i := 0; i < b.N; i++ {
		echo3([]string{"Hello", "World", "this", "is", "Dog"}, &buff)
	}
}
