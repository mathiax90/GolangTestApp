package main
import (
	"log"
	"fmt"
	"os"
)
func main(){
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
	log.Println("Hello world!")
	log.Fatalln("Fatal!")
	fmt.Println("Hello")
}