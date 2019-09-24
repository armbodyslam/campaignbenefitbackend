package common

import (
	"fmt"
	"os"
	"strings"

	cf "github.com/spf13/viper"
)

var tagenv = os.Getenv("ENVAPP")

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

	cf.SetConfigName("dbconfig")
	cf.AddConfigPath("./common")
	cf.AutomaticEnv()

	// แปลง _ underscore ใน env เป็น . dot notation ใน viper
	cf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// อ่าน config
	err := cf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	dbInfo.DBName = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".dbname"))
	dbInfo.Username = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".username"))
	dbInfo.Password = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".password"))
	dbInfo.HostIP = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".hostip"))

	return dbInfo
}

// GetDatasourceName for connect db
func GetDatasourceName(profilename string) string {
	var dbInfo DatabaseInfo
	dbInfo = DBReadConfig(profilename)

	var constr string
	constr = dbInfo.Username + "/" + dbInfo.Password + "@" + dbInfo.DBName

	return constr
}

// GetDatasourceNameMongo for connect mongodb
func GetDatasourceNameMongo(profilename string) MongoDBInfo {
	var mgDbInfo MongoDBInfo
	cf.SetConfigName("dbconfig")
	cf.AddConfigPath("./common")
	cf.AutomaticEnv()

	// แปลง _ underscore ใน env เป็น . dot notation ใน viper
	cf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// อ่าน config
	err := cf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	mgDbInfo.User = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".username"))
	mgDbInfo.Password = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".password"))
	mgDbInfo.Host = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".host"))
	mgDbInfo.Port = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".port"))
	mgDbInfo.Database = fmt.Sprintf("%v", cf.Get(tagenv+"."+profilename+".dbname"))

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
