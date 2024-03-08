# Assigment 01

Description :
Buatlah sebuah service berupa CLI untuk menampilkan data teman-teman kalian dikelas. Contohnya, ketika kalian menjalankan perintah `go run biodata.go 1` maka data yang akan muncul adalah data teman kalian dengan absen no 1.
Data yang harus ditampilkan yaitu:

- Nama
- Alamat
- Pekerjaan
- Alasan memilih kelas Golang Gunakanlah struct dan function untuk menampilkan data tersebut.

Kalian bisa menggunakan `os.Args` untuk mendapatkan argument pada terminal

Send github link (repository link)

## Data

Untuk datanya sebagian menggunakan [mackaroo](https://www.mockaroo.com/). Dan namanya didapatkan dari google form absen.

## Perintah

```terminal
go run main.go -data=int
```

- data=int _(required)_

  Pilih data dari 0-54

## Example

input: `go run main.go -data=1`

output:

```terminal
Nama : FAIZAL AMRI
Alamat : 5 Katie Avenue
Pekerjaan : Editor
Alasan : Praesent id massa id nisl venenatis lacinia. Aenean sit amet justo. Morbi ut odio.
```
