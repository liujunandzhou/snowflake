package idgen

import "time"

const MAX_SEQ = 1024
const FIRST_43 uint64 = 0x7fffffffffffffff
const MID_10 uint16 = 0x03ff
const LAST_12 uint16 = 0x0fff

//id生成器
type IdMaker struct {
	seq      uint16 //最大1024
	mid      uint16 //中间的id
	lasttime uint64 //最后一次时间戳
	getter   Getter //获取借口
}

//获取中间的uniqid,可以实现接口
type Getter interface {
	Get() uint16
}

func getCurMsec() uint64 {

	nanoSecs := time.Now().UnixNano()
	return uint64(nanoSecs / 1000 / 1000)
}

func buildUint64(stamp uint64, mid_10 uint16, last_12 uint16) uint64 {

	var res uint64 = 0

	res = (stamp << 22 & FIRST_43) | uint64((mid_10&MID_10)<<12) | uint64(last_12&LAST_12)

	return res
}

func NewIdMaker(get Getter) *IdMaker {

	idMaker := &IdMaker{0, 0, 0, get}

	idMaker.mid = get.Get()

	return idMaker
}

func (this *IdMaker) waitNextMs() uint64 {

	var cur uint64 = 0

	for {
		cur = getCurMsec()

		if cur > this.lasttime {
			break
		}
	}

	this.lasttime = cur

	return cur
}

//不考虑线程安全性
func (this *IdMaker) GenId() uint64 {

	this.seq = this.seq + 1

	if this.seq >= MAX_SEQ {
		this.seq = 0
		this.waitNextMs()
	}

	return buildUint64(getCurMsec(), this.mid, this.seq)
}
