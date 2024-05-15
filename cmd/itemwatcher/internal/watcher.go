package itemwatcher

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/memory"
	"github.com/hectorgimenez/d2go/pkg/nip"
)

type Watcher struct {
	gr                     *memory.GameReader
	rules                  []nip.Rule
	alreadyNotifiedItemIDs []itemFootprint
}

type itemFootprint struct {
	detectedAt time.Time
	area       area.ID
	position   data.Position
	name       item.Name
	quality    item.Quality
}

func (fp itemFootprint) Match(area area.ID, i data.Item) bool {
	return fp.area == area && fp.position == i.Position && fp.name == i.Name && fp.quality == i.Quality
}

func NewWatcher(gr *memory.GameReader, rules []nip.Rule) *Watcher {
	return &Watcher{gr: gr, rules: rules}
}

func (w *Watcher) Start(ctx context.Context) error {
	w.alreadyNotifiedItemIDs = make([]itemFootprint, 0)
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
			for _, i := range d.Inventory.ByLocation(item.LocationGround) {
				for _, r := range w.rules {
					res, err := r.Evaluate(i)
					if err != nil {
						log.Printf("error evaluating rule: %v", err)
						continue
					}
					if res == nip.RuleResultNoMatch {
						continue
					}
				}

				found := false
				for _, fp := range w.alreadyNotifiedItemIDs {
					if fp.Match(d.PlayerUnit.Area, i) {
						found = true
						break
					}
				}
				if found {
					continue
				}

				log.Printf("%s: Item detected: %s. Quality: %s", time.Now().Format(time.RFC3339), i.Name, i.Quality.ToString())

				w.alreadyNotifiedItemIDs = append(w.alreadyNotifiedItemIDs, itemFootprint{
					detectedAt: time.Now(),
					area:       d.PlayerUnit.Area,
					position:   i.Position,
					name:       i.Name,
					quality:    i.Quality,
				})
				speaker.Play(audioBuffer.Streamer(0, audioBuffer.Len()))
			}

			// Cleanup after 10 minute AND out of range
			purgedNotifiedItems := make([]itemFootprint, 0)
			for _, t := range w.alreadyNotifiedItemIDs {
				found := false
				for _, it := range d.Inventory.ByLocation(item.LocationGround) {
					if t.Match(d.PlayerUnit.Area, it) {
						found = true
					}
				}
				if found || time.Since(t.detectedAt) < time.Minute*10 {
					purgedNotifiedItems = append(purgedNotifiedItems, t)
				}
			}
			w.alreadyNotifiedItemIDs = purgedNotifiedItems
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
