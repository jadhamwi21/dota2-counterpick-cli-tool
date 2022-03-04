package utils

import (
	"fmt"
	"strings"
)

func HeroNameFormatter(heroName string) string {
	return strings.ToLower(strings.Join(strings.Split(heroName, " "), "-"))
}

func HeroCountersFormatter(heroCounters string) []string {

	hero_counters := strings.Split(heroCounters, ",")
	for index := range hero_counters {
		hero_counters[index] = strings.ToLower(strings.TrimSpace(hero_counters[index]))
	}
	return hero_counters
}

func SingleHeroCountersFormatter(counters []string) string {
	return strings.Join(counters, "\n")

}

func MultipleHeroesCountersFormater(heroes []string, countersList map[string][]string) string {
	multipleHeroesCountersString := ""
	for _, heroName := range heroes {
		multipleHeroesCountersString += fmt.Sprintf("%v :\n%v\n------------------------\n", heroName, strings.Join(countersList[heroName], "\n"))
	}
	return multipleHeroesCountersString
}
