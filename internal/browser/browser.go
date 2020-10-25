package browser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// JSONBrowse is only for json browse
func parser(body []byte) *G {
	var repJSON *G
	err := json.Unmarshal(body, &repJSON)
	CheckErr(err)
	return repJSON
}

// Browser use given gallery id and token to fetch comic info
// return a gmetadata struct pointer.
func Browser(galleryID string, token string) (*G, error) {
	gd, err := NewGData(galleryID, token)
	if gd == nil {
		return nil, &InvalidKeyError{
			Gid:       galleryID,
			ErrorInfo: err.Error(),
		}
	}

	req, err := http.NewRequest("POST", "https://api.e-hentai.org/api.php", bytes.NewReader(gd))
	// post request.
	CheckErr(err)

	client := NewHTTPClient()
	resp, err := client.Do(req)
	CheckErr(err)

	// read and store bytes.
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)

	g := parser(body)
	for _, gmd := range g.Gmd {
		if gmd.Err != "" {
			log.Printf("[Input Error]%d's token is invalid or this gid is not existed.", gmd.Gid)
		}
	}
	return g, nil
}
