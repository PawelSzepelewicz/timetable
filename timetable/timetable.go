package timetable

import (
	"fmt"
	"time"
	"math"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/alxshelepenok/timetable/utils"
)

type Timetable struct {
	Monday []*Subject `json:"monday"`
	Tuesday []*Subject `json:"tuesday"`
	Wednesday []*Subject `json:"wednesday"`
	Thursday []*Subject `json:"thursday"`
	Friday []*Subject `json:"friday"`
	Saturday []*Subject `json:"saturday"`
	Sunday []*Subject `json:"sunday"`
}

type Subject struct {
	Name string `json:"name"`
	Guru string `json:"guru"`
	Time string `json:"time"`
	Room string `json:"room"`
	Weeks []int `json:"weeks"`
}

func New(path string) (*Timetable, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	t := &Timetable{}

	err = json.Unmarshal(file, t)
	if err != nil {
		return nil, err
	}
	
	return t, nil
}

func (t *Timetable) getTimetable(now time.Time) (string, int) {
	weekday := now.Weekday().String()
 
	var day []*Subject

	switch(weekday) {
	case "Monday":
		day = t.Monday
	case "Tuesday":
		day = t.Tuesday
	case "Wednesday":
		day = t.Wednesday
	case "Thursday":
		day = t.Thursday
	case "Friday":
		day = t.Friday
	case "Saturday":
		day = t.Saturday
	case "Sunday":
		day = t.Sunday
	}

	var initYear int
	if (now.Month() >= time.January && now.Month() <= time.July) {
		initYear = now.Year() - 1
	} else {
		initYear = now.Year()
	}

	init := time.Date(initYear, time.September, 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())

	_, weeksFromNow := now.ISOWeek()
	_, weeksFromInit := init.ISOWeek()

	weeks := float64((53 - weeksFromInit) + weeksFromNow)
	numberWeek := int(math.Mod(weeks, 4))
	if (numberWeek == 0) {
		numberWeek = 4;
	}

	var subjects []string

	for _, subject := range day {
		if (utils.Contains(subject.Weeks, numberWeek)) {
			if ((len(subject.Room) == 0) && (len(subject.Guru) == 0) && (len(subject.Time) == 0) ) {
				subjects = append(subjects, fmt.Sprintf("%s", subject.Name))
			} else if ((len(subject.Room) == 0) && (len(subject.Guru) == 0)) {
				subjects = append(subjects, fmt.Sprintf("%s: %s", subject.Time, subject.Name))
			} else if (len(subject.Room) == 0) {
				subjects = append(subjects, fmt.Sprintf("%s: %s, %s", subject.Time, subject.Name, subject.Guru))
			} else {

				subjects = append(subjects, fmt.Sprintf("%s: %s, %s (%s)", subject.Time, subject.Name, subject.Guru, subject.Room))
			}
		}
	}
	
	return strings.Join(subjects, "\n"), len(subjects)
}

func (t *Timetable) Today() (string, int) {
	now := time.Now()

	return t.getTimetable(now)
}

func (t *Timetable) Nextday() (string, int) {
	now := time.Now().AddDate(0, 0, 1)

	return t.getTimetable(now)
}

func (t *Timetable) TodayWeek() int {
	now := time.Now()

	var initYear int
	if (now.Month() >= time.January && now.Month() <= time.July) {
		initYear = now.Year() - 1
	} else {
		initYear = now.Year()
	}

	init := time.Date(initYear, time.September, 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())

	_, weeksFromNow := now.ISOWeek()
	_, weeksFromInit := init.ISOWeek()

	weeks := float64((53 - weeksFromInit) + weeksFromNow)
	numberWeek := int(math.Mod(weeks, 4))
	if (numberWeek == 0) {
		numberWeek = 4;
	}

	return numberWeek
}

func (t *Timetable) NextdayWeek() int {
	now := time.Now().AddDate(0, 0, 1)

	var initYear int
	if (now.Month() >= time.January && now.Month() <= time.July) {
		initYear = now.Year() - 1
	} else {
		initYear = now.Year()
	}

	init := time.Date(initYear, time.September, 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())

	_, weeksFromNow := now.ISOWeek()
	_, weeksFromInit := init.ISOWeek()

	weeks := float64((53 - weeksFromInit) + weeksFromNow)
	numberWeek := int(math.Mod(weeks, 4))
	if (numberWeek == 0) {
		numberWeek = 4;
	}

	return numberWeek
}