package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Athletes struct {
	time  int
	place int
}

type SortAthletes struct {
	p *Athletes
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t := 0
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		n := 0
		fmt.Fscan(in, &n)

		if n == 1 {
			fmt.Fprintln(out, 1)
			help := 0
			fmt.Fscan(in, &help)
			continue
		}

		aa := make([]Athletes, n, n)
		ss := make([]SortAthletes, 0, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &aa[j].time)
			i := sort.Search(len(ss), func(i int) bool { return ss[i].p.time >= aa[j].time })
			ss = append(ss[:i], append([]SortAthletes{{&aa[j]}}, ss[i:]...)...)
		}

		place := 1
		ss[0].p.place = 1
		help := 1
		for i := 1; i < len(ss); i++ {
			help++
			if ss[i].p.time != ss[i-1].p.time && ss[i].p.time != ss[i-1].p.time+1 {
				place = help
			}
			ss[i].p.place = place
		}

		for i := 0; i < len(aa); i++ {
			fmt.Fprint(out, aa[i].place, " ")
		}
		fmt.Fprintln(out)
	}
}

/*func addss(ss []SortAthletes, athlete *Athletes) []SortAthletes {
	i := sort.Search(len(ss), func(i int) bool { return ss[i].p.time >= athlete.time })
	ss = append(ss[:i], append([]SortAthletes{SortAthletes{athlete}}, ss[i:]...)...)
	return ss
}*/
/*
func main1() {

	nameFile := "50"
	fileReadQuestion, err := os.Open("C:/Users/Zver/Documents/GitHub/StudyGo/Ozon Contest/05.2024/05_2024 Training/Competition/tests/" + nameFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReadQuestion.Close()

	/*fileReadAnswer, err := os.Open("C:/Users/Zver/Documents/GitHub/StudyGo/Ozon Contest/05.2024/05_2024 Training/Right queue/Tests/" + nameFile + ".a")
	if err != nil {
		log.Fatal(err)
	}
	defer fileReadAnswer.Close()

	filePrint, err := os.Create("C:/Users/Zver/Documents/GitHub/StudyGo/Ozon Contest/05.2024/05_2024 Training/Competition/tests/Print.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer filePrint.Close()

	in := bufio.NewReader(fileReadQuestion)
	//inAnswer := bufio.NewReader(fileReadAnswer)
	out := bufio.NewWriter(filePrint)
	defer out.Flush()

	//var in *bufio.Reader
	//in = bufio.NewReader(os.Stdin)
	//var out *bufio.Writer
	//out = bufio.NewWriter(os.Stdout)
	//defer out.Flush()

	var aa []Athletes
	var ss []SortAthletes
	ss = ss
	t := 0
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		n := 0
		fmt.Fscan(in, &n)

		if n == 1 {
			fmt.Fprintln(out, 1)
			continue
		}

		aa = make([]Athletes, n, n)
		ss := make([]SortAthletes, 0, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &aa[j].time)
		}
		for j := 0; j < n; j++ {
			ss = addss(ss, &aa[j])
		}

		place := 1
		ss[0].p.place = 1
		help := 1
		for i := 1; i < len(ss); i++ {
			help++
			if ss[i].p.time != ss[i-1].p.time && ss[i].p.time != ss[i-1].p.time+1 {
				place = help
			}
			ss[i].p.place = place
		}

		for i := 0; i < len(aa); i++ {
			fmt.Fprint(out, aa[i].place, " ")
		}
		fmt.Fprintln(out)
	}

}
*/
