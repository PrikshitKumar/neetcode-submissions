func carFleet(target int, position []int, speed []int) int {

	type Car struct {
		pos int
		time float64
	}

	cars := []Car{}

	for i:=0; i<len(position); i++ {
		time := float64(target - position[i]) / float64(speed[i])
		cars = append(cars, Car {
			pos: position[i],
			time: time,
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	maxTime := 0.0

	for _, car := range cars {
		if car.time > maxTime {
			fleets++
			maxTime = car.time
		}
	}

	return fleets

}
