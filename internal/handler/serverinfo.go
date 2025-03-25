package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"paddex.net/cs2-rcon-go/gameinfo"
)

func (h *Handler) serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	reqBody := []struct {
		Name     string
		Addr     string
		Port     int
		Password string
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
		slog.String("Handler", "serverInfoHandler"),
		slog.String("Data", fmt.Sprintf("%+v", reqBody)),
	)

	type serverinfo struct {
		Name    string
		Addr    string
		Port    int
		Map     string
		Players string
	}
	response := []serverinfo{}

	for _, s := range reqBody {
		client, err := gameinfo.NewClient(s.Addr, s.Port)
		if err != nil {
			h.ErrorLogger.Error(
				err.Error(),
			)
		}
		info, err := client.GetServerInfo()
		if err != nil {
			h.ErrorLogger.Error(
				err.Error(),
			)
		}
		h.ErrorLogger.Debug(
			"Gameinfo",
			slog.String("Data", fmt.Sprintf("%+v", info)),
		)
		i := serverinfo{
			Name:    info.Name,
			Addr:    s.Addr,
			Port:    int(info.Port),
			Map:     info.MapName,
			Players: fmt.Sprintf("%d/%d", info.Players, info.MaxPlayers),
		}

		response = append(response, i)
	}

	res, err := json.Marshal(response)
	if err != nil {
		h.ErrorLogger.Error(err.Error())
	}

	w.Write(res)
}
