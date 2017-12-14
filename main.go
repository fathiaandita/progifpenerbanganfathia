package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type datapenerbangan struct {
	Asal string
	Tujuan string
	Harga int
	Maskapai string
	Index int
}

type dataindex struct {
	Index int
	Keterangan string
}

func main() {
	port := 8080 //web service akan dijalankan pada port 8080
	http.HandleFunc("/datapenerbangan/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				s := r.URL.Query().Get("Parameter")

				if (strings.Compare("alldata", s) == 0) {
					GetDataPenerbangan(w,r) 
				} else {
					GetIndexDesc(w,r)
				}

			default:
				http.Error(w,"invalid",405)
		}
	})
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}


func GetDataPenerbangan(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/data_penerbangan")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	x:= datapenerbangan{}

	rows, err := db.Query("SELECT * FROM tabel_penerbangan")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&x.Asal, &x.Tujuan, &x.Harga, &x.Maskapai, &x.Index)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&x)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetIndexDesc(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/data_penerbangan")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	x:= dataindex{}

	rows, err := db.Query("SELECT * FROM tabel_index")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&x.Index, &x.Keterangan)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&x)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}


