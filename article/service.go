package article

type Service struct {
	db *DBStruct
}

type ServiceInterface interface {
	PostArticleInService(article Article)
	GetArticleInService(articleId string)
	newService()
}

func newService(dbs *DBStruct) *Service {
	return &Service{
		db: dbs,
	}
}

func (service *Service) PostArticleInService(article Article) (Article, error) {

	article, err := service.db.insertArticle(article)
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

func (service *Service) GetArticleInService(articleId int) (Article, error) {

	article, err := service.db.findArticle(articleId)

	if err != nil {
		return Article{}, err
	} else {
		return article, nil
	}
}

func (service *Service) UpdateArticleInService(id int, article Article) (Article, error) {
	article, err := service.db.updateArticle(id, article)
	if err != nil {
		return Article{}, err
	}
	return article, err
}
func (service *Service) DeleteArticleInService(id int) error {
	err := service.db.deleteArticle(id)
	return err
}
