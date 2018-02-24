package main

import (
	"fmt"
	"student"
)

func main() {
	
	println("欢迎使用学生管理系统，请选择数字键进行操作")
	println("新增学生 1")
	println("删除学生 2")
	println("修改学生信息 3")
	println("查看所有学生 4")

	input := new(int)

	fmt.Scan(input)
	
	switch *input {
		case 1:
			student.Add()
		case 2:

		case 3:

		case 4:
			
		default:
			println("输入无效，请重新输入")
		}
}