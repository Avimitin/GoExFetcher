package browser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// JSONBrowse is only for json browse
func parse(body []byte) (*G, error) {
	var repJSON G
	err := json.Unmarshal(body, &repJSON)
	if err != nil {
		return nil, fmt.Errorf("unmarshall %s: %v", body, err)
	}
	return &repJSON, nil
}

// Browser use given gallery id and token to fetch comic info
// return a gmetadata struct pointer.
func Browser(galleryID string, token string) (*G, error) {
	gdata := NewGData(0)
	gdata.SetComic(galleryID, token)
	gd, err := json.Marshal(&gdata)
	if err != nil {
		return nil, fmt.Errorf("Browser: marshal %+v: %v", gdata, err)
	}

	req, err := http.NewRequest("POST", "https://api.e-hentai.org/api.php", bytes.NewReader(gd))
	// post request.
	if err != nil {
		return nil, fmt.Errorf("Browser: new request %v", err)
	}

	client := NewHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Browser: do request %v got %v", req, err)
	}

	// read and store bytes.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Browser: read response %v", err)
	}

	g, err := parse(body)
	if err != nil {
		return nil, fmt.Errorf("Browser: parse %v got %w", body, err)
	}
	var gmdWithoutErr []GMetaData
	for _, gmd := range g.Gmd {
		if gmd.Err != "" {
			log.Printf("[Input Error]%d's token is invalid or this gid is not existed.", gmd.Gid)
			continue
		}
		gmdWithoutErr = append(gmdWithoutErr, gmd)
	}
	if len(gmdWithoutErr) == 0 {
		return nil, errors.New("Browser: all given comic are invalid")
	}
	g.Gmd = gmdWithoutErr
	return g, nil
}
