package database

import "os"

// MysqlAddress address of MySQL database
func MysqlAddress() string {
	value, exists := os.LookupEnv("MYSQL_ADDRESS")
	if !exists {
		value = "35.240.223.151"
	}
	return value
}
