package models

import "time"

type StructLogin struct {
	AccessToken          string `json:"accessToken"`
	UserID               string `json:"userId"`
	UserEmail            string `json:"userEmail"`
	Name                 string `json:"name"`
	IsAdmin              bool   `json:"isAdmin"`
	ShouldChangePassword bool   `json:"shouldChangePassword"`
}

type StructServerInfo struct {
	Photos      int   `json:"photos"`
	Videos      int   `json:"videos"`
	Usage       int64 `json:"usage"`
	UsageByUser []struct {
		UserID   string `json:"userId"`
		UserName string `json:"userName"`
		Photos   int    `json:"photos"`
		Videos   int    `json:"videos"`
		Usage    int    `json:"usage"`
	} `json:"usageByUser"`
}

type StructDiskInfo struct {
	DiskAvailable       string  `json:"diskAvailable"`
	DiskSize            string  `json:"diskSize"`
	DiskUse             string  `json:"diskUse"`
	DiskAvailableRaw    int64   `json:"diskAvailableRaw"`
	DiskSizeRaw         int64   `json:"diskSizeRaw"`
	DiskUseRaw          int64   `json:"diskUseRaw"`
	DiskUsagePercentage float64 `json:"diskUsagePercentage"`
}

type StructServerVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type StructAllUsers []struct {
	ID                   string    `json:"id"`
	Email                string    `json:"email"`
	Name                 string    `json:"name"`
	CreatedAt            time.Time `json:"createdAt"`
	ProfileImagePath     string    `json:"profileImagePath"`
	ShouldChangePassword bool      `json:"shouldChangePassword"`
	IsAdmin              bool      `json:"isAdmin"`
	DeletedAt            time.Time `json:"deletedAt"`
	OauthID              string    `json:"oauthId"`
}

type StructCustomUser struct {
	Email   string
	ID      string
	Name    string
	IsAdmin bool
}
