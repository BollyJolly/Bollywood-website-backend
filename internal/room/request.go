package room

type CreateRoomRequest struct {
	PlaylistID string `json:"playlistId" binding:"required"`
}