package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ISectionService interface {
	Get() ([]model.Section, error)
	GetById(id int) (model.Section, error)
	Post(section *model.Section) (model.Section, error)
	Update(id int, section model.Section) (model.Section, error)
	Delete(id int) error
}
