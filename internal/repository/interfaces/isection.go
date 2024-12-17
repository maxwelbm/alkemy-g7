package interfaces



type ISectionRepo interface {
	Get() (map[int]model.Section, error)
	GetById(id int) (model.Section, error)
	Post(section model.Section) (model.Section, error)
	Update(id int, section model.Section) (model.Section, error)
	Delete(id int) error
}