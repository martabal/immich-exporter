package prom

import (
	"immich-exp/models"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type Gauge []struct {
	name  string
	help  string
	value float64
}

func SendBackMessagePreference(
	result *models.StructServerInfo,
	result2 *models.StructAllUsers,
	result3 *models.StructAllJobsStatus,
	r *prometheus.Registry,
) {

	gauges := Gauge{
		{"total photos", "The total number of photos", float64((*result).Photos)},
		{"total videos", "The total number of videos", float64((*result).Videos)},
		{"total usage", "The max number of active torrents allowed", float64((*result).Usage)},
		{"number users", "The total number of users", float64(len((*result).UsageByUser))},
	}

	register(gauges, r)

	user_info := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_info",
		Help: "All infos about users",
	}, []string{"videos", "photos", "uid", "usage", "name"})

	user_usage := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_usage",
		Help: "The usage of the user",
	}, []string{"uid", "name"})
	user_photos := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_photos",
		Help: "The number of photo of the user",
	}, []string{"uid", "name"})
	user_videos := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_user_videos",
		Help: "The number of videos of the user",
	}, []string{"uid", "name"})

	job_count := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "immich_job_count",
		Help: "The item count in the job",
	}, []string{"status", "job_name"})

	r.MustRegister(user_info)
	r.MustRegister(user_usage)
	r.MustRegister(user_videos)
	r.MustRegister(user_photos)
	r.MustRegister(job_count)

	for i := 0; i < len((*result).UsageByUser); i++ {
		var myuser = GetName((*result).UsageByUser[i].UserID, result2)
		user_info.With(prometheus.Labels{"videos": strconv.Itoa((*result).UsageByUser[i].Videos), "photos": strconv.Itoa((*result).UsageByUser[i].Photos), "uid": (*result).UsageByUser[i].UserID, "usage": strconv.Itoa(int((*result).UsageByUser[i].Usage)), "name": myuser.Name}).Inc()
		user_photos.With(prometheus.Labels{"uid": (*result).UsageByUser[i].UserID, "name": myuser.Name}).Set(float64((*result).UsageByUser[i].Photos))
		user_usage.With(prometheus.Labels{"uid": (*result).UsageByUser[i].UserID, "name": myuser.Name}).Set(float64((*result).UsageByUser[i].Usage))
		user_videos.With(prometheus.Labels{"uid": (*result).UsageByUser[i].UserID, "name": myuser.Name}).Set(float64((*result).UsageByUser[i].Videos))
	}

	setJobStatusCounts(job_count, "background_task", &result3.BackgroundTask)
	setJobStatusCounts(job_count, "clip_encoding", &result3.ClipEncoding)
	setJobStatusCounts(job_count, "library", &result3.Library)
	setJobStatusCounts(job_count, "metadata_extraction", &result3.MetadataExtraction)
	setJobStatusCounts(job_count, "migration", &result3.Migration)
	setJobStatusCounts(job_count, "object_tagging", &result3.ObjectTagging)
	setJobStatusCounts(job_count, "recognize_faces", &result3.RecognizeFaces)
	setJobStatusCounts(job_count, "search", &result3.Search)
	setJobStatusCounts(job_count, "sidecar", &result3.Sidecar)
	setJobStatusCounts(job_count, "storage_template_migration", &result3.StorageTemplateMigration)
	setJobStatusCounts(job_count, "thumbnail_generation", &result3.ThumbnailGeneration)
	setJobStatusCounts(job_count, "video_conversion", &result3.VideoConversion)
}

func setJobStatusCounts(job_count *prometheus.GaugeVec, jobName string, result *models.StructJobStatus) {
	job_count.With(prometheus.Labels{"status": "active", "job_name": jobName}).Set(float64(result.JobCounts.Active))
	job_count.With(prometheus.Labels{"status": "completed", "job_name": jobName}).Set(float64(result.JobCounts.Completed))
	job_count.With(prometheus.Labels{"status": "failed", "job_name": jobName}).Set(float64(result.JobCounts.Failed))
	job_count.With(prometheus.Labels{"status": "delayed", "job_name": jobName}).Set(float64(result.JobCounts.Delayed))
	job_count.With(prometheus.Labels{"status": "waiting", "job_name": jobName}).Set(float64(result.JobCounts.Waiting))
	job_count.With(prometheus.Labels{"status": "paused", "job_name": jobName}).Set(float64(result.JobCounts.Paused))
}

func SendBackMessageserverVersion(result *models.StructServerVersion, r *prometheus.Registry) {

	version := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "immich_app_version",
		Help: "Immich version",
		ConstLabels: map[string]string{
			"version": strconv.Itoa((*result).Major) + "." + strconv.Itoa((*result).Minor) + "." + strconv.Itoa((*result).Patch),
		},
	})
	version.Set(1)
	r.MustRegister(version)

}

func GetName(result string, result2 *models.StructAllUsers) models.StructCustomUser {
	var myuser models.StructCustomUser
	for i := 0; i < len(*result2); i++ {
		if (*result2)[i].ID == result {

			myuser.ID = (*result2)[i].ID
			myuser.Name = (*result2)[i].Name
			myuser.Email = (*result2)[i].Email
			myuser.IsAdmin = (*result2)[i].IsAdmin
		}

	}
	return myuser
}

func register(gauges Gauge, r *prometheus.Registry) {
	for _, gauge := range gauges {
		name := "immich_app_" + strings.Replace(gauge.name, " ", "_", -1)
		help := gauge.help
		g := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: name,
			Help: help,
		})
		r.MustRegister(g)
		g.Set(gauge.value)
	}
}
