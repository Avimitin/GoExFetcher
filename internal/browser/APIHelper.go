package browser

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// CheckErr will check error, if error is not nil, program will print a log and exit.
func CheckErr(err error) {
	if err != nil {
		log.Printf("[Request Error]%s", err.Error())
		os.Exit(-1)
	}
}

// NewHTTPClient return a new client.
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
	}
}

// NewGData generate gdata for post message.
// Require manga's id and token, return a parsed json byte message.
// Exp:
// With a link like this: https://e-hentai.org/g/618395/0439fa3666/
// it's id is 618396 and it's token is 0439fa3666
func NewGData(id string, token string) ([]byte, error) {
	gl := [][]string{{id, token}}
	gd := GData{
		Method:    "gdata",
		GidList:   gl,
		NameSpace: 1,
	}
	g, err := json.Marshal(gd)
	if err == nil {
		return g, nil
	}
	log.Printf("Error occur when parsing posted data, INFO: %v\n", err)
	return nil, err
}

// Reg will use pattern to fetch id and token in user's input
func Reg(url string) string {
	pattern := regexp.MustCompile(`\d+/\w+`)
	if pattern != nil {
		results := pattern.FindAllStringSubmatch(url, -1)
		if len(results) == 0 {
			return ""
		}
		return results[0][0]
	}
	return ""
}

// GetGIDAndToken will parse the target and return parsed args
func GetGIDAndToken(target string) []string {
	result := strings.Split(target, "/")
	if len(result) != 2 {
		return nil
	}
	return result
}

// Beautify will output a beautilfy gmetadata struct
func Beautify(result *GMetaData) string {
	bt, _ := json.Marshal(result)
	var out bytes.Buffer
	json.Indent(&out, bt, "", "\t")
	return out.String()
}
