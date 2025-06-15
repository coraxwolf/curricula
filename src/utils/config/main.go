/*
Package config holds global configuration variables for the application.
These are read at startup from a configuration file in the user's config directory.
These can be changed via the config set command with the flags provided with a value.
These can be cleared via the config clear [flags] command.
*/
package config

import "fmt"

// Version and build information. These variables are set via ldflags at build time.
var (
	VERSION      string = "v0.0.0"
	BUILD_NUMBER string = "0"
	BUILD_DATE   string = "1970-01-01T00:00:00Z"
	BUILD_STATE  string = ""
	BUILD_STATUS string = "dirty"
)

// Application Global Variables
var (
	APP_NAME           string = "Curricula"
	APP_STATE          string = fmt.Sprintf("%s (%s)", BUILD_STATE, BUILD_STATUS)
	APP_DATE           string = BUILD_DATE
	KEY_NAME           string = "curricula_key"
	VERSION_NO         string = fmt.Sprintf("%s-%s (%s)", VERSION, BUILD_NUMBER, BUILD_STATUS)
	FLG_BASE_URL       string = "http://localhost:8080"
	FLG_API_TOKEN      string = "changeme"
	FLG_OUTPUT_FORMAT  string = "json"    // json, csv
	FLG_OUPUT_FILE     string = "results" // default to stdout
	FLG_LOG_DIR        string = "./logs"  // default to ./logs
	FLG_START_TIME     string = ""        // Datetime for Date Range
	FLG_END_TIME       string = ""        // Datetime for Date Range
	FLG_DURATION       string = "720h"    // Duration string for Date Range (default: 30d (720h)). uses months (mo), weeks (w), days (d), hours (h), minutes (m)
	FLG_USER_ID        int    = 0         // User ID to filter by
	FLG_MAX_ITEMS      int    = -1        // Max items to return (-1 for unlimited)
	FLG_MAX_PAGES      int    = -1        // Max pages to retrieve (-1 for unlimited)
	FLG_PER_PAGE       int    = 100       // Items per page (default: 100, max: 100)
	FLG_BODY_DATA_FILE string = ""        // file containing the body json data (for POST/PUT/PATCH requests)
)
