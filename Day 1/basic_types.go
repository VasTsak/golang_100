package main 

import (
	"fmt"
	"math"
)

var (
	v_bool bool = false
	v_string string = "I am a string variable."
	// int -> signed integer, which means they can take both positive and negative values
	// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	v_int int = -5 // 16 bits, it cannot take numbers greater than 8 digits, (1111111111111111 -> 65536)
	// you should be mindful when you define how many bits you require for the integer to have because 
	// you may waste memory or run out of memory.
	v_int8 int8 = -5 // 8 bits, it cannot take numbers greater than 8 digits, (11111111 -> 256)
	v_int16 int16 = -5 // 16 bits, it cannot take numbers greater than 8 digits, (1111111111111111 -> 65536)
	v_int32 int32 = -5 // 32 bits, it cannot take numbers greater than 32 digits, (11111111111111111111111111111111 -> 4294967296)
	v_int64 int64 = -5 // 64 bits, it cannot take numbers greater than 32 digits, (1111111111111111111111111111111111111111111111111111111111111111 -> 18446744073709551616)
	// uint -> unsigned integer, which means they can take only positive or zero.
	v_uint  uint 5 // 16bits
	// you should be mindful when you define how many bits you require for the integer to have because 
	// you may waste memory or run out of memory.
	v_int8 uint8 = 5 // 8 bits
	v_int16 uint16 = 5 // 16 bits
	v_int32 uint32 = 5 // 32 bits
	v_int64 uint64 = 5 // 64 bits

	v_byte byte = 15 // alias for uint8
	v_rune rune = 123124124 // alias for int32, represents a Unicode code point

	v_float32 float32 = 1234.56789
	v_float64 float64 = 1234.56789

	v_complx64 complex64 = math.cmplx.Sqrt(-5 + 12i)
	v_complx128 complex128 = math.cmplx.Sqrt(-5 + 12i)

)

func main() {
	fmt.Printf("Type: %T Value: %v\n", v_bool, v_bool)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
