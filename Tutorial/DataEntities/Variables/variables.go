package main

import "fmt"

func main() {

	/*
		Primitive types

		· Boolean (bool)
		· Strings
		· Signed integers (int)
			· 8 bits (int8)
			· 16 bits (int16)
			· 32 bits (int32) == int
			· 64 bits (int64)
		· Unsigned integers (uint)
			· 8 bits (uint8)
			· 16 bits (uint16)
			· 32 bits (uint32) == uint
			· 64 bits (uint64)
		· Floats
			· 32 bits (float32)
			· 64 bits (float64)
		· Complex
			· 64 bits (complex64)
			· 128 bits (complex128)
		· Specials
			· Byte (byte) == int8
			· Rune (rune) == int32 -> It represents a Unicode code point
	*/

	/*
		¡IMPORTANT!

		Go asign zero-valued values by default to every variable declared but not initialized.
		The zero-valued values are different for each data type:

		· Rune = ''
		· Integer = 0
		· Float = 0.0
		· String = ""
		· Complex = 0 + 0i
		· Boolean = False
	*/

	/* Examples */

	// To declare a variable we have to use 'var' following this sintaxis -> var var_name type
	var my_rune rune
	var my_string string
	var my_complex complex64

	fmt.Printf("Default value for my_rune: %c\n", my_rune)
	fmt.Printf("Default value for my_string: %s\n", my_string)
	fmt.Printf("Default value for my_complex: %g\n", my_complex)

	// Obvioysly we can modify the value of a variable
	my_rune = 'F'
	my_complex = -54 + 3i
	my_string = "A simple stirng"

	fmt.Printf("New my_rune value: %c\n", my_rune)
	fmt.Printf("New my_string value: %s\n", my_string)
	fmt.Printf("New my_complex value: %g\n", my_complex)

	// It is possible to initialize a variable inline
	var my_bool bool = false

	// Moreover we can create a variable without specify the type, Go infers it value
	var a_variable = 12.6543

	fmt.Printf("my_bool: %t\n", my_bool)
	fmt.Printf("a_variable: %f\n", a_variable)
}
