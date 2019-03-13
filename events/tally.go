package events

const TallyEventName = "tally"

type TallyEvent struct {
	Program bool
	Preview bool
}
