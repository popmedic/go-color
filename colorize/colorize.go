package colorize

import (
	"sync"
)

type IColorize interface {
	Color(string) string
	Start() string
	End() string
	Add(...IColorize) IColorize
}

type Colorize struct {
	lock  sync.RWMutex
	start string
	end   string
}

func NewColorize(start, end string, colorize ...IColorize) IColorize {
	n := &Colorize{
		lock:  sync.RWMutex{},
		start: start,
		end:   end,
	}
	return n.Add(colorize...)
}

func (cc *Colorize) Color(s string) string {
	return cc.start + s + cc.end
}

func (cc *Colorize) Add(colorize ...IColorize) IColorize {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	for _, c := range colorize {
		cc.start = cc.start + c.Start()
	}

	e := ""
	for i := len(colorize) - 1; i >= 0; i-- {
		e = e + colorize[i].End()
	}
	cc.end = e + cc.end
	return cc
}

func (cc *Colorize) Start() string {
	cc.lock.RLock()
	defer cc.lock.RUnlock()

	return cc.start
}

func (cc *Colorize) End() string {
	cc.lock.RLock()
	defer cc.lock.RUnlock()

	return cc.end
}
