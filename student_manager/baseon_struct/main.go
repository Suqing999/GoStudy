package main

import (
	"fmt"
	"os"
)

//全局变量
var (
	allStudent map[int]*student //声明变量
)

type student struct {
	id   int
	name string
}

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAll() {
	//打印所有学生
	for _, v := range allStudent {
		fmt.Printf("学号:%d   名字:%s\n", v.id, v.name)
	}
}

func insertStu() {
	//输入学生学号
	var (
		id   int
		name string
	)
	fmt.Println("请输入学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入名字:")
	fmt.Scanln(&name)

	//追加进map
	newStu := newStudent(id, name)
	allStudent[id] = newStu
	fmt.Println("添加完成")
}

func deleteStu() {
	//输入学生学号
	var (
		id int
	)
	fmt.Println("请输入学号:")
	fmt.Scanln(&id)
	//追用delete函数删除
	delete(allStudent, id)
	fmt.Println("删除完成")
}

func main() {
	var inputNum int
	//初始化学生
	allStudent = make(map[int]*student, 32)
	allStudent[0] = newStudent(0, "wx0")
	allStudent[1] = newStudent(1, "wx1")
	allStudent[2] = newStudent(2, "wx2")
	allStudent[3] = newStudent(3, "wx3")
	allStudent[4] = newStudent(4, "wx4")
	allStudent[5] = newStudent(5, "wx5")

	for {

		//菜单
		fmt.Println("学生管理系统")
		fmt.Println("############")
		fmt.Println(`
1、查看所有学生
2、新增学生
3、删除学生
4、退出
	`)
		fmt.Println("############")
		fmt.Println("输入选项：")
		//选择
		fmt.Scanln(&inputNum)
		//执行函数
		switch inputNum {
		case 1:
			showAll()
		case 2:
			insertStu()
		case 3:
			deleteStu()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入错误!")
		}

	}
}
