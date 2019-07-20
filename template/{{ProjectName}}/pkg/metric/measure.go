package metric

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var gc = cache.New(5*time.Minute, 10*time.Minute)

func StartRecord(key string, startTime *time.Time) {

	if startTime != nil {
		gc.Set(key, *startTime, cache.DefaultExpiration)
	} else {
		gc.Set(key, time.Now(), cache.DefaultExpiration)
	}
}

func CommitRecord(key string) *time.Duration {
	value, found := gc.Get(key)
	if found {
		elapsed := time.Since(value.(time.Time))
		return &elapsed
	} else {
		return nil
	}
}
