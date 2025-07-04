package main

import "fmt"

func main() {
	// program pertama saya
    fmt.Println("Hello world")

	// komentar dalam golang
	/*
		ini adalah komentar
		yang lebih panjang
		dan lebih mudah
		dibaca
	*/

	// variable dalam golang
	var nama string = "John"
	var nama2 string = "Alfaizuna"
	var umur int = 20
	var isStudent bool = true

	var lastName string
    lastName = "wick"

	fmt.Println(nama, umur, isStudent)
	fmt.Println(lastName, nama2)

	// constant dalam golang
	const firstName string = "John"
	const lastName2 string = "wick"

	fmt.Println(firstName, lastName2)

	// multiple variable dalam golang
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Println(first, second, third)

	// multiple variable dalam golang dengan tipe data yang sama
	var first2, second2, third3 int
	first2, second2, third3 = 1, 2, 3

	fmt.Println(first2, second2, third3)

	// tipe data dalam golang
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// boolean dalam golang
	var condition bool = true
	fmt.Println("is it permitted?", condition)

	// string dalam golang
	var message string = "Hello world"
	fmt.Println(message)

	// array dalam golang
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// array dalam golang dengan tipe data yang sama 
	var fruits = [4]string{"apple", "banana", "cherry", "date"}
	fmt.Println("Jumlah element \t\t:", len(fruits))
	fmt.Println("Isi semua element \t:", fruits)

	// array dalam golang dengan tipe data yang sama
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Println("Jumlah element \t\t:", len(numbers))
	fmt.Println("Isi semua element \t:", numbers)

	// perulangan dalam golang
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// perulangan dalam golang dengan array
	var fruits2 = [4]string{"apple", "banana", "cherry", "date"}
	for i := 0; i < len(fruits2); i++ {
		fmt.Println(fruits2[i])
	}

	// perulangan dalam golang dengan array
	var fruits3 = [4]string{"apple", "banana", "cherry", "date"}
	for index, fruit := range fruits3 {
		fmt.Printf("index: %d, fruit: %s\n", index, fruit)
	}

	// perulangan dalam golang dengan array dengan key dan value
	var fruits4 = [4]string{"apple", "banana", "cherry", "date"}
	for index, fruit := range fruits4 {
		fmt.Printf("index: %d, fruit: %s\n", index, fruit)
	}

	// perulangan dalam golang dengan break dan continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		if i > 8 {
			break
		}
	}

	// perulangan dalam golang bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	
	// buat segitiga dengan perulangan bersarang di dalam golang 
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// slice dalam golang
	var fruits5 = []string{"apple", "banana", "cherry", "date"}
	fmt.Println(fruits5[0])

	// slice dalam golang dengan tipe data yang sama
	var numbers2 = []int{1, 2, 3, 4}
	fmt.Println(numbers2[0])

	var fruits6 = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits6[0:2]

	fmt.Println(newFruits) // ["apple", "grape"]
	
	// A.16.4. Fungsi len() dalam golang 
	var fruits7 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits7)) // 4

	// A.16.5. Fungsi cap() dalam golang
	var fruits8 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(cap(fruits8)) // 4

	// A.16.5. Fungsi cap() dalam golang dengan slice
	var fruits9 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(cap(fruits9)) // 4

	// A.16.6. Fungsi append() dalam golang dengan slice
	var fruits10 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(append(fruits10, "orange")) // ["apple", "grape", "banana", "melon", "orange"]

	// A.16.7. Fungsi copy() dalam golang dengan slice
	var fruits11 = []string{"apple", "grape", "banana", "melon"}
	var fruits12 = []string{"orange", "pineapple", "durian"}
	copy(fruits11, fruits12)
	fmt.Println(fruits11) // ["orange", "pineapple", "durian"]
	fmt.Println(fruits12) // ["orange", "pineapple", "durian"]

	// A.16.8. Fungsi delete() dalam golang dengan map
	var person = map[string]string{
		"name": "John",
		"age": "20",
		"address": "Jakarta",
	}

	fmt.Println(person["name"])

	// Pengaksesan Elemen Slice Dengan 3 Indeks dalam golang
	var zfruits = []string{"apple", "grape", "banana"}
	var aFruits = zfruits[0:2]
	var bFruits = zfruits[0:2:2]

	fmt.Println(zfruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(zfruits)) // len: 3
	fmt.Println(cap(zfruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
	
	// map dalam golang
	var person2 = map[string]string{
		"name": "John",
		"age": "20",
		"address": "Jakarta",
	}

	fmt.Println(person2["name"])
	
	// Iterasi Item Map Menggunakan for - range dalam golang 
	var person3 = map[string]string{
		"name": "John",
		"age": "20",
		"address": "Jakarta",
	}

	for key, value := range person3 {
		fmt.Println(key, ":", value)
	}

	// Menghapus Item Map dalam golang
	var person4 = map[string]string{
		"name": "John",
		"age": "20",
		"address": "Jakarta",
	}

	delete(person4, "age")
	fmt.Println(person4)

	// Deteksi Keberadaan Item Dengan Key Tertentu dalam golang
	var person5 = map[string]string{
		"name": "John",
		"age": "20",
		"address": "Jakarta",
	}

	fmt.Println(person5["name"])
	
	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi Slice & Map dalam golang
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["name"], chicken["gender"])
	}

}