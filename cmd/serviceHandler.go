package cmd

type Services struct {
	Services []Service `json:"services"`
}

type Service struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	IsRunning bool
	observers []Observer
}

func (s *Service) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Service) notifyObservers() {
	for _, observer := range s.observers {
		observer.Notify(s.Name, s.IsRunning)
	}
}

func (s *Service) SetIsRunning(isRunning bool) {
	if s.IsRunning != isRunning {
		s.IsRunning = isRunning
		s.notifyObservers()
	}
}
