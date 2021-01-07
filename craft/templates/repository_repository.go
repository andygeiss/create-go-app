package templates

// RepositoryGo ...
var RepositoryGo = `

package repository

// Repository ...
type Repository interface {
	Add(id string, value interface{})
	Error() error
	FindByID(id string) (item interface{})
	FindByKeyValue(key string, val interface{}) (item interface{})
	Remove(id string)
}
`
