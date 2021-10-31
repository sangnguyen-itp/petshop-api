package user_data

import (
	"errors"
	"petshop/models"
)

var users = []models.User{
	{
		ID: "ps-1",
		FirstName: "Nguyen",
		LastName:  "A",
		Avatar:    "",
		Disabled:  false,
	},
	{
		ID: "ps-2",
		FirstName: "Nguyen",
		LastName:  "B",
		Avatar:    "",
		Disabled:  false,
	},
	{
		ID: "ps-3",
		FirstName: "Nguyen",
		LastName:  "C",
		Avatar:    "",
		Disabled:  false,
	},
	{
		ID: "ps-4",
		FirstName: "Nguyen",
		LastName:  "D",
		Avatar:    "",
		Disabled:  false,
	},
	{
		ID: "ps-5",
		FirstName: "Nguyen",
		LastName:  "E",
		Avatar:    "",
		Disabled:  false,
	},
}

var userData map[string]models.User

func init() {
  	userData = make(map[string]models.User, 0)
  	for _, user := range users {
  		userData[user.ID] = user
	}
}

func getUserStorage() map[string]models.User {
	return userData
}

func Find(id string) (models.User, error) {
	s := getUserStorage()
	user, found := s[id]
	if !found {
		return models.User{}, errors.New("record not found")
	}
	return user, nil
}

func FindAll(filterIDs []string) ([]models.User, error) {
	if len(filterIDs) == 0 {
		return users, nil
	}
	var result []models.User
	for _, id := range filterIDs {
		user, err := Find(id)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}