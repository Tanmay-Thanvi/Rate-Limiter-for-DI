package service

type Level int

const (
	Organization = iota
	Resource
	API
	User
)

var Levels = []Level{Organization, Resource, API, User}
