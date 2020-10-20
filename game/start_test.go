package game

import (
	"fmt"
	"testing"
)

func TestServer_Serve(t *testing.T) {
	Start()

}

func Test1(t *testing.T) {

	a := []int{1,2,3,4,5}
	a = DeleteSlice(a,1)
	fmt.Println(a)
	a = DeleteSlice(a,2)
	fmt.Println(a)
	a = DeleteSlice(a,3)
	fmt.Println(a)
	a = DeleteSlice(a,4)
	fmt.Println(a)
	a = DeleteSlice(a,5)
	fmt.Println(a)
}
