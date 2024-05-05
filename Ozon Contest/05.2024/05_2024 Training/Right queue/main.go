package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main1() {

	/*var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()*/

	for n := 1; n <= 27; n++ {
		nameFile := strconv.Itoa(n)

		fileReadQuestion, err := os.Open("C:/Users/Zver/Documents/GitHub/StudyGo/Yandex Contex/05_2024 Training/Tests/" + nameFile)
		if err != nil {
			log.Fatal(err)
		}
		defer fileReadQuestion.Close()

		fileReadAnswer, err := os.Open("C:/Users/Zver/Documents/GitHub/StudyGo/Yandex Contex/05_2024 Training/Tests/" + nameFile + ".a")
		if err != nil {
			log.Fatal(err)
		}
		defer fileReadAnswer.Close()

		filePrint, err := os.Create("C:/Users/Zver/Documents/GitHub/StudyGo/Yandex Contex/05_2024 Training/Tests/Print.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer filePrint.Close()

		inQuestion := bufio.NewReader(fileReadQuestion)
		inAnswer := bufio.NewReader(fileReadAnswer)
		out := bufio.NewWriter(filePrint)
		defer out.Flush()

		col, bLen := 0, 0

		fmt.Fscan(inQuestion, &col)
		for i := 0; i < col; i++ {
			fmt.Fscan(inQuestion, &bLen)
			m := make(map[byte]int)

			b := make([]byte, 0, bLen)
			fmt.Fscan(inQuestion, &b)

			answer := ""
			fmt.Fscan(inAnswer, &answer)

			if b[0] == 'Z' || b[len(b)-1] == 'X' {
				if "No" != answer {
					fmt.Println("test ", i, " ->", b)
				}
				//fmt.Fprintln(out, "No")
				continue
			}

			for i := 0; i < len(b); i++ {
				m[b[i]]++
			}
			if m['X'] > (m['Y']+m['Z']) || m['Z'] > (m['X']+m['Y']) || m['Y'] > (m['X']+m['Z']) {
				if "No" != answer {
					fmt.Println("test ", i, " ->", b)
				}
				//fmt.Fprintln(out, "No")
				continue
			}
			colX := 0
			colY := 0
			colZ := 0

			t := true

			for j := 0; j < len(b); j++ {
				if b[j] == 'X' {
					colX++
					continue
				}
				if b[j] == 'Y' {
					if colX > 0 && m['X'] >= (m['Y']+m['Z']) { //colX?
						if m['Z'] >= 0 { //m['Z'] == 0 || m['Z'] == 1 { //|| m['X'] > m['Z']
							colX--
							m['X']--
							m['Y']--
							continue
						} else {
							colY++
							continue
						}
					} else {
						if m['Z'] == 0 { //
							t = false
							break
						} else {
							//colY--
							colZ++
							m['Z']--
							m['Y']--
						}
					}
					continue
				}
				if b[j] == 'Z' && colZ == 0 {
					if colY > 0 && m['Y'] > m['X'] {
						colY--
						//colZ--
						m['Z']--
						m['Y']--
					} else if colX > 0 {
						colX--
						//colZ--
						m['Z']--
						m['X']--
					} else {
						t = false
						break
					}
				} else {
					colZ--
				}
				if m['X'] > (m['Y']+m['Z']) || m['Z'] > (m['X']+m['Y']) || m['Y'] > (m['X']+m['Z']) {
					break
				}
			}

			test := ""
			if m['X'] == m['Y'] && t { //colX == colY {
				//fmt.Fprintln(out, "YES")
				test = "Yes"
			} else {
				//fmt.Fprintln(out, "NO")
				test = "No"
			}

			if test != answer {
				fmt.Println("test ", i, " ->", string(b), "my ", test, " need ", answer)
				break
			}
		}
		//fmt.Println(n, "-> done")
	}

}
func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	col, bLen := 0, 0

	fmt.Fscan(in, &col)

	s := make([]string, 0, col)

	for i := 0; i < col; i++ {
		fmt.Fscan(in, &bLen)
		m := make(map[byte]int)

		b := make([]byte, 0, bLen)
		fmt.Fscan(in, &b)

		if b[0] == 'Z' || b[len(b)-1] == 'X' {
			fmt.Fprintln(out, "No")
			continue
		}

		for i := 0; i < len(b); i++ {
			m[b[i]]++
		}
		if m['X'] > (m['Y']+m['Z']) || m['Z'] > (m['X']+m['Y']) || m['Y'] > (m['X']+m['Z']) {
			fmt.Fprintln(out, "No")
			continue
		}
		colX := 0
		colY := 0
		colZ := 0

		t := true

		for j := 0; j < len(b); j++ {
			if b[j] == 'X' {
				colX++
				continue
			}
			if b[j] == 'Y' {
				if colX > 0 && m['X'] >= (m['Y']+m['Z']) { //colX?
					if m['Z'] >= 0 { //m['Z'] == 0 || m['Z'] == 1 { //|| m['X'] > m['Z']
						colX--
						m['X']--
						m['Y']--
						continue
					} else {
						colY++
						continue
					}
				} else {
					if m['Z'] == 0 {
						t = false
						break
					} else {
						colZ++
						m['Z']--
						m['Y']--
					}
				}
				continue
			}
			if b[j] == 'Z' && colZ == 0 {
				if colY > 0 && m['Y'] > m['X'] {
					colY--
					m['Z']--
					m['Y']--
				} else if colX > 0 {
					colX--
					m['Z']--
					m['X']--
				} else {
					t = false
					break
				}
			} else {
				colZ--
			}
			if m['X'] > (m['Y']+m['Z']) || m['Z'] > (m['X']+m['Y']) || m['Y'] > (m['X']+m['Z']) {
				break
			}
		}

		if m['X'] == m['Y'] && t { //colX == colY {
			s = append(s, "Yes") //fmt.Fprintln(out, "YES")
		} else {
			s = append(s, "No") //fmt.Fprintln(out, "NO")
		}
	}
	for _, v := range s {
		fmt.Fprintln(out, v)
	}

}
