// Dependencies :
// 	github.com/hotei/bmp
// How to install:
// 	go get github.com/hotei/bmp

// USAGE:
// package main

// import (
// 	"CaptchaParser"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	reader, err := os.Open("captcha.bmp")
// 	if err != nil {
// 		log.Fatal("File error")
// 	}
// 	output := captcha.GetCaptcha(reader)
// 	fmt.Println(output)
// }

package captcha

import (
	"github.com/hotei/bmp"
	"image"
	"os"
)

func pixelLoad(img image.Image) [][]int {
	b := img.Bounds()
	pix := [][]int{}
	for y := b.Min.Y; y < b.Max.Y; y++ {
		row := []int{}
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if r == 65535 && g == 65278 && b == 57311 && a == 65535 {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
		}
		pix = append(pix, row)
	}
	return pix
}

func match_img(rx, ry int, pix1 [][]int, mask1 [][]string) int {
	flag := 1
	breakflag := 0
	yy := len(mask1)
	count := 0
	for y := 0; y < yy; y++ {
		for x := 0; x < len(mask1[y]); x++ {
			if ry+y < 24 && rx+x < 129 {
				if mask1[y][x] == "1" {
					if pix1[ry+y][rx+x] == 1 {
						count = count + 1
					} else {
						flag = 0
						breakflag = 1
						break
					}
				}
			} else {
				flag = 0
				breakflag = 1
				break
			}
		}
		if breakflag == 1 {
			break
		}
	}
	if count == 0 {
		flag = 0
	}
	return flag
}

func skip(start, end []int, y int) int {
	flag := 0
	for i := 0; i < len(start); i++ {
		if y >= start[i] && y <= end[i] {
			flag = 1
			break
		}
	}
	return flag
}

func sort(sorter []int, captcha []string) {
	for i := 0; i < len(sorter); i++ {
		less := sorter[i]
		swap := 0
		ls := i
		for k := i; k < len(sorter); k++ {
			if sorter[k] < less {
				less = sorter[k]
				ls = k
				swap = 1
			}
		}
		if swap == 1 {
			sorter[i], sorter[ls] = sorter[ls], sorter[i]
			captcha[i], captcha[ls] = captcha[ls], captcha[i]
		}
	}
}

func GetCaptcha(file *os.File) string {
	keys := map[string][][]string{
		"0": [][]string{
			[]string{"0", "0", "0", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "1", "1", "0"},
			[]string{"1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "0", "0", "0", "1", "1", "0", "1", "1"},
			[]string{"1", "1", "0", "0", "1", "1", "0", "0", "1", "1"},
			[]string{"1", "1", "0", "1", "1", "0", "0", "0", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1"},
			[]string{"0", "1", "1", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "0", "0", "0"},
		},
		"1": [][]string{
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
		},
		"2": [][]string{
			[]string{"0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"1", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
		},
		"3": [][]string{
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
		},
		"4": [][]string{[]string{"0", "0", "0", "0", "0", "0", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0", "0"},
		},
		"5": [][]string{[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"}}, "6": [][]string{[]string{"0", "0", "0", "0", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "0", "0", "0"}}, "7": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "0", "0", "0", "0"}}, "8": [][]string{[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "0", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"}}, "9": [][]string{[]string{"0", "0", "0", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "0", "1", "1", "1"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "0", "0", "0", "0"}}, "A": [][]string{[]string{"0", "0", "0", "0", "1", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "1", "1"}}, "B": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"}}, "C": [][]string{[]string{"0", "0", "0", "0", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "0", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "1"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "0", "1", "1"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1", "1", "1", "0"}}, "D": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0", "0"}}, "E": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"}}, "F": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"}}, "G": [][]string{[]string{"0", "0", "0", "0", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1", "1", "1", "1", "0", "0"}}, "H": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"}}, "I": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1"}}, "J": [][]string{[]string{"0", "0", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "0", "0"}}, "K": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"}}, "L": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1"}}, "M": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "0", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "1", "1", "0", "1", "1", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "1", "1", "1", "1", "1", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"}}, "N": [][]string{[]string{"1", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "1", "1", "1", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "1", "1", "1", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"}}, "O": [][]string{[]string{"0", "0", "0", "0", "1", "1", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1", "1", "0", "0", "0", "0"}}, "P": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0"}}, "Q": [][]string{[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "1", "1"},
			[]string{"1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "1", "1"},
			[]string{"1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "1", "1"},
			[]string{"1", "1", "0", "0", "0", "0", "0", "1", "0", "0", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "0", "1", "1"},
			[]string{"0", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "1", "0"}}, "R": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "1", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"}}, "S": [][]string{[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "1", "1", "0"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "0"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "0", "0", "0"}}, "T": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"}}, "U": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "1", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "0", "0", "0"}}, "V": [][]string{[]string{"1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1", "0", "0", "0", "0"}}, "W": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "1", "1", "0", "1", "1", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "0", "1", "1", "0", "1", "1", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "1", "1", "1", "0", "1", "1", "1", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "1", "1", "1", "0", "1", "1", "1", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "1", "1", "0", "0", "0", "1", "1", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "0", "1", "1", "0", "0", "0", "1", "1", "0", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1", "1", "0", "0"}}, "X": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "0", "1", "1", "1"}}, "Y": [][]string{[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "1", "1"},
			[]string{"0", "1", "1", "1", "0", "0", "0", "1", "1", "1", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "1", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0"}}, "Z": [][]string{[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "0", "1", "1", "1", "1"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1", "1", "0"},
			[]string{"0", "0", "0", "0", "0", "1", "1", "1", "0", "0"},
			[]string{"0", "0", "0", "0", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "1", "0", "0", "0"},
			[]string{"0", "0", "0", "1", "1", "1", "0", "0", "0", "0"},
			[]string{"0", "0", "1", "1", "1", "0", "0", "0", "0", "0"},
			[]string{"0", "1", "1", "1", "1", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "0", "0", "0", "0", "0", "0"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
			[]string{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"}},
	}
	order := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	img, _ := bmp.Decode(file)
	pix := pixelLoad(img)
	newpix := [][]int{}
	for y := 0; y < 24; y++ {
		t := []int{}
		for x := 0; x < 129; x++ {
			temp := pix[y][x]
			if y != 0 && y != 24 {
				if pix[y+1][x] == 0 && temp == 1 && pix[y-1][x] == 0 {
					pix[y][x] = 0
				}
			}
			t = append(t, pix[y][x])
		}
		newpix = append(newpix, t)
	}
	xoff := 20
	yoff := 2
	skipstart := []int{}
	skipend := []int{}
	sorter := []int{}
	captcha := []string{}
	for l := 0; l < 36; l++ {
		mask := keys[order[l]]
		f := 0
		for y := yoff; y < 24; y++ {
			for x := xoff; x < 129; x++ {
				if skip(skipstart, skipend, x) == 0 {
					if match_img(x, y, newpix, mask) != 0 {
						skipstart = append(skipstart, x)
						skipend = append(skipend, x+len(mask[0]))
						sorter = append(sorter, x)
						captcha = append(captcha, order[l])
						f = f + 1
					}
				}
			}
		}
		if f == 6 {
			break
		}
	}
	sort(sorter, captcha)
	result := ""
	for cap := range captcha {
		result = result + captcha[cap]
	}
	return result

}
