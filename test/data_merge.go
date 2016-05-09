package main

import "fmt"
import "time"

func Uint642Bstr(data uint64) string {

	var buffer []byte

	buffer = make([]byte, 64)

	var mask uint64 = 1
	for i := 0; i < 64; i++ {
		if data&mask == 1 {
			buffer[63-i] = '1'
		} else {
			buffer[63-i] = '0'
		}

		data = data >> 1
	}

	return string(buffer)
}

func getCurMsec() uint64 {

	nanoSecs := time.Now().UnixNano()

	return uint64(nanoSecs / 1000 / 1000)
}

func main() {

	fmt.Printf("%d\n", getCurMsec())

	var a uint64 = getCurMsec()
	var b uint16 = 10
	var c uint16 = 0

	var first_42 uint64 = 0x7fffffffffffffff
	var mid_10 uint16 = 0x03ff
	var last_12 uint16 = 0x0fff

	var res uint64 = 0

	res = (a << 22 & first_42) | uint64((b&mid_10)<<12) | uint64(c&last_12)

	strRes := Uint642Bstr(res)

	fmt.Printf("res=%d\n", res)

	fmt.Printf("%s\n", strRes)

	fmt.Printf("first_42:%s\n", strRes[0:42])
	fmt.Printf("mid_10:%s\n", strRes[42:52])
	fmt.Printf("lst_12:%s\n", strRes[52:])

}
