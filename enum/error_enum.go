package enum

import (
	"errors"
)

var (
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrRigsterUserError = errors.New("注册用户异常")
	// ErrNotImplemented not implemented
)
