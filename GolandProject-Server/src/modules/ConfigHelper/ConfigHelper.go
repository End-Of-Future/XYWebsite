package ConfigHelper

import (
	"GolandProject/src/modules/Log"
	"os"
	"path/filepath"
	"strings"
)

type Conf struct {
	content []string
	path    string
}

func NewConf(path string) Conf {
	p, err := filepath.Abs(path)
	chk(err)
	f, err := os.ReadFile(p)
	chk(err)
	res := strings.Split(string(f), "\n")
	return Conf{content: res, path: path}
}

func (c Conf) FindQuery(rname string) string {
	for _, val := range c.content {
		if strings.HasPrefix(val, "["+rname+"]") {
			return strings.TrimLeft(val, "["+rname+"]")
		}
	}
	return ""
}

func chk(err error) {
	if err != nil {
		Log.Fatal(err)
	}
}
