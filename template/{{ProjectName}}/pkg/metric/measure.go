package metric

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Measure struct {
	gc *cache.Cache
}

func NewMeasure() *Measure {
	return &Measure{cache.New(5*time.Minute, 10*time.Minute)}
}

func (m *Measure) StartRecord(key string, startTime *time.Time) {

	if startTime != nil {
		m.gc.Set(key, *startTime, cache.DefaultExpiration)
	} else {
		m.gc.Set(key, time.Now(), cache.DefaultExpiration)
	}
}

func (m *Measure) CommitRecord(key string) *time.Duration {
	value, found := m.gc.Get(key)
	if found {
		elapsed := time.Since(value.(time.Time))
		return &elapsed
	} else {
		return nil
	}
}
