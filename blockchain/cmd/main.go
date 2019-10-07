package main
import "blockchain/core"
func main() {
	bc:= demochain.NewBlockchain()
	for i:=1;i<100;i++{
	bc.SendData("Send 1 Coin to Jacky")
	bc.Print()
	}
}