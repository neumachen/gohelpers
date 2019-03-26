package gblcodes

const (
	// AppHealthHealthy ...
	AppHealthHealthy appCode = "app_health.healthy"
	// AppHealthFailedDep ...
	AppHealthFailedDep appCode = "app_health.failed_dep"
)

var (
	appHealthMessages = map[AppCode]string{
		AppHealthHealthy:   "Application is healthy",
		AppHealthFailedDep: "Application dependency failure. See additional error message",
	}
)
