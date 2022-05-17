package handler

import (
	"database/sql"
	"fmt"
	"my-clean-rchitecture/app"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Connect() (*migrate.Migrate, error) {
	db, err := sql.Open("mysql", app.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "执行迁移，连接数据库失败")
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "执行迁移，获取driver失败")
	}

	return migrate.NewWithDatabaseInstance(
		"file:./migrations",
		"mysql", driver)
}

func Up(c *cli.Context) error {
	m, err := Connect()
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "迁移UP失败")
	}

	return nil
}

func Down(c *cli.Context) error {
	fmt.Println("开始执行迁移回滚")
	m, err := Connect()
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "迁移DOWN失败")
	}

	return nil
}
