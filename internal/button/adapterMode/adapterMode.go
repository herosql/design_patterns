package adapterMode

import "fmt"

/*
使用适配器模式对代码进行改进

这时,Button和Dialer都符合开闭原则.

如果要求能够用一个按钮控制多个设备，比如按钮按下进行拨号的同时，还需要扬声器根据不同按钮发出不同声音，将来还需要根据不同按钮点亮不同颜色的灯。按照当前设计，可能需要在适配器中调用多个设备，增加设备要修改适配器代码，又不符合开闭原则了。
*/
const SendButton = -99

type ButtonServer interface {
	ButtonPressed(token int) error
}

type Dialer struct {
}

type DigitButtonDialerAdapter struct {
	dialer Dialer
}

type SendButtonDialerAdapter struct {
	dialer Dialer
}

func NewDigitButtonDialerAdapter(dialer Dialer) DigitButtonDialerAdapter {
	return DigitButtonDialerAdapter{dialer: dialer}
}

func NewSendButtonDialerAdapter(dialer Dialer) SendButtonDialerAdapter {
	return SendButtonDialerAdapter{dialer: dialer}
}

func (d DigitButtonDialerAdapter) ButtonPressed(token int) error {
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
		d.dialer.EnterDigit(token)
	case SendButton:
		d.dialer.Dial()
	default:
		return fmt.Errorf("unknown button pressed: token=%d", token)
	}
	return nil
}

func (s SendButtonDialerAdapter) ButtonPressed(token int) error {
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
