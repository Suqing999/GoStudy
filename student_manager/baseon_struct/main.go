package main

import (
	"fmt"
	"os"
)

//全局变量声明
var sm studentManager

func menu() {
	fmt.Println("学生管理系统")
	fmt.Println("############")
	fmt.Println("1、查看所有")
	fmt.Println("2、增加学生")
	fmt.Println("3、删除学生")
	fmt.Println("4、修改学生")
	fmt.Println("5、退出系统")
	fmt.Println("############")
}

func main() {
	sm = studentManager{
		allStudent: make(map[int]student, 32),
	}
	for {
		menu()
		//用户输入
		fmt.Println("请输入选项:")
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			sm.showAll()
		case 2:
			sm.addStu()
		case 3:
			sm.deleteStu()
		case 4:
			sm.updateStu()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入错误！")
		}

	}
}
