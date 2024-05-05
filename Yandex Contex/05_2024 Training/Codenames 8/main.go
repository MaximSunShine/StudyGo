package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)

	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t, n, b, r, f := 0, 0, 0, 0, 0
	//n, b, r, f := 6, 2, 2, 6
	//n, b, r, f := 9, 1, 1, 9
	//w := n - b - r - 1
	//w = w
	fmt.Fscan(in, &t)
	for col := 0; col < t; col++ {
		fmt.Fscan(in, &n, &b, &r, &f)
		ss := make([]string, n, n)
		mBlack := make(map[string]struct{})

		//	ss = []string{"aaaa", "ccbbaa", "waaw", "abc", "caaccaa", "bcaa"}
		//ss = []string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
		strMax := ""
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &ss[i])
			if len(strMax) < len(ss[i]) {
				strMax = ss[i]
			}
		}
		f--
		for i := 1; i <= len(ss[f]); i++ {
			for j := 0; j+i <= len(ss[f]); j++ {
				if _, inMap := mBlack[ss[f][j:j+i]]; !inMap {
					mBlack[ss[f][j:j+i]] = struct{}{}
				}
			}
		}

		mAll := make(map[string]struct{})
		for i := 0; i < n; i++ {
			mAll[ss[i]] = struct{}{}
		}
		delete(mAll, ss[f])

		mBlueRed := make(map[string]int)
		for k := 0; k < b; k++ {
			mHelp := make(map[string]struct{})
			if _, inMap := mBlack[ss[k]]; inMap {
				continue
			}
			for i := len(ss[k]); i > 0; i-- {
				for j := 0; j+i <= len(ss[k]); j++ {
					if _, inMap := mBlack[ss[k][j:j+i]]; inMap {
						continue
					}
					if _, inMap := mAll[ss[k][j:j+i]]; inMap {
						continue
					}
					if _, inMap := mHelp[ss[k][j:j+i]]; !inMap {
						mHelp[ss[k][j:j+i]] = struct{}{}
						mBlueRed[ss[k][j:j+i]]++
					}
				}
			}
		}
		for k := b; k < b+r; k++ {
			mHelp := make(map[string]struct{})
			if _, inMap := mBlack[ss[k]]; inMap {
				continue
			}
			for i := len(ss[k]); i > 0; i-- {
				for j := 0; j+i <= len(ss[k]); j++ {
					if _, inMap := mBlack[ss[k][j:j+i]]; inMap {
						continue
					}
					if _, inMap := mAll[ss[k][j:j+i]]; inMap {
						continue
					}
					if _, inMap := mBlueRed[ss[k][j:j+i]]; !inMap {
						continue
					}
					if _, inMap := mHelp[ss[k][j:j+i]]; !inMap {
						mHelp[ss[k][j:j+i]] = struct{}{}
						mBlueRed[ss[k][j:j+i]]--
					}
				}
			}
		}

		Max := 0
		str := ""
		for key, val := range mBlueRed {
			if Max < val {
				Max = val
				str = key
			}
		}
		if Max > 0 {
			fmt.Fprintln(out, str, Max)
		} else if len(strMax) < 10 {
			fmt.Fprintln(out, strMax+"a", Max)
		} else {
			fmt.Fprintln(out, "tkhapjiabb ", Max)
		}
	}

}
