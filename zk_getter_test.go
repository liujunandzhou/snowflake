package idgen_test

import "testing"
import "time"
import "fmt"
import . "github.com/liujunandzhou/snowflake"

func TestZkGetter(t *testing.T) {
	host := "localhost:2181"
	duration := time.Second * 5

	ins := NewZkGen(host, duration)

	for i := 0; i < 2000; i++ {
		fmt.Println(ins.Get())
	}
}
