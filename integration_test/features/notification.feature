Feature: Sending messages to the grpc service
	As a client of the service API
	To understand that data is coming in and being recorded
	I want to receive events from the corresponding queue
	
	Scenario: getting an event from the queue for a new event
		When I send a new event with "myevent", "description", "2020-Feb-22", "2020-Feb-23"
		Then the error should be nil

	Scenario: deleting the received event
			When I send the event id
			Then the error should be nil
