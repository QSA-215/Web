package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title			string
	FeaturedPosts	[]featuredPostData
	MostRecentPosts	[]mostRecentPostData
}

type postPage struct {
	Title	string
	Content	string
}

type featuredPostData struct {
	Title		string `db:"title"`
	Subtitle	string `db:"subtitle"`
	ImgModifier	string `db:"image_url"`
	Author		string `db:"author"`
	AuthorImg	string `db:"author_url"`
	PublishDate	string `db:"publish_date"`
}

type mostRecentPostData struct {
	Title		string `db:"title"`
	Subtitle	string `db:"subtitle"`
	ImgModifier	string `db:"image_url"`
	Author		string `db:"author"`
	AuthorImg	string `db:"author_url"`
	PublishDate	string `db:"publish_date"`
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPosts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)	// В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return	// Не забываем завершить выполнение ф-ии
		}

		mostRecentPosts, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		data := indexPage{
			Title:				"Escape",
			FeaturedPosts:		featuredPosts,
			MostRecentPosts:	mostRecentPosts,
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := postPage{
		Title:   "The Road Ahead",
		Content: "The road ahead might be paved - it might not be.",
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url
		FROM
			post
		WHERE featured = 1
	`

	var featuredPosts []featuredPostData

	err := db.Select(&featuredPosts, query)
	if err != nil {
		return nil, err
	}

	return featuredPosts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url
		FROM
			post
		WHERE featured = 0
	`
	var mostRecentPosts []mostRecentPostData

	err := db.Select(&mostRecentPosts, query)
	if err != nil {
		return nil, err
	}

	return mostRecentPosts, nil
}