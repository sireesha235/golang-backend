package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Servertime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	zone, offset := t.Zone()
	fmt.Println(zone, offset)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["Server Epoch Time"] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	resp["Server Time Zone"] = zone
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}

func handleRequests() {
	http.HandleFunc("/", health)
	http.HandleFunc("/servertime", Servertime)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {
	handleRequests()
}
