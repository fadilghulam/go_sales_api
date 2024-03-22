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

func newProductBranch(db *gorm.DB, opts ...gen.DOOption) productBranch {
	_productBranch := productBranch{}

	_productBranch.productBranchDo.UseDB(db, opts...)
	_productBranch.productBranchDo.UseModel(&model.ProductBranch{})

	tableName := _productBranch.productBranchDo.TableName()
	_productBranch.ALL = field.NewAsterisk(tableName)
	_productBranch.BranchID = field.NewInt16(tableName, "branch_id")
	_productBranch.ProductID = field.NewInt16(tableName, "product_id")

	_productBranch.fillFieldMap()

	return _productBranch
}

type productBranch struct {
	productBranchDo

	ALL       field.Asterisk
	BranchID  field.Int16
	ProductID field.Int16

	fieldMap map[string]field.Expr
}

func (p productBranch) Table(newTableName string) *productBranch {
	p.productBranchDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p productBranch) As(alias string) *productBranch {
	p.productBranchDo.DO = *(p.productBranchDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *productBranch) updateTableName(table string) *productBranch {
	p.ALL = field.NewAsterisk(table)
	p.BranchID = field.NewInt16(table, "branch_id")
	p.ProductID = field.NewInt16(table, "product_id")

	p.fillFieldMap()

	return p
}

func (p *productBranch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *productBranch) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 2)
	p.fieldMap["branch_id"] = p.BranchID
	p.fieldMap["product_id"] = p.ProductID
}

func (p productBranch) clone(db *gorm.DB) productBranch {
	p.productBranchDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p productBranch) replaceDB(db *gorm.DB) productBranch {
	p.productBranchDo.ReplaceDB(db)
	return p
}

type productBranchDo struct{ gen.DO }

type IProductBranchDo interface {
	gen.SubQuery
	Debug() IProductBranchDo
	WithContext(ctx context.Context) IProductBranchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProductBranchDo
	WriteDB() IProductBranchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProductBranchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProductBranchDo
	Not(conds ...gen.Condition) IProductBranchDo
	Or(conds ...gen.Condition) IProductBranchDo
	Select(conds ...field.Expr) IProductBranchDo
	Where(conds ...gen.Condition) IProductBranchDo
	Order(conds ...field.Expr) IProductBranchDo
	Distinct(cols ...field.Expr) IProductBranchDo
	Omit(cols ...field.Expr) IProductBranchDo
	Join(table schema.Tabler, on ...field.Expr) IProductBranchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProductBranchDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProductBranchDo
	Group(cols ...field.Expr) IProductBranchDo
	Having(conds ...gen.Condition) IProductBranchDo
	Limit(limit int) IProductBranchDo
	Offset(offset int) IProductBranchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProductBranchDo
	Unscoped() IProductBranchDo
	Create(values ...*model.ProductBranch) error
	CreateInBatches(values []*model.ProductBranch, batchSize int) error
	Save(values ...*model.ProductBranch) error
	First() (*model.ProductBranch, error)
	Take() (*model.ProductBranch, error)
	Last() (*model.ProductBranch, error)
	Find() ([]*model.ProductBranch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProductBranch, err error)
	FindInBatches(result *[]*model.ProductBranch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProductBranch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProductBranchDo
	Assign(attrs ...field.AssignExpr) IProductBranchDo
	Joins(fields ...field.RelationField) IProductBranchDo
	Preload(fields ...field.RelationField) IProductBranchDo
	FirstOrInit() (*model.ProductBranch, error)
	FirstOrCreate() (*model.ProductBranch, error)
	FindByPage(offset int, limit int) (result []*model.ProductBranch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProductBranchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p productBranchDo) Debug() IProductBranchDo {
	return p.withDO(p.DO.Debug())
}

func (p productBranchDo) WithContext(ctx context.Context) IProductBranchDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productBranchDo) ReadDB() IProductBranchDo {
	return p.Clauses(dbresolver.Read)
}

func (p productBranchDo) WriteDB() IProductBranchDo {
	return p.Clauses(dbresolver.Write)
}

func (p productBranchDo) Session(config *gorm.Session) IProductBranchDo {
	return p.withDO(p.DO.Session(config))
}

func (p productBranchDo) Clauses(conds ...clause.Expression) IProductBranchDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productBranchDo) Returning(value interface{}, columns ...string) IProductBranchDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productBranchDo) Not(conds ...gen.Condition) IProductBranchDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productBranchDo) Or(conds ...gen.Condition) IProductBranchDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productBranchDo) Select(conds ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productBranchDo) Where(conds ...gen.Condition) IProductBranchDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productBranchDo) Order(conds ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productBranchDo) Distinct(cols ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productBranchDo) Omit(cols ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productBranchDo) Join(table schema.Tabler, on ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productBranchDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productBranchDo) RightJoin(table schema.Tabler, on ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productBranchDo) Group(cols ...field.Expr) IProductBranchDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productBranchDo) Having(conds ...gen.Condition) IProductBranchDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productBranchDo) Limit(limit int) IProductBranchDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productBranchDo) Offset(offset int) IProductBranchDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productBranchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProductBranchDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productBranchDo) Unscoped() IProductBranchDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productBranchDo) Create(values ...*model.ProductBranch) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productBranchDo) CreateInBatches(values []*model.ProductBranch, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productBranchDo) Save(values ...*model.ProductBranch) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productBranchDo) First() (*model.ProductBranch, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBranch), nil
	}
}

func (p productBranchDo) Take() (*model.ProductBranch, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBranch), nil
	}
}

func (p productBranchDo) Last() (*model.ProductBranch, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBranch), nil
	}
}

func (p productBranchDo) Find() ([]*model.ProductBranch, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProductBranch), err
}

func (p productBranchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProductBranch, err error) {
	buf := make([]*model.ProductBranch, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productBranchDo) FindInBatches(result *[]*model.ProductBranch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productBranchDo) Attrs(attrs ...field.AssignExpr) IProductBranchDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productBranchDo) Assign(attrs ...field.AssignExpr) IProductBranchDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productBranchDo) Joins(fields ...field.RelationField) IProductBranchDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productBranchDo) Preload(fields ...field.RelationField) IProductBranchDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productBranchDo) FirstOrInit() (*model.ProductBranch, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBranch), nil
	}
}

func (p productBranchDo) FirstOrCreate() (*model.ProductBranch, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBranch), nil
	}
}

func (p productBranchDo) FindByPage(offset int, limit int) (result []*model.ProductBranch, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productBranchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productBranchDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productBranchDo) Delete(models ...*model.ProductBranch) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productBranchDo) withDO(do gen.Dao) *productBranchDo {
	p.DO = *do.(*gen.DO)
	return p
}
