package gql

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type Resolver struct {
	DB *gorm.DB
	mu sync.Mutex
}
