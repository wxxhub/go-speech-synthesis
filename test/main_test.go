package test

import (
	"bytes"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/wxxhub/go-speech-synthesis/synthesis/voice"
	"io"
	"os"
	"testing"
	"time"
)

func TestMains(t *testing.T) {

	f, err := os.Open("/Users/wxx/GolandProjects/go-speech-synthesis/voice/1.wav")
	if err != nil {
		panic(fmt.Sprintf("couldn't open audio file - %v", err))
	}

	//f2, err := os.Open("/Users/wxx/GolandProjects/go-speech-synthesis/moss2.mp3")
	//if err != nil {
	//	panic(fmt.Sprintf("couldn't open audio file - %v", err))
	//}

	data, _ := io.ReadAll(f)
	s1 := bytes.NewReader(data)
	s2 := bytes.NewReader(data)

	//mossAll, _ := io.ReadAll(f2)
	//mossData := bytes.NewReader(mossAll)
	d, format, err := wav.Decode(s1)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	d2, _, err := wav.Decode(s2)
	if err != nil {
		panic(err)
	}
	defer d2.Close()

	s3 := bytes.NewReader(voice.GetVoice(voice.VOICE_NI3))
	mossd, _, err := wav.Decode(s3)
	if err != nil {
		panic(err)
	}
	defer mossd.Close()

	fmt.Println(d)
	fmt.Println(format)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	//w := sync.WaitGroup{}
	c := make(chan struct{})

	speaker.Play(beep.Seq(mossd, beep.Callback(func() {
		fmt.Println("done1")
		c <- struct{}{}
	})))

	<-c
	//speaker.Init(12000, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(d2, beep.Callback(func() {
		fmt.Println("done2")
		c <- struct{}{}
	})))

	<-c
	//w.Wait()
}
