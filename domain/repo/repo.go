package repo

import (
	"gorm.io/gorm"
	"my-clean-rchitecture/domain/repo/cat"
	"my-clean-rchitecture/domain/repo/user"
)

type Repos struct {
	User user.RepoInterface
	Cat  cat.RepoInterface
}

func NewRepos(db *gorm.DB) Repos {
	var rs Repos
	rs.User = user.NewUserRepo(db)
	rs.Cat = cat.NewCatRepo(db)

	return rs
}
