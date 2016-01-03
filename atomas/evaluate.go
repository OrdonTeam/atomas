package atomas
import (
	"container/list"
)

const (
	PLUS_SIGN  int = iota
)

func EvaluateBoard(arrayBoard []int) (int, []int) {
	score := 0
	multiplier := 1
	board := toList(arrayBoard)
	score, multiplier, board = lookForPossibleCombinations(board, multiplier)
	return score * multiplier, toArray(board)
}

func lookForPossibleCombinations(board *list.List, multiplier int) (int, int, *list.List) {
	score := 0
	for e := board.Front(); e != nil; e = e.Next() {
		if e.Value == PLUS_SIGN && shouldMergeElements(board, e) {
			score, multiplier, board = combineElements(board, e, multiplier)
		}
	}
	return score, multiplier, board
}

func combineElements(board *list.List, element *list.Element, multiplier int) (int, int, *list.List) {
	next := nextWithLoop(board, element)
	prev := prevWithLoop(board, element)
	surroundingValue := next.Value.(int)
	score := surroundingValue * 2
	element.Value = Max(surroundingValue, element.Value.(int)) + 1
	board.Remove(prev)
	board.Remove(next)
	if (shouldMergeElements(board, element)) {
		partialScore, multiplier, board := combineElements(board, element, multiplier + 1)
		return partialScore + score, multiplier, board
	} else if (aNewMergeEmerged(board)) {
		partialScore, multiplier, board := lookForPossibleCombinations(board, multiplier + 1)
		return partialScore + score, multiplier, board
	}else{
		return score, multiplier, board
	}
}

func shouldMergeElements(board *list.List, element *list.Element) bool {
	return (board.Len() > 2 && isSurroundingSame(board, element) && theyAreNotPluses(board, element))
}

func aNewMergeEmerged(board *list.List) bool {
	for e := board.Front(); e != nil; e = e.Next() {
		if e.Value == PLUS_SIGN && shouldMergeElements(board, e) {
			return true
		}
	}
	return false
}

func theyAreNotPluses(board *list.List, element *list.Element) bool {
	return nextWithLoop(board, element).Value != PLUS_SIGN
}

func nextWithLoop(board *list.List, element *list.Element) *list.Element {
	if (element.Next() != nil ) {
		return element.Next()
	}else {
		return board.Front()
	}
}

func prevWithLoop(board *list.List, element *list.Element) *list.Element {
	if (element.Prev() != nil ) {
		return element.Prev()
	}else {
		return board.Back()
	}
}

func isSurroundingSame(board *list.List, element *list.Element) bool {
	return nextWithLoop(board, element).Value == prevWithLoop(board, element).Value
}

func toList(board []int) *list.List {
	result := list.New()
	for _, element := range board {
		result.PushBack(int(element))
	}
	return result
}

func toArray(board *list.List) []int {
	array := make([]int, board.Len())
	i := 0
	for e := board.Front(); e != nil; e = e.Next() {
		array[i] = e.Value.(int)
		i += 1
	}
	return array
}
