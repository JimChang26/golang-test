package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
測試用專案
*/
func main() {

	b()
	/*
		fmt.Println(add(2, 3))
		fmt.Println(double(4))
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			k := strconv.Itoa(3)
			c.JSON(200, gin.H{
				"message": "pong" + k,
			})
		})
		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	*/
	wg := new(sync.WaitGroup)
	wg.Add(1)
	fmt.Println(addTest(5))
	go safelyDo(wg)
	wg.Wait()

	fmt.Println(calculatedArea(5, -4))

	area, _ := calculatedArea(10, 10)

	fmt.Println(area)

	var a [4]int
	fmt.Println("a")
	fmt.Println(a)
	fmt.Println("&a")
	fmt.Println(&a)
	fmt.Println(&a[0])
	fmt.Println(&a[1])
	fmt.Println(&a[2])
	fmt.Println(&a[3])

	var b = new([4]int)
	var c []int
	fmt.Println("&B")
	fmt.Println(&b)
	fmt.Println("B")
	fmt.Println(b)
	fmt.Println("C")
	fmt.Println(c)

	test(&a)

	fmt.Println(a)

	testDefer()
}

func test(pointer *[4]int) {
	fmt.Println(&pointer)
	fmt.Println(&pointer[0])
	pointer[0] = 5
}

func testDefer() {

	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%d ", i)
		}()
	}
	fmt.Println("for end")
}

func calculatedArea(length int, width int) (area int, fail error) {

	defer func() {
		area = -1
	}()
	if length <= 0 || width <= 0 {
		fail = errors.New("長寬必須大於0")
		return
	}
	area = length * width
	return
}

func safelyDo(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("work failed:", err)
		}
		wg.Done()
	}()
	do()
}
func do() {
	fmt.Println("do do do do do")
	//panic("die?")
}

func addTest(a int) (returnValue int) {
	defer func() {
		returnValue = 10
	}()
	return a
}

func add(a int, b int) int {
	return a + b
}

func double(a int) int {
	return a * a
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
