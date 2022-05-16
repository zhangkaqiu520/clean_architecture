package repo

import (
	"gorm.io/gorm"
	"my-clean-rchitecture/domain/repo/cat"
	"my-clean-rchitecture/domain/repo/goods"
	"my-clean-rchitecture/domain/repo/user"
)

type Repos struct {
	User  user.RepoInterface
	Cat   cat.RepoInterface
	Goods goods.RepoInterface
}

func NewRepos(db *gorm.DB) Repos {
	var rs Repos
	rs.User = user.NewUserRepo(db)
	rs.Cat = cat.NewCatRepo(db)
	rs.Goods = goods.NewGoodsRepo(db)

	return rs
}
