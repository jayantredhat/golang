package main

import "fmt"

func main() {
	fmt.Println("Hello Go, I am starting the learning Go...")
	fmt.Println("..........................................")

}

/*
Go Syntax
A Go file consists of the following parts:
Package declaration
Import packages
Functions
Statements and expressions
Look at the following code, to understand it better:
## Example explained
Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.
Line 2: import ("fmt") lets us import files included in the fmt package.
Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.
Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.
Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".
Note: In Go, any executable code belongs to the main package.
Go Statements
fmt.Println("Hello Go, I am starting the learning Go..") is a statement.
In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".
Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).
The left curly bracket { cannot come at the start of a line.
Run the following code and see what happens:
Go Compact Code
You can write more compact code, like shown below (this is not recommended because it makes the code more difficult to read):
Example
package main; import ("fmt"); func main() { fmt.Println("Hello Go, I am starting the learning Go...");}
*/
