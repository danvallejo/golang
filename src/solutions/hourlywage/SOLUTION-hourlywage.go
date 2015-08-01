package main

import "fmt"

const OvertimeHours = 40.0

func main() {
	hours := 0.0
	fmt.Print("Enter employee hours:")
	fmt.Scanf("%f\n", &hours)

	hourlyWage := 0.0
	fmt.Print("Enter employee hourly wage:")
	fmt.Scanf("%f\n", &hourlyWage)

	overtimeHours := 0.0
	overtimeWage := hourlyWage * 1.5

	if hours > OvertimeHours {
		overtimeHours = hours - OvertimeHours
		hours = OvertimeHours
	}

	regularWages := hours * hourlyWage
	overtimeWages := overtimeHours * overtimeWage

	wage := regularWages + overtimeWages

	fmt.Printf("Regular  %v\n", regularWages)
	fmt.Printf("Overtime %v\n", overtimeWages)

	fmt.Printf("Wage is  %v\n", wage)
}
