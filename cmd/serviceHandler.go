package cmd

type Services struct {
	Services []Service `json:"services"`
}

type Service struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}
