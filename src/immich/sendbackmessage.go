package immich

import (
	"immich-exporter/src/models"
	"strconv"
)

func Sendbackmessagepreference(result *models.Users, result2 *models.AllUsers) string {
	total := "# HELP immich_app_number_users The number of immich users\n# TYPE immich_app_max_active_downloads gauge\nimmich_app_nb_users " + strconv.Itoa(len((*result).UsageByUser)) + "\n"
	total = total + "# HELP immich_app_total_photos The total number of photos\n# TYPE immich_app_total_photos gauge\nimmich_app_total_photos " + strconv.Itoa((*result).Photos) + "\n"
	total = total + "# HELP immich_app_total_videos The total number of videos\n# TYPE immich_app_total_videos gauge\nimmich_app_total_videos " + strconv.Itoa((*result).Videos) + "\n"
	total = total + "# HELP immich_app_total_usage The usage of the user\n# TYPE immich_app_total_usage gauge\nimmich_app_total_usage " + strconv.Itoa(int((*result).UsageRaw)) + "\n"
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
	return total + immich_user_info + immich_user_videos + immich_user_photos + immich_user_usageRaw
}

func Sendbackmessageserverversion(result *models.ServerVersion) string {

	return "# HELP immich_app_version The current immich version\n# TYPE immich_app_version gauge\nimmich_app_version" + `{version="` + strconv.Itoa((*result).Major) + "." + strconv.Itoa((*result).Minor) + "." + strconv.Itoa((*result).Patch) + `",} 1.0` + "\n"
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
