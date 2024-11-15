package types

import (
	"errors"
	"fmt"
)

var (
	ErrNameTooLong      = fmt.Errorf("internal/database: name length exceeds %d bytes", NameMaxLen)
	ErrEmailTooLong     = fmt.Errorf("internal/database: email length exceeds %d bytes", EmailMaxLen)
	ErrUserEmailExists  = errors.New("internal/database: user with this email exists")
	ErrUserIdNotExists  = errors.New("internal/database: user with this id not exists")
	ErrUserFKConstraint = errors.New("internal/database: update or delete users violates foreign key constraint media_content table")

	ErrTitleTooLong            = fmt.Errorf("internal/database: title length exceeds %d bytes", TitleMaxLen)
	ErrDescriptionTooLong      = fmt.Errorf("internal/database: description length exceeds %d bytes", DescriptionMaxLen)
	ErrMediaContentIdNotExists = errors.New("internal/database: media content with this id not exists")
)
