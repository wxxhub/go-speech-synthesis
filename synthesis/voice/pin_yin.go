package voice

import (
	"github.com/go-ego/gpy"
)

type PinYin struct {
	args gpy.Args
}

func InitPinYin() *PinYin {
	a := gpy.NewArgs()
	a.Style = gpy.Tone3

	return &PinYin{
		args: a,
	}
}

func (p *PinYin) Pinyin(content string) []string {
	return gpy.LazyPinyin(content, p.args)
}
