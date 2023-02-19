package synthesis

import (
	"bytes"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/wxxhub/go-speech-synthesis/player"
	"github.com/wxxhub/go-speech-synthesis/synthesis/voice"
	"regexp"
)

type Synthesis struct {
	running    bool
	sampleRate beep.SampleRate

	pinyin *voice.PinYin
	player *player.Player
}

func InitSynthesis(sampleRate int) *Synthesis {

	return &Synthesis{
		sampleRate: beep.SampleRate(sampleRate),
		player:     player.InitPlayer(sampleRate),
		pinyin:     voice.InitPinYin(),
	}
}

func (s *Synthesis) Append(content string) {

	sep := "[s,！？｡＂＃＄％＆＇（）＊＋，－／：；＜＝＞＠［＼］＾＿｀｛｜｝～｟｠｢｣､、〃》「」『』【】〔〕〖〗〘〙〚〛〜〝〞〟〰〾〿–—‘’‛“”„‟…‧﹏.]+"
	reg, _ := regexp.Compile(sep)

	//lines := strings.Split(content, sep)
	lines := reg.Split(content, -1)
	fmt.Println(lines)
	for _, line := range lines {
		fmt.Println(s.pinyin.Pinyin(line))
		s.player.Append(s.getBuffer(s.pinyin.Pinyin(line)))
	}

}

func (s *Synthesis) Running() bool {
	return s.running
}

func (s *Synthesis) Close() {

}

func (s *Synthesis) getBuffer(pinyinList []string) *beep.Buffer {
	b := beep.NewBuffer(beep.Format{SampleRate: s.sampleRate})
	for _, pinyin := range pinyinList {
		vb := bytes.NewBuffer(voice.GetVoice(voice.VOICE(pinyin)))

		stream, _, err := wav.Decode(vb)
		if err == nil {
			b.Append(stream)
		} else {
			fmt.Println(err)
		}

	}

	return b
}
