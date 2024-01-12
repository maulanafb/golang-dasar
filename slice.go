package main

import "fmt"

func main() {
	// ini di bawah adalah array
	names := [...]string{"eko", "Kurniawan", "Khanedi", "Joko", "budi", "nugraha"}
	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// contoh append  Child

	// ini adalah array
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")

	fmt.Println(daySlice2)
	fmt.Println(daySlice1)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Maulana"
	newSlice[1] = "Maulana"

	newSlice2 := append(newSlice, "Maulana", "Fatih", "Fatih", "Fatih", "Fatih")
	newSlice2[0] = "Budi"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// Slice akan lebih sering digunakan
}
