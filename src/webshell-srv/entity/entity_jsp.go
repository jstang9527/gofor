package entity

type JSPEntity struct {
	Target  string
	Command string
}

func (e *JSPEntity) RunCmdWithOutput() (string, error) {
	return "", nil
}
