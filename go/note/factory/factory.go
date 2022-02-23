package factory

type mes struct {
	C   string
	pwd string
}

func NewMes() *mes {
	return &mes{}
}

func (m *mes) SetPwd(p string) {
	m.pwd = p
}
