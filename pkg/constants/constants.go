package constants

var (
	// Server Props
	ServerPort = 9010
	// Key to Lookup from ENV
	DBHostKey     = "DB_HOST"
	DBPortKey     = "DB_PORT"
	DBDatabaseKey = "DB_DATABASE"
	DBUserKey     = "DB_USER"
	// Defaults for local testing
	SQL_USER_LOCAL     = "root"
	SQL_HOSTNAME_LOCAL = "127.0.0.1"
	SQL_PORT_LOCAL     = "3306"
	SQL_DB_LOCAL       = "database"
	// Other defaults for conn
	SQL_CONN_METHOD      = "tcp"
	SQL_ADDITIONAL_PROPS = "charset=utf8&parseTime=True&loc=Local"
	SQL_TABLE_NAME       = "lists"
)
