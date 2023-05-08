package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	electricCooker := new(ElectricCooker)
	electricCookerInvoker := new(ElectricCookerInvoker)

	// 蒸饭
	stramRiceCommand := NewStreamRiceCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(stramRiceCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 煮粥
	cookCongeeCommand := NewCookCongeeCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(cookCongeeCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 停止
	shutdownCommand := NewShutdownCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(shutdownCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	t.Fail()
}
