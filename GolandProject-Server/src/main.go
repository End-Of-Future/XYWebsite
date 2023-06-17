package main

import (
	"GolandProject/src/modules/ConfigHelper"
	"GolandProject/src/modules/Log"
	"fmt"
	"net/http"
)

func Items(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	chk(err)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	Log.Success("CONNECTED")
	c := ConfigHelper.NewConf("./src/dbs/strings.config")
	s := r.FormValue("mode")
	switch s {
	case "title":
		Log.Send("Parameters following:")
		for k, v := range r.Form {
			Log.Send("KEY:", k, " VALUE:", v)
		}
		_, err = fmt.Fprintf(w, c.FindQuery("title")) //标题
		chk(err)
	case "t":
		fmt.Println("Parameters following:")
		for k, v := range r.Form {
			fmt.Println("KEY:", k, " VALUE:", v)
		}
		_, err = fmt.Fprintf(w, c.FindQuery("top")) //导航栏
		chk(err)
	case "l":
		fmt.Println("Parameters following:")
		for k, v := range r.Form {
			fmt.Println("KEY:", k, " VALUE:", v)
		}
		_, err := fmt.Fprintf(w, c.FindQuery("l-"+r.FormValue("path"))) //左目录
		chk(err)
	default:
		_, err := fmt.Fprintf(w, "Unknown")
		chk(err)
	}
}

func main() {
	http.HandleFunc("/items", Items)
	fmt.Println("Served")
	err := http.ListenAndServe(":2047", nil)
	chk(err)
}
func GetItem(path string) string {
	switch path {
	case "index":
		return "项A 项B 项C 项D"
	default:
		return "Unknown request"
	}
}

func chk(err error) {
	if err != nil {
		Log.Fatal(err)
	}
}
