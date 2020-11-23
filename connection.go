package lorm

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var DefaultConnection = "default"

func init()  {
	for k,v := range dataBaseConfig{
		dsn := v.FormatDSN()
		connection,_ := sql.Open("mysql",dsn)
		connections[k] = connection
	}
}

var (
	dataBaseConfig =  map[string]mysql.Config{
		"default" : {
			User:                    "ms_pre_code",
			Passwd:                  "Rk9jCQOVX8GXgrW3",
			Net:                     "tcp",
			Addr:                    "127.0.0.1:3308",
			DBName:                  "class",
			Params:                  nil,
			Collation:               "",
			Loc:                     nil,
			MaxAllowedPacket:        0,
			ServerPubKey:            "",
			TLSConfig:               "",
			Timeout:                 0,
			ReadTimeout:             0,
			WriteTimeout:            0,
			AllowAllFiles:           false,
			AllowCleartextPasswords: false,
			AllowNativePasswords:    false,
			AllowOldPasswords:       false,
			CheckConnLiveness:       false,
			ClientFoundRows:         false,
			ColumnsWithAlias:        false,
			InterpolateParams:       false,
			MultiStatements:         false,
			ParseTime:               false,
			RejectReadOnly:          false,
		},
	}
	connections = map[string]*sql.DB{}
)

func getConnection(name string)*sql.DB  {
	_,ok := connections[name]
	if !ok {
		name = DefaultConnection
	}
	return connections[name]
}