package serviceB

type server struct {
	name string
}

func MessageService() *server {
	return &server{
		name: "Server",
	}
}
