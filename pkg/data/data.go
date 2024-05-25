package data

type AgeMap map[string]int

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
		"0 - 9":   8853,
		"10 - 19": 8205,
		"20 - 29": 7550,
		"30 - 39": 12240,
		"40 - 49": 10188,
		"50 - 59": 8408,
		"60 - 69": 7671,
		"70 - 79": 3113,
		"80 - +":  1424,
	},
	Married: 63,
	Race: map[string]float32{
		"slavs":        91.00,
		"middle_asian": 2.00,
		"exotic":       0.0000034,
		"other":        6.00,
	},
	Height: map[float32]float32{
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
	Salary: {
		"автомобильный бизнес": ISpeISphereDescription{
			Salary: ISphereRange{
				Start: 46,
				Mid:   92,
				End:   66,
			},
			Age: ISpereRange{
				Start: 20,
				Mid:   42,
				End:   53,
			},
		},
		"административный персонал": {
			"salary": {
				"start": 35,
				"mid":   62,
				"end":   43,
			},
			"age": {
				"start": 18,
				"mid":   36,
				"end":   64,
			},
		},
		"банки": {
			"salary": {
				"start": 47,
				"mid":   109,
				"end":   63,
			},
			"age": {
				"start": 19,
				"mid":   35,
				"end":   64,
			},
		},
		"безопасность": {
			"salary": {
				"start": 47,
				"mid":   79,
				"end":   61,
			},
			"age": {
				"start": 21,
				"mid":   41,
				"end":   59,
			},
		},
		"бухгалтерия": {
			"salary": {
				"start": 39,
				"mid":   96,
				"end":   69,
			},
			"age": {
				"start": 19,
				"mid":   43,
				"end":   61,
			},
		},
		"высший менеджмент": {
			"salary": {
				"start": 57,
				"mid":   198,
				"end":   164,
			},
			"age": {
				"start": 21,
				"mid":   45,
				"end":   59,
			},
		},
		"гос служба и НКО": {
			"salary": {
				"start": 46,
				"mid":   82,
				"end":   66,
			},
			"age": {
				"start": 22,
				"mid":   40,
				"end":   50,
			},
		},
		"домашний персонал": {
			"salary": {
				"start": 68,
				"mid":   63,
				"end":   58,
			},
			"age": {
				"start": 30,
				"mid":   44,
				"end":   59,
			},
		},
		"закупки": {
			"salary": {
				"start": 63,
				"mid":   102,
				"end":   81,
			},
			"age": {
				"start": 23,
				"mid":   41,
				"end":   50,
			},
		},
		"информационные технологии": {
			"salary": {
				"start": 43,
				"mid":   146,
				"end":   93,
			},
			"age": {
				"start": 18,
				"mid":   43,
				"end":   59,
			},
		},
		"искусство и медиа": {
			"salary": {
				"start": 36,
				"mid":   72,
				"end":   63,
			},
			"age": {
				"start": 18,
				"mid":   41,
				"end":   55,
			},
		},
		"маркетинг и реклама": {
			"salary": {
				"start": 34,
				"mid":   108,
				"end":   88,
			},
			"age": {
				"start": 18,
				"mid":   41,
				"end":   51,
			},
		},
		"медицина и фармацевтика": {
			"salary": {
				"start": 43,
				"mid":   87,
				"end":   74,
			},
			"age": {
				"start": 20,
				"mid":   35,
				"end":   59,
			},
		},
		"продажи": {
			"salary": {
				"start": 30,
				"mid":   81,
				"end":   59,
			},
			"age": {
				"start": 17,
				"mid":   38,
				"end":   61,
			},
		},
		"производство и сельское хозяйство": {
			"salary": {
				"start": 44,
				"mid":   93,
				"end":   87,
			},
			"age": {
				"start": 20,
				"mid":   41,
				"end":   64,
			},
		},
		"рабочий персонал": {
			"salary": {
				"start": 38,
				"mid":   53,
				"end":   38,
			},
			"age": {
				"start": 19,
				"mid":   36,
				"end":   61,
			},
		},
		"транспорт,логистика": {
			"salary": {
				"start": 40,
				"mid":   67,
				"end":   53,
			},
			"age": {
				"start": 19,
				"mid":   40,
				"end":   62,
			},
		},
		"туризм": {
			"salary": {
				"start": 40,
				"mid":   69,
				"end":   56,
			},
			"age": {
				"start": 18,
				"mid":   38,
				"end":   55,
			},
		},
		"тренинги и управление персоналом": {
			"salary": {
				"start": 43,
				"mid":   91,
				"end":   76,
			},
			"age": {
				"start": 20,
				"mid":   42,
				"end":   55,
			},
		},
		"фитнес и красота": {
			"salary": {
				"start": 41,
				"mid":   72,
				"end":   64,
			},
			"age": {
				"start": 18,
				"mid":   41,
				"end":   52,
			},
		},
		"юристы": {
			"salary": {
				"start": 46,
				"mid":   119,
				"end":   102,
			},
			"age": {
				"start": 21,
				"mid":   38,
				"end":   52,
			},
		},
		"консультирование": {
			"salary": {
				"start": 51,
				"mid":   133,
				"end":   142,
			},
			"age": {
				"start": 21,
				"mid":   34,
				"end":   41,
			},
		},
		"наука и образование": {
			"salary": {
				"start": 35,
				"mid":   67,
				"end":   65,
			},
			"age": {
				"start": 19,
				"mid":   49,
				"end":   59,
			},
		},
		"страхование": {
			"salary": {
				"start": 58,
				"mid":   70,
				"end":   76,
			},
			"age": {
				"start": 23,
				"mid":   34,
				"end":   48,
			},
		},
	},
}

func (s *IStats) calcAvgSalaryPerAge(age int) {
	// for sphere, v := range s.Salary {
	// 	isDecl :=
	// }
}
