package main

import "fmt"

type AnimalFarm struct {
	Name     string
	Weight   int
	Consumes int
	Produces int
}
type Farm struct {
	Animals []AnimalFarm
}

func (m *Farm) TotalNumbers() int {
	return len(m.Animals)
}

func (m *Farm) TotalWeight() int {
	if len(m.Animals) == 0 {
		return 0
	}
	sum := 0
	for _, v := range m.Animals {
		sum += v.Weight
	}
	return sum
}

func (m *Farm) TotalConsumes() int {
	if len(m.Animals) == 0 {
		return 0
	}
	sum := 0
	for _, v := range m.Animals {
		sum += v.Consumes
	}
	return sum

}

func (m *Farm) TotalProduces() int {
	if len(m.Animals) == 0 {
		return 0
	}
	sum := 0
	for _, v := range m.Animals {
		sum += v.Produces
	}
	return sum
}

//func (m *Farm) EqualWeight() int {
//	if len(m.Animals) == 0 {
//		return 0
//	}

//}

func main() {
	p := Farm{}
	p.Animals = []AnimalFarm{
		{Name: "Sygyr", Weight: 150, Consumes: 100, Produces: 50},
		{Name: "Goyun", Weight: 25, Consumes: 50, Produces: 10},
		{Name: "Geci", Weight: 15, Consumes: 20, Produces: 5},
		{Name: "Towuk", Weight: 5, Consumes: 5, Produces: 10},
	}
	totalNumbers := p.TotalNumbers()
	totalWeight := p.TotalWeight()
	totalConsumes := p.TotalConsumes()
	totalProduces := p.TotalProduces()
	totalProfit := p.TotalConsumes() - p.TotalProduces()

	fmt.Printf("Total Numbers: %d\n", totalNumbers)
	fmt.Printf("Total Weight: %d\n", totalWeight)
	fmt.Printf("Total Consumes: %d\n", totalConsumes)
	fmt.Printf("Total Produces: %d\n", totalProduces)
	fmt.Printf("Total Profit: %d\n", totalProfit)

}
