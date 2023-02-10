package models

import "time"

type Login struct {
	AccessToken          string `json:"accessToken"`
	UserID               string `json:"userId"`
	UserEmail            string `json:"userEmail"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	IsAdmin              bool   `json:"isAdmin"`
	ShouldChangePassword bool   `json:"shouldChangePassword"`
}

type Users struct {
	Photos      int `json:"photos"`
	Videos      int `json:"videos"`
	UsageByUser []struct {
		UserID   string `json:"userId"`
		Videos   int    `json:"videos"`
		Photos   int    `json:"photos"`
		UsageRaw int64  `json:"usageRaw"`
		Usage    string `json:"usage"`
	} `json:"usageByUser"`
	UsageRaw int64  `json:"usageRaw"`
	Usage    string `json:"usage"`
}

type ServerInfo struct {
	DiskAvailable       string  `json:"diskAvailable"`
	DiskSize            string  `json:"diskSize"`
	DiskUse             string  `json:"diskUse"`
	DiskAvailableRaw    int64   `json:"diskAvailableRaw"`
	DiskSizeRaw         int64   `json:"diskSizeRaw"`
	DiskUseRaw          int64   `json:"diskUseRaw"`
	DiskUsagePercentage float64 `json:"diskUsagePercentage"`
}

type ServerVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type AllUsers []struct {
	ID                   string    `json:"id"`
	Email                string    `json:"email"`
	FirstName            string    `json:"firstName"`
	LastName             string    `json:"lastName"`
	CreatedAt            time.Time `json:"createdAt"`
	ProfileImagePath     string    `json:"profileImagePath"`
	ShouldChangePassword bool      `json:"shouldChangePassword"`
	IsAdmin              bool      `json:"isAdmin"`
	DeletedAt            time.Time `json:"deletedAt"`
	OauthID              string    `json:"oauthId"`
}

type CustomUser struct {
	Email     string
	ID        string
	FirstName string
	LastName  string
	IsAdmin   bool
}
