package helper

type Biodata struct {
	Nama, Alamat,
	Pekerjaan, Alasan string
}

type mapOfBiodata map[int]Biodata

func (biodata Biodata) GetBiodata() string {
	return "Nama : " + biodata.Nama + "\nAlamat : " + biodata.Alamat + "\nPekerjaan : " + biodata.Pekerjaan + "\nAlasan : " + biodata.Alasan
}

var Data =  mapOfBiodata{
	1: Biodata{
		Nama:      "John",
		Alamat:    "Medan",
		Pekerjaan: "Mahasiswa baik",
		Alasan:    "Karena golang simpel",
	},
	2: Biodata{
		Nama:      "Kevin",
		Alamat:    "Medan",
		Pekerjaan: "Mahasiswa nakal",
		Alasan:    "Karena golang asik",
	},
	3: Biodata{
		Nama:      "Timothy",
		Alamat:    "Jakarta",
		Pekerjaan: "Mahasiswa rajin",
		Alasan:    "Karena golang mudah",
	},
	4: Biodata{
		Nama:      "Nathan",
		Alamat:    "Tarutung",
		Pekerjaan: "Mahasiswa",
		Alasan:    "Karena golang menyenangkan",
	},
	5: Biodata{
		Nama:      "Andre",
		Alamat:    "Porsea",
		Pekerjaan: "Mahasiswa pro",
		Alasan:    "Sudah expert",
	},
}

// var Data = map[int]Biodata{
// 	1: Biodata{
// 		Nama:      "John",
// 		Alamat:    "Medan",
// 		Pekerjaan: "Mahasiswa baik",
// 		Alasan:    "Karena golang simpel",
// 	},
// 	2: Biodata{
// 		Nama:      "Kevin",
// 		Alamat:    "Medan",
// 		Pekerjaan: "Mahasiswa nakal",
// 		Alasan:    "Karena golang asik",
// 	},
// 	3: Biodata{
// 		Nama:      "Timothy",
// 		Alamat:    "Jakarta",
// 		Pekerjaan: "Mahasiswa rajin",
// 		Alasan:    "Karena golang mudah",
// 	},
// 	4: Biodata{
// 		Nama:      "Nathan",
// 		Alamat:    "Tarutung",
// 		Pekerjaan: "Mahasiswa",
// 		Alasan:    "Karena golang menyenangkan",
// 	},
// 	5: Biodata{
// 		Nama:      "Andre",
// 		Alamat:    "Porsea",
// 		Pekerjaan: "Mahasiswa pro",
// 		Alasan:    "Sudah expert",
// 	},
// }
