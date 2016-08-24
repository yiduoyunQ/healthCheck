package main

import (
	"flag"
	"time"
)

var (
	timeout     time.Duration
	readTimeout time.Duration
	mysqlDriver = "mysql"
	version     = "v1.0.0"

	fConfigFile  string
	fUser        string
	fPassword    string
	fDBName      string
	fTableName   string
	fTimeout     time.Duration
	fReadTimeout time.Duration
)

func init() {
	flag.StringVar(&fConfigFile, "default-file", "/DBAASDAT/my.cnf", "db config file")
	flag.StringVar(&fUser, "root-user", "check", "db username")
	flag.StringVar(&fPassword, "root-password", "123.com", "db user's password")
	flag.StringVar(&fDBName, "default-db", "dbaas_check", "db name")
	flag.StringVar(&fTableName, "default-table", "chk", "db tableName")
	flag.DurationVar(&fTimeout, "time-out", 5, "check timeout,x s")
	flag.DurationVar(&fReadTimeout, "read-time-out", 5, "db session timeout,x s")

	timeout = fTimeout * time.Second
	readTimeout = fReadTimeout * time.Second

	flag.Parse()
}
