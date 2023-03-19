package synthesis

import (
	"bytes"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/wxxhub/go-speech-synthesis/synthesis/voice"
	"github.com/wxxhub/go-speech-synthesis/worker/worker_interface"
	"regexp"
	"sync"
)

type Synthesis struct {
	sampleRate beep.SampleRate

	pinyin *voice.PinYin
	worker worker_interface.Worker

	count int

	runningMux sync.RWMutex
}

func InitSynthesis(sampleRate int, w worker_interface.Worker) *Synthesis {
	s := &Synthesis{
		sampleRate: beep.SampleRate(sampleRate),
		pinyin:     voice.InitPinYin(),
		count:      0,
	}

	w.SetFinishCallBack(s.call)
	s.worker = w
	return s
}

func (s *Synthesis) call() {
	s.runningMux.Lock()
	s.count -= 1
	s.runningMux.Unlock()
}

func (s *Synthesis) addRun() {
	s.runningMux.Lock()
	s.count += 1
	s.runningMux.Unlock()
}

func (s *Synthesis) Append(content string) {

	sep := "[s,！？｡＂＃＄％＆＇（）＊＋，－／：；＜＝＞＠［＼］＾＿｀｛｜｝～｟｠｢｣､、〃》「」『』【】〔〕〖〗〘〙〚〛〜〝〞〟〰〾〿–—‘’‛“”„‟…‧﹏.]+"
	reg, _ := regexp.Compile(sep)

	//lines := strings.Split(content, sep)
	lines := reg.Split(content, -1)
	fmt.Println(lines)
	for _, line := range lines {
		//fmt.Println(s.pinyin.Pinyin(line))
		s.addRun()
		s.worker.Append(s.getBuffer(s.pinyin.Pinyin(line)))
	}

}

func (s *Synthesis) Running() bool {
	s.runningMux.RLock()
	defer s.runningMux.RUnlock()
	return s.count > 0
}

func (s *Synthesis) Close() {
	s.worker.Close()
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

func (s *Synthesis) Delivery() *bytes.Buffer {
	return nil
}
