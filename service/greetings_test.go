package greetings

import (
	"context"
	"testing"

	"github.com/slofurno/greetings/service/greetingspb"
)

func TestCreateGreeting(t *testing.T) {
	svc := &greeter{}
	ctx := context.Background()

	greetings, err := svc.ListGreetings(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(greetings.Greetings) != 0 {
		t.Fatalf("expected 0 greetings, got: %d", len(greetings.Greetings))
	}

	if _, err := svc.CreateGreeting(ctx, &greetingspb.Greeting{}); err == nil {
		t.Fatal("expected error creating empty greeting")
	}

	expectedGreeting := "test greeting!"
	if _, err := svc.CreateGreeting(ctx, &greetingspb.Greeting{
		Greeting: expectedGreeting,
	}); err != nil {
		t.Fatal(err)
	}

	greetings, err = svc.ListGreetings(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(greetings.Greetings) != 1 {
		t.Fatalf("expected 1 greetings, got: %d", len(greetings.Greetings))
	}

	actualGreeting := greetings.Greetings[0].Greeting
	if actualGreeting != expectedGreeting {
		t.Fatalf("expected greeting: %s, got: %s", expectedGreeting, actualGreeting)
	}
}
