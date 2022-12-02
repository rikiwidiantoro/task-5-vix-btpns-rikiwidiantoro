# task-5-vix-btpns-rikiwidiantoro
task akhir VIX Full Stack Developer di BTPN Syariah

## Deskripsi

Berdasarkan data yang telah diolah oleh tim Data Analysts, bahwa untuk meningkatkan engagement user pada aplikasi m-banking adalah meningkatkan aspek memiliki user pada aplikasi tersebut. Saran yang diberikan oleh tim data analysts adalah membentuk fitur personalize user, salah satunya adalah memungkinkan user dapat mengupload gambar untuk dijadikan foto profilnya. Tim developer bertanggung jawab untuk mengembangkan fitur ini, dan kalian diberikan tugas untuk merancang API pada fitur upload, dan menghapus gambar. Beberapa ketentuannya antara lain :


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

Photos Endpoint
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


B. Tools :
Tools yang dapat kalian gunakan :

● Gin Gonic Framework : https://github.com/gin-gonic/gin
● Gorm : https://gorm.io/index.html
● JWT Go : https://github.com/dgrijalva/jwt-go
● Go Validator : http://github.com/asaskevich/govalidator

Untuk database, gunakanlah server SQL open source seperti MySQL, PostgreSQL, atau Microsoft SQL Server.
