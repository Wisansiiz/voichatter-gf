// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChannelDao is the data access object for table channel.
type ChannelDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ChannelColumns // columns contains all the column names of Table for convenient usage.
}

// ChannelColumns defines and stores column names for table channel.
type ChannelColumns struct {
	ChannelId    string //
	ChannelName  string //
	ServerId     string //
	Type         string //
	CreationDate string //
	CreateUserId string //
}

// channelColumns holds the columns for table channel.
var channelColumns = ChannelColumns{
	ChannelId:    "channel_id",
	ChannelName:  "channel_name",
	ServerId:     "server_id",
	Type:         "type",
	CreationDate: "creation_date",
	CreateUserId: "create_user_id",
}

// NewChannelDao creates and returns a new DAO object for table data access.
func NewChannelDao() *ChannelDao {
	return &ChannelDao{
		group:   "default",
		table:   "channel",
		columns: channelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChannelDao) Columns() ChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
