package main

import (
	"fmt"
	"regexp"
)

type Symbol struct {
	ch   rune
	next []*Symbol
}

func add(symbol *Symbol, r []rune) {
	for _, val := range r {
		symbol.next = append(symbol.next, &Symbol{val, []*Symbol{}})
	}
}

func isNumber1(s string) bool {

	head := &Symbol{'*', []*Symbol{}}
	r := []rune{'+', '-', '1', '.'} //First enter
	add(head, r)

	help := head.next[0] //+...
	r = []rune{'1', '.'} //+1... || +.  ...
	add(help, r)

	help = head.next[0].next[0] //+1...
	r = []rune{'!', 'e', '.'}   //+1end || +1. ... || +1e...
	add(help, r)

	help = head.next[0].next[0].next[0] // +1e...
	r = []rune{'+', '-', '1'}           //	+1e+... || +1e-... || +1e1...
	add(help, r)

	help = head.next[0].next[0].next[0].next[0] // +1e+...
	r = []rune{'1'}                             //	+1e+1...
	add(help, r)

	help = head.next[0].next[0].next[0].next[0].next[0] // +1e+1...
	r = []rune{'!', '1'}                                //	+1e+1end || +1e+1...
	add(help, r)

	help = head.next[0].next[0].next[1] // +1. ...
	r = []rune{'!', '1', 'e'}           //	+1.end || +1.1... || +1.e...
	add(help, r)

	help = head.next[0].next[1] //+. ...
	r = []rune{'1'}             //+.1...
	add(help, r)

	help = head.next[1]           //-...
	help.next = head.next[0].next //-1 || -.

	help = head.next[2]       // 1...
	r = []rune{'!', '.', 'e'} //1end || 1. || 1e
	add(help, r)

	help = head.next[3] // . ...
	r = []rune{'1'}     // .1
	add(help, r)

	return true
}

func isNumber2(s string) bool {

	rStr := []string{
		`^([+]|[-])?[\d]+\.[\d]+e([+]|[-])?[\d]+$`,
		`^([+]|[-])?\.[\d]+e([+]|[-])?[\d]+$`,
		`^([+]|[-])?[\d]+\.e([+]|[-])?[\d]+$`,
		`^([+]|[-])?[\d]+\.[\d]+$`,
		`^([+]|[-])?\.[\d]+$`,
		`^([+]|[-])?[\d]+\.$`}

	for step := 0; step < 36; step++ {
		str := ""
		fmt.Scanln(&str)
		for _, v := range rStr {
			re := regexp.MustCompile(v)
			for i, match := range re.FindAllString(str, -1) {
				fmt.Println(match, "found at index", i, "->", step)
			}
		}
	}
	return true
}

var rStr = []string{
	`^([+]|[-])?[\d]+([e]|[E])([+]|[-])?[\d]+$`,
	`^([+]|[-])?[\d]+\.[\d]+([e]|[E])([+]|[-])?[\d]+$`,
	`^([+]|[-])?\.[\d]+([e]|[E])([+]|[-])?[\d]+$`,
	`^([+]|[-])?[\d]+\.([e]|[E])([+]|[-])?[\d]+$`,
	`^([+]|[-])?[\d]+\.[\d]+$`,
	`^([+]|[-])?\.[\d]+$`,
	`^([+]|[-])?[\d]+\.$`,
	`^([+]|[-])?[\d]+$`}

func isNumber(s string) bool {

	for _, v := range rStr {
		re := regexp.MustCompile(v)
		if re.MatchString(s) {
			return true
		}
	}
	return false
}

func main() {

	for step := 0; step < 45; step++ {
		str := ""
		fmt.Scanln(&str)
		fmt.Println(step, isNumber(str))
	}

	//fmt.Println(isNumber2("12345679"))
}
