package data

import (
	"fmt"
	"strconv"
	"strings"
)

type AgeMap map[string]int
type HeightMap map[float32]float32

type IStats struct {
	Total   int
	Age     AgeMap
	Married int
	Race    map[string]float32
	Height  map[float32]float32
	Salary  ISphere
}

type ISphere map[string]ISphereDescription
type ISphereRange struct {
	Start int
	Mid   int
	End   int
}
type ISphereDescription struct {
	Salary ISphereRange
	Age    ISphereRange
}

var Stats = IStats{
	Total: 146150789,
	Age: AgeMap{
		"10 - 19": 3745,
		"20 - 24": 3653,
		"25 - 29": 8582,
		"30 - 34": 6370,
		"35 - 39": 6150,
		"40 - 44": 5450,
		"45 - 49": 4690,
		"50 - 54": 4179,
		"55 - 59": 4621,
	},
	Married: 63,
	Race: map[string]float32{
		"slavs":        91.00,
		"middle_asian": 2.00,
		"exotic":       0.0000034,
		"other":        6.00,
	},
	Height: HeightMap{
		165.00: 0.14827,
		166.00: 0.66827,
		167.00: 1.17827,
		168.00: 1.68827,
		169.00: 2.19827,
		170.00: 2.71827,
		171.00: 3.22827,
		172.00: 3.73827,
		173.00: 4.25827,
		174.00: 5.74827,
		175.00: 7.14827,
		176.00: 8.54827,
		177.00: 9.94827,
		178.00: 9.60827,
		179.00: 8.10827,
		180.00: 6.60827,
		181.00: 5.10827,
		182.00: 4.01827,
		183.00: 3.54827,
		184.00: 3.07827,
		185.00: 2.61827,
		186.00: 2.14827,
		187.00: 1.67827,
		188.00: 1.20827,
		189.00: 0.74827,
		190.00: 0.27827,
		191.00: 0.005,
		192.00: 0.005,
		193.00: 0.005,
		194.00: 0.005,
		195.00: 0.005,
		196.00: 0.005,
		197.00: 0.005,
		198.00: 0.005,
		199.00: 0.005,
		200.00: 0.005,
	},
	Salary: ISphere{
		"автомобильный бизнес": ISphereDescription{
			Salary: ISphereRange{
				Start: 46,
				Mid:   92,
				End:   66,
			},
			Age: ISphereRange{
				Start: 20,
				Mid:   42,
				End:   53,
			},
		},
		"административный персонал": {
			Salary: ISphereRange{
				Start: 35,
				Mid:   62,
				End:   43,
			},
			Age: ISphereRange{
				Start: 18,
				Mid:   36,
				End:   64,
			},
		},
		"банки": {
			Salary: ISphereRange{
				Start: 47,
				Mid:   109,
				End:   63,
			},
			Age: ISphereRange{
				Start: 19,
				Mid:   35,
				End:   64,
			},
		},
		"безопасность": {
			Salary: ISphereRange{
				Start: 47,
				Mid:   79,
				End:   61,
			},
			Age: ISphereRange{
				Start: 21,
				Mid:   41,
				End:   59,
			},
		},
		"бухгалтерия": {
			Salary: ISphereRange{
				Start: 39,
				Mid:   96,
				End:   69,
			},
			Age: ISphereRange{
				Start: 19,
				Mid:   43,
				End:   61,
			},
		},
		"высший менеджмент": {
			Salary: ISphereRange{
				Start: 57,
				Mid:   198,
				End:   164,
			},
			Age: ISphereRange{
				Start: 21,
				Mid:   45,
				End:   59,
			},
		},
		"гос служба и НКО": {
			Salary: ISphereRange{
				Start: 46,
				Mid:   82,
				End:   66,
			},
			Age: ISphereRange{
				Start: 22,
				Mid:   40,
				End:   50,
			},
		},
		"домашний персонал": {
			Salary: ISphereRange{
				Start: 68,
				Mid:   63,
				End:   58,
			},
			Age: ISphereRange{
				Start: 30,
				Mid:   44,
				End:   59,
			},
		},
		"закупки": {
			Salary: ISphereRange{
				Start: 63,
				Mid:   102,
				End:   81,
			},
			Age: ISphereRange{
				Start: 23,
				Mid:   41,
				End:   50,
			},
		},
		"информационные технологии": {
			Salary: ISphereRange{
				Start: 43,
				Mid:   146,
				End:   93,
			},
			Age: ISphereRange{
				Start: 18,
				Mid:   43,
				End:   59,
			},
		},
		"искусство и медиа": {
			Salary: ISphereRange{
				Start: 36,
				Mid:   72,
				End:   63,
			},
			Age: ISphereRange{
				Start: 18,
				Mid:   41,
				End:   55,
			},
		},
		"маркетинг и реклама": {
			Salary: ISphereRange{
				Start: 34,
				Mid:   108,
				End:   88,
			},
			Age: ISphereRange{
				Start: 18,
				Mid:   41,
				End:   51,
			},
		},
		"медицина и фармацевтика": {
			Salary: ISphereRange{
				Start: 43,
				Mid:   87,
				End:   74,
			},
			Age: ISphereRange{
				Start: 20,
				Mid:   35,
				End:   59,
			},
		},
		"продажи": {
			Salary: ISphereRange{
				Start: 30,
				Mid:   81,
				End:   59,
			},
			Age: ISphereRange{
				Start: 17,
				Mid:   38,
				End:   61,
			},
		},
		"производство и сельское хозяйство": {
			Salary: ISphereRange{
				Start: 44,
				Mid:   93,
				End:   87,
			},
			Age: ISphereRange{
				Start: 20,
				Mid:   41,
				End:   64,
			},
		},
		"рабочий персонал": {
			Salary: ISphereRange{
				Start: 38,
				Mid:   53,
				End:   38,
			},
			Age: ISphereRange{
				Start: 19,
				Mid:   36,
				End:   61,
			},
		},
		"транспорт,логистика": {
			Salary: ISphereRange{
				Start: 40,
				Mid:   67,
				End:   53,
			},
			Age: ISphereRange{
				Start: 19,
				Mid:   40,
				End:   62,
			},
		},
		"туризм": {
			Salary: ISphereRange{
				Start: 40,
				Mid:   69,
				End:   56,
			},
			Age: ISphereRange{
				Start: 18,
				Mid:   38,
				End:   55,
			},
		},
		"тренинги и управление персоналом": {
			Salary: ISphereRange{
				Start: 43,
				Mid:   91,
				End:   76,
			},
			Age: ISphereRange{
				Start: 20,
				Mid:   42,
				End:   55,
			},
		},
		"фитнес и красота": {
			Salary: ISphereRange{
				Start: 41,
				Mid:   72,
				End:   64,
			},
			Age: ISphereRange{
				Start: 18,
				Mid:   41,
				End:   52,
			},
		},
		"юристы": {
			Salary: ISphereRange{
				Start: 46,
				Mid:   119,
				End:   102,
			},
			Age: ISphereRange{
				Start: 21,
				Mid:   38,
				End:   52,
			},
		},
		"консультирование": {
			Salary: ISphereRange{
				Start: 51,
				Mid:   133,
				End:   142,
			},
			Age: ISphereRange{
				Start: 21,
				Mid:   34,
				End:   41,
			},
		},
		"наука и образование": {
			Salary: ISphereRange{
				Start: 35,
				Mid:   67,
				End:   65,
			},
			Age: ISphereRange{
				Start: 19,
				Mid:   49,
				End:   59,
			},
		},
		"страхование": {
			Salary: ISphereRange{
				Start: 58,
				Mid:   70,
				End:   76,
			},
			Age: ISphereRange{
				Start: 23,
				Mid:   34,
				End:   48,
			},
		},
	},
}

