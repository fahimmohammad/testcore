package article

//Export DB
/*type DBSession struct {
	Session   *mgo.Session
	DbName    string
	TableName string
}*/

type Article struct {
	ID          int32  `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	IsPublished bool   `json:"ispublished" bson:"ispublished"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
}

/*const (
	DBname    = "testDb"
	Tablename = "article"
)*/
