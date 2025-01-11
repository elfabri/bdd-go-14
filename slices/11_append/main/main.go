/*
We've been asked to total costs per day, in a given period.

Complete the getCostsByDay() function using the append() function.
It accepts a slice of cost structs and returns a slice of float64s,
where each float64 represents the total cost for a day.

The length of the returned slice should be inferred from
the highest numbered day field. Some days may have multiple costs,
while others may have none. Include all actual and implied totals
in the returned slice, starting with day '0'. Use the append() function
to dynamically increase the size of the returned slice when needed.

Example
Given this input:

[]cost{
    {0, 4.0},
    {1, 2.1},
    {5, 2.5},
    {1, 3.1},
}

We expect this result:

[]float64{
    4.0, // first day
    5.2, // 2.1 + 3.1
    0.0, // intermediate days with no costs
    0.0, // ...
    0.0, // ...
    2.5, // last day
}
*/

package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {

    costByDay := []float64{0.0}
    d := 0
    v := 0.0

    for i:=0; i<len(costs); i++ {

        d = costs[i].day
        v = costs[i].value

        if d >= len(costByDay) {

            for j:=len(costByDay); j<=d; j++ {
                costByDay = append(costByDay, 0.0)
            }
        }

        costByDay[d] += v
    }

	return costByDay
}
