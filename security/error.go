package saas_security

import "errors"

var (
	NoTeamMemberErr  = errors.New("no team member")
	NoTeamOwnerError = errors.New("no team owner")
)
