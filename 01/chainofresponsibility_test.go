package chainofresponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	boardingProcessor := buildBoardingProcessorChain()
	passenger := newPassenger("李四", true)
	boardingProcessor.ProcessFor(passenger)
	if !passenger.isCompleteForBoarding {
		t.Fatalf("乘客：%s，登机流程未完成", passenger.name)
	}
}

func newPassenger(name string, hasLuggage bool) *Passenger {
	return &Passenger{
		name:       name,
		hasLuggage: hasLuggage,
	}
}

func buildBoardingProcessorChain() BoardingProcessor {
	completeBoardingNode := &completeBoardingProcessor{}

	securityCheckNode := &securityCheckProcessor{}
	securityCheckNode.SetNextProcessor(completeBoardingNode)

	identityCheckNode := &identityCheckProcessor{}
	identityCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckInNode := &luggageCheckInProcessor{}
	luggageCheckInNode.SetNextProcessor(identityCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckInNode)

	return boardingPassNode
}
