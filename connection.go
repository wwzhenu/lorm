package lorm

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var DefaultConnection = "default"

func init()  {
	for k,v := range dataBaseConfig{
		dsn := v.FormatDSN()
		connection,_ := sql.Open("mysql",dsn)
		err := connection.Ping()
		if err != nil {
			log.Println("连接失败",dsn)
			panic(err)
		}
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
			DBName:                  "base_info",
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
			AllowNativePasswords:    true,
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

func GetConnection(name string)*sql.DB  {
	_,ok := connections[name]
	if !ok {
		name = DefaultConnection
	}
	return connections[name]
}