package main

import (
	"github.com/euroledger/my-logger2"
)

func main() {
	mylogger.LogInfo("This is an info message")
	mylogger.LogWarning("warning message")

	mylogger.LogError("Its broke!")

}