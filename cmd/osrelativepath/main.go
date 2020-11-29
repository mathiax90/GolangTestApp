package main
import (	
	"fmt"
	"os"
)
func main(){
	fmt.Println("Hello")
	file1, err1 := os.Create("../test.xml")
	_ = file1
	_ = err1 
}