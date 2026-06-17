package probes

import (
	"encoding/json"
)

type siteHealth struct {
	Probe
}

var data *map[string]interface{}

func (siteHealth *siteHealth) Initialize() {
	log = _init()
	log.Info("Site Healt Probe: Initialized")
}

func HandleRecvString(recv []byte) {
	err := json.Unmarshal(recv, data)
	_checkerr(err)
}
