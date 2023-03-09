package main

import (
	"fmt"
)

func main() {
	for i :=0; i<5; i++ {
		fmt.Println("Nilai i = ",i) //menampilkan perulangan untuk nilai i yang dimulai dari 0 sampai 4
		}
		
		var j int = 0;
		for j<5 {
			fmt.Println("Nilai j = ",j) //menampilkan perulangan untuk nilai j yang dimulai dari 0 sampai 4 (kurang dari 5)
			j++
		}

		for pos, char := range "САШАРВО" {
 			fmt.Printf("character %#U starts at byte position %d\n", char, pos) //menampilkan perulangan untuk posisi byte dari range char "САШАРВО"
		}
		
		for {
			if j++; j <= (10) {
				fmt.Println("Nilai j = ",j) //menampilkan perulangan untuk lanjutan nilai j sebelumnya yang dimulai dari 6 sampai 10 (kurang sama dengan 10)
				}else{
					break
				}
		}
}

 
