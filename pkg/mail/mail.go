package mail

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Text    []byte
	Html    []byte
}

type Mail struct {
	Driver Driver
}
