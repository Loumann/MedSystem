package repository

import (
	"somename/models"
)

var waitingUsers models.WaitingUsers

func (r *Repository) AppendWaitUser(user *models.User) error {
	waitingUsers.Lock()

	waitingUsers.Items = append(waitingUsers.Items, *user)

	waitingUsers.Unlock()

	return nil
}

func (r *Repository) GetWaitingUsers() ([]models.User, error) {
	return waitingUsers.Items, nil
}

func (r *Repository) RemoveWaitingUser(userID int) error {
	waitingUsers.Lock()

	var users []models.User

	for _, item := range waitingUsers.Items {
		if item.ID != userID {
			waitingUsers.Items = append(users, item)
		}
	}
	waitingUsers.Items = users

	waitingUsers.Unlock()

	return nil
}
