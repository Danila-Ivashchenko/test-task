package hasher

import "go-kafka/internal/domain/model"

type Hasher interface {
	Set(user *model.User) error
	Get(id int64) (*model.User, error)
	Delete(id int64) error
	Update(user *model.User) error
}
