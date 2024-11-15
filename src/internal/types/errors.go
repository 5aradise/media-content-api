package types

import (
	"errors"
	"fmt"
)

var (
	ErrNameTooLong     = fmt.Errorf("internal/database: name length exceeds %d bytes", NameMaxLen)
	ErrEmailTooLong    = fmt.Errorf("internal/database: email length exceeds %d bytes", EmailMaxLen)
	ErrUserEmailExists = errors.New("internal/database: user with this email exists")
	ErrUserIdNotExists = errors.New("internal/database: user with this id not exists")
)
