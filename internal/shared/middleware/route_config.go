package middleware

type RouteConfig struct {
	Path        string
	Method      string
	Roles       []string
	Description string
}

var RouteConfigs = []RouteConfig{
	// Auth routes (no authentication required)
	{Path: "/api/auth/login", Method: "POST", Roles: []string{}, Description: "User login"},
	{Path: "/api/auth/register", Method: "POST", Roles: []string{}, Description: "User registration"},
	{Path: "/api/auth/refresh", Method: "POST", Roles: []string{}, Description: "Token refresh"},
	{Path: "/api/auth/logout", Method: "POST", Roles: []string{}, Description: "User logout"},

	// User management (admin only)
	{Path: "/api/admin/users", Method: "POST", Roles: []string{"admin"}, Description: "Create user"},
	{Path: "/api/admin/users/:id", Method: "GET", Roles: []string{"admin"}, Description: "Get user by ID"},
	{Path: "/api/admin/users", Method: "GET", Roles: []string{"admin"}, Description: "Get all users"},

	// Course management (admin and teachers)
	{Path: "/api/courses", Method: "POST", Roles: []string{"admin", "teacher"}, Description: "Create course"},
	{Path: "/api/courses", Method: "GET", Roles: []string{"admin", "teacher"}, Description: "Get all courses"},
	{Path: "/api/courses/:id", Method: "GET", Roles: []string{"admin", "teacher"}, Description: "Get course by ID"},
	{Path: "/api/courses/:id", Method: "PUT", Roles: []string{"admin", "teacher"}, Description: "Update course"},
	{Path: "/api/courses/:id", Method: "DELETE", Roles: []string{"admin", "teacher"}, Description: "Delete course"},

	// Classroom management (admin and teachers)
	{Path: "/api/classrooms", Method: "POST", Roles: []string{"admin", "teacher"}, Description: "Create classroom"},
	{Path: "/api/classrooms", Method: "GET", Roles: []string{"admin", "teacher"}, Description: "Get all classrooms"},
	{Path: "/api/classrooms/:id", Method: "GET", Roles: []string{"admin", "teacher"}, Description: "Get classroom by ID"},
	{Path: "/api/classrooms/location/:location", Method: "GET", Roles: []string{"admin", "teacher"}, Description: "Get classrooms by location"},
	{Path: "/api/classrooms/:id", Method: "PUT", Roles: []string{"admin", "teacher"}, Description: "Update classroom"},
	{Path: "/api/classrooms/:id", Method: "DELETE", Roles: []string{"admin", "teacher"}, Description: "Delete classroom"},

	// QR Code management (teachers only)
	{Path: "/api/qrcodes", Method: "POST", Roles: []string{"teacher"}, Description: "Create QR code"},
	{Path: "/api/qrcodes/:id", Method: "GET", Roles: []string{"teacher"}, Description: "Get QR code by ID"},
	{Path: "/api/qrcodes/code/:code", Method: "GET", Roles: []string{"teacher"}, Description: "Get QR code by code"},
	{Path: "/api/qrcodes/course/:courseId", Method: "GET", Roles: []string{"teacher"}, Description: "Get course QR codes"},
	{Path: "/api/qrcodes/course/:courseId/active", Method: "GET", Roles: []string{"teacher"}, Description: "Get active QR code"},
	{Path: "/api/qrcodes/:id", Method: "PUT", Roles: []string{"teacher"}, Description: "Update QR code"},
	{Path: "/api/qrcodes/:id", Method: "DELETE", Roles: []string{"teacher"}, Description: "Delete QR code"},

	// Attendance management (teachers and students)
	{Path: "/api/attendance", Method: "POST", Roles: []string{"teacher", "student"}, Description: "Create attendance"},
	{Path: "/api/attendance/:id", Method: "GET", Roles: []string{"teacher", "student"}, Description: "Get attendance by ID"},
	{Path: "/api/attendance/student/:studentId", Method: "GET", Roles: []string{"teacher", "student"}, Description: "Get student attendance"},
	{Path: "/api/attendance/course/:courseId", Method: "GET", Roles: []string{"teacher", "student"}, Description: "Get course attendance"},
	{Path: "/api/attendance/:id", Method: "PUT", Roles: []string{"teacher", "student"}, Description: "Update attendance"},
	{Path: "/api/attendance/:id", Method: "DELETE", Roles: []string{"teacher", "student"}, Description: "Delete attendance"},
}

func GetRouteConfig(path, method string) *RouteConfig {
	for _, config := range RouteConfigs {
		if config.Path == path && config.Method == method {
			return &config
		}
	}
	return nil
}

func IsPublicRoute(path, method string) bool {
	config := GetRouteConfig(path, method)
	return config != nil && len(config.Roles) == 0
} 