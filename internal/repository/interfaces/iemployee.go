package interfaces


type IEmployeeRepo interface {
	Get() (map[int]model.Employee, error)
	GetById(id int) (model.Employee, error)
	Post(employee model.Employee) (model.Employee, error)
	Update(id int, employee model.Employee) (model.Employee, error)
	Delete(id int) error
}
