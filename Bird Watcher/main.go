package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, numberOfBirds := range birdsPerDay {
		total += numberOfBirds
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0

	if len(birdsPerDay) <= 7 && week > 0 {
		for _, numberOfBirds := range birdsPerDay {
			total += numberOfBirds
		}
		return total
	}

	for _, numberOfBirds := range birdsPerDay[7*(week-1) : 7*week+1] {
		total += numberOfBirds
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
