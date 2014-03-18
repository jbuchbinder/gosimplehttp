package gosimplehttp

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
)

// MultipartComponenter interface defines MIME sections of an HTTP POST
// request.
type MultipartComponenter interface {
	// Encode instantiates the POST section by rendering the output to
	// the multipart.Writer object which was passed to SetWriter.
	Encode() error

	// SetWriter sets the multipart.Writer which is used to render the
	// section. It is called before the Encode() method.
	SetWriter(multipart.Writer)
}

// MpFile is a MultpartComponenter implementation.
type MpFile struct {
	name     string
	filename string
	filetype string
	writer   multipart.Writer
}

func (s *MpFile) Encode() error {
	file, err := os.Open(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			EscapeQuotes(s.name),
			EscapeQuotes(filepath.Base(s.filename))))
	h.Set("Content-Type", s.filetype)
	part, err := s.writer.CreatePart(h)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	return nil
}

func (s *MpFile) SetWriter(w multipart.Writer) {
	s.writer = w
}

// PostFile creates a MultipartComponenter instance exposing a file
// for a POST request.
func PostFile(k, n, t string) *MpFile {
	p := &MpFile{name: k, filename: n, filetype: t}
	return p
}

// MpData is a MultpartComponenter implementation.
type MpData struct {
	name     string
	data     []byte
	filetype string
	writer   multipart.Writer
}

func (s *MpData) Encode() error {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"`,
			EscapeQuotes(s.name)))
	h.Set("Content-Type", s.filetype)
	part, err := s.writer.CreatePart(h)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(s.data)

	_, err = io.Copy(part, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *MpData) SetWriter(w multipart.Writer) {
	s.writer = w
}

// PostData creates a MultipartComponenter instance exposing a file
// for a POST request.
func PostData(k string, d []byte, t string) *MpData {
	p := &MpData{name: k, data: d, filetype: t}
	return p
}

// MpValue is a MultpartComponenter implementation.
type MpValue struct {
	name   string
	value  string
	writer multipart.Writer
}

func (s *MpValue) Encode() error {
	err := s.writer.WriteField(s.name, s.value)
	return err
}

func (s *MpValue) SetWriter(w multipart.Writer) {
	s.writer = w
}

// PostValue creates a MultipartComponenter instance exposing a parameter
// for a POST request.
func PostValue(k, v string) *MpValue {
	p := &MpValue{name: k, value: v}
	return p
}
