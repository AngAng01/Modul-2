# MODUL DUA
 ## SOAL 2A nomer 1
   Program di atas adalah program Perpindahan Nilai String dalam Tiga Variabel. Fungsi utama program tersebut adalah memindahkan nilai dari tiga variabel string satu ke variabel lainnya dan mencetak output sebelum dan setelah perpindahan nilai tersebut. Yang dimana Pada Program diatas pada nilai satu sementara disimpan dalam variabel temp, nilai dua di pindahkan ke satu, nilai tiga dipindahkan ke dua, dan nilai yang terdapat pada temp di pindah ke tiga
   
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

      
## SOAL 2A nomer 2
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
   kode diatas adalah sebuah deklarasi 1 variabel string (tahun) yang digunakan untuk menyimpan inputan dan satu variabel (kabisat) untuk menentukan data yang diinputkan itu berstatus true atau false

      ```go
         	fmt.Print("Tahun : ")
	         fmt.Scan(&tahun)
      ```
   kode diatas adalah kode yang digunakan untuk menginput sebuah inputan tahun. 

      ```go
         	if tahun % 400 == 0  || tahun % 4 == 0 && tahun % 100 != 0  {
           		kabisat = true
           		fmt.Print("Kabisat : ", kabisat)
         	} else {
           		kabisat = false
           		fmt.Print("Kabisat : ", kabisat)
         	}	
      ```
   kode diatas adalah kode untuk suatu kondisi tertentu 'if else', seperti diatas kode tersebut akan menginputkan kabisat = true apabila inputan tahun habis dibagi 400 atau habis dibagi 4 dan tidak habis dibagi 100.

   ## SOAL 2A nomer 3
   Program di atas bertujuan untuk menghitung volume dan luas permukaan sebuah bola, dengan menerima input berupa jari-jari bola dari pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Program meminta pengguna untuk memasukkan nilai jejari bola ('r').
      - Menghitung volume bola menggunakan rumus '4/3 * π * r * r * r'
      - Menghitung luas permukaan bola menggunakan rumus '4 * π * r * r'
      - Mencetak hasil perhitungan 'volume' dan 'luas' permukaan bola.

   ## Code Explanation
      ```go
         var r, phi, volume, luas  float64
      ```
   kode diatas adalah sebuah deklarasi 4 variabel float64 yang digunakan untuk r = jari-jari bola, phi = konstanta matematika yang bernilai sekitar 3.1415926535, volume = Deklarasi rumus volume bola, luas = Deklarasi rumus luas permukaan bola.

      ```go
         	fmt.Print("Jejari : ")
	         fmt.Scan(&r)
      ```
   kode diatas adalah kode yang digunakan untuk menginput sebuah inputan jari-jari dari suatu bola. 

      ```go
         	phi = 3.1415926535
         	volume = 4.0 / 3.0 * phi * r * r * r
         	luas = 4 * phi * r * r 
      ```
   kode diatas adalah kode untuk Menginisialisasi nilai phi, Menghitung volume bola menggunakan rumus, dan Menghitung luas permukaan bola menggunakan rumus.

      ```go
         	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
      ```
   kode diatas adalah kode yang digunakan untuk mengouputkan hasil dari volume bola dan luas permukaan bola


   ## SOAL 2A nomer 4
   Program di atas adalah program konversi suhu dari Celsius ke tiga skala suhu lainnya, yaitu Reamur, Fahrenheit, dan Kelvin.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Mendeklarasikan variabel 'R', 'F', 'K', dan suhu dengan tipe data float64 untuk menyimpan hasil konversi dan input suhu.
      - Program meminta pengguna untuk memasukkan temperatur dalam derajat Celsius melalui perintah fmt.Print dan menyimpan nilai tersebut dalam variabel suhu.
      - Menghitung derajat Reamur dengan rumus 'suhu * ( 4 / 5 )'
      - Menghitung derajat Fahrenheit dengan rumus '( 9 / 5 ) * suhu + 32'
      - Menghitung derajat Kelvin dengan rumus 'suhu + 273'
      - Program mencetak hasil konversi suhu dalam derajat Reamur, Fahrenheit, dan Kelvin dengan format dua angka di belakang koma menggunakan fmt.Printf.

   ## Code Explanation
      ```go
         var R, F, K, suhu float64
      ```
   kode diatas adalah sebuah deklarasi 4 variabel float64 yang digunakan untuk R = Reamur, F = Fahrenheit, K = Kelvin, dan suhu = suhu Celcius.

      ```go
         	fmt.Print("Temperatur Celcius : ")
	         fmt.Scan(&suhu)
      ```
   kode diatas adalah kode yang digunakan untuk menginput sebuah suhu dengan satuan Celcius. 

      ```go
         	R = suhu * ( 4.00 / 5.00 )
         	F = ( 9.00 / 5.00 ) * suhu + 32
         	K = suhu + 273
      ```
   kode diatas adalah kode untuk Menghitung konversi Celcius ke Reamur, Menghitung konversi Celcius ke Fahrenheit, dan Menghitung konversi Celcius ke Kelvin.

      ```go
         	fmt.Printf("Derajat Reamur : %.2f\n", R)
         	fmt.Printf("TDerajat Fahrenheit : %.2f\n", F)
         	fmt.Printf("Derajat Kelvin : %.2f\n", K)
      ```
   kode diatas adalah kode yang digunakan untuk mengouputkan hasil dari konversi suhu Celcius ke tiga skala suhu lainnya, yaitu Reamur, Fahrenheit, dan Kelvin.


   ## SOAL 2A nomer 5
   Program di atas adalah program yang menerima input berupa angka dan karakter, lalu menampilkan karakter yang sesuai dengan nilai ASCII dari angka tersebut.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi variabel 'int1', 'int2', 'int3', 'int4', 'int5' bertipe int untuk menyimpan 5 nilai integer yang dimasukkan oleh pengguna.
      - Deklarasi variabel 'char1', 'char2', 'char3' bertipe rune untuk menyimpan 3 karakter yang akan dimasukkan oleh pengguna.
      - Program mengonversi nilai integer menjadi karakter ASCII yang sesuai dan mencetaknya ke layar menggunakan 'fmt.Printf'.
      - Program meminta pengguna untuk memasukkan 3 karakter dengan menggunakan 'fmt.Scanf' untuk membacanya secara langsung.

   ## Code Explanation
      ```go
         var int1, int2, int3, int4, int5 int
      ```
   kode diatas adalah sebuah deklarasi 5 variabel int yang digunakan untuk menginputkan sebuah angka yang akan dieksekusi.
   
      ```go
         	var char1, char2, char3 rune
      ```
   kode diatas adalah sebuah deklarasi 3 variabel karakter Unicode yang digunakan untuk menginputkan sebuah huruf kapital yang akan dieksekusi.

      ```go
         	fmt.Println("Masukan 5 buah nilai antara 32 - 127 : ")
         	fmt.Scan(&int1, &int2, &int3, &int4, &int5)
         	fmt.Printf("%c %c %c %c %c\n", int1, int2, int3, int4, int5)
      ```
   Kode di atas adalah bagian dari program yang meminta pengguna untuk memasukkan lima nilai integer dalam rentang 32 hingga 127. Nilai-nilai ini mewakili kode ASCII untuk karakter yang dapat dicetak. Setelah menerima input, program akan mengonversi masing-masing nilai integer menjadi karakter ASCII yang sesuai dan mencetaknya ke layar. 

      ```go
         	fmt.Scanln()
      ```
   kode diatas adalah kode yang digunakan untuk membaca input dari pengguna hingga karakter newline (tekanan Enter). Fungsi ini biasanya digunakan untuk menunggu input tambahan dari pengguna sebelum melanjutkan eksekusi program.

      ```go
         	fmt.Println("Masukkan 3 karakter:")
         	fmt.Scanf("%c%c%c",&char1,&char2,&char3) 
         	fmt.Printf("%c%c%c\n", char1, char2, char3)
      ```
   Kode di atas meminta pengguna untuk memasukkan tiga karakter setelah mencetak pesan "Masukkan 3 karakter:". Fungsi fmt.Scanf("%c%c%c", &char1, &char2, &char3) membaca ketiga karakter dan menyimpannya dalam variabel char1, char2, dan char3. Kemudian, fmt.Printf("%c%c%c\n", char1, char2, char3) mencetak ketiga karakter tersebut ke layar.



   ## SOAL 2B nomer 1
   Program di atas adalah program yang melakukan beberapa percobaan untuk memeriksa input warna yang dimasukkan oleh pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Variabel keberhasilan bertipe 'boolean' yang digunakan untuk menentukan apakah semua percobaan input warna berhasil.
      - Menggunakan loop 'for' yang berjalan sebanyak 5 kali untuk meminta pengguna memasukkan 4 warna dalam setiap percobaan.
      - Menggunakan 'fmt.Scanln' untuk membaca input 4 warna dari pengguna dan menyimpannya dalam variabel 'warna1', 'warna2', 'warna3', dan 'warna4'.
      - Memeriksa apakah keempat warna yang dimasukkan sesuai dengan urutan yang benar: "merah", "kuning", "hijau", dan "ungu". Jika tidak, variabel keberhasilan diatur menjadi 'false'.
      - Setelah semua percobaan, program mencetak hasil 'keberhasilan' berdasarkan nilai variabel keberhasilan. Jika semua percobaan berhasil, mencetak "KEBERHASILAN: True"; jika tidak, mencetak "KEBERHASILAN: False".

   ## Code Explanation
      ```go
         keberhasilan := true 
      ```
   kode di atas adalah bagian dari program yang mendeklarasikan dan menginisialisasi variabel bernama keberhasilan. 

      ```go
         	var warna1, warna2, warna3, warna4 string
      ```
   kode diatas adalah sebuah deklarasi 4 variabel string yang digunakan untuk menginputkan sebuah warna yang akan dicek.

      ```go
         	for i := 1; i <= 5; i++ { 
          		fmt.Print("Percobaan ", i, " : ")
          		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)
        	 }
      ```
   Kode di atas adalah sebuah loop yang berjalan sebanyak 5 kali, di mana setiap iterasi mencetak pesan "Percobaan i :" untuk menunjukkan nomor percobaan dan meminta pengguna untuk memasukkan empat warna. Input warna disimpan dalam variabel warna1, warna2, warna3, dan warna4.

      ```go
         	if !(warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu") {
			          keberhasilan = false
          }
      ```
   Kode di atas adalah pernyataan kondisi yang memeriksa apakah input warna yang dimasukkan oleh pengguna sesuai dengan urutan yang benar.

      ```go
          	if keberhasilan {
          		  fmt.Println("KEBERHASILAN: True")
          	} else {
          		  fmt.Println("KEBERHASILAN: False")
          	}
       ```
   Kode di atas adalah pernyataan kondisi yang menentukan output berdasarkan nilai variabel keberhasilan.


   ## SOAL 2B nomer 2
   Program di atas adalah program dalam bahasa Go yang meminta pengguna untuk memasukkan nama bunga secara berulang hingga pengguna memasukkan kata "selesai" lalu mencetaknya.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi Variabel :
          1. 'inputBunga' : Variabel untuk menyimpan nama bunga yang dimasukkan oleh pengguna.
          2. 'daftarBunga' : Variabel string untuk menyimpan daftar nama bunga yang telah dimasukkan.
          3. 'jumlahBunga' : Variabel integer untuk menghitung jumlah bunga yang dimasukkan.
      - Menggunakan loop 'for' yang berjalan terus menerus hingga pengguna memasukkan "selesai".
      - Memeriksa apakah 'inputBunga' sama dengan "selesai" untuk menghentikan loop.
      - Menyusun Daftar Bunga: Menambahkan nama bunga yang dimasukkan ke dalam 'daftarBunga' setiap kali pengguna memasukkan nama bunga.

   ## Code Explanation
      ```go
         var inputBunga,daftarBunga string
	        var jumlahBunga int 
      ```
   kode di atas adalah sebuah deklarasi 2 variabel string yang digunakan untuk menginputkan sebuah bunga dan menyimpan daftar bunga, dan satu variabel int yang digunakan untuk menghitung jumlah bunga yang diinput.

      ```go
         	for {
           		jumlahBunga++
           		fmt.Print("Bunga ", jumlahBunga , ": ")
           		fmt.Scan(&inputBunga)
           
           		if inputBunga == "selesai"{
           			  break
           		} 
           		daftarBunga+= " - " + inputBunga
         	}
      ```
   Kode di atas adalah bagian dari program yang digunakan untuk meminta pengguna memasukkan nama bunga secara berulang hingga pengguna memasukkan kata "selesai".
   
      ```go
         	fmt.Print("Pita :", daftarBunga)
	         fmt.Print("\nBunga : ", jumlahBunga-1)
      ```
   Kode di atas adalah bagian dari program yang digunakan untuk mencetak hasil akhir setelah pengguna selesai memasukkan nama bunga. 


   ## SOAL 2B nomer 3
   Program di atas adalah program yang digunakan untuk menghitung dan menentukan apakah sepeda motor akan "oleng" berdasarkan berat belanjaan yang dimasukkan pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi Variabel 'berat1' dan 'berat2' bertipe float64 untuk menyimpan berat belanjaan di dua kantong.
      - Deklarasi Variabel 'oleng' bertipe bool untuk menandakan apakah sepeda motor akan oleng.
      - Menggunakan loop 'for' tanpa kondisi untuk meminta input berulang kali hingga pengguna memasukkan nilai negatif.
      - Memeriksa apakah berat yang dimasukkan kurang dari 0 dan menghentikan proses jika ya.
      - Menghitung selisih antara kedua berat dan total berat belanjaan.
      - Memeriksa apakah total berat lebih dari 150 dan menghentikan proses jika ya.
      - Memeriksa selisih berat dan mengatur nilai 'oleng' berdasarkan kriteria tertentu.
      - Mencetak apakah sepeda motor pak Andi akan oleng atau tidak berdasarkan nilai 'oleng'.

   ## Code Explanation
      ```go
         var berat1, berat2 float64
	        var oleng bool
      ```
   kode di atas adalah sebuah deklarasi 2 variabel float yang digunakan untuk menginputkan sebuah berat , dan satu variabel boolean yang digunakan untuk menentukan true or falsenya suatu kondisi.

      ```go
         	for {
           		fmt.Print("\nMasukkan berat belanjaan di kedua kantong: ")
           		fmt.Scan(&berat1, &berat2)
         	}
      ```
   Bagian kode di atas merupakan bagian dari program yang meminta pengguna untuk memasukkan berat belanjaan di dua kantong secara berulang.
   
      ```go
         	if berat1 < 0 || berat2 < 0 {
         			fmt.Print("Proses Selesai")
         			break
         	}
      ```
   Kode di atas merupakan kondisi yang digunakan untuk memeriksa nilai berat belanjaan yang dimasukkan oleh pengguna. 

     ```go
          selisih := berat1 - berat2 
		        total := berat1 + berat2
     ```
   Kode di atas merupakan kode yang bertanggung jawab untuk melakukan perhitungan dasar terhadap berat belanjaan yang dimasukkan. Dengan menghitung selisih dan total, program dapat mengambil keputusan lebih lanjut mengenai stabilitas sepeda motor berdasarkan input yang diberikan oleh pengguna. 

     ```go
          if total > 150 {
          			fmt.Print("Proses Selesai.")
          			break
        		} else if selisih >= 9 || -selisih >= 9 {
        			  oleng = true
        		} else {
        			  oleng = false
        		}
     ```
   Bagian kode ini berfungsi untuk menentukan apakah proses input harus dihentikan berdasarkan total berat yang dimasukkan, dan untuk mengevaluasi apakah sepeda motor akan "oleng" berdasarkan selisih berat antara kedua kantong. 

     ```go
          fmt.Print("Sepeda motor pak Andi akan oleng: ", oleng)
     ```
   Bagian kode ini berfungsi untuk memberikan umpan balik (Output) kepada pengguna mengenai stabilitas sepeda motor pak Andi berdasarkan hasil perhitungan berat belanjaan yang telah dimasukkan sebelumnya. 


   ## SOAL 2B nomer 4
   Program di atas adalah program yang menghitung nilai dari fungsi matematika tertentu berdasarkan input yang diberikan oleh pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi variabel 'k' untuk menyimpan nilai input dari pengguna dan 'hasil' sebagai variabel untuk menyimpan hasil perhitungan, diinisialisasi dengan nilai 1.0.
      - Meminta pengguna untuk memasukkan nilai K melalui fungsi 'fmt.Print' dan 'fmt.Scan'.
      - Menggunakan loop 'for' untuk menghitung nilai 'f(K)' berdasarkan rumus yang diberikan dengan melakukan perhitungan di dalam setiap iterasi.
      - Mencetak hasil akhir dari perhitungan 'f(K)' dengan format 10 angka desimal menggunakan 'fmt.Printf'.

   ## Code Explanation
      ```go
         var k int
	        var hasil float64 = 1.0 
      ```
   kode di atas adalah sebuah deklarasi 1 variabel int yang digunakan untuk menyimpan nilai yang dimasukkan oleh pengguna , dan satu variabel float yang digunakan untuk menyimpan hasil perhitungan yang dilakukan selama iterasi loop.
   
      ```go
         	fmt.Print("Nilai K = ")
	         fmt.Scan(&k)
      ```
   Bagian kode di atas merupakan bagian dari program yang meminta pengguna untuk memasukkan nilai k
   
      ```go
         	for i := 0; i <= k; i++ {
           		term := float64((4*i + 2) * (4*i + 2)) / float64((4*i + 1) * (4*i + 3))
           		hasil *= term
         	}
      ```
   Bagian kode ini adalah sebuah loop for yang bertujuan untuk menghitung nilai fungsi secara bertahap dengan menambahkan hasil dari setiap iterasi ke variabel hasil. 

     ```go
          fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
     ```
   Kode di atas merupakan kode yang berfungsi untuk mencetak nilai dari variabel hasil ke layar dengan format tertentu. 

    
   ## SOAL 2C nomer 1
   Program di atas adalah program yang menghitung biaya pengiriman parsel berdasarkan berat dalam gram yang dimasukkan oleh pengguna. Program ini memecah berat parsel menjadi kilogram dan gram, menghitung biaya untuk setiap bagian, lalu menghitung total biaya pengiriman. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Program meminta pengguna memasukkan berat parsel dalam gram menggunakan fmt.Scan untuk menerima input.
      - Berat yang dimasukkan dibagi menjadi dua bagian kilogram ('kg') dan sisa gram ('gram') melalui operasi pembagian dan modulus.
      - Biaya pengiriman dihitung berdasarkan berat dalam kilogram dan gram, di mana biaya kilogram tetap dan biaya gram bervariasi tergantung berat parsel.
      - Program mencetak rincian biaya dan total biaya pengiriman ke layar.

   ## Code Explanation
      ```go
         var beratGram, costGram, total, costKg int
      ```
   kode di atas adalah sebuah deklarasi 4 variabel int yang digunakan untuk beratGram = Menyimpan input berat parsel dalam satuan gram, costGram = Menyimpan biaya pengiriman untuk bagian berat yang tersisa dalam gram (di bawah 1 kilogram), total = Menyimpan hasil perhitungan total biaya pengiriman (gabungan biaya untuk kilogram dan gram), costKg = Menyimpan biaya pengiriman untuk bagian berat dalam satuan kilogram.
    
      ```go
         	fmt.Print("Berat Parsel (gram) : ")
	         fmt.Scan(&beratGram)
      ```
   Bagian kode di atas merupakan bagian dari program yang meminta pengguna untuk memasukkan nilai berat parsel dalam gram.
   
      ```go
         	kg := beratGram / 1000
	         gram := beratGram % 1000
         	}
      ```
   Bagian kode ini bertujuan menghitung jumlah kilogram dengan membagi nilai beratGram dengan 1000 (karena 1 kilogram sama dengan 1000 gram). Hasilnya disimpan dalam variabel kg dan Menghitung sisa berat dalam gram dengan menggunakan operator modulus (%), yaitu menghitung sisa hasil bagi dari pembagian berat gram dengan 1000. Hasilnya disimpan dalam variabel gram.

     ```go
          fmt.Print("Detail berat : ", kg ," kg + ", gram , " gr")
     ```
   Kode di atas merupakan kode yang berfungsi untuk mencetak rincian berat parsel dalam satuan kilogram dan gram ke layar.

     ```go
          costKg = kg * 10000
     ```
   Kode di atas merupakan kode yang berfungsi untuk menghitung biaya pengiriman berdasarkan berat dalam satuan kilogram.

     ```go
          if kg < 10 {
           		if gram < 500 {
           			  costGram = gram * 15
           		} else {
           			  costGram = gram * 5
           		}
         	}
     ```
   Kode di atas merupakan kode yang berfungsi untuk menghitung biaya pengiriman berdasarkan berat sisa dalam gram ketika berat total parsel kurang dari 10 kilogram. 

     ```go
          total = costKg + costGram
         	fmt.Print("\nDetail biaya : ", "Rp. ", costKg , " + ", "Rp. ", costGram )
         	fmt.Print("\nTotal Biaya : Rp. ", total)
     ```
   Kode di atas merupakan kode yang berfungsi untuk menghitung dan menampilkan total biaya pengiriman berdasarkan rincian biaya yang telah dihitung sebelumnya. 


   ## SOAL 2C nomer 2
   Program di atas adalah program yang berfungsi untuk menentukan nilai huruf (grade) berdasarkan nilai akhir mata kuliah yang dimasukkan oleh pengguna. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi Variabel 'nam' float ini digunakan untuk menyimpan nilai akhir mata kuliah yang dimasukkan oleh pengguna dalam format angka desimal.
      - Deklarasi Variabel 'nmk' string  ini digunakan untuk menyimpan nilai huruf (grade) yang akan ditampilkan kepada pengguna.
      - Menggunakan struktur kontrol if, else if, dan else untuk menentukan nilai huruf berdasarkan nilai akhir yang dimasukkan
      - Mencetak nilai huruf yang sesuai berdasarkan input yang diberikan oleh pengguna.
      - Menerima input dari pengguna dan menyimpannya ke dalam variabel

   ## Code Explanation
      ```go
         var nam float64
	        var nmk string
      ```
   kode di atas adalah sebuah deklarasi 1 variabel float yang digunakan untuk menyimpan nilai akhir mata kuliah yang dimasukkan oleh pengguna dalam format angka desimal, dan 1 variabel string yang digunakan untuk menyimpan nilai huruf (grade) yang akan ditentukan berdasarkan nilai akhir yang diberikan.
   
      ```go
         	fmt.Print("Nilai akhir mata kuliah : ")
	         fmt.Scanln(&nam)
      ```
   Bagian kode di atas merupakan bagian dari program yang meminta pengguna untuk memasukkan nilai akhir mata kuliah.
   
      ```go
         	if nam > 80 {
         		  nmk = "A"
         	} else if nam > 72.5 {
         		  nmk = "AB"
         	} else if nam > 65 {
         		  nmk = "B"
         	} else if nam > 57.5 {
         		  nmk = "BC"
         	} else if nam > 50 {
         		  nmk = "C"
         	} else if nam > 40 {
         		  nmk = "D"
         	} else {
         		  nmk = "E"
         	}
      ```
   Bagian kode ini adalah serangkaian pernyataan kondisional yang digunakan untuk menentukan nilai huruf (grade) berdasarkan nilai akhir mata kuliah (nam) yang telah dimasukkan oleh pengguna.

     ```go
          fmt.Println("Nilai mata kuliah :", nmk)
     ```
   Kode di atas merupakan perintah untuk mencetak output ke layar. Kode ini menampilkan pesan "Nilai mata kuliah :" diikuti oleh nilai huruf (nmk) yang telah ditentukan sebelumnya berdasarkan nilai akhir mata kuliah yang dimasukkan oleh pengguna.

   ## Jawaban dari pertanyaan-pertanyaan 
   a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
      = Keluaran dari program tersebut jika diinputkan nilai 80.1 adalah A, dan eksekusi program sesuai dengan spesifikasi soal
      
   b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
      = Kesalahan terdapat pada serangkaian percabangan, Untuk nilai nam, seharusnya menggunakan else if setelah setiap kondisi agar     program tidak memberikan dua nilai pada satu nam. Misalnya, jika nam adalah 85, seharusnya hanya mencetak "A" dan bukan juga "AB" atau "B". jadi seharusnya menggunakan else-if bukan if-if.
      
   c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.
       = 
       
       ```go
       package main

        import "fmt"
        
        func main() {
        	var nam float64
        	var nmk string
        
        	fmt.Print("Nilai akhir mata kuliah : ")
        	fmt.Scanln(&nam)
        
        	if nam > 80 {
        		nmk = "A"
        	} else if nam > 72.5 {
        		nmk = "AB"
        	} else if nam > 65 {
        		nmk = "B"
        	} else if nam > 57.5 {
        		nmk = "BC"
        	} else if nam > 50 {
        		nmk = "C"
        	} else if nam > 40 {
        		nmk = "D"
        	} else {
        		nmk = "E"
        	}
        
        	fmt.Println("Nilai mata kuliah :", nmk)
        }
        ```


   ## SOAL 2C nomer 3
   Program di atas adalah program yang bertujuan untuk menentukan apakah sebuah bilangan bulat yang dimasukkan oleh pengguna adalah bilangan prima atau tidak. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Variabel 'bil' dan 'faktor' dideklarasikan untuk menyimpan bilangan yang dimasukkan pengguna serta jumlah faktor dari bilangan tersebut.
      - Program meminta pengguna untuk memasukkan bilangan yang akan diperiksa dengan menggunakan fungsi fmt.Scan.
      - Terdapat loop 'for' yang menghitung faktor dari bilangan tersebut dengan memeriksa apakah hasil bagi dari bilangan dengan angka literasi ('i') adalah 0. Jika iya, angka tersebut adalah faktor dari bilangan.
      - Setelah menghitung faktor, program menggunakan pernyataan 'if' untuk memeriksa apakah bilangan tersebut memiliki tepat dua faktor. Jika ya, bilangan tersebut dianggap prima dan mencetak hasil "Prima : true", jika tidak mencetak "Prima : false".
      - Program mencetak hasil berupa faktor-faktor dari bilangan dan apakah bilangan tersebut adalah bilangan prima atau tidak.

   ## Code Explanation
      ```go
         var bil int
	        var faktor int
      ```
   Kode di atas adalah deklarasi dua variabel, variabel bil bertipe data int digunakan untuk menyimpan bilangan bulat yang akan diperiksa apakah merupakan bilangan prima atau tidak. Variabel faktor, juga bertipe int, digunakan untuk menghitung jumlah faktor dari bilangan tersebut.
   
      ```go
         	fmt.Print("Bilangan : ")
	         fmt.Scan(&bil )
      ```
   Bagian kode di atas merupakan bagian dari program yang meminta pengguna untuk memasukkan nilai bilangan yang akan dieksekusi.
   
      ```go
         	if bil > 1{
          		for i := 1; i <= bil; i++ {
            			if bil % i == 0 {
            				fmt.Print(i, " ")
            				faktor++
            			}
          		}
         		
         	}
      ```
   Kode di atas adalah bagian dari program yang digunakan untuk menemukan dan mencetak semua faktor dari bilangan yang diberikan, serta menghitung jumlah faktor tersebut.

     ```go
          if faktor == 2 {
           		fmt.Print("\nPrima : true")
         	} else {
         		  fmt.Print("\nPrima : false")
         	}
     ```
   Kode di atas berfungsi untuk menentukan apakah bilangan yang diberikan adalah bilangan prima atau bukan, berdasarkan jumlah faktor yang dihitung sebelumnya. 



   

