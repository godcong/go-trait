package trait

import (
	"bytes"
	"compress/zlib"
	"github.com/rakyll/statik/fs"
	"io"
	"os"
	"path"
)

// CompressToFile ...
func CompressToFile(filename string, b []byte) error {
	buff := bytes.NewBuffer(b)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	w := zlib.NewWriter(file)
	if _, err := io.Copy(w, buff); err != nil {
		return err
	}
	w.Flush()
	defer w.Close()
	return nil
}

// DecompressFromFile ...
func DecompressFromFile(filename string) ([]byte, error) {
	buff := bytes.Buffer{}
	file, err := os.OpenFile(filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r, err := zlib.NewReader(file)
	if err == nil {
		if _, err := io.Copy(&buff, r); err == nil {
			return buff.Bytes(), nil
		}
	}
	defer r.Close()
	return nil, err
}

// DecompressFromStatik ...
func DecompressFromStatik(filename string) ([]byte, error) {
	buff := bytes.Buffer{}
	sfs, err := fs.New()
	if err != nil {
		return nil, err
	}

	file, err := sfs.Open(path.Join("/", filename))
	if err != nil {
		return nil, err
	}
	r, err := zlib.NewReader(file)
	if err == nil {
		if _, err := io.Copy(&buff, r); err == nil {
			return buff.Bytes(), nil
		}
	}
	defer r.Close()
	return nil, err
}
