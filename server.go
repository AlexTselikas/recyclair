package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func main() {
	db, _ = sql.Open("postgres", "user=####### password=###### dbname=##### host=127.0.0.1 port=5432 sslmode=disable")
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/setbins", setbinHandler)
	mux.HandleFunc("/getbins", getbinHandler)
	mux.HandleFunc("/deletebin", deleteBinHandler)
	//http.HandleFunc("/setbins", setbinHandler)
	http.HandleFunc("/test", test)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500", "http://localhost:5500"},
	})
	handler := cors.Default().Handler(mux)
	handler = c.Handler(handler)
	http.ListenAndServe(":8081", handler)
}
func getbinHandler(w http.ResponseWriter, r *http.Request) {
	var Bins []BinInfo
	swLat, err1 := strconv.ParseFloat(r.FormValue("sw_lat"), 64)
	swLon, err2 := strconv.ParseFloat(r.FormValue("sw_lon"), 64)
	neLat, err3 := strconv.ParseFloat(r.FormValue("ne_lat"), 64)
	neLon, err4 := strconv.ParseFloat(r.FormValue("ne_lon"), 64)
	if swLat == 0 || swLon == 0 || neLat == 0 || neLon == 0 || err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sqlStatement := "SELECT * FROM bin_data WHERE bin_location_x >= $1 AND bin_location_y >= $2 AND bin_location_x <= $3 AND bin_location_y <= $4"
	rows, err := db.Query(sqlStatement, swLat, swLon, neLat, neLon)
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
	sqlStatement := `SELECT exists (SELECT 1 FROM bin_data WHERE bin_location_x = $1 AND bin_location_y=$2 LIMIT 1);`
	var result bool
	xPos, _ := strconv.ParseFloat(r.FormValue("xPos"), 64)
	yPos, _ := strconv.ParseFloat(r.FormValue("yPos"), 64)
	row := db.QueryRow(sqlStatement, xPos, yPos)
	err := row.Scan(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	if result == true {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	sqlStatement = `
	INSERT INTO bin_data (bin_type,bin_location_x,bin_location_y,enabled)
	VALUES ($1,$2,$3,$4)`
	binType, _ := strconv.ParseInt(r.FormValue("binType"), 10, 32)
	binEnabled := r.FormValue("binEnabled")
	enabled := false
	if binEnabled == "true" {
		enabled = true
	} else if binEnabled == "false" {
		enabled = false
	}
	_, err = db.Exec(sqlStatement, int32(binType), xPos, yPos, enabled)
	if err != nil {
		panic(err)
	}
}
func deleteBinHandler(w http.ResponseWriter, r *http.Request) {
	binId, _ := strconv.ParseInt(r.FormValue("id"), 10, 32)
	sqlStatement := `UPDATE bin_data SET enabled = FALSE WHERE id = $1`
	_, err := db.Exec(sqlStatement, binId)
	if err != nil {
		checkErr(err)
	}
}
func test(w http.ResponseWriter, r *http.Request) {
	xPos, _ := strconv.ParseFloat(r.FormValue("xPos"), 64)
	yPos, _ := strconv.ParseFloat(r.FormValue("yPos"), 64)
	binType, _ := strconv.ParseInt(r.FormValue("binType"), 10, 32)
	binEnabled := r.FormValue("binEnabled")
	fmt.Println(xPos)
	fmt.Println(yPos)
	fmt.Println(binType)
	fmt.Println(binEnabled)
	w.Write([]byte("y position:" + r.FormValue("yPos")))

}
func checkErr(err error) {
	if err != nil {
		println(err.Error)
	}
}
