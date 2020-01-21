package main

import (
	"fmt"
	"go_demo/mysqlt"
)

type SUKA int

const (
	suka1 SUKA = iota
	suka2
	suka3
	suka4
	suka5
)

func dealData() (int, int, string) {
	a := 88
	b := 99
	c := `dsds`
	return a, b, c
}
func main() {

	a, _, _ := dealData()
	_, b, _ := dealData()
	c := map[string]int{"map1": 001, "map2": 002}
	d := c["map1"]
	f := &a
	e := *f
	var g int
	H := [3]int{1, 2, 3}
	sdf := [3]interface{}{1, "sdf", 8788.222}

	for k, v := range H {
		fmt.Print("bianli", k, v)
	}

	fmt.Println("a=", a, ",", "b=", b, ",", "c=", c, ",", "d=", d)
	fmt.Println("e=", e, ",", "f=", f, ",", "g=", g, ",", "d=", d)
	fmt.Println("sika", suka1, suka2, suka3, suka4, suka5)

	p := person{name: "dada", age: 18}
	p.eat()

	m := mouse{name: "鼠标"}
	m.start()
	m.end()

	u := Upan{name: "u盘"}
	u.start()
	u.end()

	testIn(m)
	testIn(u)

	fmt.Println(sdf)


	fmt.Println("================MySQL================")

	connect := new(mysqlt.Connect)
	connect.UserName="root"
	connect.Password="123456"
	connect.Ip="*****"
	connect.Port="****"
	connect.DbName="***"
	mysqlt.ConnectMysql(connect)

}


type person struct {
	name string
	age  int
}

func (p person) eat() {
	fmt.Printf("%s 吃饭了,今年 %d 岁/n", p.name, p.age)
}

type USB interface {
	start()
	end()
}
type TSB interface {
	start()
	end()
}
type mouse struct {
	name string
}
type Upan struct {
	name string
}

func (m mouse) start() {
	fmt.Println(m.name, "开始工作")
}
func (m mouse) end() {
	fmt.Println(m.name, "结束工作")
}
func (m Upan) start() {
	fmt.Println(m.name, "开始工作")
}
func (m Upan) end() {
	fmt.Println(m.name, "结束工作")
}
func testIn(usb USB) {
	usb.start()
	usb.end()
}
