package venicebeach

type UserInfo struct {
	User struct {
		Id                            int         `json:"id"`
		Imageurl                      interface{} `json:"imageurl"`
		Firstname                     string      `json:"firstname"`
		Lastname                      string      `json:"lastname"`
		Level                         int         `json:"level"`
		LevelName                     string      `json:"levelName"`
		LevelImage                    string      `json:"levelImage"`
		FirebaseUid                   string      `json:"firebaseUid"`
		FollowerCount                 int         `json:"follower_count"`
		FollowingCount                int         `json:"following_count"`
		Email                         string      `json:"email"`
		Sex                           int         `json:"sex"`
		Gender                        string      `json:"gender"`
		Balance                       interface{} `json:"balance"`
		Birthday                      int         `json:"birthday"`
		AgeVisible                    bool        `json:"age_visible"`
		FacebookConnected             bool        `json:"facebook_connected"`
		LastnameObfuscated            bool        `json:"lastname_obfuscated"`
		NotifyCourseChangeViaEmail    bool        `json:"notify_course_change_via_email"`
		NotifyBookingOverviewViaEmail bool        `json:"notify_booking_overview_via_email"`
		NotificationBadgeCount        int         `json:"notificationBadgeCount"`
		BadgeCount                    int         `json:"badgeCount"`
		TermsAccepted                 bool        `json:"termsAccepted"`
		Language                      string      `json:"language"`
		ShowRatingAt                  interface{} `json:"showRatingAt"`
		RetryShowRatingCount          int         `json:"retryShowRatingCount"`
		RatedAt                       interface{} `json:"ratedAt"`
		OccurenceJoinDeniedTill       interface{} `json:"occurenceJoinDeniedTill"`
		CountryCode                   string      `json:"countryCode"`
		ShowScheduledPosts            bool        `json:"showScheduledPosts"`
		Links                         []struct {
			Rel    string `json:"rel"`
			Uri    string `json:"uri"`
			Method string `json:"method"`
			Params struct {
				Message    string `json:"message,omitempty"`
				Image      string `json:"image,omitempty"`
				Video      string `json:"video,omitempty"`
				Apptours   string `json:"apptours,omitempty"`
				Searchterm string `json:"searchterm,omitempty"`
			} `json:"params,omitempty"`
		} `json:"links"`
		Address struct {
			Street       string `json:"street"`
			StreetNumber string `json:"streetNumber"`
			Postcode     string `json:"postcode"`
			City         string `json:"city"`
			Links        []struct {
				Rel    string `json:"rel"`
				Uri    string `json:"uri"`
				Method string `json:"method"`
			} `json:"links"`
		} `json:"address"`
		Contact struct {
			Mobile        interface{} `json:"mobile"`
			Phone         string      `json:"phone"`
			BusinessPhone interface{} `json:"business_phone"`
			Links         []struct {
				Rel    string `json:"rel"`
				Uri    string `json:"uri"`
				Method string `json:"method"`
			} `json:"links"`
		} `json:"contact"`
		AccountInfo struct {
			AccountHolder string `json:"accountHolder"`
			BankName      string `json:"bankName"`
			Iban          string `json:"iban"`
			Bic           string `json:"bic"`
			Links         []struct {
				Rel    string `json:"rel"`
				Uri    string `json:"uri"`
				Method string `json:"method"`
			} `json:"links"`
		} `json:"accountInfo"`
		TenantUser struct {
			Level              string        `json:"level"`
			HomeStudioId       int           `json:"homeStudioId"`
			CustomerId         string        `json:"customerId"`
			ContractStart      int           `json:"contractStart"`
			ContractEnd        int           `json:"contractEnd"`
			CreatedAt          int           `json:"createdAt"`
			Public             bool          `json:"public"`
			TutorialFinished   bool          `json:"tutorialFinished"`
			NotifyActinatePost bool          `json:"notify_actinate_post"`
			NotifyTenantPost   bool          `json:"notify_tenant_post"`
			Links              []interface{} `json:"links"`
			Studio             Studio        `json:"studio"`
		} `json:"tenantUser"`
		DeactivatedFeatures []struct {
			Feature string `json:"feature"`
		} `json:"deactivatedFeatures"`
		UserApptourFeatures []struct {
			Key  string `json:"key"`
			Seen int    `json:"seen"`
		} `json:"userApptourFeatures"`
		AccessTypeSettings struct {
			TimelineInteraction     bool `json:"timelineInteraction"`
			Mediathek               bool `json:"mediathek"`
			AttendStudioClasses     bool `json:"attendStudioClasses"`
			AttendLivestreamClasses bool `json:"attendLivestreamClasses"`
			BookTimeslots           bool `json:"bookTimeslots"`
			EnableStrike            bool `json:"enableStrike"`
			ReferFriend             bool `json:"referFriend"`
			AutoCheckinCourses      bool `json:"autoCheckinCourses"`
			UsersConnections        bool `json:"usersConnections"`
		} `json:"accessTypeSettings"`
		AccessTypes []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"accessTypes"`
	} `json:"user"`
	Studios struct {
		Data []Studio `json:"data"`
	} `json:"studios"`
	TimelineCategories struct {
		Data []struct {
			Id   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"data"`
	} `json:"timelineCategories"`
	EnableUsersConnections bool `json:"enableUsersConnections"`
}
