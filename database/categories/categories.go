package categories

import (
	"modding_server/database"
)

type Category struct {
	Id   int
	Name string
}

const table string = "categories"

func getArgs(ctg *Category) []any {
	return []any{&ctg.Id, &ctg.Name}
}

func Create(name string) (Category, error) {
	var ctg Category = Category{Name: name}
	args := getArgs(&ctg)
	err := database.Create(table, "name", "$1", args)
	return ctg, err
}

func Read(id int) (Category, error) {
	var ctg Category = Category{Id: id}
	args := getArgs(&ctg)
	err := database.Read(table, args)
	return ctg, err
}

func Update(category Category) (Category, error) {
	var ctg Category = category
	args := getArgs(&ctg)
	err := database.Update(table, "name = $2", args)
	return ctg, err
}

func Delete(id int) (Category, error) {
	var ctg Category = Category{Id: id}
	args := getArgs(&ctg)
	err := database.Delete(table, args)
	return ctg, err
}

func GetMods(id int) ([]int, error) {
	return database.GetList("mods_categories", "mod_id", "category_id", id)
}
