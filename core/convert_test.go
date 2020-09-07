package core

import (
	"fmt"
	"testing"
)

func TestIntToBytes(t *testing.T) {

	var c chan int

	fmt.Println(nil == c)
	select {
	case c<-1:
		
	}
	fmt.Println("oo")


}
