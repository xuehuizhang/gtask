package tenantDao

import (
	_ "codego/com.yiqibishe/dao_v2"
	"codego/com.yiqibishe/model"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	err := Add(&model.Nav{
		Name:   "Java",
		Url:    "/java",
		Status: 1,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestDel(t *testing.T) {

}

func TestGetOne(t *testing.T) {
	one, err := GetOne(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(one)
}

func TestRaw(t *testing.T) {
	list, err := Raw("select * from nav where id>?", []interface{}{0})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func TestCount(t *testing.T) {
	count, err := Count("id >?", []interface{}{0})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}
