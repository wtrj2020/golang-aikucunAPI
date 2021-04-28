package tools

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoodsIdRepeat(t *testing.T) {
	type args struct {
		goodsIDs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoodsIdRepeat(tt.args.goodsIDs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoodsIdRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuid(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Guid(tt.args.userId); got != tt.want {
				t.Errorf("Guid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5_salt(t *testing.T) {
	type args struct {
		str  string
		salt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5_salt(tt.args.str, tt.args.salt); got != tt.want {
				t.Errorf("Md5_salt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand(t *testing.T) {

fmt.Println(	Rand(99999))
}