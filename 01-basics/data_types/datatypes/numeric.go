package datatypes

import (
	"fmt"
	"math"
	"unsafe"
)

// NumericExample 展示Go中的数值类型
func NumericExample() {
	// 整数类型
	fmt.Println("=== 整数类型 ===")

	// 有符号整数
	var i8 int8 = 127                   // -128 到 127
	var i16 int16 = 32767               // -32768 到 32767
	var i32 int32 = 2147483647          // -2147483648 到 2147483647
	var i64 int64 = 9223372036854775807 // -9223372036854775808 到 9223372036854775807
	var i int = 42                      // 32位系统为int32，64位系统为int64

	// 无符号整数
	var ui8 uint8 = 255                    // 0 到 255
	var ui16 uint16 = 65535                // 0 到 65535
	var ui32 uint32 = 4294967295           // 0 到 4294967295
	var ui64 uint64 = 18446744073709551615 // 0 到 18446744073709551615
	var ui uint = 42                       // 32位系统为uint32，64位系统为uint64

	// 特殊整数类型
	var uptr uintptr = 0xc82000c290 // 足够存放指针的无符号整数类型

	// 打印整数类型的值和大小
	fmt.Printf("int8: %d, 大小: %d 字节\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("int16: %d, 大小: %d 字节\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("int32: %d, 大小: %d 字节\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("int64: %d, 大小: %d 字节\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("int: %d, 大小: %d 字节\n", i, unsafe.Sizeof(i))

	fmt.Printf("uint8: %d, 大小: %d 字节\n", ui8, unsafe.Sizeof(ui8))
	fmt.Printf("uint16: %d, 大小: %d 字节\n", ui16, unsafe.Sizeof(ui16))
	fmt.Printf("uint32: %d, 大小: %d 字节\n", ui32, unsafe.Sizeof(ui32))
	fmt.Printf("uint64: %d, 大小: %d 字节\n", ui64, unsafe.Sizeof(ui64))
	fmt.Printf("uint: %d, 大小: %d 字节\n", ui, unsafe.Sizeof(ui))

	fmt.Printf("uintptr: %x, 大小: %d 字节\n", uptr, unsafe.Sizeof(uptr))

	// 浮点类型
	fmt.Println("\n=== 浮点类型 ===")

	var f32 float32 = 3.14159265358979323846
	var f64 float64 = 3.14159265358979323846

	// 打印浮点类型的值、精度和大小
	fmt.Printf("float32: %.10f, 大小: %d 字节\n", f32, unsafe.Sizeof(f32))
	fmt.Printf("float64: %.20f, 大小: %d 字节\n", f64, unsafe.Sizeof(f64))

	// 浮点数特殊值
	fmt.Printf("float64 最大值: %g\n", math.MaxFloat64)
	fmt.Printf("float64 最小值: %g\n", math.SmallestNonzeroFloat64)
	fmt.Printf("正无穷: %f\n", math.Inf(1))
	fmt.Printf("负无穷: %f\n", math.Inf(-1))
	fmt.Printf("非数值 (NaN): %f\n", math.NaN())

	// 复数类型
	fmt.Println("\n=== 复数类型 ===")

	var c64 complex64 = complex(float32(1.5), float32(2.5)) // 由两个float32组成
	var c128 complex128 = complex(2.5, 3.5)                 // 由两个float64组成

	// 打印复数的值、实部、虚部和大小
	fmt.Printf("complex64: %v, 大小: %d 字节\n", c64, unsafe.Sizeof(c64))
	fmt.Printf("complex64 的实部: %f, 虚部: %f\n", real(c64), imag(c64))

	fmt.Printf("complex128: %v, 大小: %d 字节\n", c128, unsafe.Sizeof(c128))
	fmt.Printf("complex128 的实部: %f, 虚部: %f\n", real(c128), imag(c128))

	// 复数运算
	var c1 = complex(2, 3)
	var c2 = complex(4, 5)

	// 加法
	fmt.Printf("(%v) + (%v) = %v\n", c1, c2, c1+c2)

	// 减法
	fmt.Printf("(%v) - (%v) = %v\n", c1, c2, c1-c2)

	// 乘法
	fmt.Printf("(%v) * (%v) = %v\n", c1, c2, c1*c2)

	// 除法
	fmt.Printf("(%v) / (%v) = %v\n", c1, c2, c1/c2)

	// 整数运算和类型转换
	fmt.Println("\n=== 整数运算和类型转换 ===")

	a := 10
	b := 3

	fmt.Printf("%d + %d = %d\n", a, b, a+b)  // 加法
	fmt.Printf("%d - %d = %d\n", a, b, a-b)  // 减法
	fmt.Printf("%d * %d = %d\n", a, b, a*b)  // 乘法
	fmt.Printf("%d / %d = %d\n", a, b, a/b)  // 整数除法
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // 取余

	// 位运算
	fmt.Printf("%d & %d = %d\n", a, b, a&b)   // 按位与
	fmt.Printf("%d | %d = %d\n", a, b, a|b)   // 按位或
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)   // 按位异或
	fmt.Printf("%d << %d = %d\n", a, b, a<<b) // 左移
	fmt.Printf("%d >> %d = %d\n", a, b, a>>b) // 右移

	// 类型转换
	var intVal int = 42
	var float64Val float64 = float64(intVal)
	var uint8Val uint8 = uint8(intVal)
	var int32Val int32 = int32(float64(intVal))

	fmt.Printf("int 值: %d\n", intVal)
	fmt.Printf("转换为 float64: %f\n", float64Val)
	fmt.Printf("转换为 uint8: %d\n", uint8Val)
	fmt.Printf("转换为 int32: %d\n", int32Val)

	// 注意：小类型转向大类型安全，大类型转向小类型可能丢失精度或溢出
	var bigVal int32 = 1000000
	var smallVal uint8 = uint8(bigVal) // 溢出

	fmt.Printf("大值: %d, 转为uint8后: %d (发生溢出)\n", bigVal, smallVal)
}
