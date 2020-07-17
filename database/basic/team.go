package go_saas_basic_database

import (
	"github.com/go-saas/go-saas/model"
	go_saas_security "github.com/go-saas/go-saas/security"
	"github.com/jinzhu/gorm"
)

// CreateTeam creates team and retrieves the new team with all their users
func (database *Database) CreateTeam(team *go_saas_model.Team) (*go_saas_model.Team, error) {
	return team, database.GetConnection().
		Create(team).
		Preload("TeamUsers.Team").
		Preload("TeamUsers.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("users.id, users.name, users.email")
		}).
		Find(team).
		Error
}

// DeleteTeamUsers deletes own teamUser
// FIXME: Delete all other teamUsers if user is team creator
func (database *Database) DeleteTeamUsers(user *go_saas_model.User, team *go_saas_model.Team) error {
	return database.GetConnection().
		Exec("DELETE FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ?", team.GetId(), user.GetId()).
		Error
}

func (database *Database) DeleteTeam(user *go_saas_model.User, team *go_saas_model.Team) error {
	return database.GetConnection().
		Exec("DELETE FROM teams WHERE teams.id = ? AND teams.user_id = ?", team.GetId(), user.GetId()).
		Error
}

func (database *Database) IsTeamMember(team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, database.GetConnection().
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}

func (database *Database) IsTeamOwner(team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, database.GetConnection().
		Raw("SELECT COUNT(*) FROM team_users WHERE team_users.team_id = ? AND team_users.user_id = ? AND team_users.role = ? LIMIT 1", team.GetId(), user.GetId(), go_saas_security.RoleTeamOwner).
		Count(&count).
		Error
}

func (database *Database) IsTeamCreator(team *go_saas_model.Team, user *go_saas_model.User) (bool, error) {
	var count int

	return count == 1, database.GetConnection().
		Raw("SELECT COUNT(*) FROM teams WHERE teams.id = ? AND teams.user_id = ? LIMIT 1", team.GetId(), user.GetId()).
		Count(&count).
		Error
}
