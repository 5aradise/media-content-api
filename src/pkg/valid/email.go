package valid

import (
	"regexp"
)

var emailRe = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

func Email(email string) bool {
	return emailRe.MatchString(email)
}
