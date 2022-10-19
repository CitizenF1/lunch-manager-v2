package main

import (
	"lunch-manager/internal/bot"
)

func main() {
	// ====Route log to file====
	// f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	//Bot Start
	bot.StartBot()
}
