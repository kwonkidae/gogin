package db

import (
	"log"
	"os"
	"time"

	"encoding/json"

	"gopkg.in/mgo.v2"
)

type configuration struct {
	MongoDBHost, MongoDBUser, MongoDBPwd, Database string
}

var appConfig configuration

func init() {
	loadAppConfig()
	createDbSession()
}

func loadAppConfig() {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	appConfig = configuration{}
	err = decoder.Decode(&appConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

// 몽고 디비 세션 선언
var session *mgo.Session

// 세션을 얻어오는 함수
func GetSession() *mgo.Session {
	log.Println(appConfig)
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{appConfig.MongoDBHost},
			Username: appConfig.MongoDBUser,
			Password: appConfig.MongoDBPwd,
			Database: appConfig.Database,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

// 몽고 디비 세션 생성하는 함수.
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{appConfig.MongoDBHost},
		Username: appConfig.MongoDBUser,
		Password: appConfig.MongoDBPwd,
		Database: appConfig.Database,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[CreateDbSession]: %s\n", err)
	}
}
