package main

import "fmt"

var dizi = [9][9]int{
	{7, 1, 0, 4, 0, 9, 0, 0, 2},
	{0, 8, 0, 5, 0, 0, 0, 0, 0},
	{9, 0, 0, 0, 3, 0, 0, 1, 0},
	{8, 3, 9, 0, 0, 2, 1, 0, 4},
	{0, 0, 7, 0, 4, 0, 2, 0, 0},
	{4, 0, 1, 3, 0, 0, 7, 9, 5},
	{0, 5, 0, 0, 7, 0, 0, 0, 8},
	{0, 0, 0, 0, 0, 5, 0, 3, 0},
	{6, 0, 0, 1, 0, 3, 0, 4, 7},
}

func sayikontrolu(x, y, sayi int) bool {
	for i := 0; i < 9; i++ {
		if dizi[x][i] == sayi || dizi[i][y] == sayi {
			return false
		}
	}
	kutux, kutuy := x/3*3, y/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if dizi[i+kutux][j+kutuy] == sayi {
				return false
			}
		}
	}
	return true
}

func sudokucoz() bool {

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if dizi[x][y] == 0 {
				for sayi := 1; sayi <= 9; sayi++ {
					if sayikontrolu(x, y, sayi) {
						dizi[x][y] = sayi
						if sudokucoz() {
							return true
						}
						dizi[x][y] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func sudouyazdir() {
	for _, x := range dizi {
		for _, sayi := range x {
			fmt.Printf("%d ", sayi)
		}
		fmt.Println()
	}
}

func main() {
	if sudokucoz() {
		sudouyazdir()
	} else {
		fmt.Println("error")
	}
}
