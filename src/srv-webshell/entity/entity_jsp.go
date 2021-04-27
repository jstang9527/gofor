package entity

type JSPEntity struct {
	Target  string
	Command string
}

func (e *JSPEntity) RunCmdWithOutput(cmd string) (string, error) {
	return "", nil
}
