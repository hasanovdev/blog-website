package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/hasanovdev/personal-blog/config"
	"github.com/hasanovdev/personal-blog/models"

	_ "github.com/go-sql-driver/mysql"
)

var temp *template.Template
var db *sql.DB

func init() {
	temp = template.Must(template.ParseGlob("templates/*.html"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "services.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	db = config.ConnectDB()
	if r.Method != http.MethodPost {
		temp.ExecuteTemplate(w, "contact.html", nil)
		return
	}
	user := models.User{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	if r.FormValue("submit") == "SEND" {
		_, err := db.Exec("insert into blog_users(name,email,message) values(?,?,?)", user.Name, user.Email, user.Message)
		if err != nil {
			temp.Execute(w, struct {
				Success bool
				Msg     string
			}{true, err.Error()})
		} else {
			temp.Execute(w, struct {
				Success bool
				Msg     string
			}{true, "Message sent!"})
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/services", servicesHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Server running at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
