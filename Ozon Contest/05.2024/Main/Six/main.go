package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	/*in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()*/
	dir := "C:\\Users\\Zver\\Documents\\GitHub\\StudyGo\\Ozon Contest\\05.2024\\Main\\Six\\tests\\"
	list, _ := os.ReadDir(dir)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	m := make(map[int]int)
	bb := make([]int, 0, 30)
	for numTest := 0; numTest < len(list); numTest += 2 {
		dataTest, answerTest := list[numTest].Name(), list[numTest+1].Name()

		fileReadData, err := os.Open(dir + dataTest)
		if err != nil {
			log.Fatal(err)
		}

		fileReadAnswer, err := os.Open(dir + answerTest)
		if err != nil {
			log.Fatal(err)
		}
		//your code is here

		in := bufio.NewReader(fileReadData)
		tests := bufio.NewReader(fileReadAnswer)

		t := 0
		fmt.Fscan(in, &t)

		for i := 0; i < t; i++ {
			n, k := 0, 0
			fmt.Fscan(in, &n, &k)
			b := 0
			fmt.Fscan(in, &b)

			for bi := 0; bi < b; bi++ {
				help := 0
				fmt.Fscan(in, &help)
				help = 1 << help
				if _, ok := m[help]; !ok {
					bb = append(bb, help)
				}
				m[help]++
			}
			sort.Slice(bb, func(i, j int) bool { return bb[i] > bb[j] })

			colMove := 1
		loop:
			for ; len(m) > 0; colMove++ {
				for colCar := 0; colCar < n; colCar++ {
					kg := 0
					/*help := k - kg
					l := len(strconv.FormatInt(int64(help), 2)) - 1
					help = 1 << l

					for help > 0 {

						if _, ok := m[help]; ok {
							kg += help
							m[help]--
							if m[help] == 0 {
								delete(m, help)
								if len(m) == 0 {
									break loop
								}
							}
							if kg == k {
								break
							}

							help = k - kg
							l = len(strconv.FormatInt(int64(help), 2)) - 1
							help = 1 << l

						} else {
							help = help >> 1
						}
					}*/
					for i := 0; i < len(bb); i++ {
						if _, ok := m[bb[i]]; ok {
							if kg+bb[i] <= k {
								kg += bb[i]
								m[bb[i]]--
								if m[bb[i]] == 0 {
									delete(m, bb[i])
									bb = append(bb[:i], bb[i+1:]...)
									if len(m) == 0 {
										break loop
									}
								}
								if kg == k {
									break
								}
								i--
							}
						}
					}
				}
			}
			rez := 0
			fmt.Fscan(tests, &rez)

			if colMove != rez {
				fmt.Fprintln(out, dataTest, i, " -> ", colMove, rez)
				fileReadData.Close()
				fileReadAnswer.Close()
				return
			}
		}

		//must be in the end
		fileReadData.Close()
		fileReadAnswer.Close()
	}
	fmt.Fprintln(out, "Done")

}
