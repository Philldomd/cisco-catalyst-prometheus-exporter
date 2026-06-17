package probes

import (
	"log/slog"
)

var log *slog.Logger

type Probe struct{}
type Probes interface {
	Initialize()
}

var probe = Probe{}

func _checkerr(err error) bool {
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}

func _init() *slog.Logger {
	return log
}

func Init(lg *slog.Logger) {
	log = lg
	p1 := &siteHealth{probe}
	l_probes := []Probes{Probes(p1)}
	for _, p := range l_probes {
		p.Initialize()
	}
}
