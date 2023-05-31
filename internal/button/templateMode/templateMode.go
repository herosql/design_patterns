package templateMode

import (
	"fmt"
)

/*
使用模板方法模式实现开闭原则

如果业务要求按下按钮的时候，除了控制设备，按钮本身还需要执行一些操作，完成一些成员变量的状态更改，不同按钮类型进行的操作和记录状态各不相同。按照当前设计可能又要在 Button 的 press 方法中增加 switch/case 了。
*/
const SendButtonNumber = -99

type ButtonListener interface {
	ButtonPressed(token int) error
}

type OnPressInter interface {
	OnPress()
}

type SendButton struct {
	Button
}

func (s SendButton) OnPress() {
	println(".........OnPress..........")
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
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		d.dialer.EnterDigit(token)
	case SendButtonNumber:
		d.dialer.Dial()
	default:
		return fmt.Errorf("unknown button pressed: token=%d", token)
	}
	return nil
}

func (s SendButtonDialerAdapter) ButtonPressed(token int) error {
	if token == SendButtonNumber {
		s.dialer.Dial()
	}
	return nil
}

type Button struct {
	OnPressInter
	token           int
	buttonListeners []ButtonListener
}

func (m *Button) Press() error {
	m.OnPress()
	for _, buttonListener := range m.buttonListeners {
		err := buttonListener.ButtonPressed(m.token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Button) AddListener(buttonListener ButtonListener) {
	m.buttonListeners = append(m.buttonListeners, buttonListener)
}

func NewButton(token int) Button {
	return Button{token: token, buttonListeners: []ButtonListener{}, OnPressInter: SendButton{}}
}

func (b *Button) GetToken() int {
	return b.token
}

func (m Dialer) EnterDigit(digit int) {
	fmt.Printf("enter digit:%d\n", digit)
}

func (m Dialer) Dial() {
	fmt.Println("dialing...")
}

type Phone struct {
	dialer       Dialer
	digitButtons []Button
	sendButton   Button
}

func NewPhone() Phone {
	dialer := Dialer{}
	var digitButtons [10]Button

	for i := 0; i < len(digitButtons); i++ {
		digitButtons[i] = NewButton(i)
		digitButtons[i].AddListener(NewDigitButtonDialerAdapter(dialer))
	}

	digitButtonsSlice := digitButtons[:]
	sendButton := NewButton(SendButtonNumber)
	sendButton.AddListener(NewSendButtonDialerAdapter(dialer))
	return Phone{dialer: dialer, digitButtons: digitButtonsSlice, sendButton: sendButton}
}
