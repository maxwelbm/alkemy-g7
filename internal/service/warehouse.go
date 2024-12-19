package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
)

type WareHouseDefault struct {
	rp repository.WareHouseMap
}

func NewWareHoureService(rp repository.WareHouseMap) *WareHouseDefault {
	return &WareHouseDefault{rp: rp}
}

func (s *WareHouseDefault) GetAllWareHouse() (w map[int]model.WareHouse, err error) {
	w, err = s.rp.GetAllWareHouse()
	return
}
