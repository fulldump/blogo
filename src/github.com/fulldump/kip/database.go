package kip

import (
	"sort"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Database struct {
	addrs   string
	name    string // TODO: rename to `name`
	session *mgo.Session
}

var MONGO_DIAL_TIMEOUT = 0 * time.Second
var MONGO_SYNC_TIMEOUT = 3 * time.Second
var MONGO_SOCKET_TIMEOUT = 3 * time.Second

var sessions_by_addrs = map[string]*mgo.Session{}

func NewDatabase(mongourl string) (*Database, error) {

	url, url_err := mgo.ParseURL(mongourl)
	if nil != url_err {
		panic(url_err)
	}

	addrs := normalize_addrs(url.Addrs)

	// Initialize Database
	db := &Database{
		addrs:   addrs,
		name:    url.Database,
		session: nil,
	}

	// Check if a session already exists
	if previous_session, exists := sessions_by_addrs[addrs]; exists {
		db.session = previous_session
		return db, nil
	}

	// Create new session
	session, err := mgo.DialWithTimeout(mongourl, MONGO_DIAL_TIMEOUT)
	if nil != err {
		return nil, err
	}

	// Enable autoreconnect and autoreset
	session.SetMode(mgo.Eventual, true)

	// Set timeouts
	session.SetSyncTimeout(MONGO_SYNC_TIMEOUT)
	session.SetSocketTimeout(MONGO_SOCKET_TIMEOUT)

	sessions_by_addrs[addrs] = session

	db.session = session

	return db, nil
}

func Close(addrs string) {
	if session, exists := sessions_by_addrs[addrs]; exists {
		session.Close()
	}
}

func CloseAll() {
	for i, session := range sessions_by_addrs {
		delete(sessions_by_addrs, i)
		session.Close()
	}
}

func (d *Database) C(collection string) *mgo.Collection {
	return d.session.DB(d.name).C(collection)
}

func (d *Database) GetName() string {
	return d.name
}

func (d *Database) Clone() *Database {
	return &Database{
		addrs:   d.addrs,
		name:    d.name,
		session: d.session.Clone(),
	}
}

func (d *Database) Close() {
	d.session.Close()
}

func normalize_addrs(addrs []string) string {

	sort.Strings(addrs)

	return strings.Join(addrs, ",")
}
