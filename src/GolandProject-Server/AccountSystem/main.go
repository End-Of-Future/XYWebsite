package AccountSystem

import (
	"GolandProject-Server/AccountSystem/SQLiteHelper"
	"GolandProject-Server/modules/Log"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

const (
	OPERATE_SELECT = iota + 1
	OPERATE_INSERT
	OPERATE_UPDATE
	OPERATE_DELETE
)

//TODO:Finish those two functions to complete user system

func Test(mode int) {
	SQLiteHelper.Init()
	switch mode {
	case OPERATE_SELECT:
		s := []any{"", 0, ""}
		err := SQLiteHelper.SliceSelect(s, "SELECT * FROM users;")
		chk(err)
		Log.Send("s[0]: ", s[0], "s[1]: ", s[1], "s[2]: ", s[2])
	case OPERATE_INSERT:
		SQLiteHelper.Insert()
	case OPERATE_UPDATE:
		SQLiteHelper.Update("name", "I_NOT_FOUND", 1)
	case OPERATE_DELETE:
		SQLiteHelper.Delete()
	default:
		Log.Error("Unknown mode")
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {

}
func SignUp(w http.ResponseWriter, r *http.Request) {

}

func chk(err error) {
	if err != nil {
		Log.Fatal(err)
	}
}
