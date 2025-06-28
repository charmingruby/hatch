package mqtt

import "strings"

func TopicMatchesFilter(filter, topic string) bool {
	if filter == "" && topic == "" {
		return true
	}

	if filter == "" || topic == "" {
		return false
	}

	filterSegments := strings.Split(filter, "/")
	topicSegments := strings.Split(topic, "/")

	filterLen := len(filterSegments)
	topicLen := len(topicSegments)

	for i := range filterLen {
		filterSegment := filterSegments[i]

		if filterSegment == "#" {
			if i == filterLen-1 {
				return true
			}

			for j := i; j < topicLen; j++ {
				if TopicMatchesFilter(strings.Join(filterSegments[i+1:], "/"), strings.Join(topicSegments[j:], "/")) {
					return true
				}
			}
			return false
		}

		if i >= topicLen {
			return false
		}

		topicSegment := topicSegments[i]

		if filterSegment == "+" {
			if topicSegment == "" {
				return false
			}

			continue
		}

		if filterSegment != topicSegment {
			return false
		}
	}

	return filterLen == topicLen
}
