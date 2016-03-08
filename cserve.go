package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func lacc(r *http.Request) {
	fmt.Println(r)
}
func outcmd(c string, w http.ResponseWriter) {
	out, err := exec.Command(c).Output()
	if err != nil {
		io.WriteString(w, "AAAAaaaah")
	}
	io.WriteString(w, string(out))

}
func cam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Expires", time.Now().Format(http.TimeFormat))
	w.Header().Set("Cache-Control=", "no-cache")
	out, err := exec.Command("v4l2grab", "-q", "35", "-W", "120", "-H", "140", "-o", "./o.jpg").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	out, err = ioutil.ReadFile("./o.jpg")
	io.WriteString(w, string(out))
	if err != nil {
		io.WriteString(w, "AAAAaaaah")
	}
	lacc(r)

}
func last(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "====================\nBattery Status\n=================\n")
	outcmd("acpi", w)
	io.WriteString(w, "====================\nIinterfaces\n=================\n")
	outcmd("ifconfig", w)
	io.WriteString(w, "====================\nWireless Iinterfaces\n=================\n")
	outcmd("iwconfig", w)
	io.WriteString(w, "====================\nLast Logins\n=================\n")
	outcmd("last", w)
	lacc(r)
}
func notfound(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		w.WriteHeader(404)
		http.ServeFile(w, r, "./e.html")
		lacc(r)
	} else {
		index(w, r)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
	lacc(r)
}
func index2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/var/www"+r.URL.Path)
	if r.RequestURI == "/" {
		io.WriteString(w, "<pre>\n")
		http.ServeFile(w, r, "/etc/issue.net")
	}
	fmt.Print("/var/www/" + r.URL.Path)
	lacc(r)
}
func tailable(w http.ResponseWriter, r *http.Request) {
	fmt.Println("begin stream")
	w.Header().Set("Content-Type", "audio/x-raw")
	w.Header().Set("Expires", time.Now().Format(http.TimeFormat))
	w.Header().Set("Cache-Control=", "no-cache")
	out, err := exec.Command("tail", "-f", "/var/www/source.wav").Output()
	//TODO: this must return  a reader and not a string in order to stream
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(w, string(out))
	if err != nil {
		io.WriteString(w, "AAAAaaaah")
	}
	lacc(r)
}
func lt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Did you mean <a href=\"http://laptop.swiley.net:9980/\">laptop.swiley.net:9980</a>?") //TODO:autmatic temporary redirect
	lacc(r)
}
func main() {
	http.HandleFunc("/stats/", last)
	http.HandleFunc("/cam/", cam)
	http.HandleFunc("/stream.wav", tailable)
	http.HandleFunc("/", index2)
	http.HandleFunc("laptop.swiley.net/", lt)
	http.ListenAndServe(":80", nil)
}
