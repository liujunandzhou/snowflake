package idgen

//最大1024
type RandomGen struct {
	cur uint16
}

const MAX_VALUE = 1024

func NewRandomGen() *RandomGen {
	ins := new(RandomGen)
	ins.cur = 0

	return ins
}

func (this *RandomGen) Get() uint16 {

	ret := this.cur

	this.cur = this.cur + 1

	if this.cur >= MAX_VALUE {
		this.cur = 0
	}

	return ret
}
