package main

import "fmt"

var (
	input_str   string
	input_int   int
	input_float float32
)

type Candidate struct {
	name, location string
	age            int
	Education
	Address
}
type Education struct {
	degree, college, yearOfPassing string
}

type Address struct {
	streetName, City, State string
}

func main() {
	cand := new(Candidate)
	cand.getCandidateDetails()
	cand.displayCandidateDetails()
	fmt.Println("Press any key to exit")
	var selection string
	n, err := fmt.Scanf("%s\n", &selection)
	if err != nil || n != 1 {
		fmt.Println(err)
	}
}

func (cand *Candidate) getCandidateDetails() int {

	fmt.Println("Create Candidate Details")
	fmt.Println("Enter Candidate Name")
	cand.name = getStringINput()
	fmt.Println("Enter Candidate Location")
	cand.location = getStringINput()
	fmt.Println("Enter Candidate Age")
	cand.age = getIntINput()

	fmt.Println("Enter latest degree")
	degree_l := getStringINput()
	fmt.Println("Enter last college studied")
	college_l := getStringINput()
	fmt.Println("Enter Passed out year")
	yearOfPassing_l := getStringINput()
	cand.addEducation(degree_l, college_l, yearOfPassing_l)

	fmt.Println("Enter Street Name")
	streetName_l := getStringINput()
	fmt.Println("Enter City Name")
	city_l := getStringINput()
	fmt.Println("Enter State Name")
	state_l := getStringINput()
	cand.addAddress(streetName_l, city_l, state_l)
	return 0
}

// func (c *Candidate) addCandidate(street, city, state) int {
// 	c.addAddress(street, city, state)
// 	c.addEducation(deg, college, yop)
// 	return 0
// }

func (e *Education) addEducation(deg, college, yop string) int {
	e.degree = deg
	e.college = college
	e.yearOfPassing = yop
	return 0
}

func (a *Address) addAddress(street, city, state string) int {
	a.City = city
	a.State = state
	a.streetName = street
	return 0
}

func (c *Candidate) displayCandidateDetails() int {
	fmt.Println("Details to verify")
	for i := 1; i < 50; i++ {
		fmt.printf("*")
	}
	fmt.printf("\n")
	fmt.Println("Pre filled details are")
	fmt.Println(c)
	return 0
}

func getStringINput() string {
	n, err := fmt.Scanf("%s\n", &input_str)
	if err != nil || n != 1 {
		fmt.Println(err)
		return ""
	}
	return input_str

}
func getIntINput() int {
	n, err := fmt.Scanf("%d\n", &input_int)
	if err != nil || n != 1 {
		fmt.Println(err)
		return 0
	}
	return input_int
}
func getFloatINput() float32 {
	n, err := fmt.Scanf("%s\n", &input_float)
	if err != nil || n != 1 {
		fmt.Println(err)
		return 0.0
	}
	return input_float
}
