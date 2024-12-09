package minigo

import (
	"fmt"
	"log"
)

var S = fmt.Sprintf
var SN = fmt.Sprintln

var F = fmt.Fprintf
var FN = fmt.Fprintln

var P = fmt.Printf
var PN = fmt.Println

var E = fmt.Errorf

func FE(format string, err error) {
	if err != nil {
		log.Fatalf(format, err)
	}
}

func N1[T any](first T, rest ...any) T {
	if len(rest) > 0 {
		err, ok := rest[0].(error)
		if ok && err != nil {
			log.Fatal(err)
		}
	}
	return first
}

func N2[T any](_, second T, rest ...any) T {
	if len(rest) > 0 {
		err, ok := rest[0].(error)
		if ok && err != nil {
			log.Fatal(err)
		}
	}
	return second
}

func N3[T any](_, _ any, third T, rest ...any) T {
	if len(rest) > 0 {
		err, ok := rest[0].(error)
		if ok && err != nil {
			log.Fatal(err)
		}
	}
	return third
}
