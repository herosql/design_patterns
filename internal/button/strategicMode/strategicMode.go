package strategicMode

import "fmt"

/*
使用策略模式对代码进行改进
将拨号抽象成一个策略,通过接口进行隔离

这时Button符合开闭原则,但Dialer不符合开闭原则
*/
const SendButton = -99

type ButtonServer interface {
	ButtonPressed(token int) error
}

type Dialer struct {
}

func (d Dialer) ButtonPressed(token int) error {
	switch token {
	case 0:
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
		d.EnterDigit(token)
	case SendButton:
		d.Dial()
	default:
		return fmt.Errorf("unknown button pressed: token=%d", token)
	}
	return nil
}

type Button struct {
	token        int
	buttonServer ButtonServer
}

func (m *Button) Press() error {
	return m.buttonServer.ButtonPressed(m.token)
}

func NewButton(token int, buttonServer ButtonServer) Button {
	return Button{token: token, buttonServer: buttonServer}
}

func (b *Button) GetToken() int {
	return b.token
}

func (m Dialer) EnterDigit(digit int) {
	fmt.Printf("enter digit:%d", digit)
}

func (m Dialer) Dial() {
	fmt.Println("dialing...")
}
