package database

import (
	"github.com/vinitkantrod/rapido-coupons/logging"
	mgo "gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	// DBname : database name string
	DBname string
	// MDB : mongodb database session
	MDB *mgo.Database
	// UserCollection :
	UserCollection *mgo.Collection
	// RiderCollection :
	RiderCollection *mgo.Collection
	// OrderCollection :
	OrderCollection *mgo.Collection
)

// Init initialized the database collection session
func Init(host, dbname string) {
	DBname = dbname
	session, err := mgo.Dial(host)
	MDB = session.DB(DBname)
	if err != nil {
		logging.Error.Println("Error connecting database : ", err)
		// panic(err)
	}
	// UserCollection : referral collection
	UserCollection = MDB.C("users")
	// RiderCollection : referral Usage collection
	RiderCollection = MDB.C("riders")
	// OrderCollection : Referral Metadata collection
	OrderCollection = MDB.C("orders")
	// ReferralRuleCollection : Referral Rule collection
}

func close() {
	session.Close()
}

// // Sample Queries -------------------------

// func Get(collection *mgo.Collection, id string, i interface{}) {
// 	collection.FindId(bson.ObjectIdHex(id)).One(i)
// }

// func Get2(collection *mgo.Collection, id bson.ObjectId, i interface{}) {
// 	collection.FindId(id).One(i)
// }
