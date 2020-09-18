package go_saas_security

import "errors"

var (
	NoTeamUserErr    = errors.New("no team user")
	NoTeamCreatorError = errors.New("no team creator")
	NoTeamRoleError    = errors.New("no team role")
)
