package main

import "fmt"

const NMAX int = 100

type tabGame [NMAX] string

func main(){
	var dataGame tabGame
	var nDataGame, option int

	for option!=5{
		Menu()
		fmt.Print("Masukan Opsi : ")
		fmt.Scan(&option)
		switch option{
		case 1: InputGame(&dataGame, &nDataGame)
		case 2: OuputGame(dataGame, nDataGame)
		case 3: Rename(&dataGame, nDataGame)
		case 4: Delete(&dataGame, &nDataGame)
		case 5: fmt.Println("Program Selesai, TERIMA KASIH")
		default : fmt.Println("Opsi tidak valid")
		}
	}
}

func Menu(){
	fmt.Println("\t===Menu===")
	fmt.Println("1-Input Judul Game")
	fmt.Println("2-List Judul Game")
	fmt.Println("3-Ganti Judul Game")
	fmt.Println("4-Delete Judul Game")
	fmt.Println("5-Exit")
}

func InputGame(DataGame *tabGame, n *int){	
	var inputLama, inputBaru int 
	inputLama += *n
	fmt.Print("Masukan jumlah judul game yang ingin di input : ")
	fmt.Scan(&inputBaru)
	inputBaru += *n
	
	for inputLama<inputBaru{
		fmt.Scan(&DataGame[inputLama])
		inputLama++
	}
	*n = inputBaru
	fmt.Println("Data Berhasil Ditambahkan")
}

// =============================================================================== Opsi simple GPT
// func InputGame(DataGame *tabGame, n *int){	
// 	var  inputBaru int 
// 	fmt.Print("Masukan jumlah judul game yang ingin di input : ")
// 	fmt.Scan(&inputBaru)
	
// 	for i:=0;i<inputBaru;i++{
// 		fmt.Scan(&DataGame[*n])
// 		*n++
// 	}
// 	fmt.Println("Data Berhasil Ditambahkan")
// }
// ===============================================================================


func OuputGame(DataGame tabGame, n int){
	for i:=0;i<n;i++{
		fmt.Printf("%s ", DataGame[i])
	}
	fmt.Println()
}

func SeqSearch(DataGame tabGame, n int, cari string)int{
	var idx , i int
	idx = -1
	i = 0

	for i<n&&idx==-1{
		if DataGame[i] == cari{
			idx = i
		}
		i++
	}
	return idx
}

func Rename(dataGame *tabGame, n int){
	var Mencari int 
	var namaLama, namaBaru string
	fmt.Print("Masukan nama yang ingin di rename : ")
	fmt.Scan(&namaLama)
	Mencari = SeqSearch(*dataGame, n, namaLama)

	if Mencari != -1{
		fmt.Print("Masukan nama baru : ")
		fmt.Scan(&namaBaru)
		dataGame[Mencari] = namaBaru
		fmt.Println("Judul berhasil di ganti.")
	} else{
		fmt.Println("Judul tidak ditemukan")
	}
}

func Delete(dataGame *tabGame, n *int){
	var Mencari, i int
	var HapusNama string
	fmt.Print("Masukan Judul yang ingin dihapus : ")
	fmt.Scan(&HapusNama)
	Mencari = SeqSearch(*dataGame, *n, HapusNama)

	if Mencari != -1{
		i = Mencari
		for i < *n-1{
			dataGame[i] = dataGame[i+1]
			i++
		}
		dataGame[*n-1] = ""	
		*n--
		fmt.Println("Judul berhasil di hapus")
	} else{
		fmt.Println("Judul tidak ditemukan.")
	}
}