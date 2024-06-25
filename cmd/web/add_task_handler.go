package web

import (
	"net/http"
	"strconv"
	"time"

	"github.com/heitorfreitasferreira/go-project-manager/internal/models"
	usecases "github.com/heitorfreitasferreira/go-project-manager/internal/use-cases"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {

	projectId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	startDate, err := time.Parse("2006-01-02T15:04", r.FormValue("start_date"))
	if err != nil {
		http.Error(w, "Invalid Start Date", http.StatusBadRequest)
	}
	endDate, err := time.Parse("2006-01-02T15:04", r.FormValue("end_date"))
	if err != nil {
		http.Error(w, "Invalid End Date", http.StatusBadRequest)
	}
	status := models.NotStarted

	in := usecases.CreateTaskIn{
		ProjectId:   projectId,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		StartDate:   &startDate,
		EndDate:     &endDate,
		Owner:       "self",
		Status:      &status,
	}
	usecases.CreateTask.Execute(in)
}
