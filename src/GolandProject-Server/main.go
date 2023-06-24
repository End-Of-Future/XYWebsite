package main

import (
	"GolandProject-Server/AccountSystem"
	"GolandProject-Server/modules/ConfigHelper"
	"GolandProject-Server/modules/Log"
	"fmt"
	"net/http"
	"path/filepath"
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
	case "c":
		fmt.Println("Parameters following:")
		for k, v := range r.Form {
			fmt.Println("KEY:", k, " VALUE:", v)
		}
		p, err := filepath.Abs("")
		chk(err)
		_, err = fmt.Fprintf(w, c.GetFile(p+"\\src\\Contents\\"+r.FormValue("path"))) //左目录
		chk(err)
	default:
		_, err := fmt.Fprintf(w, "Unknown")
		chk(err)
	}
}

func Init() {
	http.HandleFunc("/items", Items)
	http.HandleFunc("/signin", AccountSystem.SignIn)
	http.HandleFunc("/signup", AccountSystem.SignUp)
	fmt.Println("Server-Served")
	err := http.ListenAndServe(":2047", nil)
	chk(err)
}

func Test() {
	AccountSystem.Test(1)
}

func main() {
	//Init()
	Test()
}

func chk(err error) {
	if err != nil {
		Log.Fatal(err)
	}
}
