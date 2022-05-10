package progressbar

import (
	"fmt"
	"strings"
	"sync"
)

type ProgressBar struct {
	max     float64
	current float64
	lock    sync.Mutex
	addChan chan int
}

func NewProgressBar(max float64) *ProgressBar {
	p := &ProgressBar{
		max:     max,
		addChan: make(chan int),
	}
	return p
}

func (p *ProgressBar) Show() {
	go func() {
		fmt.Printf("[%-101s] %d%%\r", "", 0)
		for i := range p.addChan {
			fmt.Printf("[%-101s] %d%%", strings.Repeat("=", i)+">", i)
			if i == 100 {
				fmt.Printf("\n")
			} else {
				fmt.Printf("\r")
			}
		}
	}()
}

func (p *ProgressBar) Add(n float64) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.current >= p.max {
		return
	}
	p.current += n

	i := int((p.current / p.max) * 100)
	if i > 100 {
		i = 100
	}
	p.addChan <- i
	if p.current >= p.max {
		close(p.addChan)
	}
}
