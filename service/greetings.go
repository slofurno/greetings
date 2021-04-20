package greetings

import (
	"context"
	"math/rand"
	"sync"

	"github.com/slofurno/greetings/service/greetingspb"

	"github.com/twitchtv/twirp"
)

type greeter struct {
	mu sync.Mutex

	greetings []string
}

func (s *greeter) GetGreeting(ctx context.Context, _ *greetingspb.Empty) (*greetingspb.Greeting, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	i := rand.Intn(len(s.greetings))

	return &greetingspb.Greeting{
		Greeting: s.greetings[i],
	}, nil
}

func (s *greeter) ListGreetings(ctx context.Context, _ *greetingspb.Empty) (*greetingspb.GreetingList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	greetings := []*greetingspb.Greeting{}
	for _, greeting := range s.greetings {
		greetings = append(greetings, &greetingspb.Greeting{
			Greeting: greeting,
		})
	}

	return &greetingspb.GreetingList{
		Greetings: greetings,
	}, nil
}

func (s *greeter) CreateGreeting(ctx context.Context, greeting *greetingspb.Greeting) (*greetingspb.Greeting, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(greeting.Greeting) < 2 {
		return nil, twirp.InvalidArgumentError("Greeting", "greeting required")
	}

	s.greetings = append(s.greetings, greeting.Greeting)
	return greeting, nil
}

func (s *greeter) Error(ctx context.Context, _ *greetingspb.Empty) (*greetingspb.GreetingList, error) {
	return nil, twirp.NewError(twirp.Unavailable, "taking a nap ...")
}

func New() *greeter {
	return &greeter{
		greetings: []string{"hello!", "hey!"},
	}
}
