package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect
	"github.com/spf13/viper"
)

//EntityManager contains database connection and CRUD methods
type EntityManager struct {
	db *gorm.DB
}

//ContentStore has CRUD methods for Content
type ContentStore struct {
	em *EntityManager
}

//NewEntityManager creates a new instance of an EntityManager and sets up a database connection
func NewEntityManager() (em *EntityManager) {
	gormDB, err := gorm.Open("mysql", viper.GetString("database_dsn"))
	if err != nil {
		panic("failed to connect database")
	}
	// defer gormDB.Close()
	em = &EntityManager{}
	em.db = gormDB

	return em
}

//AutoMigrate updates database structures
func (e *EntityManager) AutoMigrate() {
	e.db.AutoMigrate(&Content{})
}

//NewContentStore creates a new CRUD store for Content
func (e *EntityManager) NewContentStore() *ContentStore {
	store := ContentStore{
		em: NewEntityManager(),
	}

	return &store
}

//FindAll retrieves all content structs
func (s *ContentStore) FindAll() (contents []Content, err error) {
	err = s.em.db.Find(&contents).Error
	return contents, err
}

//FindOneBySlug retrieves a single Content by slug
func (s *ContentStore) FindOneBySlug(slug string) (content Content, err error) {
	err = s.em.db.First(&content, "slug = ?", slug).Error
	return content, err
}

//Persist persists a Content struct
func (s *ContentStore) Persist(content *Content) error {
	return s.em.db.Create(content).Error
}
