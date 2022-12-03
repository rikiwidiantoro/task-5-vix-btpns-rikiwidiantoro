# task-5-vix-btpns-rikiwidiantoro
task akhir VIX Full Stack Developer di BTPN Syariah batch November 2022

## Deskripsi
Berdasarkan data yang telah diolah oleh tim Data Analysts, bahwa untuk meningkatkan engagement user pada aplikasi m-banking adalah meningkatkan aspek memiliki user pada aplikasi tersebut. Saran yang diberikan oleh tim data analysts adalah membentuk fitur personalize user, salah satunya adalah memungkinkan user dapat mengupload gambar untuk dijadikan foto profilnya. Tim developer bertanggung jawab untuk mengembangkan fitur ini, dan kalian diberikan tugas untuk merancang API pada fitur upload, dan menghapus gambar.


### Buatlah API menggunakan bahasa GoLang yang sesuai dengan ketentuan dan kebutuhan diatas!
Pada task akhir VIX Full Stack Developer ini kalian diarahkan untuk membentuk API berdasarkan kasus yang telah diberikan. Pada kasus ini, kalian diinstruksikan untuk membuat API untuk mengupload dan menghapus gambar. API yang kalian bentuk adalah POST, GET, PUT, dan DELETE.

A. Ketentuan API :

Pada bagian User Endpoint :
  1. POST : /users/register, dan gunakan atribut berikut ini
      - ID (primary key, required)
      - Username (required)
      - Email (unique & required)
      - Password (required & minlength 6)
      - Relasi dengan model Photo (Gunakan constraint cascade)
      - Created At (timestamp)
      - Updated At (timestamp)
  2. GET: /users/login
      - Using email & password (required)
  3. PUT : /users/:userId (Update User)
  4. DELETE : /users/:userId (Delete User)

Photos Endpoint :
  1. POST : /photos
    - ID
    - Title
    - Caption
    - PhotoUrl
    - UserID
    - Relasi dengan model User
  2. GET : /photos
  3. PUT : /photoId
  4. DELETE : /:photoId


B. Tools yang dapat kalian gunakan :
  - Gin Gonic Framework : https://github.com/gin-gonic/gin
  - Gorm : https://gorm.io/index.html
  - JWT Go : https://github.com/dgrijalva/jwt-go
  - Go Validator : http://github.com/asaskevich/govalidator

Untuk database, gunakanlah server SQL open source seperti MySQL, PostgreSQL, atau Microsoft SQL Server.


## Hasil
Dapat membuat Restfull API untuk photo menggunakan bahasa GoLang, yaitu dapat menambah(POST), mengambil(GET) single maupun all photo, mengubah(PUT) dan menghapus(DELETE) photo sesuai Endpoint yang ada di requirement diatas.

### Fitur yang belum selesai :
1. Fitur User & Authentication
2. karena tidak menggunakan user maka field UserID dan Relasi model user tidak digunakan pada endpoint POST photo
3. semua code di jalankan pada satu file yaitu main.go dan tidak menggunakan struktur folder pada requirement karena hanya dapat membuat Restfull API untuk photo saja.

### Dependensi yang digunakan :
1. Gorilla mux: For creating routes and HTTP handlers. link : go get github.com/gorilla/mux
2. gorm: An ORM tool for MySQL. link : go get github.com/jinzhu/gorm
3. mysql: The MySQL driver. link : go get github.com/go-sql-driver/mysql

### Tools yang digunakan :
- Bahasa Pemrograman GoLang & MySQL
- XAMPP
- Visual Studio Code
- Postman

### Referensi Belajar
- https://www.youtube.com/watch?v=gmP50RUd6YA
- https://levelup.gitconnected.com/build-a-rest-api-using-go-mysql-gorm-and-mux-a02e9a2865ee

