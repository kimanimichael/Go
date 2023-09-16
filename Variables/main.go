package main

import "fmt"

func main()  {
	//Normal assignment
	var username string = "mk"
	fmt.Println(username)

	//Short assignment
	num := 10
	fmt.Println(num)

	//Print type of variable
	price := 3.0
	fmt.Printf("The type of the price variable is %T\n" ,price)

	//Same line declaration
	averageSales, bestEmployer := 400, "MK"
	fmt.Println(averageSales, bestEmployer)

	//Type casting
	averageWorkhours := 7.9
	averageWorkhoursInt := int8(averageWorkhours)
	fmt.Println("Work hours are:", averageWorkhoursInt)

	//Const variables

	//Short assignment not supported for const variables
	// const premiumPlanName := "Premium"

	const premiumPlanName = "Premium"
	const basicPlanName  = "Access"
	fmt.Println(premiumPlanName, basicPlanName)

	//String formatting
	const name = "MK"
	openRate := 30.5
	msg := fmt.Sprintf("Hell0 %v. Your open rate is %v percent", name, openRate)
	fmt.Println(msg)
}