package go_saas_security_basic

import (
	"github.com/go-saas/go-saas/model"
)

func (security *Security) IsTeamInvitation(teamInvitation *go_saas_model.TeamInvitation) (bool, error) {
	return security.GetDatabase().IsTeamInvitation(security.GetDatabase().GetConnection(), teamInvitation)
}
