package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ISectionRepo interface {
	Get() ([]model.Section, error)
	GetByID(id int) (model.Section, error)
	Post(section *model.Section) (model.Section, error)
	Update(id int, section *model.Section) (model.Section, error)
	Delete(id int) error
	CountProductBatchesBySectionID(id int) (countProdBatches model.SectionProductBatches, err error)
	CountProductBatchesSections() (countProductBatches []model.SectionProductBatches, err error)
}
