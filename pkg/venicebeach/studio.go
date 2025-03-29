package venicebeach

type Studio struct {
	Id                     int         `json:"id"`
	Name                   string      `json:"name"`
	Street                 string      `json:"street"`
	Postcode               string      `json:"postcode"`
	City                   string      `json:"city"`
	Phone                  string      `json:"phone"`
	EnableCourseJoin       bool        `json:"enableCourseJoin"`
	EnableCourses          bool        `json:"enableCourses"`
	EnableCheckinTimeslots bool        `json:"enableCheckinTimeslots"`
	Subscribed             bool        `json:"subscribed"`
	Favorited              bool        `json:"favorited"`
	HomeStudio             bool        `json:"homeStudio"`
	Email                  string      `json:"email"`
	VideoUrl               interface{} `json:"videoUrl"`
	ShopUrl                interface{} `json:"shopUrl"`
	ProfileIconUrl         interface{} `json:"profileIconUrl"`
	AppointmentUrl         interface{} `json:"appointmentUrl"`
	ExternalCoursesUrl     interface{} `json:"externalCoursesUrl"`
	PhoneBooking           interface{} `json:"phone_booking"`
	Order                  interface{} `json:"order"`
	CheckedInUsers         struct {
		Current int         `json:"current"`
		Max     interface{} `json:"max"`
		Min     interface{} `json:"min"`
	} `json:"checkedInUsers"`
	Links []struct {
		Rel    string `json:"rel"`
		Uri    string `json:"uri"`
		Method string `json:"method"`
	} `json:"links"`
	OpeningHours []struct {
		Day  string `json:"day"`
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"openingHours"`
	Images []struct {
		Id        int    `json:"id"`
		Thumbnail string `json:"thumbnail"`
		Image     string `json:"image"`
	} `json:"images"`
	Employees []interface{} `json:"employees"`
}
