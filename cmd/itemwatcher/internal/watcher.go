package itemwatcher

import (
	"context"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/hectorgimenez/d2go/pkg/itemfilter"
	"github.com/hectorgimenez/d2go/pkg/memory"
	"github.com/hectorgimenez/d2go/pkg/nip"
	"log"
	"os"
	"time"
)

type Watcher struct {
	gr                     *memory.GameReader
	rules                  []nip.Rule
	alreadyNotifiedItemIDs map[int]interface{}
}

func NewWatcher(gr *memory.GameReader, rules []nip.Rule) *Watcher {
	return &Watcher{gr: gr, rules: rules}
}

func (w *Watcher) Start(ctx context.Context) error {
	w.alreadyNotifiedItemIDs = make(map[int]interface{}, 0)
	audioBuffer, err := initAudio()
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			time.Sleep(100 * time.Millisecond)

			d := w.gr.GetData()
			for _, i := range d.Items.Ground {
				if !itemfilter.Evaluate(i, w.rules) {
					continue
				}

				if _, found := w.alreadyNotifiedItemIDs[int(i.UnitID)]; found {
					continue
				}

				log.Printf("%s: Item detected: %s. Quality: %s", time.Now().Format(time.RFC3339), i.Name, i.Quality.ToString())

				w.alreadyNotifiedItemIDs[int(i.UnitID)] = nil
				speaker.Play(audioBuffer.Streamer(0, audioBuffer.Len()))
			}

			// Cleanup the already notified list when item is out of range
			for k := range w.alreadyNotifiedItemIDs {
				found := false
				for _, i := range d.Items.Ground {
					if int(i.UnitID) == k {
						found = true
					}
				}
				if !found {
					delete(w.alreadyNotifiedItemIDs, k)
				}
			}
		}
	}
}

func initAudio() (*beep.Buffer, error) {
	f, err := os.Open("assets/ching.wav")
	if err != nil {
		return nil, err
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return nil, err
	}
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	return buffer, nil
}