func (s *IStats) calcAvgSalaryPerAge(age int) (float64, float64) {
	var salary float32
	var count int
	var manSalaryMultiplier = 1.15 //there is stats about ppl in general, mens salary a bit higher
	maxSalary := 0
	for _, v := range s.Salary {
		isSalaryChanging := (v.Salary.Mid - v.Salary.End) > 0
		isShouldSalaryRise := age > v.Age.Start && age <= v.Age.Mid
		var step float32

		if isSalaryChanging && isShouldSalaryRise {
			step = (float32(v.Salary.Mid - v.Salary.Start)) / (float32(v.Age.Mid - v.Age.Start))
			salary += float32(v.Salary.Start) + step*float32((age-v.Age.Start))

			if v.Salary.Mid > maxSalary {
				maxSalary = v.Salary.Mid
			}
		} else if isSalaryChanging && !isShouldSalaryRise {
			step = (float32(v.Salary.Mid - v.Salary.End)) / (float32(v.Age.End - v.Age.Mid))
			salary += float32(v.Salary.Mid) - step*float32((age-v.Age.Mid))
			if v.Salary.End > maxSalary {
				maxSalary = v.Salary.End
			}
		} else if !isSalaryChanging {
			salary += float32(v.Salary.End) //берем любое значение
			if v.Salary.End > maxSalary {
				maxSalary = v.Salary.End
			}
		}
		count += 1
	}
	avgSalary := (float64(salary) * manSalaryMultiplier) / float64(count)
	return avgSalary * 1000, float64(maxSalary) * manSalaryMultiplier
}

