package student

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

func saveToRedis(studentName string,studentInfo string) {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")

	if err != nil {
		fmt.Println("connect to redis error",err)
		return
	}

	defer conn.Close()

	_,err = conn.Do("SET",studentName,studentInfo)

	if err == nil {
		fmt.Println("save successed")
	}else {
		fmt.Println("save failed",err)
	}
}

type Student struct {
	Name string
	Age int
	Sex int
	Gradle int
	Class int
}

func Add() {
	name := new(string)
	age := new(int)
	sex := new(int)
	gradle := new(int)
	class := new(int)

	fmt.Println("请输入姓名")
	fmt.Scan(name)

	fmt.Println("请输入年龄")
	fmt.Scan(age)

	fmt.Println("请输入性别 1男2女")
	fmt.Scan(sex)

	fmt.Println("请输入年级(1-6)")
	fmt.Scan(gradle)

	fmt.Println("请输入班级(1-10)")
	fmt.Scan(class)

	student := Student{*name,*age,*sex,*gradle,*class}

	fmt.Println(student)

	studentInfo,_ := json.Marshal(student)

	saveToRedis(*name,string(studentInfo))
}

func Delete() {

}

func Edit() {

}

func Query() {

}