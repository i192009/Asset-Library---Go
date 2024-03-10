package configs

import (
	"github.com/joho/godotenv"
	"gitlab.zixel.cn/go/framework/config"
	"time"
)

var (
	_           = godotenv.Load()
	Timeout     = 600 * time.Second
	ServiceName = "assetLibary"
	ServiceUUID = "UUID-af7455-98ab372-652c1bb"

	RealinkHttpAddr = config.GetString("http.connections.realink", "http://localhost:6080")
	RealinkMarks    = config.GetString("realink.marks", "local")

	ObsEndPoint    = config.GetString("obs.endpoint", "")
	ObsAccessKey   = config.GetString("obs.access_key", "")
	ObsSecretKey   = config.GetString("obs.secret_key", "")
	ObsDlBucket    = config.GetString("obs.dl_bucket", "")
	ObsUlBucket    = config.GetString("obs.ul_bucket", "")
	ObsUlFolder    = config.GetString("obs.ul_folder", "")
	ObsScopeId     = config.GetInt("obs.scope_id", 0)
	ObsInstanceMap = config.GetObject("obs.instance")
)
