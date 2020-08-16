package breadth

import (
	"container/list"
	"fmt"
)

type Graph = map[string][]string

func Breath(graph Graph, start string, search string) string {
	done := map[string]bool{}
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() != 0 {
		qElem := queue.Front()
		elem := qElem.Value.(string)
		moreDogs, ok := graph[elem]
		if !ok {
			panic(fmt.Sprintf("Must be on graph %q !", elem))
		}
		if hasSearchElem(moreDogs, search) {
			return elem
		}
		done[elem] = true
		for _, d := range moreDogs {
			if _, ok := done[d]; ok {
				continue
			}
			queue.PushBack(d)
		}
		queue.Remove(qElem)
	}
	return ""
}

func hasSearchElem(dogs []string, search string) bool {
	for _, d := range dogs {
		if d == search {
			return true
		}
	}
	return false
}
