# Publisher

A library for Golang that allows subscribing for a specific topic and receive data through a channel.

# Usage

Create a new instance of Publisher:

```go
publisher := broadcast.New[string]()
```

## Subscribe topic

Routines can subscribe specific topics by passing passing the corresponding topic to the Subscribe method:

```go
subscription := publisher.Subscribe("topic")
```

## Publishing data

Publishers can send data to all subscribers of a specific topic by calling the Publish method:

```go
publisher.Publish("topic", data)
```

## Deregistering a Subscriber

Subscribers can deregister from a topic by calling the Deregister method:

```go
publisher.Deregister("topic", subscription)
```

## Concurrency

The publisher uses a RWMutex to synchronize access to the subscriber list and ensure that data races do not occur.

```go

func main() {
	publisher := New[uint]()
	counterTopic := "count"

	go func() {
		stop := false
		subscription := publisher.Subscribe(counterTopic)

		for !stop {
			select {
			case no := <-subscription:
				fmt.Printf("Counter: %d \n", no)
			case <-time.After(time.Second * 2):
				stop = true
			}
		}
	}()

	time.Sleep(time.Second)

	go func() {
		for counter := uint(0); counter <= 3; counter++ {
			publisher.Publish(counterTopic, counter)
			time.Sleep(time.Second)
		}
	}()

	for counter := uint(4); counter <= 8; counter++ {
		publisher.Publish(counterTopic, counter)
		time.Sleep(time.Second)
	}
}
```
Prints:
```
Counter: 4 
Counter: 0 
Counter: 1 
Counter: 5 
Counter: 6 
Counter: 2 
Counter: 7 
Counter: 3 
Counter: 8 
```