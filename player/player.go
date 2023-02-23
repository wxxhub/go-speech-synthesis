package player

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"sync"
	"time"
)

type Player struct {
	cancel chan struct{}

	input       chan *beep.Buffer
	runningChan chan struct{}

	callFunc func()
}

func InitPlayer(sampleRate int, callFunc func()) *Player {
	s := beep.SampleRate(sampleRate)
	speaker.Init(s, s.N(time.Millisecond))
	p := &Player{
		cancel:   make(chan struct{}, 1),
		input:    make(chan *beep.Buffer, 1000),
		callFunc: callFunc,
	}

	go p.run()
	return p
}

func (p *Player) run() {
	for {
		select {
		case <-p.cancel:
			speaker.Clear()
			speaker.Close()
			return
		case b := <-p.input:
			p.play(b)
		}
	}
}

func (p *Player) play(b *beep.Buffer) {
	w := sync.WaitGroup{}
	w.Add(1)
	speaker.Play(beep.Seq(b.Streamer(0, b.Len()), beep.Callback(func() {
		p.callFunc()
		w.Done()
	})))
	w.Wait()
}

func (p *Player) Append(b *beep.Buffer) {
	select {
	case p.input <- b:
	default:

	}
}

func (p *Player) Close() {
	p.cancel <- struct{}{}
}
