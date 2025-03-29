package venicebeach

type Client interface {
	Ping() error
	UserInfo() (UserInfo, error)
	Timeline() (Timeline, error)
	GetStudioInfo(studioId int) (Studio, error)
}
