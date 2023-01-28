package broadcast

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPublishSubscribe(t *testing.T) {
	t.Parallel()

	publisher := New[string]()
	topicA := "a"
	topicB := "b"
	expectedSubAMsg := "A"
	expectedSubBMsg := "B"

	go func() {
		subscriptionA := publisher.Subscribe(topicA)
		actualSubAMsg := ""
		stop := false

		for !stop {
			select {
			case msg := <-subscriptionA:
				actualSubAMsg = msg
			case <-time.After(time.Second * 2):
				stop = true
			}
		}

		require.Equal(t, expectedSubAMsg, actualSubAMsg)
	}()

	go func() {
		subscriptionA := publisher.Subscribe(topicA)
		subscriptionB := publisher.Subscribe(topicB)
		actualSubAMsg := ""
		actualSubBMsg := ""
		stop := false

		for !stop {
			select {
			case msg := <-subscriptionA:
				actualSubAMsg = msg
			case msg := <-subscriptionB:
				actualSubBMsg = msg
			case <-time.After(time.Second * 2):
				stop = true
			}
		}

		require.Equal(t, expectedSubAMsg, actualSubAMsg)
		require.Equal(t, expectedSubBMsg, actualSubBMsg)
	}()

	time.Sleep(time.Second)

	publisher.Publish(topicA, expectedSubAMsg)
	publisher.Publish(topicB, expectedSubBMsg)
}
