package venicebeach

type Timeline struct {
	Timeline struct {
		Data []struct {
			Id                     int         `json:"id"`
			Title                  string      `json:"title"`
			Description            string      `json:"description"`
			DescriptionRich        string      `json:"descriptionRich"`
			Cta                    *string     `json:"cta"`
			Link                   *string     `json:"link"`
			PreviewImageUrl        *string     `json:"previewImageUrl"`
			ImageUrl               *string     `json:"imageUrl"`
			VideoUrl               interface{} `json:"videoUrl"`
			VideoPreviewUrl        interface{} `json:"videoPreviewUrl"`
			VideoRatio             interface{} `json:"videoRatio"`
			Subscribed             bool        `json:"subscribed,omitempty"`
			CommentsCount          int         `json:"comments_count"`
			HighfivesCount         int         `json:"highfives_count"`
			Mine                   bool        `json:"mine,omitempty"`
			Highfived              bool        `json:"highfived"`
			CreatedAt              int         `json:"createdAt,omitempty"`
			ScheduledAt            int         `json:"scheduledAt,omitempty"`
			Tenant                 bool        `json:"tenant,omitempty"`
			Global                 bool        `json:"global,omitempty"`
			DisableComments        bool        `json:"disableComments"`
			DisableShare           bool        `json:"disableShare"`
			Public                 bool        `json:"public,omitempty"`
			DisableCommentsMessage interface{} `json:"disableCommentsMessage"`
			Pinned                 bool        `json:"pinned,omitempty"`
			SharedPostDeleted      bool        `json:"sharedPostDeleted,omitempty"`
			Links                  []struct {
				Rel    string `json:"rel"`
				Uri    string `json:"uri"`
				Method string `json:"method"`
			} `json:"links"`
			Type   string `json:"type,omitempty"`
			Studio Studio `json:"studio,omitempty"`
		} `json:"data"`
		Meta struct {
			Cursor struct {
				Current interface{} `json:"current"`
				Prev    interface{} `json:"prev"`
				Next    string      `json:"next"`
				Count   int         `json:"count"`
			} `json:"cursor"`
		} `json:"meta"`
	} `json:"timeline"`
	Classes struct {
		Data []interface{} `json:"data"`
	} `json:"classes"`
	Videos struct {
		Data []struct {
			Id              int    `json:"id"`
			Title           string `json:"title"`
			Description     string `json:"description"`
			VideoUrl        string `json:"videoUrl"`
			New             bool   `json:"new"`
			PreviewImage    string `json:"previewImage"`
			DurationSeconds int    `json:"durationSeconds"`
			Links           []struct {
				Rel    string `json:"rel"`
				Uri    string `json:"uri"`
				Method string `json:"method"`
			} `json:"links"`
			Trainers []struct {
				Id    int    `json:"id"`
				Name  string `json:"name"`
				Links []struct {
					Rel    string `json:"rel"`
					Uri    string `json:"uri"`
					Method string `json:"method"`
				} `json:"links"`
			} `json:"trainers"`
		} `json:"data"`
	} `json:"videos"`
	VideoSettings struct {
		Name  interface{}   `json:"name"`
		Image string        `json:"image"`
		Links []interface{} `json:"links"`
	} `json:"videoSettings"`
}
