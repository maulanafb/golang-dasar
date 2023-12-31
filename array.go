package main

import "fmt"

func main() {

	// Di Go (Golang), setelah array dideklarasikan dan diinisialisasi dengan suatu ukuran tertentu, ukuran tersebut tidak dapat diubah. Artinya, Anda tidak dapat menambah atau mengurangi elemen array setelah array tersebut dibuat. Array di Go memiliki ukuran yang tetap dan tidak dapat diubah setelah dideklarasikan.
	var names [3]string
	names[0] = "Maulana"
	names[1] = "Fatih"
	names[2] = "Bilqisthi"
	// Jika Anda memerlukan koleksi data yang dapat tumbuh atau menyusut, lebih disarankan untuk menggunakan slice. Slice adalah tipe data dinamis yang dibangun di atas array dan dapat diubah ukurannya sesuai kebutuhan.
	for _, name := range names {
		fmt.Print(name + " ")
	}
	fmt.Println()

	var values = [3]int{1, 2, 3}

	// hanya bisa jika langsung di tentukan datanya
	var takDitentukan = [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	values[2] = 55
	fmt.Println(values)
	fmt.Println(len(values))

	fmt.Println(takDitentukan)
}
