package models

import "sync"

type Counter struct {
	TotalSum int
	Mu       sync.Mutex
}
