package go_saas_security

import "errors"

var (
	NoTeamMemberErr    = errors.New("no team member")
	NoTeamCreatorError = errors.New("no team creator")
	NoTeamRoleError    = errors.New("no team role")
)
