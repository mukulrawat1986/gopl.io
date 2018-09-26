package main

import "testing"

func BenchmarkIsAnagram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAnagram("I Am Dot In Place", "A Decimal Point")
	}

}

func BenchmarkIsAnagram1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAnagram1("I Am Dot In Place", "A Decimal Point")
	}

}

func BenchmarkIsAnagram2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAnagram2("I Am Dot In Place", "A Decimal Point")
	}

}
