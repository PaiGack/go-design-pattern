package command

// 做饭指令接口
type CookCommand interface {
	Execute() string
}

// 蒸饭指令
type stramRiceCommand struct {
	electricCooker *ElectricCooker
}

func NewStreamRiceCommand(electricCooker *ElectricCooker) *stramRiceCommand {
	return &stramRiceCommand{
		electricCooker: electricCooker,
	}
}

func (s *stramRiceCommand) Execute() string {
	s.electricCooker.SetFire("中")
	s.electricCooker.SetPressure("正常")
	return "蒸饭：" + s.electricCooker.Run("30分钟")
}

type cookCongeeCommand struct {
	electricCooker *ElectricCooker
}

func NewCookCongeeCommand(electricCooker *ElectricCooker) *cookCongeeCommand {
	return &cookCongeeCommand{
		electricCooker: electricCooker,
	}
}

func (s *cookCongeeCommand) Execute() string {
	s.electricCooker.SetFire("大")
	s.electricCooker.SetPressure("强")
	return "煮粥：" + s.electricCooker.Run("45分钟")
}

// 停止指令
type shutdownCommand struct {
	electricCooker *ElectricCooker
}

func NewShutdownCommand(electricCooker *ElectricCooker) *shutdownCommand {
	return &shutdownCommand{
		electricCooker: electricCooker,
	}
}

func (s *shutdownCommand) Execute() string {
	return s.electricCooker.Shutdown()
}

// 电饭煲指令触发器
type ElectricCookerInvoker struct {
	cookCommand CookCommand
}

// 设置指令
func (e *ElectricCookerInvoker) SetCookCommand(cookCommand CookCommand) {
	e.cookCommand = cookCommand
}

// 执行指令
func (e *ElectricCookerInvoker) ExecuteCookCommand() string {
	return e.cookCommand.Execute()
}
