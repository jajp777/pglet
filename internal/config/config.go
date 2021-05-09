package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const (

	// general settings
	defaultServerPort              = 5000
	serverPort                     = "SERVER_PORT"
	forceSSL                       = "FORCE_SSL"
	defaultWebSocketMaxMessageSize = 65535
	wsMaxMessageSize               = "WS_MAX_MESSAGE_SIZE"

	// pages/sessions
	defaultPageLifetimeMinutes = 1440
	defaultAppLifetimeMinutes  = 60
	defaultCookieName          = "pglet"
	pageLifetimeMinutes        = "PAGE_LIFETIME_MINUTES"
	appLifetimeMinutes         = "APP_LIFETIME_MINUTES"
	checkPageIP                = "CHECK_PAGE_IP" // unauthenticated clients only
	limitPagesPerHour          = "LIMIT_PAGES_PER_HOUR"
	limitSessionsPerHour       = "LIMIT_SESSIONS_PER_HOUR"
	limitSessionSizeBytes      = "LIMIT_SESSION_SIZE_BYTES"
	reservedAccountNames       = "RESERVED_ACCOUNT_NAMES"
	reservedPageNames          = "RESERVED_PAGE_NAMES"
	allowRemoteHostClients     = "ALLOW_REMOTE_HOST_CLIENTS"
	hostClientsAuthToken       = "HOST_CLIENTS_AUTH_TOKEN"

	// redis
	defaultRedisMaxIdle   = 5
	defaultRedisMaxActive = 10
	redisAddr             = "REDIS.ADDR"
	redisPassword         = "REDIS.PASSWORD"
	redisMaxIdle          = "REDIS.MAX_IDLE"
	redisMaxActive        = "REDIS.MAX_ACTIVE"

	// auth
	cookieName         = "COOKIE_NAME"
	cookieSecrets      = "COOKIE_SECRETS"
	githubClientID     = "GITHUB_CLIENT_ID"
	githubClientSecret = "GITHUB_CLIENT_SECRET"
	azureClientID      = "AZURE_CLIENT_ID"
	azureClientSecret  = "AZURE_CLIENT_SECRET"

	// security
	masterSecretKey        = "MASTER_SECRET_KEY"
	defaultMasterSecretKey = "master_secret_key"
)

var (
	defaultCookieSecrets = []string{"secret_hash secret_encrypt"}
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if runtime.GOOS == "windows" {
		viper.AddConfigPath(filepath.Join(os.Getenv("ProgramData"), "pglet"))
		viper.AddConfigPath(filepath.Join(os.Getenv("USERPROFILE"), ".pglet"))
	} else {
		viper.AddConfigPath("/etc/pglet")
		viper.AddConfigPath("$HOME/.config/pglet")
	}
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("error reading config file: %s", err))
		}
	}
	viper.SetEnvPrefix("pglet")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// general
	viper.SetDefault(serverPort, defaultServerPort)
	viper.SetDefault(wsMaxMessageSize, defaultWebSocketMaxMessageSize)

	// pages/sessions
	viper.SetDefault(pageLifetimeMinutes, defaultPageLifetimeMinutes)
	viper.SetDefault(appLifetimeMinutes, defaultAppLifetimeMinutes)

	// redis
	viper.SetDefault(redisMaxIdle, defaultRedisMaxIdle)
	viper.SetDefault(redisMaxActive, defaultRedisMaxActive)

	// auth
	viper.SetDefault(cookieName, defaultCookieName)
	viper.SetDefault(cookieSecrets, defaultCookieSecrets)

	// security
	viper.SetDefault(masterSecretKey, defaultMasterSecretKey)
}

func ServerPort() int {
	return viper.GetInt(serverPort)
}

func MaxWebSocketMessageSize() int {
	return viper.GetInt(wsMaxMessageSize)
}

func ForceSSL() bool {
	return viper.GetBool(forceSSL)
}

func AllowRemoteHostClients() bool {
	return viper.GetBool(allowRemoteHostClients)
}

func HostClientsAuthToken() string {
	return viper.GetString(hostClientsAuthToken)
}

func RedisAddr() string {
	return viper.GetString(redisAddr)
}

func RedisPassword() string {
	return viper.GetString(redisPassword)
}

func RedisMaxIdle() int {
	return viper.GetInt(redisMaxIdle)
}

func RedisMaxActive() int {
	return viper.GetInt(redisMaxActive)
}

func PageLifetimeMinutes() int {
	return viper.GetInt(pageLifetimeMinutes)
}

func AppLifetimeMinutes() int {
	return viper.GetInt(appLifetimeMinutes)
}

func CheckPageIP() bool {
	return viper.GetBool(checkPageIP)
}

func ReservedAccountNames() []string {
	return viper.GetStringSlice(reservedAccountNames)
}

func ReservedPageNames() []string {
	return viper.GetStringSlice(reservedPageNames)
}

func LimitPagesPerHour() int {
	return viper.GetInt(limitPagesPerHour)
}

func LimitSessionsPerHour() int {
	return viper.GetInt(limitSessionsPerHour)
}

func LimitSessionSizeBytes() int {
	return viper.GetInt(limitSessionSizeBytes)
}

// Auth

func CookieName() string {
	return viper.GetString(cookieName)
}

func CookieSecrets() []string {
	return viper.GetStringSlice(cookieSecrets)
}

func GithubClientID() string {
	return viper.GetString(githubClientID)
}

func GithubClientSecret() string {
	return viper.GetString(githubClientSecret)
}

func AzureClientID() string {
	return viper.GetString(azureClientID)
}

func AzureClientSecret() string {
	return viper.GetString(azureClientSecret)
}

// Security

func MasterSecretKey() string {
	return viper.GetString(masterSecretKey)
}
