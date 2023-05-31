package adapterMode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButton(t *testing.T) {
	var buttonServer ButtonServer = NewDigitButtonDialerAdapter(Dialer{})
	button := NewButton(1, buttonServer)
	err := button.Press()
	assert.Equal(t, nil, err, "Press() 的预期结果为 %s", err)
}
