package generator

// ParallelOperation enables parallel operations with gg
// by creating one context per goroutine and merging all
// contexts as soon as they're done.

// This greatly improves speed in generators with a lot of tasks.
// In the collage generator where each block (cover + text),
// parallel operations were 7-9x faster than a single threaded
// drawing environment (running collage 3x3).

// Also, it uses image.Image, which means that you're not limited to gg.
// As long as your task returns image.Image it'll work.
import (
	"github.com/fogleman/gg"
	"image"
)

type TaskResponse struct {
	Image  image.Image
	DrawAt [2]int
	Mask   bool
}

type ParallelImageOperator struct {
	Base    gg.Context
	ToRecv  int
	Channel chan TaskResponse
}

func NewPIO(x, y int) ParallelImageOperator {
	return ParallelImageOperator{
		ToRecv:  0,
		Base:    *gg.NewContext(x, y),
		Channel: make(chan TaskResponse),
	}
}

// Adds a task
func (p *ParallelImageOperator) RunTask(f func(chan TaskResponse)) {
	p.ToRecv = p.ToRecv + 1
	go func() {
		f(p.Channel)
	}()
}

func (p *ParallelImageOperator) RunTaskWithData(f func(chan TaskResponse, interface{}), b interface{}) {
	p.ToRecv = p.ToRecv + 1
	go func(a interface{}) {
		f(p.Channel, a)
	}(b)
}

func (p ParallelImageOperator) WaitAndDraw() {
	go func() {
		var v *TaskResponse = nil
		for {
			r, more := <-p.Channel
			if more {
				if r.Mask && p.ToRecv != 1 {
					v = &r
					continue
				}
				p.Base.DrawImage(r.Image, r.DrawAt[0], r.DrawAt[1])
				p.ToRecv = p.ToRecv - 1

				if p.ToRecv == 1 && v != nil {
					p.Base.DrawImage(v.Image, v.DrawAt[0], v.DrawAt[1])
					p.ToRecv = 0
				}
				if p.ToRecv == 0 {
					return
				}
			} else {
				return
			}
		}
	}()

	for {
		if p.ToRecv == 0 {
			close(p.Channel)
			return
		}
	}
}

func (p ParallelImageOperator) Image() image.Image {
	return p.Base.Image()
}
