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
		"60 - 64": 4345,
		"65 - 69": 3218,
		"70 - +":  4480,
	},
	Married: 63,
	Race: map[string]float32{
		"slavs":        91.00,
		"Middle_asian": 2.00,
		"exotic":       0.0000034,
		"other":        6.00,
	},
	Height: HeightMap{
		165.00: 0.15,
		166.00: 0.67,
		167.00: 1.18,
		168.00: 1.69,
		169.00: 2.20,
		170.00: 2.72,
		171.00: 3.23,
		172.00: 3.74,
		173.00: 4.26,
		174.00: 5.75,
		175.00: 7.15,
		176.00: 8.55,
		177.00: 9.95,
		178.00: 9.61,
		179.00: 8.11,
		180.00: 6.61,
		181.00: 5.11,
		182.00: 4.02,
		183.00: 3.55,
		184.00: 3.08,
		185.00: 2.62,
		186.00: 2.15,
		187.00: 1.68,
		188.00: 1.21,
		189.00: 0.75,
		190.00: 0.28,
		191.00: 0.00,
		192.00: 0.00,
		193.00: 0.00,
		194.00: 0.00,
		195.00: 0.00,
		196.00: 0.00,
		197.00: 0.00,
		198.00: 0.00,
		199.00: 0.00,
		200.00: 0.00,
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

func (s *IStats) calcAvgSalaryPerAge(age int) float32 {
	var salary float32
	var count int
	for _, v := range s.Salary {
		isSalaryChanging := (v.Salary.Mid - v.Salary.End) > 0
		isShouldSalaryRise := age > v.Age.Start && age <= v.Age.Mid
		var step float32

		if isSalaryChanging && isShouldSalaryRise {
			step = (float32(v.Salary.Mid - v.Salary.Start)) / (float32(v.Age.Mid - v.Age.Start))
			salary = float32(v.Salary.Start) + step*float32((age-v.Age.Start))
		} else if isSalaryChanging && !isShouldSalaryRise {
			step = (float32(v.Salary.Mid - v.Salary.End)) / (float32(v.Age.End - v.Age.Mid))
			salary = float32(v.Salary.Mid) - step*float32((age-v.Age.Mid))
		} else if !isSalaryChanging {
			salary += float32(v.Salary.End) //берем любое значение
		}
		count += 1
	}
	newSalary := salary / float32(count)
	return newSalary
}

func (s *IStats) CalcChance(minAge int, maxAge int, race string, height int, money int, excludeMarried bool) float32 {
	var marriedChance float32
	if excludeMarried {
		marriedChance = .37
	} else {
		marriedChance = .63
	}

	heightChance := s.calcHeightChance(float32(height))
	ageChance := float32(s.calcAgeChance(float32(minAge), float32(maxAge))) / float32(s.Total)
	salaryChance := s.calcSalaryChance(money)
	if excludeMarried {
		marriedChance = 1 - float32(s.Married)
	} else {
		marriedChance = 1
	}
	chance := heightChance * ageChance * marriedChance * salaryChance
	return chance
}

func (s *IStats) calcSalaryChance(money int) float32 {

	var totalPeople int
	var pplWithDesiredMoney float32
	for age, ppl := range s.Age {
		ageRange := strings.Split(age, " - ")
		fmt.Println(ageRange)
		ageMin, _ := strconv.Atoi(ageRange[0])
		ageMax, _ := strconv.Atoi(ageRange[1])
		pplPerAge := float32(ppl) / float32(ageMax-ageMin+1)

		for i := ageMin; i <= ageMax; i++ {
			salary := s.calcAvgSalaryPerAge(i)
			if salary >= float32(money) {
				pplWithDesiredMoney += pplPerAge
			}
			totalPeople += int(pplPerAge)
		}
	}
	chance := pplWithDesiredMoney / float32(totalPeople)
	return chance
}

func (s *IStats) calcAgeChance(minAge, maxAge float32) float32 {
	var totalPeople int
	var peopleInChance int
	for ageRange, count := range s.Age {
		min, minOk := strconv.Atoi(string(ageRange[0]))
		max, maxOk := strconv.Atoi(string(ageRange[2]))
		totalPeople += count
		if minOk == nil && maxOk == nil && (minAge >= float32(min) && minAge <= float32(max)) || (maxAge <= float32(max) || maxAge >= float32(min)) {
			peopleInChance += count
		}
	}
	chance := float32(peopleInChance) / float32(totalPeople)
	return chance
}

func (s *IStats) calcHeightChance(value float32) float32 {
	var heightPerc float32
	for _, height := range s.Height {
		heightPerc += height
		if height == value {
			break
		}
	}
	return heightPerc
}
