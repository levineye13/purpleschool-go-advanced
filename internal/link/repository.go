package link

import "purpleschool-go/advanced/pkg/db"

type LinkRepository struct {
	db *db.Db
}

type LinkRepositoryDeps struct {
}

func NewLinkRepository(deps LinkRepositoryDeps) {
	return &LinkRepository{}
}
