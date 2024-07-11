package config

import (
	"os"
	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRINGS")

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "TBAlfian",
}

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)