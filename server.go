package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"os"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

var db *sql.DB

type BinInfo struct {
	BinType      int
	BinLocationX float64
	BinLocationY float64
	BinEnabled   bool
	BinId        int
}

var binID int = 27

func main() {
	var postgres_user string
	postgres_user = os.Getenv("POSTGRES_USER")
	fmt.Println(postgres_user)
	var postgres_password string
	postgres_password = os.Getenv("POSTGRES_PASSWORD")
	fmt.Println(postgres_password)
	var postgres_db_name string
	postgres_db_name = os.Getenv("POSTGRES_DB_name")
	db, _ = sql.Open("postgres", "user="+postgres_user+" password="+postgres_password+" dbname="+postgres_db_name+" host=127.0.0.1 sslmode=disable")
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/setbins", setbinHandler)
	mux.HandleFunc("/getbins", getbinHandler)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500", "http://localhost:5500"},
	})
	handler := cors.Default().Handler(mux)
	handler = c.Handler(handler)
	http.ListenAndServe(":8081", handler)
}
func getbinHandler(w http.ResponseWriter, r *http.Request) {
	var Bins []BinInfo
	rows, err := db.Query("SELECT * FROM bin_data")
	checkErr(err)
	for rows.Next() {
		var bin_type int
		var bin_location_x float64
		var bin_location_y float64
		var enabled bool
		var id int
		rows.Scan(&bin_type, &bin_location_x, &bin_location_y, &enabled, &id)
		fmt.Printf("%d %f %f %t %d", bin_type, bin_location_x, bin_location_y, enabled, id)
		if enabled != false {
			Bins = append(Bins, BinInfo{BinType: bin_type, BinLocationX: bin_location_x, BinLocationY: bin_location_y, BinEnabled: enabled, BinId: id})
		}
		//w.Write([]byte(homepage_cache_html))
		//checkErr(err)
	}
	jsonInfo, _ := json.Marshal(Bins)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonInfo))

}
func setbinHandler(w http.ResponseWriter, r *http.Request) {
	binID = binID + 1
	sqlStatement := `
	INSERT INTO bin_data (bin_type,bin_location_x,bin_location_y,enabled,id)
	VALUES ($1,$2,$3,$4,$5)`
	xPos, _ := strconv.ParseFloat(r.FormValue("xPos"), 64)
	yPos, _ := strconv.ParseFloat(r.FormValue("yPos"), 64)
	binType, _ := strconv.ParseInt(r.FormValue("binType"), 10, 32)
	binEnabled := r.FormValue("binEnabled")
	enabled := false
	if binEnabled == "true" {
		enabled = true
	} else if binEnabled == "false" {
		enabled = false
	}
	_, err := db.Exec(sqlStatement, int32(binType), xPos, yPos, enabled, binID)
	if err != nil {
		panic(err)
	}
}
func checkErr(err error) {
	if err != nil {
		println(err.Error)
	}
}
