package main

import (
    "fmt"
    "os"
)

// Membuat  struct
type Teman struct {
    Nama        string
    Alamat      string
    Pekerjaan  string
	Alasan       string
}

func main() {
    // Membuat slice of struct untuk menyimpan data absen teman-teman di kelas
    DataTeman := []Teman{
        {
            Nama:      "Syaiful",
            Alamat:    "Bima",
            Pekerjaan: "Mahasiswa",
			Alasan: "Karena banyak dibutuhkan untuk bahasa Backend di industri terkini",
        },
        {
            Nama:      "Dana",
            Alamat:    "Probolinggo",
            Pekerjaan: "Programmer",
			Alasan: "Karena bahasa nya mudah dipelajari dan tidak menerapkan OOP",
        },
        {
            Nama:      "Fitri",
            Alamat:    "Banyuwangi",
            Pekerjaan: "Backend Developer",
			Alasan: "Karena bahasa nya memiliki fitur garbage collector",
        },
    }

    // Mengambil argument menggunakan os.Args dari cmd/terminal
    args := os.Args[1:]

    // Jika tidak ada argument yang diberikan, tampilkan pesan error dan hentikan program
    if len(args) == 0 {
        fmt.Println("Error: harus diberikan argumen berupa No. Absen")
        os.Exit(1)
    }

    // Konversi argument ke dalam bentuk integer
    nomorAbsen := args[0]
    var index int
    fmt.Sscanf(nomorAbsen, "%d", &index)

    // Jika absen di luar range data teman, tampilkan pesan error dan hentikan program
    if index < 1 || index > len(DataTeman) {
        fmt.Println("Error: Absen tidak valid")
        os.Exit(1)
    }

    // Tampilkan data teman dengan absen yang dicari dan diketik dengan os.Args
    t := DataTeman[index-1]
    fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan Memilih Golang: %s\n", t.Nama, t.Alamat, t.Pekerjaan, t.Alasan)
}