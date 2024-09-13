// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package database

import (
	"context"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/database/db"
	"github.com/goravel/framework/database/gorm"
)

// Injectors from wire.go:

//go:generate wire
func InitializeOrm(ctx context.Context, config2 config.Config, connection string) (*OrmImpl, error) {
	configImpl := db.NewConfigImpl(config2, connection)
	dialectorImpl := gorm.NewDialectorImpl(config2, connection)
	gormImpl := gorm.NewGormImpl(config2, connection, configImpl, dialectorImpl)
	queryImpl, err := gorm.BuildQueryImpl(ctx, config2, connection, gormImpl)
	if err != nil {
		return nil, err
	}
	ormImpl, err := NewOrmImpl(ctx, config2, connection, queryImpl)
	if err != nil {
		return nil, err
	}
	return ormImpl, nil
}