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
	BUILD_STATE  string = "development"
	BUILD_STATUS string = "dirty"
)

type Config struct {
	BaseURL string `json:"base_url"`
	// API Token is NEVER stored in a file
	OutputFormat string `json:"output_format"` // json, csv
	OutputFile   string `json:"output_file"`
	LogDir       string `json:"log_dir"`
	// Start and End Times are never given default values
	Duration string `json:"duration"` // Duration string for Date Range (default: 30d (720h)). uses months (mo), weeks (w), days (d), hours (h), minutes (m)
	// User ID must be specified by the user
	MaxItems int `json:"max_items"` // Max items to return (-1 for unlimited)
	MaxPages int `json:"max_pages"` // Max pages to retrieve (-1 for unlimited)
	PerPage  int `json:"per_page"`  // Items per page (default: 100, max: 100)
	// Body Data file must be specified by the user. No Default filename or path is used.
}

// Application Global Variables
var (
	APP_NAME           string = "Curricula"
	APP_STATE          string = fmt.Sprintf("%s (%s)", BUILD_STATE, BUILD_STATUS)
	APP_DATE           string = BUILD_DATE
	KEY_NAME           string = "curricula_key"
	VERSION_NO         string = fmt.Sprintf("%s-%s (%s)", VERSION, BUILD_NUMBER, BUILD_STATUS)
	BASE_URL           string = "http://localhost:8080"
	FLG_BASE_URL       string
	FLG_API_TOKEN      string
	OUTPUT_FORMAT      string = "json" // json, csv
	FLG_OUTPUT_FORMAT  string
	OUTPUT_FILE        string = "results.json" // default to stdout
	FLG_OUPUT_FILE     string
	LOG_DIR            string = "./logs" // default to ./logs
	FLG_LOG_DIR        string
	FLG_START_TIME     string
	FLG_END_TIME       string
	DURATION           string = "720h" // 30 days
	FLG_DURATION       string
	FLG_USER_ID        int
	MAX_ITEMS          int = -1 // default to -1
	FLG_MAX_ITEMS      int
	MAX_PAGES          int = -1 // default to -1
	FLG_MAX_PAGES      int
	PER_PAGE           int = 100 // default to 100
	FLG_PER_PAGE       int
	FLG_BODY_DATA_FILE string
)

func WriteConfig(file string) error {
	// Placeholder for writing config to a file
	fmt.Println("Writing Config file is not implemented yet")
	return nil
}

// ReadConfig reads the configuration from the specified file.
// File is read in and unmarshalled into a Config struct.
// from there the values are used to set the Global Variables.
func ReadConfig(file string) error {
	// Placeholder for reading config from a file
	fmt.Println("Reading Config file is not implemented yet")
	return nil
}
