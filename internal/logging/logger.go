package logging

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	subscribers map[chan string]struct{}
	mu          sync.Mutex
}

func New() *Logger {

	return &Logger{
		subscribers: make(map[chan string]struct{}),
	}
}

func (l *Logger) Printf(format string, values ...any) {

	message := fmt.Sprintf(format, values...)
	line := time.Now().Format("2006/01/02 15:04:05 ") + message

	fmt.Fprintln(os.Stdout, line)
	l.broadcast(line)
}

func (l *Logger) Subscribe() chan string {

	subscriber := make(chan string, 10)

	l.mu.Lock()
	l.subscribers[subscriber] = struct{}{}
	l.mu.Unlock()

	return subscriber
}

func (l *Logger) Unsubscribe(subscriber chan string) {

	l.mu.Lock()
	delete(l.subscribers, subscriber)
	close(subscriber)
	l.mu.Unlock()
}

func (l *Logger) broadcast(message string) {

	l.mu.Lock()
	defer l.mu.Unlock()

	for subscriber := range l.subscribers {

		select {
		case subscriber <- message:
		default:
		}
	}
}
