package main

import (
	"testing"
	"time"
)

// An auxiliary function for creating a closing channel after a specified time.
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

// Test: If one channel is transmitted, the result must be the same channel.
func TestOr_SingleChannel(t *testing.T) {
	done := make(chan interface{})
	close(done)

	result := or(done)

	select {
	case <-result:
	case <-time.After(1 * time.Second):
		t.Fatalf("or did not close when single channel closed")
	}
}

// Test: If two channels are transmitted, or should close when either one is closed.
func TestOr_TwoChannels(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		time.Sleep(100 * time.Millisecond)
		close(ch1)
	}()

	result := or(ch1, ch2)

	select {
	case <-result:
	case <-time.After(1 * time.Second):
		t.Fatalf("or did not close when one of two channels closed")
	}
}

// Test: If multiple channels are transmitted, or should close when the first one closes.
func TestOr_MultipleChannels(t *testing.T) {
	channels := []<-chan interface{}{
		sig(500 * time.Millisecond),
		sig(100 * time.Millisecond),
		sig(1 * time.Second),
	}

	start := time.Now()
	result := or(channels...)

	<-result
	duration := time.Since(start)

	if duration < 100*time.Millisecond || duration > 200*time.Millisecond {
		t.Fatalf("or closed too late or too early: %v", duration)
	}
}

// Test: or should work with an empty channel list (result nil).
func TestOr_EmptyChannels(t *testing.T) {
	result := or()
	if result != nil {
		t.Fatalf("or did not return nil for empty input")
	}
}

// Test: Testing performance with a large number of channels.
func TestOr_ManyChannels(t *testing.T) {
	var channels []<-chan interface{}
	for i := 0; i < 1000; i++ {
		channels = append(channels, sig(10*time.Millisecond))
	}

	start := time.Now()
	result := or(channels...)

	<-result
	duration := time.Since(start)

	if duration > 20*time.Millisecond {
		t.Fatalf("or took too long for many channels: %v", duration)
	}
}
