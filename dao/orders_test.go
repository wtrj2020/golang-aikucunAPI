package dao

import (
	"fmt"
	"testing"
)

func testta() {

}

func TestGetCartList(t *testing.T) {

	list := GetCartList(60)
	fmt.Println(list)
}
