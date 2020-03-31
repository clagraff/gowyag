package commands

import "github.com/clagraff/gowyag/repo"

type Init struct {
	repo repo.Repository
}

func (_ Init) Name() string { return "init" }

func (i *Init) Exec(args []string) error {
	return nil
}
