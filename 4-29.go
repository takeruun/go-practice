package main

import (
	"io"
	"log"
	"os"
)

func LoggingSetting() {
	logfile, _ := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	muitiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.SetOutput(muitiLogFile)
}

func main() {
	LoggingSetting()
	log.Println("Hello, World!")

	_, err := os.Open("non-existing")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Fatalf("Fatal")
}
