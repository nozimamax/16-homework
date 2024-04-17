package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("storage/mainfile/file.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println("Please enter something: ")
	var input string
	fmt.Scanln(&input)

	// Write the entered string to the file
	_, err = fmt.Fprintln(f, input)
	if err != nil {
		fmt.Println(err)
		return
	}

	namefile := fmt.Sprintf("storage/changes/file_%s.txt", time.Now().Format("2006_01_02_114805"))

	newfile, err := os.Create(namefile)
	if err != nil {
		log.Fatal(err)
	}
	defer newfile.Close()
	file, err := os.Open("storage/mainfile/file.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(newfile, file)
  if err != nil{
    log.Fatal(err)
  }

	fmt.Println("Data written to file successfully!")
}