package env

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/pkg/errors"
)


type HadeEnv struct {
	folder string // represent env folder
	
	maps map[string]string
}

// NewHadeEnv have two params: folder and env
// for example: NewHadeEnv("/envfolder/")
// It will read file: /envfolder/.env
// The file have format XXX=XXX
func NewHadeEnv(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewHadeEnv param error")
	}
	
	folder := params[0].(string)
	// parse .env
	file := path.Join(folder, ".env")
	_, err := os.Stat(file)
	if err != nil || os.IsNotExist(err) {
		return nil, errors.New("file " + file + " not exist:" + err.Error())
	}
	fi, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	
	hadeEnv := &HadeEnv{
		folder: folder,
		maps : map[string]string{},
	}
	br := bufio.NewReader(fi)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := bytes.SplitN(line, []byte{'='}, 2)
		if len(s) < 2 {
			continue
		}
		key := string(s[0])
		val := string(s[1])
		hadeEnv.maps[key] = val
	}
	return hadeEnv, nil
}
// AppEnv get current environment
func (en *HadeEnv) AppEnv() string {
	return en.Get("APP_ENV")
}

// AppDebug check app is debug open
func (en *HadeEnv) AppDebug() bool {
	b, err := strconv.ParseBool(en.Get("APP_DEBUG"))
	if err == nil {
		return b
	}
	return false
}

// AppURL define app url in local
func (en *HadeEnv) AppURL() string {
	return en.Get("APP_URL")
}

// IsExist check setting is exist
func (en *HadeEnv) IsExist(key string) bool {
	_, ok := en.maps[key]
	return ok
}

// Get environment setting, if not exist, return ""
func (en *HadeEnv) Get(key string) string {
	if val, ok := en.maps[key]; ok {
		return val
	}
	return ""
}

// All return all settings
func (en *HadeEnv) All() map[string]string {
	return en.maps	
}