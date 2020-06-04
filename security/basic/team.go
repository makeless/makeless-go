package go_saas_basic_security

func (security *Security) IsTeamMember(teamId uint, userId uint) (bool, error) {
	var count int

	err := security.getDatabase().GetConnection().
		Raw("SELECT COUNT(*) FROM user_teams WHERE user_teams.team_id = ? AND user_teams.user_id = ? LIMIT 1", teamId, userId).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (security *Security) IsTeamOwner(teamId uint, userId uint) (bool, error) {
	var count int

	err := security.getDatabase().GetConnection().
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", teamId, userId).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count == 1, nil
}
