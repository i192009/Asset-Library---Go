package dao

import (
	"gitlab.zixel.cn/go/framework/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var classDB *mongo.Collection
var tagDB *mongo.Collection
var configDB *mongo.Collection
var assetDB *mongo.Collection
var publicAssetDB *mongo.Collection
var privateAssetDB *mongo.Collection
var tenantAssetDB *mongo.Collection
var instanceScopeIdDB *mongo.Collection

func Init() {
	classDB = database.GetCollection("class")
	tagDB = database.GetCollection("tag")
	configDB = database.GetCollection("config")
	assetDB = database.GetCollection("asset")
	publicAssetDB = database.GetCollection("publicAsset")
	privateAssetDB = database.GetCollection("privateAsset")
	tenantAssetDB = database.GetCollection("tenantAsset")
	instanceScopeIdDB = database.GetCollection("instanceScopeID")
}
