package common

import (
	"strings"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// DatabaseInfo struct
type DatabaseInfo struct {
	Username string
	Password string
	DBName   string
	HostIP   string
}

//MongoDBInfo struct
type MongoDBInfo struct {
	User       string
	Password   string
	Host       string
	Port       string
	Database   string
	Collection string
	URL        string
}

// DBReadConfig function
func DBReadConfig(profilename string) DatabaseInfo {
	var dbInfo DatabaseInfo
	config.Load(file.NewSource(
		file.WithPath("./common/dbconfig.json"),
	))

	dbInfo.DBName = config.Get("hosts", profilename, "dbname").String("")
	dbInfo.Username = config.Get("hosts", profilename, "username").String("")
	dbInfo.Password = config.Get("hosts", profilename, "password").String("")
	dbInfo.HostIP = config.Get("hosts", profilename, "hostip").String("")

	return dbInfo
}

// GetDatasourceName for connect db
func GetDatasourceName(profilename string) string {
	var dbInfo DatabaseInfo
	dbInfo = DBReadConfig(profilename)

	var constr string
	constr = dbInfo.Username + "/" + dbInfo.Password + "@" + dbInfo.HostIP + dbInfo.DBName

	return constr
}

// GetDatasourceNameMongo for connect mongodb
func GetDatasourceNameMongo(profilename string) MongoDBInfo {
	var mgDbInfo MongoDBInfo
	config.Load(file.NewSource(
		file.WithPath("./common/dbconfig.json"),
	))

	mgDbInfo.User = config.Get("hosts", profilename, "username").String("")
	mgDbInfo.Password = config.Get("hosts", profilename, "password").String("")
	mgDbInfo.Host = config.Get("hosts", profilename, "host").String("")
	mgDbInfo.Port = config.Get("hosts", profilename, "port").String("")
	mgDbInfo.Database = config.Get("hosts", profilename, "dbname").String("")

	var url string
	if mgDbInfo.User == "" || mgDbInfo.Password == "" {
		url = "mongodb://$host:$port/$db"
	} else {
		url = "mongodb://$user:$pass@$host:$port/$db"
	}
	url = strings.Replace(url, "$user", mgDbInfo.User, -1)
	url = strings.Replace(url, "$pass", mgDbInfo.Password, -1)
	url = strings.Replace(url, "$host", mgDbInfo.Host, -1)
	url = strings.Replace(url, "$port", mgDbInfo.Port, -1)
	url = strings.Replace(url, "$db", mgDbInfo.Database, -1)

	mgDbInfo.URL = url

	return mgDbInfo
}
