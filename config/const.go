package config

const (
	ServiceName = "Whistleblower"
	Version     = "1.0.0"

	// Service envar
	envServerPort = "SERVER_PORT"

	// Database EnvVar
	envPostgresAddress     = "POSTGRES_ADDRESS"
	envPostgresName        = "POSTGRES_DATABASE"
	envPostgresUsername    = "POSTGRES_USERNAME"
	envPostgresPassword    = "POSTGRES_PASSWORD"
	envPostgresLogEnable   = "POSTGRES_LOG_ENABLED"
	envPostgresMaxOpenConn = "POSTGRES_MAX_OPEN_CONN"
	envPostgresMaxIdleConn = "POSTGRES_MAX_IDLE_CONN"

	envLogFormat = "LOG_FORMAT"
	envLogLevel  = "LOG_LEVEL"

	// Table names
	ArticleTableName   = "articles"
	ViewerTableName    = "viewers"
	CreatedAtFieldName = "created_at"
	UpdatedAtFieldName = "updated_at"
)
