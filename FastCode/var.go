package FastCode

//package main
//
//import "fmt"
//
//func main() {
//	// 1. 基本数据类型
//	var b bool = true               // 布尔型
//	var i int = 42                  // 整型
//	var i8 int8 = -128              // 8位有符号整型
//	var i16 int16 = 32767           // 16位有符号整型
//	var i32 int32 = 2147483647      // 32位有符号整型
//	var i64 int64 = 9223372036854775807 // 64位有符号整型
//	var u8 uint8 = 255              // 8位无符号整型
//	var u16 uint16 = 65535          // 16位无符号整型
//	var u32 uint32 = 4294967295     // 32位无符号整型
//	var u64 uint64 = 18446744073709551615 // 64位无符号整型
//	var uintptr uintptr = 0x12345   // 指针类型
//	var f32 float32 = 3.1415926     // 32位浮点型
//	var f64 float64 = 3.141592653589793 // 64位浮点型
//	var c complex64 = 1 + 2i        // 64位复数类型
//	var c128 complex128 = 3 + 4i    // 128位复数类型
//	var b1 byte = 'A'               // 字节类型（uint8别名）
//	var r rune = '中'               // Unicode码点（int32别名）
//	var s string = "Hello, Go"      // 字符串类型
//
//	// 2. 数组类型
//	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}       // 一维数组
//	var arr2 [3][4]int = [3][4]int{{1,2}, {3,4,5}} // 二维数组
//	var arr3 [2][2][2]int = [2][2][2]int{{{1}, {2}}, {{3}, {4}}} // 三维数组
//	var arr4 [5]string = [5]string{"a", "b"}      // 字符串数组
//
//	// 3. 切片类型
//	var slice1 []int = []int{1, 2, 3}             // 基本切片
//	var slice2 []string = make([]string, 3, 5)    // 使用make创建切片（长度3，容量5）
//	var slice3 []int = arr1[1:3]                  // 从数组创建切片
//	var slice2D [][]int = [][]int{{1, 2}, {3, 4}} // 二维切片
//
//	// 4. 映射类型
//	var map1 map[string]int = map[string]int{"one": 1, "two": 2} // 基本map
//	var map2 map[int]string = make(map[int]string, 10)           // 使用make创建map
//	var map3 map[string][]int = map[string][]int{"nums": {1,2,3}} // 嵌套切片的map
//
//	// 5. 指针类型
//	var ptr1 *int = &i                            // 指向int的指针
//	var ptr2 *string = &s                         // 指向string的指针
//	var ptr3 **int = &ptr1                        // 二级指针
//
//	// 6. 结构体类型
//	type Person struct {
//		Name string
//		Age  int
//	}
//	var p1 Person = Person{"Alice", 30}           // 结构体变量
//	var p2 *Person = &Person{"Bob", 25}           // 结构体指针
//
//	// 7. 接口类型
//	type Shape interface {
//		Area() float64
//	}
//	var shape Shape                               // 接口变量（初始为nil）
//
//	// 8. 函数类型
//	var add func(int, int) int                    // 函数类型变量
//	add = func(a, b int) int { return a + b }     // 赋值函数
//
//	// 打印部分变量示例
//	fmt.Println("基本类型示例:", i, s, r)
//	fmt.Println("一维数组:", arr1)
//	fmt.Println("二维数组:", arr2)
//	fmt.Println("切片:", slice1)
//	fmt.Println("映射:", map1)
//	fmt.Println("结构体:", p1)
//	fmt.Println("函数调用结果:", add(2, 3))
//}
//
