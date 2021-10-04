# Проект Shop Hookah

## Usage CLI:

<pre><code>$ go build -o mcli
Root command for our application, the main purpose is to help setup subcommands

  Usage:
    mcli [command]

  Available Commands:
    completion  generate the autocompletion script for the specified shell
    help        Help about any command
    migrate     используется для миграции базы данных

  Flags:
    -h, --help   help for mcli

  Use "mcli [command] --help" for more information about a command.

Use command migrate:

$ ./mcli migrate -h
Команда migrate используется для миграции базы данных: migrate < up | down >

  Usage: mcli migrate [command]

  Available Commands:
    down        Понижение версии БД до 1 версии
    up          Повышение версии БД до последней
    to-version  Повышение версии БД до определенной версии
    version     Версия миграции

  Flags:
    -h, --help   help for migrate

  Use "mcli migrate [command] --help" for more information about a command.</code></pre>