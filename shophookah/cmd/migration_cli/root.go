package migration_cli

import (
	"github.com/spf13/cobra"
	"os"
	"shophookah/pkg/logg"
	msg "shophookah/pkg/logg/message"
)

var rootCmd = &cobra.Command{
	Use:   "mcli",
	Short: "Команда root для приложения",
	Long:  `Команда root для приложения, основная цель - помочь настроить подкоманды`,
}

// Execute - использует аргументы (по умолчанию os.Args [1:]) и проходит через дерево команд,
// находя соответствующие совпадения для команд, а затем соответствующие флаги.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logg.LogE(msg.ErrCommandExecute, err.Error())
		os.Exit(1)
	}
}
