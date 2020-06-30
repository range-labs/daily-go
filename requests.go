package daily

// ListRoomsRequest contains the parameters for listing rooms.
// https://docs.daily.co/reference#list-rooms
type ListRoomsRequest struct {
	Limit        int32  `json:"limit,omitempty"`
	EndingBefore string `json:"ending_before,omitempty"`
	EndingAfter  string `json:"ending_after,omitempty"`
}

// ListRoomsResponse is the response envelope when listing rooms.
// https://docs.daily.co/reference#list-rooms
type ListRoomsResponse struct {
	TotalCount int32  `json:"total_count"`
	Rooms      []Room `json:"data"`
}

// CreateRoomRequest contains the parameters for creating a room.
// https://docs.daily.co/reference#create-room
type CreateRoomRequest struct {
	Name    *string     `json:"name,omitempty"`
	Privacy RoomPrivacy `json:"privacy,omitempty"`
	Config  *RoomConfig `json:"properties,omitempty"`
}

// CreateRoomResponse contains the newly created room.
type CreateRoomResponse struct {
	Room
}

// CreateRoomResponse contains the requested room.
type GetRoomResponse struct {
	Room
}

// UpdateRoomRequest contains the parameters for updating a room.
type UpdateRoomRequest struct {
	Privacy RoomPrivacy `json:"privacy,omitempty"`
	Config  *RoomConfig `json:"properties,omitempty"`
}

// UpdateRoomResponse contains the updated room.
type UpdateRoomResponse struct {
	Room
}
