package synthesis

import (
	"bytes"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/wxxhub/go-speech-synthesis/synthesis/voice"
	"sync"
	"testing"
	"time"
)

type Queue struct {
	streamers []beep.Streamer
}

func (q *Queue) Add(streamers ...beep.Streamer) {
	q.streamers = append(q.streamers, streamers...)
}

func (q *Queue) Stream(samples [][2]float64) (n int, ok bool) {
	// We use the filled variable to track how many samples we've
	// successfully filled already. We loop until all samples are filled.
	filled := 0
	for filled < len(samples) {
		// There are no streamers in the queue, so we stream silence.
		if len(q.streamers) == 0 {
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		// We stream from the first streamer in the queue.
		n, ok := q.streamers[0].Stream(samples[filled:])
		// If it's drained, we pop it from the queue, thus continuing with
		// the next streamer.
		if !ok {
			q.streamers = q.streamers[1:]
		}
		// We update the number of filled samples.
		filled += n
	}
	return len(samples), true
}

func (q *Queue) Err() error {
	return nil
}

func TestDemo(t *testing.T) {
	list := []voice.VOICE{voice.VOICE_NI3, voice.VOICE_HAO3, voice.VOICE_SHI4, voice.VOICE_JIE4}

	//bBuff := beep.NewBuffer(format)

	var queue Queue
	speaker.Play(&queue)
	for _, v := range list {
		b := bytes.NewBuffer(voice.GetVoice(v))
		s, format, _ := wav.Decode(b)

		//c := make(chan struct{})
		//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
		//speaker.Play(beep.Seq(s, beep.Callback(func() {
		//	fmt.Println("done2")
		//	c <- struct{}{}
		//})))
		//
		//<-c

		// And finally, we add the song to the queue.
		resampled := beep.Resample(4, format.SampleRate, format.SampleRate, s)

		speaker.Lock()
		queue.Add(resampled)
		speaker.Unlock()
		//
		//bBuff.Append(s)
	}

	time.Sleep(5 * time.Second)
	//
	//c := make(chan struct{})
	//
	//speaker.Play(beep.Seq(bBuff.Streamer(0, bBuff.Len()), beep.Callback(func() {
	//	fmt.Println("done1")
	//	c <- struct{}{}
	//})))
	//
	//<-c
}

func TestDemo2(t *testing.T) {
	list := []voice.VOICE{voice.VOICE_NI3, voice.VOICE_HAO3, voice.VOICE_SHI4, voice.VOICE_JIE4, voice.VOICE_A}

	b := bytes.NewBuffer(voice.GetVoice(voice.VOICE_NI3))
	_, format, _ := wav.Decode(b)
	bBuff := beep.NewBuffer(format)

	for _, v := range list {
		b := bytes.NewBuffer(voice.GetVoice(v))
		t.Log(b.Len())
		s, _, _ := wav.Decode(b)
		bBuff.Append(s)
	}

	for _, v := range list {
		b := bytes.NewBuffer(voice.GetVoice(v))
		t.Log(b.Len())
		s, _, _ := wav.Decode(b)
		bBuff.Append(s)
	}
	w := sync.WaitGroup{}

	w.Add(1)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/60))
	speaker.Play(beep.Seq(bBuff.Streamer(0, bBuff.Len()), beep.Callback(func() {
		fmt.Println("done1")
		w.Done()
	})))

	w.Wait()
}

func TestDemo3(t *testing.T) {
	list := []voice.VOICE{voice.VOICE_NI3, voice.VOICE_HAO3, voice.VOICE_SHI4, voice.VOICE_JIE4}

	b := bytes.NewBuffer(voice.GetVoice(voice.VOICE_NI3))
	_, format, _ := wav.Decode(b)
	bBuff := beep.NewBuffer(format)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	for _, v := range list {
		b := bytes.NewBuffer(voice.GetVoice(v))
		t.Log(b.Len())
		s, _, _ := wav.Decode(b)
		w := sync.WaitGroup{}
		w.Add(1)
		speaker.Play(beep.Seq(s, beep.Callback(func() {
			fmt.Println("done1")
			w.Done()
		})))

		w.Wait()
		bBuff.Append(s)
	}
}

func TestDemo4(t *testing.T) {
	list := []voice.VOICE{voice.VOICE_NI3, voice.VOICE_HAO3, voice.VOICE_SHI4, voice.VOICE_JIE4, voice.VOICE_A}

	b := bytes.NewBuffer(voice.GetVoice(voice.VOICE_NI3))
	_, format, _ := wav.Decode(b)
	for _, v := range list {
		b.Write(voice.GetVoice(v))
		t.Log(b.Len())

	}

	s, format, err := wav.Decode(b)
	if err != nil {
		t.Fatal(err)
	}

	w := sync.WaitGroup{}
	w.Add(1)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/60))
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		w.Done()
	})))

	w.Wait()
}

func TestSynthesis(t *testing.T) {
	s := InitSynthesis(DefaultSampleRate)
	defer s.Close()
	s.Append("你好世界，你好世界s你好世界")

	for s.Running() {
		time.Sleep(time.Second)
	}
}
