package main

import (
	"fmt"
	"slices"
)

// Solved 37 exercise

// ------------------------------
// A. Elite Flyer Qualification
// ------------------------------

// Passenger fly passenger.
type Passenger struct {
	name    string
	miles   int // Miles flown in a year
	flights int // Number of flights taken in a year
}

// Mathematical Expression for Elite Flyer Qualification:
// ∀x (P(x) → (F(x, m) ∧ m > 25000) ∨ (T(x, n) ∧ n > 25))
// Where:
// P(x): x is a passenger
// F(x, m): x flies m miles
// T(x, n): x takes n flights
// The passenger qualifies as an elite flyer if they fly more than 25,000 miles or take more than 25 flights in a year.

// qualifiesForEliteFlyer check if the passenger qualifies as an elite flyer.
func qualifiesForEliteFlyer(p Passenger) bool {
	// A passenger qualifies if they have flown more than 25,000 miles or taken more than 25 flights.
	return p.miles > 25000 || p.flights > 25
}

// ------------------------------
// B. Marathon Qualification
// ------------------------------

// Runner human runner.
type Runner struct {
	name   string
	gender string  // "male" or "female"
	time   float64 // Best previous time in hours
}

// Mathematical Expression for Marathon Qualification:
// ∀x ((M(x) ∧ T(x) < 3) → E(x)) for males
// ∀x ((W(x) ∧ Tw(x) < 3.5) → E(x)) for females
// Where:
// M(x): x is a male
// W(x): x is a female
// T(x): Time in hours for males to qualify
// Tw(x): Time in hours for females to qualify
// E(x): x qualifies for the marathon

// qualifiesForMarathon check if a runner qualifies for the marathon.
func qualifiesForMarathon(r Runner) bool {
	// For males, time should be less than 3 hours, for females less than 3.5 hours.
	if r.gender == "male" {
		return r.time < 3
	} else if r.gender == "female" {
		return r.time < 3.5
	}
	return false
}

// ------------------------------
// C. Master's Degree Qualification
// ------------------------------

// Student base student.
type Student struct {
	name   string
	hours  int    // Total course hours taken
	thesis bool   // Whether the student wrote a master's thesis
	grade  string // Highest grade received (e.g., "A", "B", etc.)
	credit int    // Credit hours taken in a semester
}

// Mathematical Expression for Master's Degree Qualification:
// ∀x (S(x) → (H(x) ≥ 60) ∨ (H(x) ≥ 45 ∧ Thesis(x)) ∧ G(x, B))
// Where:
// S(x): x is a student
// H(x): Total hours taken
// Thesis(x): x wrote a master's thesis
// G(x, B): x has a grade no lower than B in all required courses
// The student qualifies for a master's degree if they take at least 60 hours, or at least 45 hours and write a thesis,
// and their grade is no lower than "B".

// qualifiesForMasters check if a student qualifies for a master's degree
func qualifiesForMasters(s Student) bool {
	// A student qualifies if they take at least 60 hours or take at least 45 hours and write a thesis,
	// and they have a grade no lower than "B".
	return (s.hours >= 60 || (s.hours >= 45 && s.thesis)) && s.grade >= "B"
}

// ------------------------------
// D. Student with Credit Hours and All A's
// ------------------------------

// qualifiesForAllA check if a student has taken more than 21 credit hours and received all A's
func qualifiesForAllA(s Student) bool {
	// The student qualifies if they have taken more than 21 credit hours and received all A's.
	return s.credit > 21 && s.grade == "A"
}

func main() {
	// Example 1: Checking if a passenger qualifies for elite flyer status
	passenger := Passenger{name: "John", miles: 30000, flights: 30}
	fmt.Printf("Does %s qualify as an elite flyer? %v\n", passenger.name, qualifiesForEliteFlyer(passenger))

	// Example 2: Checking if a male runner qualifies for the marathon
	man := Runner{name: "Mike", gender: "male", time: 2.9}
	fmt.Printf("Does %s qualify for the marathon? %v\n", man.name, qualifiesForMarathon(man))

	// Example 3: Checking if a female runner qualifies for the marathon
	woman := Runner{name: "Anna", gender: "female", time: 3.4}
	fmt.Printf("Does %s qualify for the marathon? %v\n", woman.name, qualifiesForMarathon(woman))

	// Example 4: Checking if a student qualifies for a master's degree
	student1 := Student{name: "Alice", hours: 50, thesis: true, grade: "A", credit: 20}
	student2 := Student{name: "Bob", hours: 45, thesis: true, grade: "B", credit: 30}
	fmt.Printf("Does %s qualify for a master's degree? %v\n", student1.name, qualifiesForMasters(student1))
	fmt.Printf("Does %s qualify for a master's degree? %v\n", student2.name, qualifiesForMasters(student2))

	// Example 5: Checking if a student qualifies for all A's
	student3 := Student{name: "Charlie", hours: 60, thesis: true, grade: "A", credit: 25}
	student4 := Student{name: "David", hours: 60, thesis: true, grade: "B", credit: 22}
	fmt.Printf("Does %s have more than 21 credit hours and receive all A's? %v\n", student3.name, qualifiesForAllA(student3))
	fmt.Printf("Does %s have more than 21 credit hours and receive all A's? %v\n", student4.name, qualifiesForAllA(student4))
}

func sortArrayByParityII(nums []int) []int {
	slices.Sort()
	resp := make([]int, len(nums))

	// [4,1,2,1]
	// or
	// [4,2,1,1]
	// [1,2,1,4]
	// lastEven and lastOdd
	lastEven := 0
	lastOdd := 0

	for i := 0; i < len(nums); i++ {
		if isEven(i) {
			if isEven(nums[i]) {
				resp[i] = nums[i]
			} else {
				if isEven(nums[lastEven]) {
					resp[i] = nums[lastEven]
				}
				lastOdd = i
			}
		} else {
			if !isEven(nums[i]) {
				resp[i] = nums[i]
			} else {
				if !isEven(nums[lastOdd]) {
					resp[i] = nums[lastOdd]
				}
				lastEven = i
			}
		}

		if !isEven(lastEven) && !isEven(nums[lastOdd]) {
			resp[lastEven] = nums[lastOdd]
		}

		if isEven(lastOdd) && isEven(nums[lastEven]) {
			resp[lastOdd] = nums[lastEven]
		}

		fmt.Println(resp)
		fmt.Println(lastEven)
		fmt.Println(lastOdd)
	}

	return resp
}

func isEven(num int) bool {
	return num%2 == 0
}
