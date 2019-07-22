package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Info struct {
	Region string `json:"region_name"`
	City   string `json:"city"`
}

// Echo headers and info
type Echo struct {
	Headers http.Header `json:"headers"`
	IP      string      `json:"ip"`
	Info    Info        `json:"info"`
}

var apiURL = "api.ipstack.com"
var APIKEY = "xxx"

func main() {
	startHTTPServer()
}

func startHTTPServer() {
	setLogging()
	fmt.Printf("Starting server,  log file: info.log")

	var PORT string
	var HOST string

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "80"
	}

	if HOST = os.Getenv("HOST"); HOST == "" {
		HOST = ""
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.RemoteAddr, r.Method, r.URL)

		s := strings.Split(r.RemoteAddr, ":")

		info := getUserLocation(s[0])
		log.Printf("[%s] Local: %v", r.RemoteAddr, info)

		response := Echo{
			Headers: r.Header,
			IP:      r.RemoteAddr,
			Info:    info,
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "unable to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Connection", "close")
		w.Write(responseBytes)
	})

	log.Printf("Starting Server %s", fmt.Sprintf("%s:%s", HOST, PORT))

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), nil)
	if err != nil {
		panic(err)
	}
}

func setLogging() {
	logfile := "info.log"
	lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

	if err != nil {
		log.Fatal("OpenLogfile: os.OpenFile:", err)
	}

	log.SetOutput(lf)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func getUserLocation(ip string) Info {
	info := Info{}
	url := fmt.Sprintf("http://%s/%s?access_key=%s&format=1", apiURL, ip, APIKEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error getting info from %s", url)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading info from %v", resp)
	}
	err = json.Unmarshal(respBody, &info)
	return info
}
