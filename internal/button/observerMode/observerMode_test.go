package observerMode

import (
	"testing"
)

func TestButton(t *testing.T) {
	phone := NewPhone()
	phone.digitButtons[9].Press()
	phone.digitButtons[1].Press()
	phone.digitButtons[1].Press()
	phone.sendButton.Press()
}
