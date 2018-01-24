package articles

type Article struct {
	Id      string `bson:"_id" json:"id"`
	Title   string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`

	OwnerId         string `bson:"owner_id" json:"owner_id"`
	CreateTimestamp int64 `bson:"create_timestamp" json:"create_timestamp"`
}
