package racer

import (
	"fmt"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// One of these will finish http.Get first and close the channel
	// The first to close the channel is the case that fires
	// time.After() will send a chan that will close and end the select
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// Closing the channel signals the http.Get is complete
		close(ch)
	}()
	return ch
}

//
// The code below works fine and is easier to understand
// But measureResponseTime is used in a way that is blocking
//

/*
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
*/
