package services

import (
	"bufio"
	"os"
	"strings"

	Utils "github.com/jadhamwi21/dota2-counterpick-cli-tool/utils"
)

type CounterServiceInterface interface {
	CounterSingleHero(heroName string) []string
	CounterMultipleHeroes(heroes []string) map[string][]string
	Initialize()
}

type CounterService struct {
	CounterList map[string]([]string)
}

func (counterService CounterService) CounterMultipleHeroes(heroes []string) string {
	return Utils.MultipleHeroesCountersFormater(heroes, counterService.CounterList)
}

func (counterService CounterService) CounterSingleHero(heroName string) string {
	return Utils.SingleHeroCountersFormatter(counterService.CounterList[heroName])
}

func (counterService *CounterService) Initialize() {
	countersListFile, _ := os.Open("static/countersList.txt")
	defer countersListFile.Close()
	listScanner := bufio.NewScanner(countersListFile)
	for listScanner.Scan() {
		slicedLine := strings.Split(listScanner.Text(), ";")[0:2]
		heroName := Utils.HeroNameFormatter(slicedLine[0])
		heroCounters := Utils.HeroCountersFormatter(slicedLine[1])
		counterService.CounterList[heroName] = heroCounters
	}
}
