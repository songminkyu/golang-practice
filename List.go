package main

import "fmt"

func main() {
	// 슬라이스 선언 및 초기화
	fruits := []string{"사과", "바나나", "체리", "딸기", "포도"}

	// 리스트 출력
	fmt.Println("=== 과일 목록 ===")
	for i, fruit := range fruits {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}

	// 항목 추가
	fruits = append(fruits, "망고")
	fmt.Println("\n망고 추가 후:")
	fmt.Println(fruits)

	// 항목 삭제 (인덱스 1 삭제 - "바나나")
	fruits = append(fruits[:1], fruits[2:]...)
	fmt.Println("\n바나나 삭제 후:")
	fmt.Println(fruits)

	// 길이와 용량
	fmt.Printf("\n길이: %d, 용량: %d\n", len(fruits), cap(fruits))

	// 슬라이싱
	fmt.Println("\n처음 3개:", fruits[:3])

	// 정수 슬라이스
	numbers := []int{10, 20, 30, 40, 50}
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("\n숫자 합계:", sum)
}
