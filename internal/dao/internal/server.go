// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServerDao is the data access object for table server.
type ServerDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ServerColumns // columns contains all the column names of Table for convenient usage.
}

// ServerColumns defines and stores column names for table server.
type ServerColumns struct {
	ServerId      string //
	ServerName    string //
	CreatorUserId string //
	CreationDate  string //
	CreateDate    string //
	ServerType    string //
	ServerImgUrl  string //
	DeletedAt     string //
}

// serverColumns holds the columns for table server.
var serverColumns = ServerColumns{
	ServerId:      "server_id",
	ServerName:    "server_name",
	CreatorUserId: "creator_user_id",
	CreationDate:  "creation_date",
	CreateDate:    "create_date",
	ServerType:    "server_type",
	ServerImgUrl:  "server_img_url",
	DeletedAt:     "deleted_at",
}

// NewServerDao creates and returns a new DAO object for table data access.
func NewServerDao() *ServerDao {
	return &ServerDao{
		group:   "default",
		table:   "server",
		columns: serverColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ServerDao) Columns() ServerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
