package logging

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	logger := New()

	if logger == nil {
		t.Fatal("expected logger instance, got nil")
	}
}

func TestPrintfBroadcastsToSubscriber(t *testing.T) {
	logger := New()
	subscriber := logger.Subscribe()

	logger.Printf("hello %s", "world")

	message := <-subscriber
	if !strings.Contains(message, "hello world") {
		t.Fatalf("expected message to contain hello world, got %s", message)
	}
}

func TestUnsubscribe(t *testing.T) {
	logger := New()
	subscriber := logger.Subscribe()

	logger.Unsubscribe(subscriber)

	_, ok := <-subscriber
	if ok {
		t.Fatal("expected subscriber channel to be closed")
	}
}
