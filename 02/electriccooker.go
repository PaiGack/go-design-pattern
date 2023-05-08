package command

import "fmt"

type ElectricCooker struct {
	fire     string
	pressure string
}

// 设置火力
func (e *ElectricCooker) SetFire(fire string) {
	e.fire = fire
}

// 设置压力
func (e *ElectricCooker) SetPressure(pressure string) {
	e.pressure = pressure
}

// 持续运行指定时间
func (e *ElectricCooker) Run(duration string) string {
	return fmt.Sprintf("电饭煲设置火力为：%s，压力为：%s，持续运行：%s", e.fire, e.pressure, duration)
}

// 停止
func (e *ElectricCooker) Shutdown() string {
	return "电饭煲停止运行"
}
