package button

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButton(t *testing.T) {
	var dialer = Dialer{}
	var button = NewButton(dialer, 1)
	err := button.Press()
	assert.Equal(t, nil, err, "Press() 的预期结果为 %s", err)
}
