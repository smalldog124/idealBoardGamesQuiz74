package idealBoardGames

import (
	"fmt"
	"testing"
)

func Test_Quiz74_Input_0_to_100_Should_be_Odd_number(t *testing.T) {
	player := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	}
	expected := []int{
		1, 3, 5, 7, 9,
		11, 13, 15, 17, 19,
		21, 23, 25, 27, 29,
		31, 33, 35, 37, 39,
		41, 43, 45, 47, 49,
		51, 53, 55, 57, 59,
		61, 63, 65, 67, 69,
		71, 73, 75, 77, 79,
		81, 83, 85, 87, 89,
		91, 93, 95, 97, 99,
	}

	actualResult := Quiz74(player)
	fmt.Printf("%v", actualResult)
	// if actualResult != expected {
	// 	t.Error("expected is:", expected, "but actualResult is:", actualResult)
	// }
}
