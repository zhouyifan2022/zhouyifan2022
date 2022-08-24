package main

import (
	"fmt"
	"os"
)

//增加库存fittings信息
//显示NO编号 Model or specification型号或规格 quantity库存数量
//删除信息
var fimsDate map[int]*Fittings

type Fittings struct {
	no                   int64
	modelorspecification string
	quantity             int64
}

func init() {
	fimsDate = make(map[int]*Fittings, 100)
}
func main() {
	for {
		fmt.Println("---welcom minluck---")
		fmt.Println("1.查看配件库存\n2.添加配件入库\n3.从库内删除配件\n4.退出系统")
		fmt.Println("请输入操作号:")

		var p1 int64
		fmt.Scanln(&p1)
		fmt.Printf("你的选择的是：%d\n", p1)
		switch p1 {
		case 1:
			showms()

		case 2:
			addms()

		case 3:
			delms()

		case 4:
			os.Exit(1)

		default:
			fmt.Println("输入错误是无效的参数")

		}
	}
}
func showms() {
	fmt.Println("------------------------------")
	for k, v := range fimsDate {
		fmt.Printf("NO编号:%d\n,Model or specification型号或规格:%d\n, quantity库存数量:%d\n", k, v.modelorspecification, v.quantity)

	}
	fmt.Println("--------------------------------")

}
func addms() {
	var (
		no                   int64
		modelorspecification string
		quantity             int64
	)
	fmt.Println("请输入编号：")
	fmt.Scanln(&no)
	fmt.Println("请输入型号或规格：")
	fmt.Scanln(&modelorspecification)
	fmt.Println("请输入数量：")
	fmt.Scanln(&quantity)

	fims := newaddms(no, modelorspecification, quantity)
	fimsDate[int(no)] = fims
	fmt.Printf("编号为：%d的入库完成\n", no)

}
func delms() {
	var no int64
	fmt.Println("请输入编号")
	fmt.Scanln(&no)
	delete(fimsDate, int(no))
	fmt.Printf("编号为：%d的配件删除成功")

}
func newaddms(no int64, modelorspecification string, quantity int64) *Fittings {
	return &Fittings{
		no:                   no,
		modelorspecification: modelorspecification,
		quantity:             quantity,
	}

}
func Exit() {

	fmt.Println("欢迎再来")
}
