package utils

import (
	"github.com/schollz/progressbar/v3"
)

type ProgressBar struct {
	bar *progressbar.ProgressBar
}

func CreateProgressBar(max int, desc string) *ProgressBar {
	return &ProgressBar{
		bar: progressbar.Default(int64(max), desc),
	}
}

func (p *ProgressBar) Add(n int) error {
	return p.bar.Add(n)
}

func (p *ProgressBar) Set(n int) error {
	return p.bar.Set(n)
}

func (p *ProgressBar) Finish() error {
	return p.bar.Finish()
}
