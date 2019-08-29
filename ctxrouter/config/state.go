package config

import (
	"regexp"
)

type ConfigState struct {
	Hosts map[string]rutesByProtcols // hosts
}

type rutesByProtcols map[string][]routeRule // protocols eg: http, https...

type routeRule struct {
	Path          *regexp.Regexp
	RequestHeader rules
	CookieValues  rules
	FormValues    rules
	SessionValues rules
	CustomLogic   string
	Target        routeTarget
}

type rules struct {
	Require   map[string]values
	Forbidden map[string]values
}

type values struct {
	values []string
	Regexp *regexp.Regexp
}

type routeTarget struct {
	Addresses  []string
	Middleware []string
}
