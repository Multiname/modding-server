package files

import (
	"modding_server/database"
)

type File struct {
	Id   int
	Path string
}

const table string = "files"

func getArgs(fl *File) []any {
	return []any{&fl.Id, &fl.Path}
}

func Create(path string) (File, error) {
	var fl File = File{Path: path}
	args := getArgs(&fl)
	err := database.Create(table, "path", "$1", args)
	return fl, err
}

func Read(id int) (File, error) {
	var fl File = File{Id: id}
	args := getArgs(&fl)
	err := database.Read(table, args)
	return fl, err
}

func Update(file File) (File, error) {
	var fl File = file
	args := getArgs(&fl)
	err := database.Update(table, "path = $2", args)
	return fl, err
}

func Delete(id int) (File, error) {
	var fl File = File{Id: id}
	args := getArgs(&fl)
	err := database.Delete(table, args)
	return fl, err
}
