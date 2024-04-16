package main

import (
	"ReAce/CLI"
	"ReAce/Logger"
	"ReAce/Server"
	"fmt"
)

func banner() {
	fmt.Println(Logger.GreenText("     _                       ____   ____  \n    / \\      ___    ___     / ___| |___ \\ \n   / _ \\    / __|  / _ \\   | |       __) |\n  / ___ \\  | (__  |  __/   | |___   / __/ \n /_/   \\_\\  \\___|  \\___|    \\____| |_____|\n                                          "))

	fmt.Println(Logger.RedText("@tupler"))
	fmt.Println(Logger.YellowText("				No Game No Life~"))

}

func init() {
	banner()
}
func main() {
	Logger.ALogger = Logger.NewLogger(false)
	temp := Server.NewListeners()
	err := temp.Add("test", Server.LISTEN_TCP, "", "5555")
	if err != nil {
		fmt.Println(err)
	}
	err = temp.Add("aaaaaa", Server.LISTEN_TCP, "", "1234")
	if err != nil {
		fmt.Println(err)
	}
	Server.Server = Server.NewAceServer(temp)
	CLI.Run()
}
