// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FriendshipDao is the data access object for table friendship.
type FriendshipDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns FriendshipColumns // columns contains all the column names of Table for convenient usage.
}

// FriendshipColumns defines and stores column names for table friendship.
type FriendshipColumns struct {
	FriendshipId string // 关系id
	UserId1      string // 用户1
	UserId2      string // 用户2
	Date         string // 日期
	DeletedAt    string // 删除日期
}

// friendshipColumns holds the columns for table friendship.
var friendshipColumns = FriendshipColumns{
	FriendshipId: "friendship_id",
	UserId1:      "user_id1",
	UserId2:      "user_id2",
	Date:         "date",
	DeletedAt:    "deleted_at",
}

// NewFriendshipDao creates and returns a new DAO object for table data access.
func NewFriendshipDao() *FriendshipDao {
	return &FriendshipDao{
		group:   "default",
		table:   "friendship",
		columns: friendshipColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FriendshipDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FriendshipDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FriendshipDao) Columns() FriendshipColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FriendshipDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FriendshipDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FriendshipDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
