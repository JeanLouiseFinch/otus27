package main

import (
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func TestMain(m *testing.M) {
	<-time.After(1 * time.Minute)
	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "progress", // Замените на "pretty" для лучшего вывода
		Paths:     []string{"features"},
		Randomize: 0, // Последовательный порядок исполнения
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
