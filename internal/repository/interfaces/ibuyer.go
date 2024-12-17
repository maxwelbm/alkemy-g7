
package interfaces


type IBuyerRepo interface {
	Get() (map[int]model.Buyer, error)
	GetById(id int) (model.Buyer, error)
	Post(buyer model.Buyer) (model.Buyer, error)
	Update(id int, buyer model.Buyer) (model.Buyer, error)
	Delete(id int) error
}
