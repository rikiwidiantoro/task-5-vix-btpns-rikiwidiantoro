package main

import (
	// "fmt"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm" // untuk migrate database
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/shopspring/decimal"
)

var db *gorm.DB
var err error


//  membuat tabel product
type Photo struct {
	ID		int							`json:"id"`
	Title	string					`json:"title"`
	Caption	string				`json:"caption"`
	PhotoUrl	string			`json:"photourl"`
}

type Result struct {
	Code		int					`json:"code"`
	Data		interface{}	`json:"data"`
	Message	string			`json:"message"`
}

func main() {
    // Please define your username and password for MySQL.
    db, err = gorm.Open("mysql", "root:@/rest_photo?charset=utf8&parseTime=True")
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function

    if err != nil {
    	log.Println("Koneksi Gagal! error: ", err)
    } else { 
    	log.Println("Koneksi Berhasil!")
    }

		db.AutoMigrate(&Photo{})

		handleRequests()
}


// routing
func handleRequests() {
	log.Println("Start the development server http:localhost:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/photos", createPhotos).Methods("POST")
	myRouter.HandleFunc("/photos", getPhotos).Methods("GET")
	myRouter.HandleFunc("/{id}", getPhoto).Methods("GET")
	myRouter.HandleFunc("/{id}", updatePhoto).Methods("PUT")
	myRouter.HandleFunc("/{id}", deletePhoto).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}


// tambah photo
func createPhotos(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Create Products")

	// bikin variabel untuk menangkap payload
	payloads, _ := ioutil.ReadAll(r.Body)

	// bikin variabel struct photo
	var photo Photo

	// casting ke stuck photo
	json.Unmarshal(payloads, &photo)

	// data yg sudah ditangkap lalu di inputkan ke tabel photos
	db.Create(&photo)

	// result
	res:= Result{Code: 200, Data: photo, Message: "Berhasil menambahkan photo!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


// menampilkan semua photo
func getPhotos(w http.ResponseWriter, r *http.Request) {
	// bikin variabel
	photos := []Photo{}

	// mengambil data
	db.Find(&photos)

	// struct result
	res := Result{Code: 200, Data: photos, Message: "Berhasil menampilkan semua photo!"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}


// menampilkan single photo
func getPhoto(w http.ResponseWriter, r *http.Request) {
	// mengambil single photo
	vars := mux.Vars(r)
	photoID := vars["id"]

	// bikin variabel
	var photo Photo
	db.First(&photo, photoID)

	// struct result
	res := Result{Code: 200, Data: photo, Message: "Berhasil menampilkan single photo!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


// mengedit photo
func updatePhoto(w http.ResponseWriter, r *http.Request) {
	// mengambil single photo
	vars := mux.Vars(r)
	photoID := vars["id"]

	// bikin variabel untuk menangkap payload
	payloads, _ := ioutil.ReadAll(r.Body)

	// bikin variabel struct photo
	var photoUpdates Photo

	// casting ke stuck photo
	json.Unmarshal(payloads, &photoUpdates)

	// mengambil eksisting photo
	var photo Photo
	db.First(&photo, photoID)

	// update ke database
	db.Model(&photo).Update(photoUpdates)

	// result
	res:= Result{Code: 200, Data: photo, Message: "Berhasil mengubah photo!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


// menghapus photo
func deletePhoto(w http.ResponseWriter, r *http.Request) {
	// mengambil single photo
	vars := mux.Vars(r)
	photoID := vars["id"]

	var photo Photo

	db.First(&photo, photoID)
	db.Delete(&photo)

	// struct result
	res := Result{Code: 200, Message: "Berhasil menghapus photo!"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}