package sessions

type Session struct {
	Id                  string `bson:"_id" json:"id"`
	Cookie              string `bson:"cookie" json:"cookie"`
	CreateTimestamp     int64  `bson:"create_timestamp" json:"create_timestamp"`
	ExpirationTimestamp int64  `bson:"expiration_timestamp" json:"expiration_timestamp"`
	UserId              string `bson:"user_id" json:"user_id"`
}
