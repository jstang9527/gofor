package entity

import (
	"errors"
	"strings"
)

type LangType int

const (
	PHP_ENV = iota + 1
	JSP_ENV
	ASP_ENV
)

type LanguageEntity interface {
	RunCmdWithOutput(cmd string) (string, error)
}

func SwitchLangType(lang string) (LangType, error) {
	lang = strings.ToLower(lang)
	switch lang {
	case "php":
		return PHP_ENV, nil
	case "jsp":
		return JSP_ENV, nil
	case "asp":
		return ASP_ENV, nil
	default:
		return 0, errors.New("unsupport lang")
	}
}

func LoadLanguageEntity(ly LangType, target, command string) LanguageEntity {
	switch ly {
	case PHP_ENV:
		return &PHPEntity{Target: target, Command: command}
	case JSP_ENV:
		return &JSPEntity{Target: target, Command: command}
	default:
		return nil
	}
}
