package template

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kacperjurak/igorm"
)

type gormw struct {
	someField string
}

func New(field string) (db igorm.Gormw, err error) {
	return &gormw{
		someField: field,
	}, err
}

func (it *gormw) Close() error {
	return nil
}

func (it *gormw) DB() *sql.DB {
	return &sql.DB{}
}

func (it *gormw) New() igorm.Gormw {
	return it
}

func (it *gormw) NewScope(value interface{}) *gorm.Scope {
	return &gorm.Scope{}
}

func (it *gormw) CommonDB() gorm.SQLCommon {
	return &sql.DB{}
}

func (it *gormw) Callback() *gorm.Callback {
	return &gorm.Callback{}
}

func (it *gormw) SetLogger(log gorm.Logger) {
}

func (it *gormw) LogMode(enable bool) igorm.Gormw {
	return it
}

func (it *gormw) SingularTable(enable bool) {
}

func (it *gormw) Where(query interface{}, args ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Or(query interface{}, args ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Not(query interface{}, args ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Limit(value int) igorm.Gormw {
	return it
}

func (it *gormw) Offset(value int) igorm.Gormw {
	return it
}

func (it *gormw) Order(value string, reorder ...bool) igorm.Gormw {
	return it
}

func (it *gormw) Select(query interface{}, args ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Omit(columns ...string) igorm.Gormw {
	return it
}

func (it *gormw) Group(query string) igorm.Gormw {
	return it
}

func (it *gormw) Having(query string, values ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Joins(query string, args ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Scopes(funcs ...func(*gorm.DB) *gorm.DB) igorm.Gormw {
	return it
}

func (it *gormw) Unscoped() igorm.Gormw {
	return it
}

func (it *gormw) Attrs(attrs ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Assign(attrs ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) First(out interface{}, where ...interface{}) igorm.Gormw {
	fmt.Printf("Hey! First() method works, and the someField value is: %v\n", it.someField)
	return it
}

func (it *gormw) Last(out interface{}, where ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Find(out interface{}, where ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Scan(dest interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Row() *sql.Row {
	return &sql.Row{}
}

func (it *gormw) Rows() (*sql.Rows, error) {
	return &sql.Rows{}, nil
}

func (it *gormw) ScanRows(rows *sql.Rows, result interface{}) error {
	return nil
}

func (it *gormw) Pluck(column string, value interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Count(value interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Related(value interface{}, foreignKeys ...string) igorm.Gormw {
	return it
}

func (it *gormw) FirstOrInit(out interface{}, where ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) FirstOrCreate(out interface{}, where ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Update(attrs ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Updates(values interface{}, ignoreProtectedAttrs ...bool) igorm.Gormw {
	return it
}

func (it *gormw) UpdateColumn(attrs ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) UpdateColumns(values interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Save(value interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Create(value interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Delete(value interface{}, where ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Raw(sql string, values ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Exec(sql string, values ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Model(value interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Table(name string) igorm.Gormw {
	return it
}

func (it *gormw) Debug() igorm.Gormw {
	return it
}

func (it *gormw) Begin() igorm.Gormw {
	return it
}

func (it *gormw) Commit() igorm.Gormw {
	return it
}

func (it *gormw) Rollback() igorm.Gormw {
	return it
}

func (it *gormw) NewRecord(value interface{}) bool {
	return false
}

func (it *gormw) RecordNotFound() bool {
	return false
}

func (it *gormw) CreateTable(values ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) DropTable(values ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) DropTableIfExists(values ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) HasTable(value interface{}) bool {
	return false
}

func (it *gormw) AutoMigrate(values ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) ModifyColumn(column string, typ string) igorm.Gormw {
	return it
}

func (it *gormw) DropColumn(column string) igorm.Gormw {
	return it
}

func (it *gormw) AddIndex(indexName string, columns ...string) igorm.Gormw {
	return it
}

func (it *gormw) AddUniqueIndex(indexName string, columns ...string) igorm.Gormw {
	return it
}

func (it *gormw) RemoveIndex(indexName string) igorm.Gormw {
	return it
}

func (it *gormw) Association(column string) *gorm.Association {
	return &gorm.Association{}
}

func (it *gormw) Preload(column string, conditions ...interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Set(name string, value interface{}) igorm.Gormw {
	return it
}

func (it *gormw) InstantSet(name string, value interface{}) igorm.Gormw {
	return it
}

func (it *gormw) Get(name string) (interface{}, bool) {
	return nil, false
}

func (it *gormw) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
}

func (it *gormw) AddForeignKey(field string, dest string, onDelete string, onUpdate string) igorm.Gormw {
	return it
}

func (it *gormw) AddError(err error) error {
	return nil
}

func (it *gormw) GetErrors() (errors []error) {
	return errors
}

func (it *gormw) RowsAffected() int64 {
	return 0
}

func (it *gormw) Error() error {
	return nil
}
