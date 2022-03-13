package main

const (
	REL_LINK = iota
	REL_A
	REL_FORM
	REL_AREA = REL_A
)

var relAttrs = map[string]map[int]bool{
	"alternate": {
		REL_LINK: true,
		REL_A:    true,
		REL_FORM: false,
	},
	"author": {
		REL_LINK: true,
		REL_A:    true,
		REL_FORM: false,
	},
	"bookmark": {
		REL_LINK: false,
		REL_A:    true,
		REL_FORM: false,
	},
	"canonical": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"dns-prefetch": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"external": {
		REL_LINK: false,
		REL_A:    true,
		REL_FORM: true,
	},
	"help": {
		REL_LINK: true,
		REL_A:    true,
		REL_FORM: true,
	},
	"icon": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"license": {
		REL_LINK: true,
		REL_A:    true,
		REL_FORM: true,
	},
	"manifest": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"modulepreload": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"next": {
		REL_LINK: true,
		REL_A:    true,
		REL_FORM: true,
	},
	"nofollow": {
		REL_LINK: false,
		REL_A:    true,
		REL_FORM: true,
	},
	"noopener": {
		REL_LINK: false,
		REL_A:    true,
		REL_FORM: true,
	},
	"noreferrer": {
		REL_LINK: false,
		REL_A:    true,
		REL_FORM: true,
	},
	"opener": {
		REL_LINK: false,
		REL_A:    true,
		REL_FORM: true,
	},
	"pingback": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"preconnect": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"prefetch": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"preload": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"prerender": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"prev": {
		REL_LINK: true,
		REL_A:    true,
		REL_FORM: true,
	},
	"search": {
		REL_LINK: true,
		REL_A:    true,
		REL_FORM: true,
	},
	"stylesheet": {
		REL_LINK: true,
		REL_A:    false,
		REL_FORM: false,
	},
	"tag": {
		REL_LINK: false,
		REL_A:    true,
		REL_FORM: false,
	},
}

func relAttributes(typ int) []string {
	var result []string
	for k, m := range relAttrs {
		if m[typ] {
			result = append(result, k)
		}
	}
	return result
}
