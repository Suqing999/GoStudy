package main

import "fmt"

type student struct {
	id   int
	name string
}

//需要有一个学生的管理者 （保留一些数据  有功能 即 结构体方法）
type studentManager struct {
	allStudent map[int]student
}

//查看学生
func (sm studentManager) showAll() {
	for _, v := range sm.allStudent {
		fmt.Printf("学号:%d   姓名:%s\n", v.id, v.name)
	}
	fmt.Println()
}

//添加学生
func (sm studentManager) addStu() {
	var (
		newid   int
		newname string
	)
	fmt.Println("请输入学生学号:")
	fmt.Scanln(&newid)
	fmt.Println("请输入学生姓名:")
	fmt.Scanln(&newname)

	newStu := student{
		id:   newid,
		name: newname,
	}

	sm.allStudent[newStu.id] = newStu

}

//删除学生
func (sm studentManager) deleteStu() {
	//输入学号
	var (
		delid   int
		delname string
	)
	fmt.Println("请输入学生学号:")
	fmt.Scanln(&delid)
	delname = sm.allStudent[delid].name
	delete(sm.allStudent, delid)
	fmt.Printf("学号为:%d的学生%s已经被删除！", delid, delname)
}

//修改学生
func (sm studentManager) updateStu() {
	//输入学号
	var (
		upid   int
		upname string
	)
	fmt.Println("请输入学生学号:")
	fmt.Scanln(&upid)

	stuObj, ok := sm.allStudent[upid]

	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println("请输入新的姓名:")
		fmt.Scanln(&upname)
		stuObj.name = upname

		//更新
		sm.allStudent[upid] = stuObj
	}

	fmt.Println("修改成功！")

}
