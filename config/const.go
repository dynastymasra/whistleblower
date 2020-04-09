package config

const (
	ServiceName = "Whistleblower"
	Version     = "1.0.0"
	RequestID   = "request_id"

	// Headers
	HeaderRequestID = "X-Request-ID"

	// Database EnvVar
	envPostgresHost        = "POSTGRES_HOST"
	envPostgresPort        = "POSTGRES_PORT"
	envPostgresName        = "POSTGRES_NAME"
	envPostgresUsername    = "POSTGRES_USERNAME"
	envPostgresPassword    = "POSTGRES_PASSWORD"
	envPostgresEnableLog   = "POSTGRES_ENABLE_LOG"
	envPostgresMaxOpenConn = "POSTGRES_MAX_OPEN_CONN"
	envPostgresMaxIdleConn = "POSTGRES_MAX_IDLE_CONN"

	envLoggerFormat = "LOGGER_FORMAT"
	envLogLevel     = "LOG_LEVEL"
)
