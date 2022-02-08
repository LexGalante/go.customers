package entities

//Entity -> base entity
type Entity interface {
	Name() string
	TableName() string
	Validate() (map[string]string, error)
}
