package browser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

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
func NewGData(id string, token string) []byte {
	gl := [][]string{{id, token}}
	gd := GData{
		Method:    "gdata",
		GidList:   gl,
		NameSpace: 1,
	}
	g, err := json.Marshal(gd)
	if err == nil {
		return g
	}
	fmt.Printf("Error occur when parsing posted data, INFO: %v\n", err)
	return nil
}

// JSONBrowse is only for json browse
func JSONBrowse() {
	client := NewHTTPClient()
	gd := NewGData("1758255", "8c7c98ff73")
	// if gd == nil this also mean some error has occur.
	if gd == nil {
		return
	}

	req, err := http.NewRequest("POST", "https://api.e-hentai.org/api.php", bytes.NewReader(gd))
	// post request.
	CheckErr(err)
	resp, err := client.Do(req)
	CheckErr(err)
	// read and store bytes.
	CheckErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	// unmarshall bytes.
	CheckErr(err)
	var repJSON G
	err = json.Unmarshal(body, &repJSON)
	CheckErr(err)
	for _, comic := range repJSON.Gmd {
		if comic.Err != "" {
			log.Printf("[Input Error]%s", comic.Err)
		}
	}
	fmt.Printf("data: \n%v", repJSON)
}

// CheckErr will check error, if error is not nil, program will print a log and exit.
func CheckErr(err error) {
	if err != nil {
		log.Printf("[Request Error]%s", err.Error())
		os.Exit(-1)
	}
}
