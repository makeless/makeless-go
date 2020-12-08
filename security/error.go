package makeless_go_security

import "errors"

var (
	NoTeamUserErr       = errors.New("no team user")
	NoTeamCreatorError  = errors.New("no team creator")
	NoTeamRoleError     = errors.New("no team role")
	NoEmailVerification = errors.New("no email verification")
	UserAlreadyExists   = errors.New("user already exists")
	UserNotDeletable    = errors.New("user not deletable")
)
