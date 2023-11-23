package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	// วิธีที่ 1 แยกออกมาให้มองเห็น
	/* var c uint32
	c = a + b
	return c */
	// ==============
	// วิธีที่ 2 สั้นๆกระชับ
	return a + b
	// ==============
}

func main() {
	var a, b, res uint32
	//fmt.Scanf("%v\n%v", &a,&b)
	// เปลี่ยนเป็นใส่ ค่าเข้าไปตรงๆเลย
	a = 2
	b = 3
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
