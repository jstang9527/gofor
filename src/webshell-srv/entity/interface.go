package entity

type LangType int

const (
	PHP_ENV = iota + 1
	JSP_ENV
	ASP_ENV
)

type LanguageEntity interface {
	RunCmdWithOutput() (string, error)
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
