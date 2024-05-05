package main

import (
	"bufio"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sort"
)

type pacient struct {
	i int
	b byte
}

type helpPacient struct {
	p *pacient
}

func main() {
	/*
		// Создаем файл для хранения профиля процессора
		f, err := os.Create("cpu.pprof")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Запуск профиля CPU
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
		// Запустите здесь код вашего приложения
	*/
	go http.ListenAndServe("0.0.0.0:8080", nil)

	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t, n, m := 0, 0, 0
	fmt.Fscan(in, &t)

	ss := make([]string, 0, n)
mainLoop:
	for i := 0; i < t; i++ {

		//ss = make([]string, 0, m)
		fmt.Fscan(in, &n, &m)

		mm := make([]pacient, m)
		helpmm := make([]helpPacient, m)
		for i, _ := range mm {
			fmt.Fscan(in, &mm[i].i)
			mm[i].b = '0'
			helpmm[i].p = &mm[i]
		}
		if m == 1 {
			//fmt.Fprintln(out, "0")
			ss = append(ss, "0")
			continue
		}
		sort.Slice(helpmm, func(i, j int) bool { return helpmm[i].p.i < helpmm[j].p.i })

		l := len(helpmm) - 1
		if len(helpmm) > 2 {
			if helpmm[0].p.i == 1 && helpmm[0].p.i == helpmm[2].p.i || helpmm[0].p.i == n && helpmm[l].p.i == helpmm[l-2].p.i {
				//fmt.Fprintln(out, "x")
				ss = append(ss, "x")
				continue
			}
		}

		col := 1
		isZeroBack := helpmm[0].p.i > 1
		needZeroForward := false
		zeroWindow := 0
		for j := 1; j < m; j++ {
			if helpmm[j].p.i-helpmm[j-1].p.i > 1 {
				if helpmm[j].p.i-helpmm[j-1].p.i > 2 || !needZeroForward {
					isZeroBack = true
					zeroWindow = j
				}
				needZeroForward = false
				col-- //?
			}

			if helpmm[j].p.i == helpmm[j-1].p.i {
				col++
			} else {
				col = 1
			}

			if needZeroForward && col == 2 {
				//fmt.Fprintln(out, "x")
				ss = append(ss, "x")
				continue mainLoop
			}
			if needZeroForward && col == 1 {
				helpmm[j].p.b = '+'
				continue
			}

			if col > 3 {
				//fmt.Fprintln(out, "x")
				ss = append(ss, "x")
				continue mainLoop
			}
			if col > 1 {
				if isZeroBack {
					for k := zeroWindow; k < j; k++ {
						helpmm[k].p.b = '-'
					}
					isZeroBack = false
					col--
				} else {
					if col == 3 {
						//fmt.Fprintln(out, "x")
						ss = append(ss, "x")
						continue mainLoop
					}
					needZeroForward = true
					helpmm[j].p.b = '+'
				}
			}
		}
		if needZeroForward && helpmm[len(helpmm)-1].p.i == n {
			//fmt.Fprintln(out, "x")
			ss = append(ss, "x")
			continue
		}
		s := ""
		for _, v := range mm {
			s += string(v.b)
		}
		//fmt.Fprintln(out, s)
		ss = append(ss, s)
	}
	for _, v := range ss {
		fmt.Fprintln(out, v)
	}
}
