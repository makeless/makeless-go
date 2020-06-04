package saas_database

import (
	"github.com/go-saas/go-saas/model"
)

func (database *Database) GetMembersTeam(users []*go_saas_model.User, teamId uint, userId uint) ([]*go_saas_model.User, error) {
	return users, database.GetConnection().
		Model(&go_saas_model.Team{
			Model:  go_saas_model.Model{Id: teamId},
			UserId: &userId,
		}).
		Related(&users, "Users").
		Error
}

func (database *Database) RemoveMemberTeam(user *go_saas_model.User, teamId uint, userId uint) error {
	return database.GetConnection().
		Model(&go_saas_model.User{
			Model: go_saas_model.Model{Id: user.GetId()},
		}).
		Association("Teams").
		Delete(&go_saas_model.Team{
			Model:  go_saas_model.Model{Id: teamId},
			UserId: &userId,
		}).
		Error
}
