package router

import (
	"html/template"
	"net/http"
	"os"
	"pogle-form/backend/data/dtos"

	"github.com/pecet3/logger"
)

type CoursePage struct {
	Course      *dtos.Course
	IsAvailable bool
}

type PageData struct {
	Courses []*CoursePage
	Title   string
}

func (r router) handleViewForm(w http.ResponseWriter, req *http.Request) {
	coursesData := PageData{
		Courses: []*CoursePage{},
		Title:   os.Getenv("FORM_TITLE"),
	}

	coursesDB, err := r.app.Data.ListCourses(req.Context())
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	for _, c := range coursesDB {
		cp := &CoursePage{
			Course:      c.ToDto(r.app.Data),
			IsAvailable: r.app.Course.CheckIfIsPlace(c.ID),
		}
		coursesData.Courses = append(coursesData.Courses, cp)
	}
	tmpl, err := template.ParseFiles("static/course_register.html")
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, coursesData)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
