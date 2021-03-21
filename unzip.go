package unzip

import (
	"archive/zip"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"
)

type Unzip struct {
	Src string
	Dest string
}

func New(src string,dest string) Unzip {

	return Unzip{src,dest}
}

func (uz Unzip) Extract() error {

	zipReader, err := zip.OpenReader(uz.Src)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	var decodeName string
	for _, f := range zipReader.File {

		if utf8.Valid([]byte(f.Name)) {
			decodeName = f.Name
		} else {
			i := bytes.NewReader([]byte(f.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			decodeName = string(content)
		}

		fpath := filepath.Join(uz.Dest, decodeName)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
