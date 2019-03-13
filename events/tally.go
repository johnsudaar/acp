package events

// Single tally
const TallyEventName = "tally"

// All tallies
const TalliesEventName = "tallies"

type TallyEvent struct {
	Program bool
	Preview bool
}

type TalliesEvent struct {
	Program []string
	Preview []string
}
