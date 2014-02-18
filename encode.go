package gosimplehttp

import (
	"bytes"
	"encoding/xml"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

// EscapeQuotes escapes a string which is to be represented in a quoted
// element. This is useful for HTML headers, as well as HTML attributes,
// when the source is wrapped in an XmlEntitiesString call.
func EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// XmlEntities performs entity replacement for text to be used in
// representations meant for display in HTML/XML documents. This method
// consumes and produces []byte types.
func XmlEntities(b []byte) []byte {
	w := bytes.NewBuffer([]byte{})
	xml.Escape(w, b)
	return w.Bytes()
}

// XmlEntitiesString performs entity replacement for text to be used in
// representations meant for display in HTML/XML documents. This method
// consumes and produces string types.
func XmlEntitiesString(s string) string {
	w := bytes.NewBuffer([]byte{})
	xml.Escape(w, []byte(s))
	return w.String()
}
