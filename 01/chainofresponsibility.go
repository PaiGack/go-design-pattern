package chainofresponsibility

import "fmt"

type BoardingProcessor interface {
	SetNextProcessor(processor BoardingProcessor)
	ProcessFor(passenger *Passenger)
}

type Passenger struct {
	name                  string
	hasBoardingPass       bool
	hasLuggage            bool
	isPassIdentityCheck   bool
	isPassSecurityCheck   bool
	isCompleteForBoarding bool
}

// 登机流程处理器基类
type baseBoardingProcessor struct {
	nextProcess BoardingProcessor
}

// SetNextProcessor 基类中统一实现设置下一个处理器方法
func (b *baseBoardingProcessor) SetNextProcessor(processor BoardingProcessor) {
	b.nextProcess = processor
}

// ProcessFor 基类中统一实现下一个处理器流转
func (b *baseBoardingProcessor) ProcessFor(passenger *Passenger) {
	if b.nextProcess != nil {
		b.nextProcess.ProcessFor(passenger)
	}
}

// boardingPassProcessor 办理登机牌处理器
type boardingPassProcessor struct {
	baseBoardingProcessor
}

func (b *boardingPassProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("===> 为旅客: %s，办理登机牌\n", passenger.name)
		passenger.hasBoardingPass = true
	}

	// 成功办理登机牌后，进入下一个流程处理
	b.baseBoardingProcessor.ProcessFor(passenger)
}

// luggageCheckInProcessor 托运行李处理器
type luggageCheckInProcessor struct {
	baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("===> 旅客: %s，未办理登机牌，不能托运行李\n", passenger.name)
		return
	}

	if passenger.hasLuggage {
		fmt.Printf("===> 为旅客: %s，办理行李托运\n", passenger.name)
	}

	// 进入下一个流程处理
	l.baseBoardingProcessor.ProcessFor(passenger)
}

// identityCheckProcessor 校验身份处理器
type identityCheckProcessor struct {
	baseBoardingProcessor
}

func (i *identityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("===> 旅客: %s，未办理登机牌，不能办理身份校验\n", passenger.name)
		return
	}

	if !passenger.isPassIdentityCheck {
		fmt.Printf("===> 为旅客: %s，核实身份信息\n", passenger.name)
		passenger.isPassIdentityCheck = true
	}

	// 进入下一个流程处理
	i.baseBoardingProcessor.ProcessFor(passenger)
}

// securityCheckProcessor 安检处理器
type securityCheckProcessor struct {
	baseBoardingProcessor
}

func (s *securityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("===> 旅客: %s，未办理登机牌，不能进行安检\n", passenger.name)
		return
	}

	if !passenger.isPassSecurityCheck {
		fmt.Printf("===> 为旅客: %s，进行安检\n", passenger.name)
		passenger.isPassSecurityCheck = true
	}

	// 进入下一个流程处理
	s.baseBoardingProcessor.ProcessFor(passenger)
}

type completeBoardingProcessor struct {
	baseBoardingProcessor
}

func (c *completeBoardingProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass ||
		!passenger.isPassIdentityCheck ||
		!passenger.isPassSecurityCheck {
		fmt.Printf("===> 旅客: %s，登机检查过程未完成，不能登机\n", passenger.name)
		return
	}

	passenger.isCompleteForBoarding = true
	fmt.Printf("===> 旅客: %s，成功登机\n", passenger.name)
}
