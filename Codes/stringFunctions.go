package main

import (
	"strings"
	"fmt"
)

var print func(a ...interface{}) (n int, err error) = fmt.Println		//We alias fmt.Println as print

func main() {

	print("Contains:  ", strings.Contains("test", "es"))
	print("Count:     ", strings.Count("test", "t"))
	print("HasPrefix: ", strings.HasPrefix("test", "te"))
	print("HasSuffix: ", strings.HasSuffix("test", "st"))
	print("Index:     ", strings.Index("test", "e"))
	print("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	print("Repeat:    ", strings.Repeat("a", 5))
	print("Replace:   ", strings.Replace("foo", "o", "0", -1))
	print("Replace:   ", strings.Replace("foo", "o", "0", 1))
	print("Split:     ", strings.Split("a-b-c-d-e", "-"))
	print("ToLower:   ", strings.ToLower("TEST"))
	print("ToUpper:   ", strings.ToUpper("test"))
	print("______________________________________________________")

	print("Length of 'test': ", len("hello"))
	print("Character:", "hello"[1])
}

/*
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST
______________________________________________________
Length of 'test':  5
Character: 101
 */