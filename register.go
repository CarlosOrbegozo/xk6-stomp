package stomp

import (
	"github.com/CarlosOrbegozo/stomp"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/stomp", stomp.New())
}
