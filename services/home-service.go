package services

func (s *Service) Home() map[string]string {
	return map[string]string{
		"data": "Hello world",
	}
}
