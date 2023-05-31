package observerMode

import (
	"fmt"
)

/*
使用观察者模式改进

以解决一对多的对象依赖关系
*/
const SendButton = -99

type ButtonListener interface {
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
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		d.dialer.EnterDigit(token)
	case SendButton:
		d.dialer.Dial()
	default:
		return fmt.Errorf("unknown button pressed: token=%d", token)
	}
	return nil
}

func (s SendButtonDialerAdapter) ButtonPressed(token int) error {
	if token == SendButton {
		s.dialer.Dial()
	}
	return nil
}

type Button struct {
	token           int
	buttonListeners []ButtonListener
}

func (m *Button) Press() error {
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
	return Button{token: token, buttonListeners: []ButtonListener{}}
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
	sendButton := NewButton(SendButton)
	sendButton.AddListener(NewSendButtonDialerAdapter(dialer))
	return Phone{dialer: dialer, digitButtons: digitButtonsSlice, sendButton: sendButton}
}
