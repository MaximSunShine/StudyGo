package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type directory struct {
	Dir     string       `json:"dir"`
	Files   []string     `json:"files"`
	Folders []*directory `json:"folders"`
}

func Hucked(d *directory) bool {
	for _, v := range d.Files {
		if len(v) > 4 && v[len(v)-5:len(v)] == ".hack" {
			return true
		}
	}
	return false
}
func findColViruses(d *directory, hacked bool) int {

	col := 0

	if hacked || Hucked(d) {
		hacked = true
		col = len(d.Files)
	}

	for i := 0; i < len(d.Folders); i++ {
		col += findColViruses(d.Folders[i], hacked)
	}

	return col
}
func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t := 0
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		d := &directory{}
		n := 0
		fmt.Fscanln(in, &n)
		var bdata []byte

		for j := 0; j < n; j++ {
			help, _, _ := in.ReadLine()
			bdata = append(bdata, help...)
		}
		if err := json.Unmarshal(bdata, &d); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintln(out, findColViruses(d, false))

	}
}

/*
231
195
242

0
195
43


*/
