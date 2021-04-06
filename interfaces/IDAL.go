package interfaces

// IDAL ...
type IDAL interface {
	Create(i interface{}) error
	Read(i interface{}) error
	Update(i interface{}) error
	Delete(i interface{}) error
	ReadAll(i interface{}) (map[string]interface{}, error)
}
