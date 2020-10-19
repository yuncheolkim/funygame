package utils

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	time.AfterFunc(time.Second*3, func() {

		fmt.Println("after")
	})

	t1 := time.Tick(time.Second)

	time.Sleep(time.Second * 3)
	for {

		select {
		case <-t1:
			{
				fmt.Println(time.Now())
			}
		}
	}
}
