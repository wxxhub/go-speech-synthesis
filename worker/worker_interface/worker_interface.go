package worker_interface

import (
	"bytes"
	"github.com/faiface/beep"
)

type Worker interface {
	Append(b *beep.Buffer)
	SetFinishCallBack(callFunc func())
	Close()
	Delivery() *bytes.Buffer
}
