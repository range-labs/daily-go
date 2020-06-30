package daily

import "time"

// DomainConfig is used when getting and setting the domain configuration.
// https://docs.daily.co/reference#get-domain-configuration
type DomainConfig struct {
	DomainName *string `json:"domain_name,omitempty"`
	Config     *Config `json:"config,omitempty"`
}

// Config options contained within the DomainConfig that can be changed by the
// user.
type Config struct {
	RedirectOnMeetingExit *string `json:"redirect_on_meeting_exit,omitempty"`
	HideDailyBranding     *bool   `json:"hide_daily_branding,omitempty"`
	HIPPAA                *bool   `json:"hipaa,omitempty"`
	IntercomAutoRecord    *bool   `json:"intercom_auto_record,omitempty"`
	Lang                  *string `json:"lang,omitempty"`
}

// Room contains information about a video location and configuration.
// https://docs.daily.co/reference#rooms
type Room struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	APICreated bool        `json:"api_created"`
	Privacy    RoomPrivacy `json:"privacy"`
	URL        string      `json:"url"`
	CreatedAt  time.Time   `json:"created_at"`
	Config     *RoomConfig `json:"config"`
}

// RoomPrivacy controls who can join a meeting.
type RoomPrivacy string

const (
	Public  RoomPrivacy = "public"
	Private RoomPrivacy = "private"
	Org     RoomPrivacy = "org"
)

// RoomConfig is the configuration for a room.
type RoomConfig struct {
	NotBefore          *int64  `json:"nbf,omitempty"` // Unix timestamp in seconds
	ExpiresAt          *int64  `json:"exp,omitempty"` // Unix timestamp in seconds
	StartVideoOff      *bool   `json:"start_video_off,omitempty"`
	StartAudioOff      *bool   `json:"start_audio_off,omitempty"`
	MaxParticipants    *int32  `json:"max_participants,omitempty"`
	AutoJoin           *bool   `json:"auto_join,omitempty"`
	EnableKnocking     *bool   `json:"enable_knocking,omitempty"`
	EnableScreenShare  *bool   `json:"enable_screenshare,omitempty"`
	EnableChat         *bool   `json:"enable_chat,omitempty"`
	OwnerOnlyBroadcast *bool   `json:"owner_only_broadcast,omitempty"`
	EnableRecording    *bool   `json:"enable_recording,omitempty"`
	EjectAtRoomExpiry  *bool   `json:"eject_at_room_exp,omitempty"`
	EjectAfterElapsed  *bool   `json:"eject_after_elapsed,omitempty"`
	Lang               *string `json:"lang,omitempty"`
}

// String returns a pointer to the string.
func String(s string) *string {
	return &s
}

// Int64 returns a pointer to the int64.
func Int64(i int64) *int64 {
	return &i
}

// Int32 returns a pointer to the int32.
func Int32(i int32) *int32 {
	return &i
}

// Timestamp returns number of seconds since epoch, consistent wih Daily's
// API expectations.
func Timestamp(t time.Time) *int64 {
	return Int64(t.Unix())
}

// Bool returns a pointer to the bool.
func Bool(b bool) *bool {
	return &b
}

// True returns a pointer to true.
func True() *bool {
	return Bool(true)
}

// False returns a pointer to false.
func False() *bool {
	return Bool(false)
}
