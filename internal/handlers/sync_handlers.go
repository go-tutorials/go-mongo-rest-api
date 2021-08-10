package handlers

import (
	"encoding/json"
	. "go-service/internal/services/sync_service"
	"net/http"
)

type SyncHandler struct {
	sync SyncService
}

type ChannelId struct {
	ChannelId string `json:"channelId,omitempty"`
	Level int `json:"level,omitempty"`
}

type PlaylistId struct {
	PlaylistId string `json:"playlistId,omitempty"`
	Level int `json:"level,omitempty"`
}

func NewSyncHandler(syncService SyncService) *SyncHandler {
	return &SyncHandler{sync: syncService}
}

func (h *SyncHandler) SyncChannel(w http.ResponseWriter, r *http.Request) {
	var channelId ChannelId
	er1 := json.NewDecoder(r.Body).Decode(&channelId)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	resultChannel, er2 := h.sync.SyncChannel(r.Context(), channelId.ChannelId)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusBadRequest)
		return
	}
	result := ""
	if resultChannel > 0 {
		result = "Sync channel successfully"
	} else {
		result = "Invalid channel to sync"
	}
	respond(w, result)
}

func (h *SyncHandler) SyncPlaylist(w http.ResponseWriter, r *http.Request) {
	var playlistId PlaylistId
	er1 := json.NewDecoder(r.Body).Decode(&playlistId)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	resultChannel, er2 := h.sync.SyncPlaylist(r.Context(), playlistId.PlaylistId, &playlistId.Level)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusBadRequest)
		return
	}
	result := ""
	if resultChannel > 0 {
		result = "Sync channel successfully"
	} else {
		result = "Invalid channel to sync"
	}
	respond(w, result)
}
