package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Row = 7
const Column = 6
const winNum  = 4

type Chessboard struct {
	//chess board
	Chessboard [Column][Row]int
	//pieces count
	PiecesCount int
	//the winner
	WinPlayer int
}

func main() {
	fmt.Println("第一题：2 => case 1")
	printEmptyChessboard()

	fmt.Println("第一题：2 => case 2")
	printMyChessboard()
	fmt.Println("===============================================")

	fmt.Println("第二题：")
	startGame()
	fmt.Println("===============================================")

	fmt.Println("第三题：")
	playGame()
}


func printEmptyChessboard() {
	var emptyChessboard Chessboard
	printChessboard(emptyChessboard)
}

func printMyChessboard() {
	var myChessboard Chessboard
	myChessboard.Chessboard[0][3], myChessboard.Chessboard[3][2] = 1, 1
	myChessboard.Chessboard[2][3], myChessboard.Chessboard[3][5] = 2, 2
	printChessboard(myChessboard)
}

func startGame() {
	var gameChessboard Chessboard
	printChessboard(gameChessboard)
	fmt.Println("====== start game ======")

	gameChessboard.PiecesCount = 0
	for gameChessboard.PiecesCount<Column*Row {
		putPieces(&gameChessboard, 1)
		putPieces(&gameChessboard, 2)
	}
	printChessboard(gameChessboard)
}

func playGame() {
	var gameChessboard Chessboard
	printChessboard(gameChessboard)
	fmt.Println("====== start game ======")

	gameChessboard.PiecesCount = 0
	for gameChessboard.PiecesCount < Column*Row {
		putPieces(&gameChessboard, 1)
		if !gameIsContinue(&gameChessboard){
			break
		}
		putPieces(&gameChessboard, 2)
		if !gameIsContinue(&gameChessboard){
			break
		}
	}
	printChessboard(gameChessboard)

	if gameChessboard.WinPlayer==0{
		fmt.Println("it's a draw")
	} else {
		fmt.Print("the winner is: player")
		fmt.Println(gameChessboard.WinPlayer)
	}
}

func gameIsContinue(chessboard *Chessboard) bool {
	for i := 0; i < Column; i++ {
		for j := 0; j < Row; j++ {
			if chessboard.Chessboard[i][j] != 0 {
				//横
				if j <= winNum-1 && chessboard.Chessboard[i][j] == chessboard.Chessboard[i][j+1] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i][j+2] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i][j+3] {
					chessboard.WinPlayer = chessboard.Chessboard[i][j]
					return false
				}

				//竖
				if i <= Column-winNum && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+1][j] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+2][j] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+3][j] {
					chessboard.WinPlayer = chessboard.Chessboard[i][j]
					return false
				}

				//对角线
				if i <= Column-winNum && j <= winNum-1 && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+1][j+1] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+2][j+2] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+3][j+3] {
					chessboard.WinPlayer = chessboard.Chessboard[i][j]
					return false
				}
				if j >= winNum-1 && i <= Column-winNum && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+1][j-1] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+2][j-2] && chessboard.Chessboard[i][j] == chessboard.Chessboard[i+3][j-3] {
					chessboard.WinPlayer = chessboard.Chessboard[i][j]
					return false
				}
			}
		}
	}

	return true
}

func putPieces(chessboard *Chessboard, playerValue int) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i = rd.Intn(Row)
	for j := Column - 1; j >= 0; j-- {
		if chessboard.Chessboard[j][i] == 0 {
			chessboard.Chessboard[j][i] = playerValue
			chessboard.PiecesCount++
			break
		}
	}
}

func printChessboard(chessboard Chessboard) {
	for i := 0; i < Column; i++ {
		for j := 0; j < Row; j++ {
			switch chessboard.Chessboard[i][j] {
			case 0:
				fmt.Print("·")
			case 1:
				fmt.Print("x")
			case 2:
				fmt.Print("o")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}