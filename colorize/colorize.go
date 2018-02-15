package colorize

import (
	"sync"
)

type IColorize interface {
	Start() string
	End() string
	Get() IColorize
	Set(IColorize)
	Add(...IColorize) IColorize
	Dup() IColorize
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

func (cc *Colorize) Get() IColorize {
	return cc.Dup()
}

func (cc *Colorize) Set(colorize IColorize) {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	cc.start = colorize.Start()
	cc.end = colorize.End()
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

func (cc *Colorize) Dup() IColorize {
	return NewColorize(cc.Start(), cc.End())
}
