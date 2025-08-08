package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalcount int
	for i := 0; i < len(birdsPerDay); i++ {
		totalcount += birdsPerDay[i]
	}
	return totalcount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var startofweek int = (week - 1) * 7
	var endofweek int = (startofweek + 7)
	var weekcount int
	for i := startofweek; i < endofweek; i++ {
		weekcount += birdsPerDay[i]
	}
	return weekcount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay

}
