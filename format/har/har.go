package format

type HARObject struct {
	Log HARLog `json:"log"`
}

type HARLog struct {
	Version string      `json:"version"`
	Creator interface{} `json:"creator"`
	Entries []Entry     `json:"entries"`
}

type Entry struct {
	Request Request `json:"request"`
}

type Request struct {
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	QueryString []QueryString `json:"queryString"`
	PostData    PostData      `json:"postData"`
}

type QueryString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PostData struct {
	MimeType string `json:"mimeType"`
	Text     string `json:"text"`
}
