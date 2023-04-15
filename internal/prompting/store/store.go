package store

import (
	"gorm.io/gorm"
	"sync"
)

//go:generate mockgen -destination mock_store.go -package store github.com/marmotedu/miniblog/internal/miniblog/store IStore,UserStore,PostStore

var (
	once sync.Once
	S    *datastore
)

// IStore 定义了 Store 层需要实现的方法
type IStore interface {
	DB() *gorm.DB
	Users() UserStore
	Posts() PostStore
}

// datastore 是 IStore 的一个具体实现
type datastore struct {
	db *gorm.DB
}

// 确保 datastore 实现了 IStore 接口
var _ IStore = (*datastore)(nil)

func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

// DB 返回存储在 datastore 中的 *gorm.DB
func (ds *datastore) DB() *gorm.DB {
	return ds.db
}

// Users 返回一个实现了 UserStore 接口的实例
func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)
}
