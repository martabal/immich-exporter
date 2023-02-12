package immich

import (
	"fmt"
	"immich-exporter/src/models"

	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

func Sendbackmessagepreference(result *models.Users, result2 *models.AllUsers, r *prometheus.Registry) {
	total_photos := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "immich_app_total_photos",
		Help: "The total number of photos",
	})
	total_videos := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "immich_app_total_videos",
		Help: "The total number of videos",
	})
	total_usage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "immich_app_total_usage",
		Help: "The total usage of disk",
	})
	total_users := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "immich_app_number_users",
		Help: "The total number of users",
	})

	r.MustRegister(total_usage)
	r.MustRegister(total_videos)
	r.MustRegister(total_photos)
	r.MustRegister(total_users)
	total_photos.Add(float64((*result).Photos))
	total_videos.Add(float64((*result).Videos))
	total_usage.Add(float64((*result).UsageRaw))
	total_users.Add(float64(len((*result).UsageByUser)))

	immich_user_videos := "# HELP immich_app_user_videos The number of videos of the user\n# TYPE immich_app_user_videos gauge\n"
	immich_user_photos := "# HELP immich_app_user_photos The number of photo of the user\n# TYPE immich_app_user_photos gauge\n"
	immich_user_usageRaw := "# HELP immich_app_user_usage The usage of the user\n# TYPE immich_app_user_usage gauge\n"
	immich_user_info := "# HELP immich_user_info All info for torrents\n# TYPE immich_user_info gauge\n"
	for i := 0; i < len((*result).UsageByUser); i++ {
		var myuser = GetName((*result).UsageByUser[i].UserID, result2)
		immich_user_info = immich_user_info + `immich_user_info{videos="` + strconv.Itoa((*result).UsageByUser[i].Videos) + `",photos="` + strconv.Itoa((*result).UsageByUser[i].Photos) + `",uid="` + (*result).UsageByUser[i].UserID + `",usage="` + strconv.Itoa(int((*result).UsageByUser[i].UsageRaw)) + `",firstname="` + myuser.FirstName + `",lastname="` + myuser.LastName + `"} 1.0` + "\n"
		immich_user_usageRaw = immich_user_usageRaw + `immich_user_usage{uid="` + (*result).UsageByUser[i].UserID + `",firstname="` + myuser.FirstName + `",lastname="` + myuser.LastName + `",} ` + strconv.Itoa(int((*result).UsageByUser[i].UsageRaw)) + "\n"
		immich_user_photos = immich_user_photos + `immich_user_photos{uid="` + (*result).UsageByUser[i].UserID + `",firstname="` + myuser.FirstName + `",lastname="` + myuser.LastName + `",} ` + strconv.Itoa((*result).UsageByUser[i].Photos) + "\n"
		immich_user_videos = immich_user_videos + `immich_user_videos{uid="` + (*result).UsageByUser[i].UserID + `",firstname="` + myuser.FirstName + `",lastname="` + myuser.LastName + `",} ` + strconv.Itoa((*result).UsageByUser[i].Videos) + "\n"
	}

}

func Sendbackmessageserverversion(result *models.ServerVersion, r *prometheus.Registry) {
	fmt.Println("test")

	version := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Immich version",
		ConstLabels: map[string]string{
			"version": strconv.Itoa((*result).Major) + "." + strconv.Itoa((*result).Minor) + "." + strconv.Itoa((*result).Patch),
		},
	})
	version.Set(1)
	r.MustRegister(version)

}

func GetName(result string, result2 *models.AllUsers) models.CustomUser {
	var myuser models.CustomUser
	for i := 0; i < len(*result2); i++ {
		if (*result2)[i].ID == result {

			myuser.ID = (*result2)[i].ID
			myuser.FirstName = (*result2)[i].FirstName
			myuser.LastName = (*result2)[i].LastName
			myuser.Email = (*result2)[i].Email
			myuser.IsAdmin = (*result2)[i].IsAdmin
		}

	}
	return myuser
}
