package makeless_go_security

import "errors"

var (
	NoTeamUserErr       = errors.New("no team user")
	NoTeamCreatorError  = errors.New("no team creator")
	NoTeamRoleError     = errors.New("no team role")
	NoEmailVerification = errors.New("no email verification")
	UserAlreadyExist    = errors.New("you are already registred")
	UserNotDeletable    = errors.New("user not deletable")
)
