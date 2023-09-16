package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0

	for _, v := range birdsPerDay {
		sum += v
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// week 1 => startDay 0, endDay 6
	// week 2 => startDay 7, endDay 13
	// week 3 => startDay 14, endDay 20

	startDay := (week - 1) * 7
	endDay := startDay + 7
	weekSlice := birdsPerDay[startDay:endDay]

	return TotalBirdCount(weekSlice)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
