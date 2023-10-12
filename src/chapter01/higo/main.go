package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"unicode"
	"unicode/utf8"
)

func StrCountChinese(s string) (count int) {
	for _, char := range s {
		if unicode.Is(unicode.Han, char) {
			count++
		}
	}
	return count
}

func StrCountByte(s string) (count int) {
	// sep=""
	count = bytes.Count([]byte(s), []byte(""))
	//byte count to str length
	count -= 1
	return count
}

func StrGetLengthCount(s string) (count int) {
	// sep=""
	count = bytes.Count([]byte(s), []byte(""))
	//byte count to str length
	count -= 1
	return count
}
func StrGetLengthRune(s string) (count int) {
	// to avoid the influence of Chinese characters
	count = utf8.RuneCountInString(s)
	return count
}
func StrGetLengthPlain(s string) (count int) {
	count = len(s)
	return count
}
func StrGetLength(s string) (count int) {
	count = StrGetLengthPlain(s)
	return count
}

// func StrGetDataTypeSize(s string) (count int) {
// 	// please use std lib unsafe only if go < 1.x
// 	count = unsafe.SizeOf(s)
// 	return count
// }

func StrCountSep(s string, sep string) (count int) {
	// sep=""
	count = bytes.Count([]byte(s), []byte(sep))
	return count
}

type StrCountCharResult struct {
	total   int
	english int
	number  int
	space   int
	other   int
	chinese int
}

// use SCCR as type alias for short
type SCCR = StrCountCharResult

// type SCCR StrCountCharResult

func StrCountCharByAscii(sl string) (sar StrCountCharResult) {
	// analyze with ascii code
	// var sl string = "112aaaaFGG123        *&^%"
	var e, s, d, c, o, t int

	for i := o; i < len(sl); i++ {
		switch {
		case 64 < sl[i] && sl[i] < 91:
			e += 1
		case 96 < sl[i] && sl[i] < 123:
			e += 1
		case 47 < sl[i] && sl[i] < 58:
			d += 1
		case sl[i] == 32:
			s += 1
		default:
			o += 1
		}
	}
	t = StrGetLengthCount(sl)
	c = StrCountChinese(sl)
	o = t - e - d - s - c

	// sar := StringAnalyzeResult{}
	sar.english = e
	sar.number = d
	sar.space = s
	sar.chinese = c
	sar.other = o
	sar.total = t
	return sar
}
func StrCountCharByRegexp(s string) (sar StrCountCharResult) {
	var rNum = regexp.MustCompile(`\d`)
	var rCharacter = regexp.MustCompile("[a-zA-Z]")
	var rBlank = regexp.MustCompile(" ")
	var specialcharacter int

	num := len(rNum.FindAllStringSubmatch(s, -1))
	character := len(rCharacter.FindAllStringSubmatch(s, -1))
	blank := len(rBlank.FindAllStringSubmatch(s, -1))
	chinese := StrCountChinese(s)
	t := StrGetLengthCount(s)

	// specialcharacter = len(s) - num - character - blank
	specialcharacter = t - num - character - blank - chinese
	// sar := StringAnalyzeResult{}
	sar.english = character
	sar.number = num
	sar.space = blank
	sar.chinese = chinese
	sar.other = specialcharacter
	sar.total = t
	return sar
}

func main() {
	// company meta info here
	fmt.Println("[info] Hello, Zero!")

	// your business logic

	// use the same package method in ./lib.go
	// pkgFunc()

	// app meta info here
	fmt.Println("[info] APP: Str Analyze Char")

	// get cmd from cli args
	args := os.Args

	var cmd string
	if len(args) >= 2 {
		cmd = args[1]
	}
	// info current cmd
	// fmt.Println("[info] cmd:", cmd)
	logo := "Zero 从零到壹"
	if cmd != "" {
		logo = cmd
	}
	// info input
	fmt.Println("[info] String: ", logo)

	// lc := StrCountChinese(logo)
	// bc := StrGetLengthCount(logo)
	// sl := StrGetLengthRune(logo)
	// oc := StrCountSep(logo, "o")
	// erc := StrCountSep(logo, "er")

	// fmt.Println("[info] String length:", sl)
	// fmt.Println("[info] String Byte Count:", bc)
	// fmt.Println("[info] Chinese Count:", lc)
	// fmt.Println("[info] o count:", oc)
	// fmt.Println("[info] er count:", erc)

	// sar := StrCountTypeByAscii(logo)
	sar := StrCountCharByRegexp(logo)
	fmt.Printf("[info] String Length: %d\n", StrGetLengthPlain(logo))
	fmt.Printf("[info] Total Characters: %d\n", sar.total)
	fmt.Printf("[info] English Characters: %d\n", sar.english)
	fmt.Printf("[info] Number Characters: %d\n", sar.number)
	fmt.Printf("[info] Space Characters: %d\n", sar.space)
	fmt.Printf("[info] Chinese Characters: %d\n", sar.chinese)
	fmt.Printf("[info] Other Characters: %d\n", sar.other)

	// use the same package method in ./pause.go
	// pause_exit_when_press_enter_or_press_any_key_then_enter()
	// pause_exit_when_typing_exit()
	// pause_exit_when_press_any_key_then_enter()

	// use ./signal.go
	// signalChan := NewShutdownSignal()
	// WaitExit(signalChan, func() {
	// 	fmt.Println("[info] server closed")
	// })

}
