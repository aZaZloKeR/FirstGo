package service

import (
	"strconv"
	"strings"
)

type ProfileVacancy struct {
	id         int32
	externalId string
}

func createProfile(body string) ProfileVacancy {
	mass := strings.Split(body, ";")
	val1, _ := strconv.ParseInt(mass[0], 10, 64)
	profileV := ProfileVacancy{
		id:         int32(val1),
		externalId: mass[1],
	}
	return profileV
}
