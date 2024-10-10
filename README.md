# MODUL DUA
 ## SOAL 2A Jawaban nomer 1
   Program di atas adalah program Perpindahan Nilai String dalam Tiga Variabel. Fungsi utama program tersebut adalah memindahkan nilai dari tiga variabel string satu ke 
   variabel lainnya dan mencetak output sebelum dan setelah perpindahan nilai tersebut. Yang dimana Pada Program diatas pada nilai satu sementara disimpan dalam variabel temp, nilai dua di pindahkan ke satu, nilai tiga dipindahkan ke dua, dan nilai yang terdapat pada temp di pindah ke tiga
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Program mendeklarasikan tiga variabel string ('satu', 'dua', 'tiga') yang akan menyimpan input dari pengguna, serta satu variabel tambahan 'temp' untuk membantu proses perpindahan nilai di antara variabel.
      
   ## Code Explanation
      ```go
               var (
                  satu, dua, tiga string
                  temp string
               )
      ```
   kode diatas adalah sebuah deklarasi 3 variabel string yang digunakan untuk menyimpan inputan dan satu variabel untuk menyimpan data sementara

   ```go
      fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
   ```
   kode diatas adalah kode yang mengoutputkan inputan awal sebelum dieksekusi

   ```go
      temp = satu 
   	satu = dua 
   	dua = tiga 
   	tiga = temp 
   ```
   kode diatas adalah kode yang mengeksekusi inputan nilai satu sementara disimpan dalam variabel temp, nilai dua di pindahkan ke satu, nilai tiga dipindahkan ke dua, dan nilai yang terdapat pada temp di pindah ke tiga

   ```go
      fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
   ```
   kode diatas adalah kode yang mengoutputkan inputan akhir setelah di eksekusi
      
## SOAL 2A Jawaban nomer 2
   Program di atas adalah Program yang bertujuan untuk menentukan apakah input berupa tahun merupakan tahun kabisat atau bukan. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Program meminta pengguna untuk memasukkan tahun ( 'tahun' )
      - Memeriksa apakah tahun tersebut habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100.
      - Mencetak 'true' jika tahunnya adalah tahun kabisat ( 'kabisat' ), jika tidak 'false'.

   ## Code Explanation
      ```go
         var tahun int
      	var kabisat bool
      ```
   kode diatas adalah sebuah deklarasi 1 variabel string (tahun)yang digunakan untuk menyimpan inputan dan satu variabel (kabisat) untuk menentukan data yang diinputkan itu berstatus true atau false

      ```go
         var tahun int
      	var kabisat bool
      ```
   kode diatas adalah sebuah deklarasi 1 variabel string (tahun)yang digunakan untuk menyimpan inputan dan satu variabel (kabisat) untuk menentukan data yang diinputkan itu berstatus true atau false
