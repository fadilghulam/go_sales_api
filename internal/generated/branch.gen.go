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

func newBranch(db *gorm.DB, opts ...gen.DOOption) branch {
	_branch := branch{}

	_branch.branchDo.UseDB(db, opts...)
	_branch.branchDo.UseModel(&model.Branch{})

	tableName := _branch.branchDo.TableName()
	_branch.ALL = field.NewAsterisk(tableName)
	_branch.ID = field.NewInt16(tableName, "id")
	_branch.Name = field.NewString(tableName, "name")

	_branch.fillFieldMap()

	return _branch
}

type branch struct {
	branchDo

	ALL  field.Asterisk
	ID   field.Int16
	Name field.String

	fieldMap map[string]field.Expr
}

func (b branch) Table(newTableName string) *branch {
	b.branchDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b branch) As(alias string) *branch {
	b.branchDo.DO = *(b.branchDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *branch) updateTableName(table string) *branch {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt16(table, "id")
	b.Name = field.NewString(table, "name")

	b.fillFieldMap()

	return b
}

func (b *branch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *branch) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 2)
	b.fieldMap["id"] = b.ID
	b.fieldMap["name"] = b.Name
}

func (b branch) clone(db *gorm.DB) branch {
	b.branchDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b branch) replaceDB(db *gorm.DB) branch {
	b.branchDo.ReplaceDB(db)
	return b
}

type branchDo struct{ gen.DO }

type IBranchDo interface {
	gen.SubQuery
	Debug() IBranchDo
	WithContext(ctx context.Context) IBranchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBranchDo
	WriteDB() IBranchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBranchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBranchDo
	Not(conds ...gen.Condition) IBranchDo
	Or(conds ...gen.Condition) IBranchDo
	Select(conds ...field.Expr) IBranchDo
	Where(conds ...gen.Condition) IBranchDo
	Order(conds ...field.Expr) IBranchDo
	Distinct(cols ...field.Expr) IBranchDo
	Omit(cols ...field.Expr) IBranchDo
	Join(table schema.Tabler, on ...field.Expr) IBranchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBranchDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBranchDo
	Group(cols ...field.Expr) IBranchDo
	Having(conds ...gen.Condition) IBranchDo
	Limit(limit int) IBranchDo
	Offset(offset int) IBranchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBranchDo
	Unscoped() IBranchDo
	Create(values ...*model.Branch) error
	CreateInBatches(values []*model.Branch, batchSize int) error
	Save(values ...*model.Branch) error
	First() (*model.Branch, error)
	Take() (*model.Branch, error)
	Last() (*model.Branch, error)
	Find() ([]*model.Branch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Branch, err error)
	FindInBatches(result *[]*model.Branch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Branch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBranchDo
	Assign(attrs ...field.AssignExpr) IBranchDo
	Joins(fields ...field.RelationField) IBranchDo
	Preload(fields ...field.RelationField) IBranchDo
	FirstOrInit() (*model.Branch, error)
	FirstOrCreate() (*model.Branch, error)
	FindByPage(offset int, limit int) (result []*model.Branch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBranchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b branchDo) Debug() IBranchDo {
	return b.withDO(b.DO.Debug())
}

func (b branchDo) WithContext(ctx context.Context) IBranchDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b branchDo) ReadDB() IBranchDo {
	return b.Clauses(dbresolver.Read)
}

func (b branchDo) WriteDB() IBranchDo {
	return b.Clauses(dbresolver.Write)
}

func (b branchDo) Session(config *gorm.Session) IBranchDo {
	return b.withDO(b.DO.Session(config))
}

func (b branchDo) Clauses(conds ...clause.Expression) IBranchDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b branchDo) Returning(value interface{}, columns ...string) IBranchDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b branchDo) Not(conds ...gen.Condition) IBranchDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b branchDo) Or(conds ...gen.Condition) IBranchDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b branchDo) Select(conds ...field.Expr) IBranchDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b branchDo) Where(conds ...gen.Condition) IBranchDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b branchDo) Order(conds ...field.Expr) IBranchDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b branchDo) Distinct(cols ...field.Expr) IBranchDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b branchDo) Omit(cols ...field.Expr) IBranchDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b branchDo) Join(table schema.Tabler, on ...field.Expr) IBranchDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b branchDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBranchDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b branchDo) RightJoin(table schema.Tabler, on ...field.Expr) IBranchDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b branchDo) Group(cols ...field.Expr) IBranchDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b branchDo) Having(conds ...gen.Condition) IBranchDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b branchDo) Limit(limit int) IBranchDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b branchDo) Offset(offset int) IBranchDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b branchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBranchDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b branchDo) Unscoped() IBranchDo {
	return b.withDO(b.DO.Unscoped())
}

func (b branchDo) Create(values ...*model.Branch) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b branchDo) CreateInBatches(values []*model.Branch, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b branchDo) Save(values ...*model.Branch) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b branchDo) First() (*model.Branch, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) Take() (*model.Branch, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) Last() (*model.Branch, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) Find() ([]*model.Branch, error) {
	result, err := b.DO.Find()
	return result.([]*model.Branch), err
}

func (b branchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Branch, err error) {
	buf := make([]*model.Branch, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b branchDo) FindInBatches(result *[]*model.Branch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b branchDo) Attrs(attrs ...field.AssignExpr) IBranchDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b branchDo) Assign(attrs ...field.AssignExpr) IBranchDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b branchDo) Joins(fields ...field.RelationField) IBranchDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b branchDo) Preload(fields ...field.RelationField) IBranchDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b branchDo) FirstOrInit() (*model.Branch, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) FirstOrCreate() (*model.Branch, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Branch), nil
	}
}

func (b branchDo) FindByPage(offset int, limit int) (result []*model.Branch, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b branchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b branchDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b branchDo) Delete(models ...*model.Branch) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *branchDo) withDO(do gen.Dao) *branchDo {
	b.DO = *do.(*gen.DO)
	return b
}