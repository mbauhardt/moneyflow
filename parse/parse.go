package parse

import (
	"regexp"
	"strings"

	"github.com/mbauhardt/moneyflow/entities"
)

func ParseTags(commandline string) []entities.Tag {
	tags := make([]entities.Tag, 0)
	tags = append(tags, ParseTagsToAdd(commandline)...)
	tags = append(tags, ParseTagsToRemove(commandline)...)
	return tags
}

func ParseTagsToAdd(commandline string) []entities.Tag {
	tagRegex, _ := regexp.Compile("\\+[a-zA-Z0-9]+(\\s|$)")
	parsedTags := tagRegex.FindAllString(commandline, -1)
	tags := make([]entities.Tag, 0)
	for _, t := range parsedTags {
		parsedName := strings.Trim(t[1:len(t)], " ")
		tags = append(tags, entities.TagToAdd(parsedName))
	}
	return tags
}

func ParseTagsToRemove(commandline string) []entities.Tag {
	tagRegex, _ := regexp.Compile("\\-[a-zA-Z0-9]+(\\s|$)")
	parsedTags := tagRegex.FindAllString(commandline, -1)
	tags := make([]entities.Tag, 0)
	for _, t := range parsedTags {
		parsedName := strings.Trim(t[1:len(t)], " ")
		tags = append(tags, entities.TagToRemove(parsedName))
	}
	return tags
}
