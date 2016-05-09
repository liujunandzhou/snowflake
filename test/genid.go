package main

import "snowflake"
import "fmt"
import "time"

func main() {
	gGetter := idgen.NewRandomGen()

	idMaker := idgen.NewIdMaker(gGetter)

	for i := 0; i < 1000000; i++ {
		fmt.Println(idMaker.GenId())
		time.Sleep(time.Second)
	}
}
