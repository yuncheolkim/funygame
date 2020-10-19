package game

import (
	"fmt"
	"testing"
	"time"
)

func TestServer_Serve(t *testing.T) {
	Start()

}

func Test1(t *testing.T) {

	time.AfterFunc(time.Second*3, func() {

		fmt.Println("after")
	})

	fmt.Println("oks")

}
