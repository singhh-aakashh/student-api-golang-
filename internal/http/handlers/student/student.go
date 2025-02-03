package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/singhh-aakashh/student-api/internal/types"
	"github.com/singhh-aakashh/student-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a student...")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err:= validator.New().Struct(student); err != nil{
			response.WriteJson(w,http.StatusBadRequest,err.Error())
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
