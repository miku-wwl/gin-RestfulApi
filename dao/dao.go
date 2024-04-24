package dao

import (
	"context"
	"restApi/config"
	"restApi/pkg/logger"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.Mysqldb)

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}

	if Db.Error != nil {
		logger.Error(map[string]interface{}{"database error": Db.Error})
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
}

var (
	MonCollection *mongo.Collection
)

func init() {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(config.MongoUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error(map[string]interface{}{"mongo.Connect error": err.Error()})
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error(map[string]interface{}{"client.Ping error": err.Error()})
	}
	databaseName := config.MongoDatabaseName
	db := client.Database(databaseName)
	collectionName := config.MongoCollectionName
	MonCollection = db.Collection(collectionName)
}
