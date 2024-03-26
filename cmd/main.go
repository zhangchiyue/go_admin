package main

import (
	"adx-admin/cmd/commands"
	"adx-admin/pkg/configer"
	"adx-admin/pkg/log"
	"adx-admin/pkg/microlog"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	Version   = "n/a"
	GitCommit = "n/a"
	GitTag    = "0"
	BuildTime = "n/a"
	GitAuthor = "none"

	Conf *configer.Config
	Log  microlog.Logger
)

func init() {
	Version = fmt.Sprintf("%s built on %s (commit: %s)", GitTag, BuildTime, GitCommit)
	commands.VersionMsg = &commands.VersionInfo{
		Tag:       GitTag,
		BuildTime: BuildTime,
		Commit:    GitCommit,
		Author:    GitAuthor,
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "server [--env local|dev|prod]",
		Short: "adx admin",
		Run: func(_ *cobra.Command, _ []string) {
			app, cleanup, err := wireApp(Conf, Log)
			if err != nil {
				panic(err)
			}
			defer cleanup()
			app.Run()
		},
		PreRunE: preEnv,
	}

	initFlags(rootCmd.PersistentFlags())

	if err := rootCmd.Execute(); err != nil {
		Log.Errorf("root cmd execute error: %v", err)
	}
}
func preEnv(cmd *cobra.Command, args []string) error {
	flags := cmd.PersistentFlags()
	if err := flags.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}
	return baseInit(flags)
}

func initFlags(flags *pflag.FlagSet) {
	flags.String("env", "dev", "environment: local|dev|prod")
}

func baseInit(flags *pflag.FlagSet) error {
	// 先初始化配置文件
	var err error
	if Conf, err = configer.Init(flags); err != nil {
		return err
	}
	// 初始化日志文件
	if Log, err = log.InitLog(Conf); err != nil {
		return err
	}
	return nil
}
