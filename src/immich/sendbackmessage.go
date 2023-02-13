package immich

import (
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
	user_info := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_info",
		Help: "All infos about users",
	}, []string{"videos", "photos", "uid", "usage", "firstname", "lastname"})

	user_usage := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_usage",
		Help: "The usage of the user",
	}, []string{"uid", "firstname", "lastname"})
	user_photos := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_photos",
		Help: "The number of photo of the user",
	}, []string{"uid", "firstname", "lastname"})
	user_videos := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_videos",
		Help: "The number of videos of the user",
	}, []string{"uid", "firstname", "lastname"})

	r.MustRegister(user_info)

	r.MustRegister(total_usage)
	r.MustRegister(total_videos)
	r.MustRegister(total_photos)
	r.MustRegister(total_users)
	r.MustRegister(user_usage)
	r.MustRegister(user_videos)
	r.MustRegister(user_photos)
	total_photos.Add(float64((*result).Photos))
	total_videos.Add(float64((*result).Videos))
	total_usage.Add(float64((*result).UsageRaw))
	total_users.Add(float64(len((*result).UsageByUser)))

	for i := 0; i < len((*result).UsageByUser); i++ {
		var myuser = GetName((*result).UsageByUser[i].UserID, result2)
		user_info.With(prometheus.Labels{"videos": strconv.Itoa((*result).UsageByUser[i].Videos), "photos": strconv.Itoa((*result).UsageByUser[i].Photos), "uid": (*result).UsageByUser[i].UserID, "usage": strconv.Itoa(int((*result).UsageByUser[i].UsageRaw)), "firstname": myuser.FirstName, "lastname": myuser.LastName}).Inc()
		user_photos.With(prometheus.Labels{"uid": (*result).UsageByUser[i].UserID, "firstname": myuser.FirstName, "lastname": myuser.LastName}).Set(float64((*result).UsageByUser[i].Photos))
		user_usage.With(prometheus.Labels{"uid": (*result).UsageByUser[i].UserID, "firstname": myuser.FirstName, "lastname": myuser.LastName}).Set(float64((*result).UsageByUser[i].UsageRaw))
		user_videos.With(prometheus.Labels{"uid": (*result).UsageByUser[i].UserID, "firstname": myuser.FirstName, "lastname": myuser.LastName}).Set(float64((*result).UsageByUser[i].Videos))
	}

}

func Sendbackmessageserverversion(result *models.ServerVersion, r *prometheus.Registry) {

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
