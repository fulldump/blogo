package articles

type Article struct {
	Id      string `bson:"_id" json:"id"`
	Title   string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
}
