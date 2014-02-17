package gosimplehttp

import (
	"bytes"
	"encoding/xml"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func XmlEntities(b []byte) []byte {
	w := bytes.NewBuffer([]byte{})
	xml.Escape(w, b)
	return w.Bytes()
}

func XmlEntitiesString(s string) string {
	w := bytes.NewBuffer([]byte{})
	xml.Escape(w, []byte(s))
	return w.String()
}
