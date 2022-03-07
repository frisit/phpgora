package main

import (
	"fmt"
	"io"
	"os"
	"time"
	// "github.com/liudng/godump"
)

func saveToFile(text string, fileName string) {
	/*
	   \n: переход на новую строку
	   \r: возврат каретки
	   \t: табуляция
	   \": двойная кавычка внутри строк
	   \\: обратный слеш
	*/
	timeNow := time.Now()

	//	file, err := os.Create(fileName)

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(timeNow.Format("2006-01-02 15:04:05") + " | " + " " + text)
	//	godump.Dump(timeNow)

	fmt.Println("Done.")

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}

}
