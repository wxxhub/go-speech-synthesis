package synthesis

import (
	"bytes"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/wxxhub/go-speech-synthesis/player"
	"github.com/wxxhub/go-speech-synthesis/synthesis/voice"
	"log"
	"regexp"
	"sync"
)

type Synthesis struct {
	sampleRate beep.SampleRate

	pinyin *voice.PinYin
	player *player.Player

	count int

	runningMux sync.RWMutex
}

func InitSynthesis(sampleRate int) *Synthesis {
	s := &Synthesis{
		sampleRate: beep.SampleRate(sampleRate),
		pinyin:     voice.InitPinYin(),
		count:      0,
	}

	s.player = player.InitPlayer(DefaultSampleRate, s.call)
	return s
}

func (s *Synthesis) call() {
	log.Println("runningChan receive")
	s.runningMux.Lock()
	s.count -= 1
	s.runningMux.Unlock()
}

func (s *Synthesis) addRun() {
	s.runningMux.Lock()
	s.count += 1
	s.runningMux.Unlock()
	log.Println(s.count)
}

func (s *Synthesis) Append(content string) {

	sep := "[s,！？｡＂＃＄％＆＇（）＊＋，－／：；＜＝＞＠［＼］＾＿｀｛｜｝～｟｠｢｣､、〃》「」『』【】〔〕〖〗〘〙〚〛〜〝〞〟〰〾〿–—‘’‛“”„‟…‧﹏.]+"
	reg, _ := regexp.Compile(sep)

	//lines := strings.Split(content, sep)
	lines := reg.Split(content, -1)
	fmt.Println(lines)
	for _, line := range lines {
		fmt.Println(s.pinyin.Pinyin(line))
		s.addRun()
		s.player.Append(s.getBuffer(s.pinyin.Pinyin(line)))
	}

}

func (s *Synthesis) Running() bool {
	s.runningMux.RLock()
	defer s.runningMux.RUnlock()
	return s.count > 0
}

func (s *Synthesis) Close() {
	s.player.Close()
}

func (s *Synthesis) getBuffer(pinyinList []string) *beep.Buffer {
	b := beep.NewBuffer(beep.Format{SampleRate: s.sampleRate, NumChannels: GlobalNumChannels, Precision: GlobalPrecision})
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
