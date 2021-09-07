package main

import (
	"fmt"
	"strings"
)

func main() {
	var pilih int
	var jumlah = 0
	var x = true
	var yt string

	for x {
		fmt.Println("====PIRAMIDA BINTANG====")
		fmt.Println("1. Piramida Segitiga")
		fmt.Println("2. Piramida Terbalik")
		fmt.Print("Masukkan Pilihan Anda: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Masukkan jumlah barisnya: ")
			fmt.Scan(&jumlah)

			for i := 1; i <= jumlah*2; i++ {
				if i%2 != 0 {
					for j := jumlah * 2; j >= i+2; j-- {
						fmt.Print(" ")
					}
					for k := 1; k <= i; k++ {
						fmt.Print(" *")
					}
					fmt.Println()
				}
			}
			fmt.Print("\nLanjutkan Program? (y/t)")
			fmt.Scan(&yt)

			yt = strings.ToLower(yt)

			if yt == "y" {
				x = true
			} else {
				x = false
			}
		case 2:
			fmt.Print("Masukkan jumlah barisnya: ")
			fmt.Scan(&jumlah)

			for i := jumlah * 2; i >= 1; i-- {
				if i%2 != 0 {
					for j := jumlah * 2; j >= i+2; j-- {
						fmt.Print(" ")
					}
					for k := i; k >= 1; k-- {
						fmt.Print(" *")
					}
					fmt.Println()
				}
			}
			fmt.Print("\nLanjutkan Program? (y/t)")
			fmt.Scan(&yt)

			yt = strings.ToLower(yt)

			if yt == "y" {
				x = true
			} else {
				x = false
			}
		}
	}

}
