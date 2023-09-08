package services

import "time"

func (s *Service) Home() (map[string]string, error) {
	result, err := s.Redis.Remember("steve", time.Minute, func() string {
		return "hello world"
	})
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"result": result,
	}, nil
}
