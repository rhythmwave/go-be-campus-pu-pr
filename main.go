package main

import (
	"embed"
	"fmt"
	"runtime"
	"time"

	cmd "github.com/sccicitb/pupr-backend/cmd"
	"github.com/sccicitb/pupr-backend/cmd/http"
	log "github.com/sccicitb/pupr-backend/infra/log"
)

var (
	// application metadata
	appName    = "pupr-backend"
	appVersion = "development"
	appCommit  = "xxxxxxx"
	goVersion  = runtime.Version()
	buildDate  = time.Now().UTC().Format("2006-01-02_15:04:05_UTC")
	buildArch  = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

//go:embed data/migrations/*.sql
var embedMigration embed.FS

//go:embed data/migrations/logs/*.sql
var embedMigrationLog embed.FS

func getAppInfo() *cmd.AppInfo {
	http.EmbedMigration = embedMigration
	http.EmbedMigrationLog = embedMigrationLog

	if appVersion == "" {
		appVersion = "0.0.1"
	}

	return &cmd.AppInfo{
		AppCommit:      appCommit,
		AppName:        appName,
		AppVersion:     appVersion,
		BuildArch:      buildArch,
		BuildDate:      buildDate,
		BuildGoVersion: goVersion,
	}
}

func init() {
	log.PrintTimestamp()
}

func main() {
	cmd.Execute(getAppInfo())
}
