package gen

import "os"

// WriteAll writes all the services to the file system
func WriteAll(services []*Service) error {
	for _, service := range services {
		err := service.Write()
		if err != nil {
			return err
		}
	}
	return nil
}

// Writes the generated code to the file system
func (s *Service) Write() error {
	result, err := s.execute()
	if err != nil {
		return err
	}
	for _, file := range result.GoFiles {
		err := os.WriteFile(file.Path, []byte(file.Content), 0644)
		if err != nil {
			return err
		}
	}
	for _, file := range result.MarkDownFiles {
		err := os.WriteFile(file.Path, []byte(file.Content), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
