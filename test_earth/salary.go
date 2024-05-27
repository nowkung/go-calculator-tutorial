package main

import (
	"fmt"
	"time"
)

type work struct {
	date time.Time
	start_hour int
	end_hour int
}

type employee struct {
	types string
	salary int
	hour_rate int
	workdays []work
}

type employeePaycheck struct {
	types string
	workhour int
	ot int
	pay int
}

func newWork(date time.Time, start_hour int, end_hour int) *work {
	return &work{
        date: date,
        start_hour: start_hour,
        end_hour: end_hour,
    }
}

func newEmployee(types string, salary int, hour_rate int) *employee{
	return &employee{
		types: types,
        salary: salary,
        hour_rate: hour_rate,
        workdays: []work{},
	}
}

func paySalary(e *employee) employeePaycheck {
	pay := 0
	workTime := 0
	ot := 0
	if e.types == "Full-Time" {
		salary_hour := e.salary/22/8
		for _, w := range e.workdays {
            workTime = w.end_hour - w.start_hour
			ot = 0 // to do
			pay += salary_hour * workTime + ((salary_hour * 2) * ot)
		}

	} else if e.types == "Outsource" {
		for _, w := range e.workdays {
            workTime = w.end_hour - w.start_hour
			ot = 0 // to do
			pay += workTime * e.hour_rate + ((e.hour_rate * 3) * ot)
		}

	} else {
		fmt.Println("Invalid employee type")
	}
	return employeePaycheck{types: e.types, workhour: workTime, ot: ot, pay: pay}
}

func main() {
	w1 := newWork(time.Now(), 9, 17)

	e := newEmployee("Outsource", 0, 500)

	e.workdays = append(e.workdays, *w1)

	fmt.Println(paySalary(e))

	w := []int{1, 2, 3, 4, 5}

	w = append(w[:3], w[2:]...)
	fmt.Println(w)
    w[2] = 20
	fmt.Println(w)
	
}