package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	t := 0
	fmt.Fscan(in, &t)
	in.ReadLine()
	s := ""
	for i := 0; i < t; i++ {
		help, _, _ := in.ReadLine()
		colRow, _ := strconv.Atoi(string(help))

		for j := 0; j < colRow; j++ {
			help, _, _ := in.ReadLine()
			s += string(help)
		}
		s += ", "
	}
	s = s[:len(s)-2]
	ss := strings.Split(s, " ")
	s = ""
	for i := 0; i < len(ss); i++ {
		s += ss[i]
	}
	ss = strings.Split(s, "\t")
	s = ""
	for i := 0; i < len(ss); i++ {
		s += ss[i]
	} //}125 { 123

	for i := 1; i < len(s); i++ { //bviaflrkh"
		if s[i] == ',' && (s[i-1] == '{' || s[i-1] == '[') {
			s = s[:i] + s[i+1:]
			i--
		}
		if (s[i] == ']' || s[i] == '}') && s[i-1]+2 == s[i] {
			r := i
			l := i - 1
			for ; s[l-1] != '[' && s[l-1] != '{' && s[l-1] != ',' && s[l-1] != ']' && s[l-1] != '}'; l-- {
			}
			if s[l-1] == ',' {
				l--
			}
			s = s[:l] + s[r+1:]
			i = l - 1
		}
	}
	s = "[" + s + "]"

	fmt.Fprintln(out, s)
	sTest := "[{\"qtfffdilio\":[\"upobhn\"]},{\"drwkqrtmn\":{\"fkudd\":\"maerj\",\"nnfi\":\"uooewq\"}},{\"sgomvjfk\":[\"mpnbukihw\"]},{\"mbigieteq\":{\"hdl\":[\"nwcokit\"]}},[{\"bviaflrkh\":\"meohj\",\"ygsbw\":\"fjj\"}],{\"kavykcjixy\":\"y\"},[{\"jgqshnndx\":\"na\"}],[{\"geoyts\":\"blksv\"}],{\"fh\":{\"uxhdn\":{\"jgtsnuoo\":[\"bdk\"]}}},[\"j\",\"mknnqxsqoc\"]]\n"
	fmt.Fprintln(out)
	fmt.Fprintln(out, sTest)

	fmt.Fprintln(out, s == sTest)

}
