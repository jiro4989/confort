package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

type (
	Config struct {
		Action string `docopt:"<action>"`
		File   string `docopt:"<file>"`
		Key    string `docopt:"<key>"`
		Value  string `docopt:"<value>"`
	}
	ErrorCode int
)

const (
	doc = `confort は設定ファイルの値を更新するだけのコマンドです。

Usage:
	confort <action> <file> <key> <value>
	confort -h | --help
	confort -v | --version

Examples:
	confort add .bashrc GOPATH '$HOME/go'
	confort add --dry-run .bashrc GOPATH '$HOME/go'
	confort add httpd.conf Listen 1234
	confort disable httpd.conf Listen 80
	confort remove httpd.conf Listen 80
	confort update httpd.conf Files.Deny 'from all'

Available actions:
	add        キーバリューを追記する
	update     キーの更新だけする
	disable    行をコメントアウトする
	remove     行を削除する

Options:
	-h --help       このヘルプを出力する。
	-v --version    バージョン情報を出力する。`
)

const (
	errorCodeOk ErrorCode = iota
	errorCodeFailedBinding
	errorCodeFailedReadingStdin
	errorCodeIllegalAlphabet
	errorCodeIllegalConverter
	errorCodeIllegalCommand
)

func main() {
	os.Exit(int(Main(os.Args)))
}

func Main(argv []string) ErrorCode {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	err := args.Bind(&config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return errorCodeFailedBinding
	}

	switch config.Action {
	case "add":
		return errorCodeOk
	case "update":
		return errorCodeOk
	case "disable":
		return errorCodeOk
	case "remove":
		return errorCodeOk
	}

	// 到達しないはず
	return 99
}
