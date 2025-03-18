package models

import "time"

type TrackDB struct {
	ID             uint   `gorm:"primaryKey"`
	Tracked        bool
	Youtube        bool
	Navidrome      string
	Title          string
	AlbumImagePath string
	AlbumImageLink string
	Duration       float64
	ArtistName     string
	ArtistID       uint `gorm:"index"`
	AlbumName      string
	AlbumID        uint `gorm:"index"`
	Path           string
	ArtistPath     string
	AlbumPath      string
}

type ArtistDB struct {
	ID          uint   `gorm:"primaryKey"`
	Tracked     bool
	Youtube     bool
	Name        string
	ImagePath   string
	ImageLink   string
	SmallImgLink string
	NbAlbums    int
	Path        string
}

type AlbumDB struct {
	ID          uint   `gorm:"primaryKey"`
	Tracked     bool
	Youtube     bool
	Name        string
	ImgPath     string
	ImgLink     string
	ImgSmallLink string
	NbTracks    int
	ArtistID    uint `gorm:"index"`
	ArtistPath  string
	Path        string
}


type QueueItem struct {
	ID        string                 `json:"id"`
	TrackID   uint                   `json:"track_id"`
	TrackName string                 `json:"track_name"`
	CreatedAt time.Time              `json:"created_at"`
}