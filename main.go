package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/term"
)

const (
	InitialSelection = BoardSideSize * BoardSideSize / 2
)

func ParseData(rawInput [BoardSideSize]string) (board *PlayingBoard) {
	board = &PlayingBoard{}

	for row_idx, whole_row := range rawInput {
		if len(whole_row) != BoardSideSize {
			panic("invalid dimensions.. should be 9")
		}

		for col_idx, val := range whole_row {
			switch {
			case val >= '1' && val <= '9':
				val, err := strconv.Atoi(string(val))
				if err != nil {
					panic("value could not be converted")
				}
				board[row_idx*BoardSideSize+col_idx] = NewPresetSlot(val)
			case val == '_':
				board[row_idx*BoardSideSize+col_idx] = NewFillableSlot()
			}

		}
	}

	// selected slot will be the one in the middle
	board[InitialSelection].SetSelected(true)

	return
}

func DrawSlot(slot Slot) {
	if slot.IsSelected() {
		if slot.IsEmpty() {
			Red("*", true)
		} else {
			Red(strconv.Itoa(slot.Value()), true)
		}
		return
	}

	if slot.IsPreset() {
		Green(strconv.Itoa(slot.Value()), false)
	} else if slot.IsEmpty() {
		fmt.Print(" ")
	} else {
		Yellow(strconv.Itoa(slot.Value()), slot.IsSelected())
	}
}

func DrawBoard(board *PlayingBoard) {
	fmt.Print("═════════════════════════════════════\n")
	for x := 0; x < BoardSideSize; x++ {
		for y := 0; y < BoardSideSize; y++ {
			if y%3 == 0 {
				fmt.Print("║ ")
			} else {
				fmt.Print("| ")
			}
			DrawSlot(board[x*BoardSideSize+y])
			fmt.Print(" ")
		}

		if x%3 == 2 {
			fmt.Print("║\n═════════════════════════════════════\n")
		} else {
			fmt.Print("║\n║-----------║-----------║-----------║\n")
		}
	}
}

func SetValue(slot Slot, value byte) {
	as_string := string(value)
	parsed, err := strconv.Atoi(as_string)
	if err == nil {
		slot.SetValue(parsed)
	}
}

func Verify(board *PlayingBoard) bool {
	for _, val := range board {
		if val.IsEmpty() {
			// not yet finished
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Hello")

	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var raw_data [BoardSideSize]string = [BoardSideSize]string{
		"_97___52_",
		"3_2_1_8__",
		"_6_4271__",
		"____3__4_",
		"7_45__6__",
		"_197__3_8",
		"9____14__",
		"8__34_27_",
		"_7_25698_",
	}
	var board *PlayingBoard = ParseData(raw_data)

	var b []byte = make([]byte, 1)
	var currentSelection int = InitialSelection
	for {
		// clear screen
		fmt.Print("\033[H\033[2J")
		// draw board
		DrawBoard(board)
		//wait for user input

		count, err := os.Stdin.Read(b)
		if count == 0 || err != nil {
			continue
		}

		newSelection := currentSelection
		switch {
		case b[0] >= '0' && b[0] <= '9':
			SetValue(board[currentSelection], b[0])

		case b[0] == 'w':
			if currentSelection-BoardSideSize >= 0 {
				newSelection = currentSelection - BoardSideSize
			}

		case b[0] == 's':
			if currentSelection+BoardSideSize < BoardSideSize*BoardSideSize {
				newSelection = currentSelection + BoardSideSize
			}

		case b[0] == 'a':
			if currentSelection-1 >= 0 && (((currentSelection - 1) / BoardSideSize) == (currentSelection / BoardSideSize)) {
				newSelection = currentSelection - 1
			}

		case b[0] == 'd':
			if (currentSelection+1 <= BoardSideSize*BoardSideSize) && (((currentSelection + 1) / BoardSideSize) == (currentSelection / BoardSideSize)) {
				newSelection = currentSelection + 1
			}
		default:
			continue
		}

		if newSelection != currentSelection {
			board[currentSelection].SetSelected(false)
			board[newSelection].SetSelected(true)
			currentSelection = newSelection
		}

		if Verify(board) {
			fmt.Println("Congrats!!!")
		}
	}
}
