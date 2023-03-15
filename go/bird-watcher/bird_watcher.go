package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, c := range birdsPerDay {
		total += c
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	lower := 7 * (week - 1)
	upper := lower + 7
	for i, c := range birdsPerDay {
		if i > lower && i < upper {
			total += c
		}
	}
	return total
}
func isOdd(i int) bool {
	if (i-1)%2 == 0 {
		return true
	}
	return false
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, c := range birdsPerDay {
		if isOdd(i) {
			birdsPerDay[i] = c + 1
		}
	}
	return birdsPerDay
}
