package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

func recursion(colours [][][]int, X1, Y1, X2, Y2 int, square int, colColours int, colPoints int) int {
	//X1, Y1 := 0, 0 // width,length
	//X2, Y2 := len(colours[0])-1, len(colours)-1
	//square := (X2 - X1 + 1) * (Y2 - Y1 + 1)
	//colPoints := 0
	//for i := 0; i < len(colours); i++ { // matrixes
	tempSquareBack := (X2 - X1 + 1) * (Y2 - Y1 + 1)
	tempSquare := 0
	xx1, yy1, xx2, yy2 := 0, 0, 0, 0
	if len(colours) == 0 {
		return square
	}

	i := 0
	//	}
	for j := 0; j < len(colours[i]); j++ { //rows
		if colours[i][j] != nil {
			for l := 0; l < len(colours[i][j]); l++ { //elements

				x1, y1, x2, y2 := 0, 0, 0, 0
				if len(colours) == colColours { // take first point
					x1, y1, x2, y2 = j, colours[i][j][l], j, colours[i][j][l]
				} else {
					x1, y1, x2, y2 = X1, Y1, X2, Y2

					if x1 > j {
						x1 = j
					}
					if y1 > colours[i][j][l] {
						y1 = colours[i][j][l]
					}

					if x2 < j {
						x2 = j
					}
					if y2 < colours[i][j][l] {
						y2 = colours[i][j][l]
					}
					if x1 > x2 {
						x1, x2 = x2, x1
					}
					if y1 > y2 {
						y1, y2 = y2, y1
					}
				}
				// если прямоугольник не изменился и остался там же где и был, после выбора последующего
				// элемента одного цвета, то искать нет смысла, будут повторы
				if tempSquare == (x2-x1+1)*(y2-y1+1) {
					if xx1 == x1 && yy1 == y1 && xx2 == x2 && yy2 == y2 {
						continue
					}
				}

				tempSquare = (x2 - x1 + 1) * (y2 - y1 + 1)
				if tempSquare < square {
					//X1, X2, Y1, Y2 = x1, x2, y1, y2
					colPoints++

					square = recursion(colours[1:], x1, y1, x2, y2, square, colColours, colPoints)
					xx1, yy1, xx2, yy2 = x1, y1, x2, y2

					if colPoints == colColours && square > tempSquare && len(colours) == 1 {
						square = tempSquare
					}

					if tempSquareBack == tempSquare {
						return square
					}
					colPoints--
				} else {
					continue
				}
			}
		}
	}

	return square
}

func main() {
	dir := ".\\tests\\"
	list, _ := os.ReadDir(dir)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	t := time.Now()
	t1 := t
	var t2 time.Time
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

		dataCol := 0 //data sets
		fmt.Fscan(in, &dataCol)
		//var colours [][][]int
		for col := 0; col < dataCol; col++ {
			n, m := 0, 0 // columns, rows
			fmt.Fscan(in, &n, &m)

			kCol := 0 //col colours
			fmt.Fscan(in, &kCol)
			colours := make([][][]int, kCol, kCol)

			for k := 0; k < kCol; k++ {
				iColourCol := 0 // col points of one colour
				fmt.Fscan(in, &iColourCol)
				colours[k] = make([][]int, n, n)
				for ii := 0; ii < iColourCol; ii++ {
					i, j := 0, 0
					fmt.Fscan(in, &i, &j)
					colours[k][i-1] = append(colours[k][i-1], j-1)
				}
			}

			square := recursion(colours, 0, 0, 0, -1, n*m, len(colours), 0)

			test := 0
			fmt.Fscan(fileReadAnswer, &test)
			//fmt.Println(square, test)
			if square != test {
				fmt.Println("Test N =", dataTest, "Failed", square, "!=", test, "Dataset", col)
				//fmt.Println(colours)
				return
			}
		}

		//must be in the end
		t2 = time.Now()
		fmt.Println("Test N =", dataTest, "is done", t2.Sub(t1).Seconds(), "sec")
		t1 = t2
		fileReadData.Close()
		fileReadAnswer.Close()
	}
	fmt.Fprintln(out, "Total time is", t2.Sub(t))
	fmt.Fprintln(out, "Middle time is", t2.Sub(t)/time.Duration(len(list)))
	fmt.Scan()
	b := bytes.ToUpper([]byte{' '})[0]
}