func (s *IStats) CalcChance(minAge int, maxAge int, race string, height int, money int, excludeMarried bool) float32 {
	var marriedChance float32
	if excludeMarried {
		marriedChance = .37
	} else {
		marriedChance = .63
	}

	heightChance := s.calcHeightChance(float32(height))
	ageChance := float32(s.calcAgeChance(float32(minAge), float32(maxAge)))
	salaryChance := s.calcSalaryChance(money)
	raceChance := s.Race[race] / 100
	if excludeMarried {
		marriedChance = 1 - float32(s.Married)
	} else {
		marriedChance = 1
	}
	fmt.Println(heightChance, " height", ageChance, " age", marriedChance, " married", salaryChance, " salary", raceChance, " race", " chanced")
	chance := heightChance * ageChance * marriedChance * salaryChance * raceChance
	chance = float32(int(chance*1000)) / 1000
	fmt.Println(heightChance, ageChance, salaryChance, raceChance, "fmt")
	return chance
}

func (s *IStats) calcSalaryChance(desiredSalary int) float32 {

	var totalPeople int
	var pplWithDesiredMoney float32
	var chance float32
	veryHighSalaryChance := 0.005 //this is uncalculated for, maybe will find stats later for men with salary >500k
	var avgSalary float64
	var maxSalary float64
	for age, ppl := range s.Age {
		ageRange := strings.Split(age, " - ")
		ageMin, _ := strconv.Atoi(ageRange[0])
		ageMax, _ := strconv.Atoi(ageRange[1])
		pplPerAge := float32(ppl) / float32(ageMax-ageMin+1)

		for i := ageMin; i <= ageMax; i++ {
			avgSalaryTemp, maxSalaryTemp := s.calcAvgSalaryPerAge(i)
			avgSalary = avgSalaryTemp
			if maxSalaryTemp > maxSalary {
				maxSalary = maxSalaryTemp
			}
			if avgSalary >= float64(desiredSalary) {
				pplWithDesiredMoney += pplPerAge
			}
			totalPeople += int(pplPerAge)
		}
		if float64(avgSalary) >= veryHighSalaryChance {
			chance += float32(veryHighSalaryChance)
		}
	}

	chance = pplWithDesiredMoney / float32(totalPeople)
	if chance == 0.00 && maxSalary > float64(desiredSalary) {
		chance = 2
	} else if chance == 0.00 && maxSalary < float64(desiredSalary) {
		chance = 0.005
	}
	return chance
}

func (s *IStats) calcAgeChance(minAge, maxAge float32) float32 {
	var totalPeople int
	var peopleInChance int

	for ageRange, count := range s.Age {
		ageSplit := strings.Split(ageRange, " - ")
		minAgeData, minOk := strconv.Atoi(string(ageSplit[0]))
		maxAgeData, maxOk := strconv.Atoi(string(ageSplit[1]))
		totalPeople += count
		if minOk == nil && maxOk == nil && ((float32(minAgeData) >= minAge && maxAge >= float32(minAgeData)) || (maxAge >= float32(maxAgeData) && minAge <= float32(maxAgeData))) {
			fmt.Println(minAge, maxAge, " min age max age")
			fmt.Println(minAgeData, maxAgeData, " minagedata maxagedata")
			fmt.Println(minAge >= float32(minAgeData), " min age >= minAgeData")
			fmt.Println(minAge <= float32(maxAgeData), " min age <= maxAgedata")
			fmt.Println("----", (minAge >= float32(minAgeData) && minAge <= float32(maxAgeData)))
			fmt.Println(maxAge >= float32(minAgeData), " max age >= minAgeData")
			fmt.Println(maxAge <= float32(maxAgeData), " max age <= maxAgeData")
			fmt.Println("---", (maxAge >= float32(minAgeData) && maxAge <= float32(maxAgeData)))
			fmt.Println("result :", (minAge >= float32(minAgeData) && minAge <= float32(maxAgeData)) || (maxAge <= float32(maxAgeData) || maxAge >= float32(minAgeData)))
			fmt.Println("-----------------------------------------------------")
			peopleInChance += count
		}
	}
	chance := float32(peopleInChance) / float32(totalPeople)
	return chance
}

func (s *IStats) calcHeightChance(height float32) float32 {
	var heightPerc float32

	for value, heightChance := range s.Height {
		if height > value {
			continue
		} else if height <= value {
			heightPerc += heightChance
		}
	}
	fmt.Println(heightPerc, " height chance")
	return heightPerc / 100
}
