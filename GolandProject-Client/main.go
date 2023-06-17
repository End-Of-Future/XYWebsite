package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	chk(err)
	if r.Method == "GET" {
		c, err := filepath.Abs(strings.TrimLeft(r.URL.Path, "/"))
		chk(err)
		s, err := readFile(c)
		if err != nil {
			s, err := readFile("404.html")
			chk(err)
			_, err = fmt.Fprintf(w, s)
			chk(err)
			return
		}
		_, err = fmt.Fprintf(w, s)
		chk(err)
	}
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("Served")
	err := http.ListenAndServe(":2048", nil)
	chk(err)
}

func readFile(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func chk(err error) {
	if err != nil {
		fmt.Println("\x1B[30mERR:\\x1B[0m" + err.Error())
	}
}
