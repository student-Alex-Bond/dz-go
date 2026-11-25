package bins

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"crearedAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(name string, private bool) (*Bin, error) {
	if name == "" {
		return nil, errors.New("ИМЕНИ НЕТ")
	}
	return &Bin{
		Id:        uuid.New().String(),
		Private:   private,
		CreatedAt: time.Now().Local(),
		Name:      name,
	}, nil
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}
