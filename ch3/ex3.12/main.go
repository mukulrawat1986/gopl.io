package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(IsAnagram("aabb", "abba"))
	fmt.Println(IsAnagram("bncxds", "ncxdsr"))
	fmt.Println(IsAnagram("I Am Dot In Place", "A Decimal Point"))
	fmt.Println(IsAnagram1("aabb", "abba"))
	fmt.Println(IsAnagram1("bncxds", "ncxdsr"))
	fmt.Println(IsAnagram1("I Am Dot In Place", "A Decimal Point"))
	fmt.Println(IsAnagram2("aabb", "abba"))
	fmt.Println(IsAnagram2("bncxds", "ncxdsr"))
	fmt.Println(IsAnagram2("I Am Dot In Place", "A Decimal Point"))

}

func IsAnagram(a, b string) bool {

	a = ToLower(SpaceMap(a))
	b = ToLower(SpaceMap(b))

	if len(a) != len(b) {
		return false
	}

	s1 := strings.Split(a, "")
	s2 := strings.Split(b, "")

	sort.Strings(s1)
	sort.Strings(s2)

	return reflect.DeepEqual(s1, s2)
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func ToLower(str string) string {
	return strings.Map(func(r rune) rune {
		return unicode.ToLower(r)
	}, str)
}

func IsAnagram1(a, b string) bool {
	a = ToLower(SpaceMap(a))
	b = ToLower(SpaceMap(b))

	if len(a) != len(b) {
		return false
	}

	m := make(map[byte]int)

	for i := 0; i < len(a); i++ {
		m[a[i]]++
		m[b[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func IsAnagram2(a, b string) bool {

	a = ToLower(SpaceMap(a))
	b = ToLower(SpaceMap(b))

	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}
