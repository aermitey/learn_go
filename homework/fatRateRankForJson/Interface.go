package main

type Rank interface {
	GetInPut() *PersonalInformation
	SaveInformation(personInfo *PersonalInformation) error
	GetRank(name string) (rank int, fatRate float64)
}
