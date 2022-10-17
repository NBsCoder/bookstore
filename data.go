package main

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultSize = 100

// model定义

func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// Shelf 书架
type Shelf struct {
	ID       int64     `gorm:"primaryKey"` // ID
	Theme    string    // 主题
	Size     int64     // 容量
	CreateAt time.Time // 生成时间
	UpdateAt time.Time // 修改时间
}

// Book 图书
type Book struct {
	ID       int64     `gorm:"primaryKey"` // ID
	Author   string    // 作者
	Title    string    // 标题
	ShelfID  int64     // 书架编号
	CreateAt time.Time // 生成时间
	UpdateAt time.Time // 修改时间
}

// ORM操作

type bookstore struct {
	db *gorm.DB
}

// ShelfCreate 增
func (b *bookstore) ShelfCreate(ctx context.Context, shelf Shelf) (*Shelf, error) {
	if len(shelf.Theme) <= 0 {
		return nil, errors.New("invalid params")
	}
	size := shelf.Size
	if size <= 0 {
		shelf.Size = defaultSize
	}
	s := Shelf{
		Theme:    shelf.Theme,
		Size:     shelf.Size,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := b.db.WithContext(ctx).Create(&s).Error
	return &s, err
}

// ShelfDelete 删
func (b *bookstore) ShelfDelete(ctx context.Context, id int64) error {
	//if id <= 0 {
	//	return errors.New("invalid params")
	//}
	//err := b.db.WithContext(ctx).Where("ID=?", id).Delete(&Shelf{}).Error

	// 根据主键来删
	return b.db.WithContext(ctx).Delete(&Shelf{}, id).Error
}

// ShelfList 列表
func (b *bookstore) ShelfList(ctx context.Context) ([]*Shelf, error) {
	var vl []*Shelf
	err := b.db.WithContext(ctx).Find(&vl).Error
	return vl, err
}

// ShelfGet 查
func (b *bookstore) ShelfGet(ctx context.Context, id int64) (*Shelf, error) {
	v := Shelf{}
	err := b.db.WithContext(ctx).First(&v, id).Error
	return &v, err
}

// BookCreate 增
func (b *bookstore) BookCreate() {

}

// BookDelete 删
func (b *bookstore) BookDelete() {

}

// BookUpdate 改
func (b *bookstore) BookUpdate() {

}

// BookGet 查
func (b *bookstore) BookGet() {

}
