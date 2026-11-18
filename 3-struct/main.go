package main

import (
	"errors"
	"time"
	"github.com/google/uuid"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func NewBin(name string, private bool) (*Bin, error) {
	if name == "" {
		return nil, errors.New("EMPTY NAME")
	}
	return &Bin{
		id:        uuid.New().String(),
		private:   private,
		createdAt: time.Now().Local(),
		name:      name,
	}, nil
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}

func main() {
}
