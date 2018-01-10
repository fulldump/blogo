package sessions

type Session struct {
	Id              string                 `bson:"_id" json:"id"`
	Cookie          string                 `bson:"cookie" json:"cookie"`
	CreateTimestamp int64                  `bson:"create_timestamp" json:"create_timestamp"`
	ExpireTimestamp int64                  `bson:"expire_timestamp" json:"expire_timestamp"`
	UserId          string                 `bson:"user_id" json:"user_id"`
	Data            map[string]interface{} `bson:"data" json:"data"`
}
