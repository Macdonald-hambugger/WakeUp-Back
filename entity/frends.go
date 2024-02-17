package entity

type Frends struct {
	Id           int64  `json:"id"`
	Nickname     string `json:"profile_nickname"`
	ProfileImage string `json:"profile_thumbnail_image"`
}

type FrendsProfileDTO struct {
	ID                    int64  `json:"id"`
	UUID                  string `json:"uuid"`
	Favorite              bool   `json:"favorite"`
	ProfileNickname       string `json:"profile_nickname"`
	ProfileThumbnailImage string `json:"profile_thumbnail_image"`
}
