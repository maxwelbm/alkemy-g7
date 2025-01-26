package service_test

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

func setupSeller() *service.SellersService {
	mocksl := new(repository.SellerMockRepository)
	mocklc := new(repository.LocalityMockRepository)

	return service.CreateServiceSellers(mocksl, mocklc)
}