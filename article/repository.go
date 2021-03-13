package article

import (
	"github.com/fahimsgit/testCore/configuration"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type DBStruct struct {
	Session   *mgo.Session
	DbName    string
	TableName string
}
type DBInterface interface {
	insertArticle()
	findArticle()
	newDB()
}

func newDB(session *mgo.Session) *DBStruct {
	return &DBStruct{
		Session:   session,
		DbName:    configuration.DbName,
		TableName: configuration.TableName,
	}
}
func (d *DBStruct) findArticle(id int, article Article) (Article, error) {

	err := d.Session.DB(d.DbName).C(d.TableName).Find(bson.M{"id": id}).One(&article)
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

func (d *DBStruct) insertArticle(article Article) (Article, error) {

	table := d.Session.DB(d.DbName).C(d.TableName)
	err := table.Insert(article)
	if err != nil {
		return Article{}, err
	} else {
		return article, nil
	}
}
