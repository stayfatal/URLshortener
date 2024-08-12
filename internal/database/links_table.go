package database

import _ "github.com/lib/pq"

type Link struct {
	Url    string `json:"url"`
	Alias  string `json:"alias"`
	Domain string
}

func (dm *DBManager) AddLink(link Link) error {
	_, err := dm.db.Exec("insert into links (url,shortened_url) values ($1,$2)", link.Url, link.Domain+link.Alias)
	if err != nil {
		return err
	}

	return nil
}

func (dm *DBManager) CheckLink(link string) (bool, error) {
	var exists bool
	err := dm.db.QueryRow("SELECT EXISTS(SELECT 1 FROM links WHERE shortened_url=$1)", link).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (dm *DBManager) GetFullUrl(shortenedUrl string) (string, error) {
	var url string
	err := dm.db.QueryRow("select url from links where shortened_url = $1", shortenedUrl).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}
