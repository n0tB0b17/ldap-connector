package models

import "time"

type DomainController struct {
	Name     string
	IPAddr   string
	Site     string
	IsGlobal bool
	LastSeen time.Time
}

type DCEvent struct {
	Type  string
	DC    DomainController
	Error error
	Time  time.Time
}
