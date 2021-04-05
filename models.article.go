// models.article.go

package main

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{}
var len_article int

// Return a list of all the articles
func getAllArticles() []article {
	id := 1
	len_article = 0
	title := "first"
	content := "first article"
	var articleList_empty = []article{}
	articleList = articleList_empty
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM article")
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	if !rows.Next() {
		return articleList_empty
	} else {
		err = rows.Scan(&id, &title, &content)
		if err != nil {
			// do something here
			log.Fatal(err)
		}
		article_t := article{
			ID:      id,
			Title:   title,
			Content: content,
		}
		articleList = append(articleList, article_t)
		len_article++
	}
	for rows.Next() {
		err = rows.Scan(&id, &title, &content)
		if err != nil {
			// do something here
			log.Fatal(err)
		}
		article_t := article{
			ID:      id,
			Title:   title,
			Content: content,
		}
		articleList = append(articleList, article_t)
		len_article++

	}
	return articleList
}

// Fetch an article based on the ID supplied
func getArticleByID(id int) (*article, error) {
	/*for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}*/
	title := "first"
	content := "first article"
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM article where ID =" + strconv.Itoa(id))
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	if !rows.Next() {
		return nil, errors.New("Article not found")
	} else {
		err = rows.Scan(&id, &title, &content)
		if err != nil {
			// do something here
			log.Fatal(err)
		}
		article_t := article{
			ID:      id,
			Title:   title,
			Content: content,
		}
		return &article_t, nil

	}

}

// Create a new article with the title and content provided
func createNewArticle(title, content string) (*article, error) {
	// Set the ID of a new article to one more than the number of articles
	len_article++
	a := article{ID: len_article, Title: title, Content: content}
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO article(Id,Title,Content) VALUES(" + strconv.Itoa(len_article) + ", '" + title + "','" + content + "')")
	if err != nil {
		log.Fatal(err)
	}
	// Add the article to the list of articles
	//articleList = append(articleList, a)

	return &a, nil
}
