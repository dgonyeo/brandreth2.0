package handler

import (
	"html/template"
	"net/http"

	"github.com/dgonyeo/brandreth2.0/config"
	"github.com/dgonyeo/brandreth2.0/db"
)

type PeoplePage struct {
	People      []*db.Person
	SearchQuery string
}

func (pp PeoplePage) IsActivePage(num int) bool {
	return num == 2
}

func (pp PeoplePage) TimeForRowStart(location int) bool {
	return location%4 == 0
}

func (pp PeoplePage) TimeForRowEnd(location int) bool {
	return location%4 == 3 || location == len(pp.People)-1
}

func (h *Handler) People(w http.ResponseWriter, req *http.Request) {
	people := h.c.GetPeople()

	t, err := template.ParseFiles(
		config.Config.Templates.Path+"templates/people.tmpl",
		config.Config.Templates.Path+"templates/stuff.tmpl")
	if err != nil {
		log.Error("Error parsing the templates: %v", err)
		return
	}
	err = t.Execute(w, PeoplePage{people, ""})
	if err != nil {
		log.Error("Error executing the templates: %v", err)
		return
	}
}
