package main

import (
	"assignment01/model"
	"encoding/csv"
	"flag"
	"log"
	"os"
)

func createPersonList(data [][]string) []model.Person {
	var persons []model.Person
	for i, line := range data {
		if i > 0 {
			var rec model.Person
			for j, field := range line {
				switch j {
				case 0:
					rec.Name = field
				case 1:
					rec.Alamat = field
				case 2:
					rec.Pekerjaan = field
				case 3:
					rec.Alasan = field
				}
			}
			persons = append(persons, rec)
		}
	}
	return persons
}

func readCSV() []model.Person {
	f, err := os.Open("data/person.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	PersonList := createPersonList(data)

	return PersonList
}

func main() {

	data := readCSV()

	var selectedPerson = flag.Int("data", 0, "Pilih data dari 0-56")

	result := ""

	flag.Parse()
	result += "Nama : " + data[*selectedPerson].Name + "\n" + "Alamat : " + data[*selectedPerson].Alamat + "\n" + "Pekerjaan : " + data[*selectedPerson].Pekerjaan + "\n" + "Alasan : " + data[*selectedPerson].Alasan + "\n"
	print(result)

}
