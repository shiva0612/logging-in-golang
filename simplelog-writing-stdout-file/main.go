package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, _ := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	logwriter := io.MultiWriter(f, os.Stdout)

	log.SetOutput(logwriter)
	log.SetPrefix("myapp: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Panic("im panic")
	log.Println("im ok")
	log.Fatalln("im fatal")
	fmt.Println("main end...")

}
