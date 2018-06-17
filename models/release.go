package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/jinzhu/gorm"
)

type Release struct {
	gorm.Model
	Raw         string    `json:"raw"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Season      int       `json:"season"`
	Episode     int       `json:"episode"`
	Size        int       `json:"size"`
	Guid        string    `json:"guid"`
	Resolution  int       `json:"resolution"`
	Team        string    `json:"team"`
	Author      string    `json:"author"`
	Verified    bool      `json:"verified"`
	Bluray      bool      `json:"bluray"`
	Uncensored  bool      `json:"uncensored"`
	Checksum    string    `json:"checksum" gorm:"unique_index"`
	Download    string    `json:"download"`
	Source      string    `json:"source"`
	Type        string    `json:"type"`
	Published   time.Time `json:"published"`
}

//func (r *Release) BeforeInsert() error {
//	fmt.Printf("ID: %#v\n", r.ID.String())
//	fmt.Printf("Checksum: %#v\n", r.Checksum)
//	return nil
//}
//
func (r *Release) BeforeSave() error {
	if r.Checksum == "" {
		r.CalculateChecksum()
	}
	return nil
}

func (r *Release) CalculateChecksum() {
	h := md5.New()
	h.Write([]byte(r.Download))
	r.Checksum = hex.EncodeToString(h.Sum(nil))
}
