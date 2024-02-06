package users

import (
	"modding_server/database"
)

type User struct {
	Id       int
	Name     string
	Login    string
	Password string
	Donation *string
}

const table string = "users"

func getArgs(usr *User) []any {
	return []any{&usr.Id, &usr.Name, &usr.Login, &usr.Password, &usr.Donation}
}

func Create(name string, login string, password string, donation *string) (User, error) {
	var usr User = User{Name: name, Login: login, Password: password, Donation: donation}
	args := getArgs(&usr)
	err := database.Create(table, "name, login, password, donation", "$1, $2, $3, $4", args)
	return usr, err
}

func Read(id int) (User, error) {
	var usr User = User{Id: id}
	args := getArgs(&usr)
	err := database.Read(table, args)
	return usr, err
}

func Update(user User) (User, error) {
	var usr User = user
	args := getArgs(&usr)
	err := database.Update(table, "name = $2, login = $3, password = $4, donation = $5", args)
	return usr, err
}

func Delete(id int) (User, error) {
	var usr User = User{Id: id}
	args := getArgs(&usr)
	err := database.Delete(table, args)
	return usr, err
}

func GetMods(id int) ([]int, error) {
	return database.GetList("mods", "id", "owner_id", id)
}
