package models

import (
	"context"

	"github.com/qiniu/qmgo"
	log "github.com/sirupsen/logrus"
	"wozaizhao.com/api/common"
)

var db *qmgo.Database

var wdjky *qmgo.Collection

var ctx context.Context

// InitDB 初始化数据库
func InitDB() {
	mongoURI := common.GetMongodbURL()
	ctx = context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoURI})
	if err != nil {
		log.Error("InitDB error", err)
	}
	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()
	db = client.Database("wzz")
	wdjky = db.Collection("wdjky")
}
