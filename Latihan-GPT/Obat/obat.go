package main

import "fmt"

const NMAX int = 100

type tabObat [NMAX] string

func main(){
	var DataObat tabObat
	var nDataObat, option int
	
	fmt.Print("Masukan Jumlah obat : ")
	fmt.Scan(&nDataObat)
	inputData(&DataObat, nDataObat)
	
	for option!=4{
		menu()
		fmt.Print("Masukan Opsi : ")
		fmt.Scan(&option)
		switch option{
		case 1: deleteData(&DataObat, &nDataObat)
		case 2: rename(&DataObat, nDataObat)
		case 3: listData(DataObat, nDataObat)
		default : break
		}
	}
}

func menu(){
	fmt.Println("\t ===Menu===")
	fmt.Println("1 - Delete Tanaman Obat")
	fmt.Println("2 - Rename Tanaman Obat")
	fmt.Println("3 - List Tanaman Obat")
	fmt.Println("4 - Exit")
}

func inputData(A *tabObat, n int){
	for i:=0;i<n;i++{
		fmt.Scan(&A[i])
	}
}

func listData(A tabObat, n int){
	for i:=0;i<n;i++{
		fmt.Printf("%s ", A[i])
	}
	fmt.Println()
}

func sequentialSearch(A tabObat, n int, x string) int{
	var cari, i int
	cari = -1
	i = 0
	for cari==-1&&i<n{
		if A[i] == x{
			cari = i
		}
		i++
	}	
	return cari
}

func deleteData(A *tabObat, n*int){
	var cari, i int
	var x string

	fmt.Print("Masukan Nama tanaman obat yang ingin di delete : ")
	fmt.Scan(&x)
	cari = sequentialSearch(*A, *n, x)
	if cari != -1{
		i = cari
		for i<*n-1{
			A[i] = A[i+1]
			i++
		}
		*n--
	} else {
		fmt.Println("Tanaman obat tidak ditemukan")
	}
}

func rename(A *tabObat, n int){
	var cari int
	var namaLama, namaBaru string

	fmt.Print("Masukan Nama Obat yang ingin diganti : ")
	fmt.Scan(&namaLama)
	cari = sequentialSearch(*A, n, namaLama)

	if cari !=-1{
		fmt.Print("Masukan Nama Obat Baru : ")
		fmt.Scan(&namaBaru)
		A[cari] = namaBaru
	} else{
		fmt.Println("Nama Obat tidak ditemukan")
	}
}
