package cmd

import (
	"fmt"
	"os"
	"strings"

	http "github.com/sccicitb/pupr-backend/cmd/http"
	logrus "github.com/sirupsen/logrus"
	cobra "github.com/spf13/cobra"
)

// AppInfo application info structure
type AppInfo struct {
	AppName        string
	AppVersion     string
	AppCommit      string
	BuildGoVersion string
	BuildArch      string
	BuildDate      string
}

var (
	// meta
	app *AppInfo

	rootCmd = &cobra.Command{
		Long:  "pupr-backend Powered By AMETORY",
		Short: "pupr-backend",
		Use:   "pupr-backend",
	}

	// version sub command
	versionCmd = &cobra.Command{
		Long: "Print version information of pupr-backend",
		Run: func(command *cobra.Command, args []string) {
			infoStr := strings.Builder{}
			infoStr.WriteString(fmt.Sprintf("%s - ametory version info:\n", app.AppName))
			infoStr.WriteString(fmt.Sprintf("Version:\t%s\n", app.AppVersion))
			infoStr.WriteString(fmt.Sprintf("Commit Hash:\t%s\n", app.AppCommit))
			infoStr.WriteString(fmt.Sprintf("Go Version:\t%s\n", app.BuildGoVersion))
			infoStr.WriteString(fmt.Sprintf("Arch:\t\t%s\n", app.BuildArch))
			infoStr.WriteString(fmt.Sprintf("Build:\t\t%s\n", strings.Replace(app.BuildDate, "_", " ", -1)))

			logrus.Println(infoStr.String())
		},
		Short: "Print version info",
		Use:   "version",
	}
)

func init() {
	err := os.Setenv("TZ", "Asia/Jakarta")
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(http.ServeHTTP())
}

// Execute run root command
func Execute(appInfo *AppInfo) {
	app = appInfo
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// GetAppInfo return application information
func GetAppInfo() *AppInfo {
	return app
}
