package main

import (
	"fmt"
	"math"

	"github.com/jfreymuth/pulse"
	"github.com/jfreymuth/pulse/proto"
)

func ListSinks(c *pulse.Client) error {
	sinks, err := c.ListSinks()
	if err != nil {
		return err
	}

	for _, sink := range sinks {
		fmt.Println(sink.ID(), ": ", sink.Name())
	}

	return nil
}

func PlaySynth(c *pulse.Client) error {
	var t, phase float32

	synth := func(out []float32) (int, error) {
		for i := range out {
			if t > 4 {
				return i, pulse.EndOfData
			}
			x := float32(math.Sin(2 * math.Pi * float64(phase)))
			out[i] = x * 0.1
			f := [...]float32{440, 550, 660, 880}[int(2*t)&3]
			phase += f / 44100
			if phase >= 1 {
				phase--
			}
			t += 1. / 44100
		}
		return len(out), nil
	}

	stream, err := c.NewPlayback(pulse.Float32Reader(synth), pulse.PlaybackLatency(.1))
	if err != nil {
		return err
	}
	defer stream.Close()

	stream.Start()
	stream.Drain()
	if stream.Error() != nil {
		return stream.Error()
	}

	// fmt.Println("Underflow:", stream.Underflow())

	return nil
}

func SetDefaultSink(c *pulse.Client, sinkName string) error {
	req := proto.SetDefaultSink{SinkName: sinkName}
	err := c.RawRequest(&req, nil)
	if err != nil {
		return err
	}

	sink, err := c.DefaultSink()
	if err != nil {
		return err
	}
	fmt.Println("Default sink: ", sink.ID())

	return nil
}

func main() {
	c, err := pulse.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	err = ListSinks(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = SetDefaultSink(c, "1__2")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = PlaySynth(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = SetDefaultSink(c, "Channel_1__Channel_2.4")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = PlaySynth(c)
	if err != nil {
		fmt.Println(err)
		return
	}
}
