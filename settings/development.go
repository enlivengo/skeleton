package settings

import (
	"github.com/enlivengo/enliven/config"
)

DevelopmentConfig := config.Config{
    "email_smtp_host":    "localhost",
    "email_smtp_auth":    "none",
    "email_from_default": "noreply@enliven.app",

    "database_driver":   "postgres",
    "database_host":     "127.0.0.1",
    "database_user":     "postgres",
    "database_dbname":   "enliven",
    "database_password": "postgres",

    "assets_static_route": "/assets/",
    "assets_static_path":  "./static/",

    "assets_statik_route": "/statik/",
}
