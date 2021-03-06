package browser

// GMetaData contain all the fileds in each response gmeta.
type GMetaData struct {
	Gid          int64    `json:"gid"`
	Token        string   `json:"token"`
	ArchiverKey  string   `json:"archiver_key"`
	Title        string   `json:"title"`
	TitleJpn     string   `json:"title_jpn"`
	Category     string   `json:"category"`
	Thumb        string   `json:"thumb"`
	Uploader     string   `json:"uploader"`
	Posted       string   `json:"posted"`
	Filecount    string   `json:"filecount"`
	Filesize     int64    `json:"filesize"`
	Expunged     bool     `json:"expunged"`
	Rating       string   `json:"rating"`
	Torrentcount string   `json:"torrentcount"`
	Tags         []string `json:"tags"`
	Err          string   `json:"error"`
}

// G is the response struct
type G struct {
	Gmd []GMetaData `json:"gmetadata"`
}

// GData is post message type.
type GData struct {
	Method    string     `json:"method"`
	GidList   [][]string `json:"gidlist"`
	NameSpace int        `json:"namespace"`
}

// NewGData generate gdata for post message.
func NewGData(ns int) *GData {
	return &GData{
		Method:    "gdata",
		NameSpace: ns,
	}
}

// SetComic append comic into GData
func (gd *GData) SetComic(id string, token string) {
	gd.GidList = append(gd.GidList, []string{id, token})
}

// InvalidKeyError is error message when an invalid key provided.
type InvalidKeyError struct {
	Gid       string
	ErrorInfo string
}

func (i *InvalidKeyError) Error() string {
	return i.ErrorInfo
}
