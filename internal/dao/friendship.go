// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"voichatter/internal/dao/internal"
)

// internalFriendshipDao is internal type for wrapping internal DAO implements.
type internalFriendshipDao = *internal.FriendshipDao

// friendshipDao is the data access object for table friendship.
// You can define custom methods on it to extend its functionality as you wish.
type friendshipDao struct {
	internalFriendshipDao
}

var (
	// Friendship is globally public accessible object for table friendship operations.
	Friendship = friendshipDao{
		internal.NewFriendshipDao(),
	}
)

// Fill with you ideas below.
