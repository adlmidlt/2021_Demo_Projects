package migration_cli

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"shophookah/internal/sh"
	l "shophookah/pkg/logg"
	msg "shophookah/pkg/logg/message"
	"strconv"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Используется для миграции базы данных",
	Long:  `Команда migrate используется для миграции базы данных: migrate < up | down >`,
}

// Init - Инициализация команд.
func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateVCmd := &cobra.Command{
		Use:   "version",
		Short: "Версия миграции",
		Long:  `Команда для информации о версии миграции`,
		Run: func(cmd *cobra.Command, args []string) {
			l.LogI("Запуск команды \"version\"")

			m, _ := NewMigrateInstance()
			verMigr, _, err := m.Version()
			l.LogI(fmt.Sprintf(msg.InfoCurrentVersMigr, verMigr))

			if verMigr == 0 {
				l.LogW(fmt.Sprintf(msg.WarnMigrVersNotExist, verMigr), err.Error())
			} else {
				l.LogI("Команда \"version\" выполнена успешно.")
			}
		},
	}
	migrateCmd.AddCommand(migrateVCmd)

	migrateDownCmd := &cobra.Command{
		Use:   "down",
		Short: "Понижение миграции БД до начальной версии",
		Long:  `Команда для понижения миграции БД до начальной версии`,
		Run: func(cmd *cobra.Command, args []string) {
			l.LogI("Запуск команды \"down\"")

			m, err := NewMigrateInstance()
			if err = m.Down(); err != nil {
				l.LogW(msg.WarnMigrFailedCommandDown, err.Error())
			} else {
				verMigr, _, _ := m.Version()
				l.LogI(fmt.Sprintf(msg.InfoCurrentVersMigr, verMigr))
				l.LogI("Команда \"down\" выполнена успешно.")
			}
		},
	}
	migrateCmd.AddCommand(migrateDownCmd)

	migrateToVCmd := &cobra.Command{
		Use:   "to-version",
		Short: "Миграция БД до определенной версии",
		Long:  `Команда для миграции БД до определенной версии (down-on version)`,
		Run: func(cmd *cobra.Command, args []string) {
			l.LogI("Запуск команды \"to-version\"")

			m, _ := NewMigrateInstance()
			verMigr := 0
			for _, ver := range args {
				num, err := strconv.Atoi(ver)
				sh.CheckError(msg.ErrConv, err)
				verMigr = num
			}
			if err := m.Migrate(uint(verMigr)); err != nil {
				l.LogW(fmt.Sprintf(msg.WarnMigrVersNotExist, verMigr), err.Error())
			} else {
				l.LogI(fmt.Sprintf(msg.InfoCurrentVersMigr, verMigr))
				l.LogI("Команда \"to-version\" выполнена успешно.")
			}
		},
	}
	migrateCmd.AddCommand(migrateToVCmd)

	migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "Повышение миграции БД до последней версии",
		Long:  `Команда для повышения миграции БД до последней версии`,
		Run: func(cmd *cobra.Command, args []string) {
			l.LogI("Запуск команды \"up\"")

			m, err := NewMigrateInstance()
			if err = m.Up(); err != nil {
				l.LogW(msg.WarnMigrFailedCommandUp, err.Error())
			} else {
				verMigr, _, _ := m.Version()
				l.LogI(fmt.Sprintf(msg.InfoCurrentVersMigr, verMigr))
				l.LogI("Команда \"up\" выполнена успешно.")
			}
		},
	}
	migrateCmd.AddCommand(migrateUpCmd)
}

// NewMigrateInstance - Новый экземпляр миграции.
func NewMigrateInstance() (*migrate.Migrate, error) {
	psqlDB := sh.CreateConnToDB()
	dbDriver, err := postgres.WithInstance(psqlDB, &postgres.Config{})
	sh.CheckError(msg.ErrCreateInstanceDB, err)

	fileSource, err := (&file.File{}).Open("file://../../internal/sh/migrations")
	sh.CheckError(msg.ErrFileWithMigrNotFound, err)

	m, err := migrate.NewWithInstance("file", fileSource, "shop_hookah", dbDriver)
	sh.CheckError(msg.ErrCreateNewMigrInstance, err)

	defer sh.CloseDB(psqlDB)
	return m, err
}
