package main

import "math/rand"

type Salary interface {
	computeSalary()
}

type Fulltime struct {
	basic 	 int
	days	 int
}


type Contracter struct {
	basic 	 int
	days	 int
}

type Freelancer struct {
	basic 	 int
	hours	 int
}

func (emp *Fulltime) computeSalary()  {
	cal := emp.basic * emp.days
	println("The total salary for fulltime:", cal)
}

func (emp *Contracter) computeSalary()  {
	cal := emp.basic * emp.days
	println("The total salary for contracter:", cal)
}

func (emp *Freelancer) computeSalary()  {
	cal := emp.basic * emp.hours
	println("The total salary for freelancer:", cal)
}

func main(){

	fulltime := Fulltime{500, 28}
	contract := Contracter{100, 28}
	freelancer := Freelancer{ 10, rand.Intn(100)}

	fulltime.computeSalary()
	contract.computeSalary()
	freelancer.computeSalary()


}