package gen

import (
	"bytes"
	"go/format"
	"io"
	"log"
	"os"
	"path"
)

func GenerateToPackage(packageName, fileName string, templateExecute func(io.Writer) error) {
	err := os.Mkdir(packageName, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	genPath := path.Join(packageName, fileName)
	f, err := os.OpenFile(genPath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)

	err = templateExecute(buf)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	w := 0
	for w < len(bs) {
		n, err := f.Write(bs)
		if err != nil {
			log.Fatal(err)
		}
		w += n
	}
}
