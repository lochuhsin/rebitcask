package models

import (
	"rebitcask/internal/dao"
)

type IMemory interface {
	Get([]byte) (dao.Entry, bool)
	Set(dao.Entry) // if the memory is in frozen state, close set operation
	GetSize() int
	GetAll() []dao.Entry // Expected order by key from small to large
}
