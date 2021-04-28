package controller

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xluohome/phonedata"
	"io"
	"os"
	"regexp"

	"strings"
	"testing"
)

func TestGetGoodsList(t *testing.T) {
	//va := time.Now().Unix()
	//fmt.Println(reflect.TypeOf(va))

	pr, err := phonedata.Find("18977155589")
	if err != nil {
		panic(err)
	}
	fmt.Print(pr)
	readtxt()
}

func readtxt() {
	fi, err := os.Open("/Users/j/Library/Mobile Documents/com~apple~Pages/Documents/jcm_.csv")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	Phonenumber := `1[3456789]\d{9}`

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str1 := strings.Split(string(a), ",")
		//fmt.Println(str1[0])
		if len(str1[0]) == 11 {
			pr, err := phonedata.Find(str1[0])
			if err != nil {
				panic(err)
			}
			fmt.Println(string(a) + "," + pr.Province + "," + pr.City + "," + pr.ZipCode + "," + pr.AreaZone)
		} else {
			re := regexp.MustCompile(Phonenumber)
			//reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
			rgx := regexp.MustCompile(Phonenumber)
			if rgx.MatchString(string(a)) == false {
				continue
			}
			allString := re.FindAllStringSubmatch(string(a), -1)

			//fmt.Println(allString[0])
			//fmt.Println(allString[0][0])

			pr, err := phonedata.Find(allString[0][0])
			if err != nil {
				panic(err)
			}

			fmt.Println(string(a) + "," + pr.Province + "," + pr.City + "," + pr.ZipCode + "," + pr.AreaZone)

		}

	}
}

func TestGoodsCats(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGoodsDetails(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestRegUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSaveGoods(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestUpload(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
