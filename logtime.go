package main

import (
	"sync"
	"time"

	"github.com/hnakamur/delayedticker"
)

type logTimeFormat int

const (
	logTimeFormatNone = iota
	logTimeFormatUTC
	logTimeFormatLocal
)

type logTimeCache struct {
	text   string
	ticker *delayedticker.DelayedTicker
	mu     sync.RWMutex
}

func NewLogTimeCache(format logTimeFormat) *logTimeCache {
	c := &logTimeCache{}
	now := time.Now()
	c.updateCache(format, now)
	go func() {
		delay := now.Truncate(time.Second).Add(time.Second).Sub(now)
		c.ticker = delayedticker.NewDelayedTicker(delay, time.Second)
		for {
			now := <-c.ticker.C
			c.updateCache(format, now)
		}
	}()
	return c
}

func (c *logTimeCache) updateCache(format logTimeFormat, t time.Time) {
	var text string
	switch format {
	case logTimeFormatUTC:
		text = t.UTC().Format(time.RFC3339)
	case logTimeFormatLocal:
		text = t.Format(time.RFC3339)
	}
	c.mu.Lock()
	c.text = text
	c.mu.Unlock()
}

func (c *logTimeCache) AppendTime(b []byte) []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return append(b, c.text...)
}

func (c *logTimeCache) Stop() {
	c.ticker.Stop()
}
