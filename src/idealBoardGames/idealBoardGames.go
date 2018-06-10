package idealBoardGames

func Quiz74(player []int) []int {
	people := []int{}
	for i := 0; i < len(player); i += 2 {
		if i+1 < len(player) {
			people = append(people, player[i])
		} else {
			people = append(people, player[i])
			people = people[1:]
		}
	}

	return people
}
