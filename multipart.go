package gosimplehttp

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
)

type multipartComponenter interface {
	Name() string
	Encode()
	SetWriter(multipart.Writer)
}

//-------------------------------------------------------------------

type mpFile struct {
	name     string
	filename string
	filetype string
	writer   multipart.Writer
}

func (s *mpFile) Name() string {
	return s.name
}

func (s *mpFile) Encode() error {
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

func (s *mpFile) SetWriter(w multipart.Writer) {
	s.writer = w
}

func PostFile(k, n, t string) *mpFile {
	p := &mpFile{name: k, filename: n, filetype: t}
	return p
}

//-------------------------------------------------------------------

type mpValue struct {
	name   string
	value  string
	writer multipart.Writer
}

func (s *mpValue) Name() string {
	return s.name
}

func (s *mpValue) Encode() {
	_ = s.writer.WriteField(s.name, s.value)
}

func (s *mpValue) SetWriter(w multipart.Writer) {
	s.writer = w
}

func PostValue(k, v string) *mpValue {
	p := &mpValue{name: k, value: v}
	return p
}
