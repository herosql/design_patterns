package button

import "fmt"

/*

需要设计一个可以通过按钮拨号的电话，核心对象是按钮和拨号器
当我们想要增加按钮类型的时候，比如，当我们需要按钮支持星号（*）和井号（#）的时候，
我们必须修改 Button 类代码；当我们想要用这个按钮控制一个密码锁而不是拨号器的时候，
因为按钮关联了拨号器，所以依然要修改 Button 类代码；当我们想要按钮控制多个设备的时候，还是要修改 Button 类代码。

耦合性较高
*/

const SendButton = -99

type Dialer struct {
}

type Button struct {
	dialer Dialer
	token  int
}

func (m *Button) Press() error {
	switch m.token {
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
		m.dialer.EnterDigit(m.token)
	case SendButton:
		m.dialer.Dial()
	default:
		return fmt.Errorf("unknown button pressed: token=%d", m.token)
	}
	return nil
}

func (m *Dialer) EnterDigit(digit int) {
	fmt.Printf("enter digit:%d", digit)
}

func (m *Dialer) Dial() {
	fmt.Println("dialing...")
}
func NewButton(dialer Dialer, token int) Button {
	return Button{dialer: dialer, token: token}
}
