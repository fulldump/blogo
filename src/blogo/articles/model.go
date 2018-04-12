package articles

type Article struct {
	Id      string `bson:"_id" json:"id"`
	Title   string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`

	// TODO: remove deprecated:
	OwnerId string `bson:"owner_id" json:"owner_id"`
	User    User   `bson:"user" json:"user"`

	CreateTimestamp int64 `bson:"create_timestamp" json:"create_timestamp"`
}

type User struct {
	Id   string `bson:"_id" json:"id"`
	Nick string `bson:"nick" json:"nick"`
	// TODO: picture

	SyncTimestamp int64 `bson:"sync_timestamp" json:"-"`
}
