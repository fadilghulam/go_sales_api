// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"go_sales_api/internal/model"
)

func newSalesman(db *gorm.DB, opts ...gen.DOOption) salesman {
	_salesman := salesman{}

	_salesman.salesmanDo.UseDB(db, opts...)
	_salesman.salesmanDo.UseModel(&model.Salesman{})

	tableName := _salesman.salesmanDo.TableName()
	_salesman.ALL = field.NewAsterisk(tableName)
	_salesman.ID = field.NewInt16(tableName, "id")
	_salesman.BranchID = field.NewInt16(tableName, "branch_id")
	_salesman.Name = field.NewString(tableName, "name")
	_salesman.CreatedAt = field.NewTime(tableName, "created_at")
	_salesman.UpdatedAt = field.NewTime(tableName, "updated_at")

	_salesman.fillFieldMap()

	return _salesman
}

type salesman struct {
	salesmanDo

	ALL       field.Asterisk
	ID        field.Int16
	BranchID  field.Int16
	Name      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (s salesman) Table(newTableName string) *salesman {
	s.salesmanDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s salesman) As(alias string) *salesman {
	s.salesmanDo.DO = *(s.salesmanDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *salesman) updateTableName(table string) *salesman {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt16(table, "id")
	s.BranchID = field.NewInt16(table, "branch_id")
	s.Name = field.NewString(table, "name")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *salesman) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *salesman) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id"] = s.ID
	s.fieldMap["branch_id"] = s.BranchID
	s.fieldMap["name"] = s.Name
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s salesman) clone(db *gorm.DB) salesman {
	s.salesmanDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s salesman) replaceDB(db *gorm.DB) salesman {
	s.salesmanDo.ReplaceDB(db)
	return s
}

type salesmanDo struct{ gen.DO }

type ISalesmanDo interface {
	gen.SubQuery
	Debug() ISalesmanDo
	WithContext(ctx context.Context) ISalesmanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISalesmanDo
	WriteDB() ISalesmanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISalesmanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISalesmanDo
	Not(conds ...gen.Condition) ISalesmanDo
	Or(conds ...gen.Condition) ISalesmanDo
	Select(conds ...field.Expr) ISalesmanDo
	Where(conds ...gen.Condition) ISalesmanDo
	Order(conds ...field.Expr) ISalesmanDo
	Distinct(cols ...field.Expr) ISalesmanDo
	Omit(cols ...field.Expr) ISalesmanDo
	Join(table schema.Tabler, on ...field.Expr) ISalesmanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISalesmanDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISalesmanDo
	Group(cols ...field.Expr) ISalesmanDo
	Having(conds ...gen.Condition) ISalesmanDo
	Limit(limit int) ISalesmanDo
	Offset(offset int) ISalesmanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISalesmanDo
	Unscoped() ISalesmanDo
	Create(values ...*model.Salesman) error
	CreateInBatches(values []*model.Salesman, batchSize int) error
	Save(values ...*model.Salesman) error
	First() (*model.Salesman, error)
	Take() (*model.Salesman, error)
	Last() (*model.Salesman, error)
	Find() ([]*model.Salesman, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Salesman, err error)
	FindInBatches(result *[]*model.Salesman, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Salesman) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISalesmanDo
	Assign(attrs ...field.AssignExpr) ISalesmanDo
	Joins(fields ...field.RelationField) ISalesmanDo
	Preload(fields ...field.RelationField) ISalesmanDo
	FirstOrInit() (*model.Salesman, error)
	FirstOrCreate() (*model.Salesman, error)
	FindByPage(offset int, limit int) (result []*model.Salesman, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISalesmanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s salesmanDo) Debug() ISalesmanDo {
	return s.withDO(s.DO.Debug())
}

func (s salesmanDo) WithContext(ctx context.Context) ISalesmanDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s salesmanDo) ReadDB() ISalesmanDo {
	return s.Clauses(dbresolver.Read)
}

func (s salesmanDo) WriteDB() ISalesmanDo {
	return s.Clauses(dbresolver.Write)
}

func (s salesmanDo) Session(config *gorm.Session) ISalesmanDo {
	return s.withDO(s.DO.Session(config))
}

func (s salesmanDo) Clauses(conds ...clause.Expression) ISalesmanDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s salesmanDo) Returning(value interface{}, columns ...string) ISalesmanDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s salesmanDo) Not(conds ...gen.Condition) ISalesmanDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s salesmanDo) Or(conds ...gen.Condition) ISalesmanDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s salesmanDo) Select(conds ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s salesmanDo) Where(conds ...gen.Condition) ISalesmanDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s salesmanDo) Order(conds ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s salesmanDo) Distinct(cols ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s salesmanDo) Omit(cols ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s salesmanDo) Join(table schema.Tabler, on ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s salesmanDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s salesmanDo) RightJoin(table schema.Tabler, on ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s salesmanDo) Group(cols ...field.Expr) ISalesmanDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s salesmanDo) Having(conds ...gen.Condition) ISalesmanDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s salesmanDo) Limit(limit int) ISalesmanDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s salesmanDo) Offset(offset int) ISalesmanDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s salesmanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISalesmanDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s salesmanDo) Unscoped() ISalesmanDo {
	return s.withDO(s.DO.Unscoped())
}

func (s salesmanDo) Create(values ...*model.Salesman) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s salesmanDo) CreateInBatches(values []*model.Salesman, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s salesmanDo) Save(values ...*model.Salesman) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s salesmanDo) First() (*model.Salesman, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salesman), nil
	}
}

func (s salesmanDo) Take() (*model.Salesman, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salesman), nil
	}
}

func (s salesmanDo) Last() (*model.Salesman, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salesman), nil
	}
}

func (s salesmanDo) Find() ([]*model.Salesman, error) {
	result, err := s.DO.Find()
	return result.([]*model.Salesman), err
}

func (s salesmanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Salesman, err error) {
	buf := make([]*model.Salesman, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s salesmanDo) FindInBatches(result *[]*model.Salesman, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s salesmanDo) Attrs(attrs ...field.AssignExpr) ISalesmanDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s salesmanDo) Assign(attrs ...field.AssignExpr) ISalesmanDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s salesmanDo) Joins(fields ...field.RelationField) ISalesmanDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s salesmanDo) Preload(fields ...field.RelationField) ISalesmanDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s salesmanDo) FirstOrInit() (*model.Salesman, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salesman), nil
	}
}

func (s salesmanDo) FirstOrCreate() (*model.Salesman, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salesman), nil
	}
}

func (s salesmanDo) FindByPage(offset int, limit int) (result []*model.Salesman, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s salesmanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s salesmanDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s salesmanDo) Delete(models ...*model.Salesman) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *salesmanDo) withDO(do gen.Dao) *salesmanDo {
	s.DO = *do.(*gen.DO)
	return s
}
