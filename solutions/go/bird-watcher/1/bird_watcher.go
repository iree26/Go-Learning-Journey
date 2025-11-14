package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.

func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for i := 0; i < len(birdsPerDay); i++ {
        sum += birdsPerDay[i]
    }
      
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekly := 0
    start := (week -1)*7
    end := start + 7
    
    for i := start; i < end ; i++ {
        weekly += birdsPerDay[i]
    }
    return weekly
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
// So add one to all even indexs
func FixBirdCountLog(birdsPerDay []int) []int {
    end := len(birdsPerDay)

    for i := 0; i < end; i++ {
        if i % 2 == 0 {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay  
}




