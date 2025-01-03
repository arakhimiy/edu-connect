package repository

import (
	"github.com/arakhimiy/edu-connect/internal/repository/sqlite"
	"github.com/pocketbase/dbx"
)

type AuthorizationI interface {
}

type I interface {
	Authorization() AuthorizationI
}

type repository struct {
	AuthorizationI
}

func (r *repository) Authorization() AuthorizationI {
	return r.AuthorizationI
}

func NewRepository(db dbx.Builder) I {
	return &repository{
		AuthorizationI: sqlite.NewAuthorization(db),
	}
}
