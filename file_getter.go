package idgen

import "io/ioutil"
import "strconv"

type FileGen struct {
	Filename string
}

func NewFileGen(filename string) *FileGen {

	ins := new(FileGen)
	ins.Filename = filename

	return ins
}

func (this *FileGen) Get() uint16 {

	buffer, err := ioutil.ReadFile(this.Filename)

	if err != nil {
		return 0
	}

	var content int

	content, err = strconv.Atoi(string(buffer))

	return uint16(content)
}
