package persistence

import (
	"github.com/cossim/coss-server/service/msg/domain/entity"
	"github.com/cossim/coss-server/service/msg/domain/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	Mr repository.MsgRepository
	db *gorm.DB
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Mr: NewMsgRepo(db),
		db: db,
	}
}

func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.GroupMessage{}, &entity.UserMessage{}, &entity.GroupMessageRead{})
}
