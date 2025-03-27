package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"paddex.net/cs2-rcon-go/rcon"
)

func (h *Handler) rconHandler(w http.ResponseWriter, r *http.Request) {
	reqBody := struct {
		Addr     string
		Port     int
		Password string
		Cmd      string
	}{}

	resultBody := struct {
		Status     int
		StatusText string
		Result     string
	}{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqBody)
	if err != nil {
		h.ErrorLogger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	h.ErrorLogger.Debug(
		"Data Received",
		slog.String("Handler", "rconHandler"),
		slog.String("Data", fmt.Sprintf("%+v", reqBody)),
	)

	client, err := rcon.NewClient(reqBody.Addr, reqBody.Port)
	if err != nil {
		h.ErrorLogger.Error(
			err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer client.Close()

	err = client.Auth(reqBody.Password)
	if err != nil {
		h.ErrorLogger.Error(
			err.Error(),
		)
		w.WriteHeader(http.StatusUnauthorized)
	}

	res, err := client.Exec(reqBody.Cmd)
	if err != nil {
		h.ErrorLogger.Error(
			err.Error(),
		)
	}
	resultBody.Status = http.StatusOK
	resultBody.StatusText = http.StatusText(http.StatusOK)
	resultBody.Result = res

	resJson, err := json.Marshal(resultBody)
	if err != nil {
		h.ErrorLogger.Error(
			err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resJson)
}
