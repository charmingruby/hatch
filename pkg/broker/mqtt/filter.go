package mqtt

import "strings"

// TopicMatchesFilter compares the topic to a filter, validating wildcards.
func TopicMatchesFilter(filter, topic string) bool {
	if filter == "" && topic == "" {
		return true
	}
	if filter == "" || topic == "" {
		return false
	}

	filterSegments := strings.Split(filter, "/")

	topicSegments := strings.Split(topic, "/")

	return matchSegments(filterSegments, topicSegments)
}

func matchSegments(filterSegments, topicSegments []string) bool {
	fLen, tLen := len(filterSegments), len(topicSegments)

	for i := range fLen {
		if i >= tLen {
			return false
		}

		fSeg := filterSegments[i]
		tSeg := topicSegments[i]

		if isMultiLevelWildcard(fSeg) {
			return handleMultiLevelWildcard(filterSegments, topicSegments, i)
		}

		if isSingleLevelWildcard(fSeg) {
			if tSeg == "" {
				return false
			}
			continue
		}

		if fSeg != tSeg {
			return false
		}
	}

	return fLen == tLen
}

func isMultiLevelWildcard(segment string) bool {
	return segment == "#"
}

func isSingleLevelWildcard(segment string) bool {
	return segment == "+"
}

func handleMultiLevelWildcard(filterSegments, topicSegments []string, index int) bool {
	if index == len(filterSegments)-1 {
		return true
	}

	for j := index; j < len(topicSegments); j++ {
		if matchSegments(filterSegments[index+1:], topicSegments[j:]) {
			return true
		}
	}

	return false
}
