package main

import (
	"html/template"
	"log"
	"net/http"
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
	Title		string
	Subtitle	string
	ImgModifier	string
	Author		string
	AuthorImg	string
	PublishDate	string
}

type mostRecentPostData struct {
	Title		string
	Subtitle	string
	ImgModifier	string
	Author		string
	AuthorImg	string
	PublishDate	string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:				"Escape",
		FeaturedPosts:		featuredPosts(),
		MostRecentPosts:	mostRecentPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
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
		Title:		"The Road Ahead",
		Content:	"The road ahead might be paved - it might not be.",
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:			"The Road Ahead",
	    	Subtitle:		"The road ahead might be paved - it might not be.",
	    	ImgModifier:	"featured-posts__post_the-road-ahead",
	    	Author:			"Mat Vogels",
	    	AuthorImg:		"static/images/Mat_Volges.png",
	    	PublishDate:	"September 25, 2015",
		},
		{
			Title:			"From Top Down",
			Subtitle:		"Once a year, go someplace you've never been before.",
			ImgModifier:	"featured-posts__post_from-top-down",
	    	Author:			"Mat Vogels",
	    	AuthorImg:		"static/images/William_Wong.png",
	    	PublishDate:	"September 25, 2015",
		},
	}
}

func mostRecentPosts() []mostRecentPostData {
	return []mostRecentPostData{
		{
			ImgModifier: "most-recent-posts__post_still-standing-tall",
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			Author:      "William Wong",
			AuthorImg:   "william-wong-face",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "most-recent-posts__post_sunny-side-up",
			Title:       "Sunny Side Upl",
			Subtitle:    "No place is ever as bad as they tell you it's going to be.",
			Author:      "Mat Vogels",
			AuthorImg:   "mat-vogels-face",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "most-recent-posts__post_water-falls",
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			Author:      "Mat Vogels",
			AuthorImg:   "mat-vogels-face",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "most-recent-posts__post_through-the-mist",
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			Author:      "William Wong",
			AuthorImg:   "william-wong-face",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "most-recent-posts__post_awaken-early",
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			Author:      "Mat Vogels",
			AuthorImg:   "mat-vogels-face",
			PublishDate: "9/25/2015",
		},
		{
			ImgModifier: "most-recent-posts__post_try-it-always",
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			Author:      "Mat Vogels",
			AuthorImg:   "mat-vogels-face",
			PublishDate: "9/25/2015",
		},
	}
}