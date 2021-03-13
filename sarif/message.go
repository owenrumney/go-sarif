package sarif

type Message struct { // https://docs.oasis-open.org/sarif/sarif/v2.1.0/csprd01/sarif-v2.1.0-csprd01.html#_Toc10540897
	Text      *string  `json:"text,omitempty"`
	Markdown  *string  `json:"markdown,omitempty"`
	Id        *string  `json:"id,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
}
