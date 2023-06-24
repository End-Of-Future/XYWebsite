package SQLiteHelper

import (
	"GolandProject-Server/modules/Log"
	"database/sql"
	"path/filepath"
)

var db *sql.DB
var ins_cmd, del_cmd *sql.Stmt

func Init() {
	p, err := filepath.Abs("")
	chk(err)
	db, err = sql.Open("sqlite3", p+"\\src\\GolandProject-Server\\dbs\\user.db")
	chk(err)
	ins_cmd, err = db.Prepare("INSERT INTO users VALUES(?,?,?);")
	chk(err)
	del_cmd, err = db.Prepare("DELETE FROM users WHERE id=?;")
	chk(err)
}

/* Notes of operating SQLite database

INSERT:
	p, err := filepath.Abs("")
	chk(err)
	db, err := sql.Open("sqlite3", p+"\\src\\GolandProject-Server\\dbs\\user.db")
	chk(err)
	cmd, err := db.Prepare("INSERT INTO users VALUES(?,?,?);") //命令根据实际需求改
	chk(err)
	res, err := cmd.Exec(1, -1, "I_NOT_FOUND") //参数根据实际需求改
	chk(err)
	rid, err := res.LastInsertId()
	chk(err)
	Log.Send("Ln: ", rid) //id为行号

UPDATE:
	p, err := filepath.Abs("")
	chk(err)
	db, err := sql.Open("sqlite3", p+"\\src\\GolandProject-Server\\dbs\\user.db")
	chk(err)
	cmd, err := db.Prepare("UPDATE users SET name=? WHERE id=?") //命令根据实际需求改
	chk(err)
	res, err := cmd.Exec("inf", 1) //参数根据实际需求改
	chk(err)
	rid, err := res.RowsAffected() //rid为行号
	chk(err)
	Log.Send("Ln: ", rid)

SELECT:
	p, err := filepath.Abs("")
	chk(err)
	db, err := sql.Open("sqlite3", p+"\\src\\GolandProject-Server\\dbs\\user.db")
	chk(err)
	rows, err := db.Query("SELECT * FROM users;")
	chk(err)
	for rows.Next() {
		//先声明再读值
		var id, lv int
		var name string
		err = rows.Scan(&id, &lv, &name)
		chk(err)
		fmt.Println("ID: ", id, " LV: ", lv, " NAME: ", name)
	}
*/

func Select(query string, args ...any) (int, int, string, error) {
	rows, err := db.Query(query, args...)
	chk(err)
	rows.Next()
	var id, lv int
	var name string
	err = rows.Scan(&id, &lv, &name)
	chk(err)
	return id, lv, name, err
}

func Insert() {
	Log.Send("INSERT")
}

func Update(col_name string, col_value any, id int) {
	upd_cmd, err := db.Prepare("UPDATE users SET " + col_name + "=? WHERE id=?;")
	chk(err)
	res, err := upd_cmd.Exec(col_value, id)
	chk(err)
	rid, err := res.RowsAffected()
	chk(err)
	Log.Send("Ln: ", rid)
}

func Delete() {
	Log.Send("DELETE")
}

func chk(err error) {
	if err != nil {
		Log.Fatal(err)
	}
}
