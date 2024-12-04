package main

import (
	"fmt"
	"os"
	"strings"
)

func upGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if i < 1 {
		return ""
	}
	return wordSearch[i-1][j]
}

func downGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if i >= len(wordSearch)-1 {
		return ""
	}

	return wordSearch[i+1][j]
}

func leftGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if j < 1 {
		return ""
	}

	return wordSearch[i][j-1]
}

func rightGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if j >= len(wordSearch[i])-1 {
		return ""
	}

	return wordSearch[i][j+1]
}

func upLeftGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if i < 1 || j < 1 {
		return ""
	}

	return wordSearch[i-1][j-1]
}

func upRightGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if i < 1 || j >= len(wordSearch[i])-1 {
		return ""
	}

	return wordSearch[i-1][j+1]
}

func downLeftGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if i >= len(wordSearch)-1 || j < 1 {
		return ""
	}

	return wordSearch[i+1][j-1]
}

func downRightGrapheme(i, j int, wordSearch [][]string) (grapheme string) {
	if i >= len(wordSearch)-1 || j >= len(wordSearch[i])-1 {
		return ""
	}

	return wordSearch[i+1][j+1]
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	inputStr = strings.ToLower(inputStr)
	inputLines := strings.Split(inputStr, "\n")
	wordSearch := make([][]string, 0)

	for _, line := range inputLines {
		wordSearch = append(wordSearch, strings.Split(line, ""))
	}

	fmt.Println("Part 1:", part1(wordSearch))
	fmt.Println("Part 2:", part2(wordSearch))

}

func part1(wordSearch [][]string) int {
	xmasFound := 0

	for i, line := range wordSearch {
		for j, letter := range line {
			if letter != "x" {
				continue
			}

			if g := upGrapheme(i, j, wordSearch); g == "m" {
				if g := upGrapheme(i-1, j, wordSearch); g == "a" {
					if g := upGrapheme(i-2, j, wordSearch); g == "s" {
						fmt.Println("Found up xmas at", i, j)
						xmasFound++
					}
				}
			}

			if g := downGrapheme(i, j, wordSearch); g == "m" {
				if g := downGrapheme(i+1, j, wordSearch); g == "a" {
					if g := downGrapheme(i+2, j, wordSearch); g == "s" {
						fmt.Println("Found down xmas at", i, j)
						xmasFound++
					}
				}
			}

			if g := leftGrapheme(i, j, wordSearch); g == "m" {
				if g := leftGrapheme(i, j-1, wordSearch); g == "a" {
					if g := leftGrapheme(i, j-2, wordSearch); g == "s" {
						fmt.Println("Found left xmas at", i, j)
						xmasFound++
					}
				}
			}

			if g := rightGrapheme(i, j, wordSearch); g == "m" {
				if g := rightGrapheme(i, j+1, wordSearch); g == "a" {
					if g := rightGrapheme(i, j+2, wordSearch); g == "s" {
						fmt.Println("Found right xmas at", i, j)
						xmasFound++
					}
				}
			}

			if g := upLeftGrapheme(i, j, wordSearch); g == "m" {
				if g := upLeftGrapheme(i-1, j-1, wordSearch); g == "a" {
					if g := upLeftGrapheme(i-2, j-2, wordSearch); g == "s" {
						fmt.Println("Found up left xmas at", i, j)
						xmasFound++
					}
				}
			}

			if g := upRightGrapheme(i, j, wordSearch); g == "m" {
				if g := upRightGrapheme(i-1, j+1, wordSearch); g == "a" {
					if g := upRightGrapheme(i-2, j+2, wordSearch); g == "s" {
						fmt.Println("Found up right xmas at", i, j)
						xmasFound++
					}
				}
			}

			if g := downLeftGrapheme(i, j, wordSearch); g == "m" {
				if g := downLeftGrapheme(i+1, j-1, wordSearch); g == "a" {
					if g := downLeftGrapheme(i+2, j-2, wordSearch); g == "s" {
						fmt.Println("Found down left xmas at", i, j)
						xmasFound++
					}
				}
			}

			if g := downRightGrapheme(i, j, wordSearch); g == "m" {
				if g := downRightGrapheme(i+1, j+1, wordSearch); g == "a" {
					if g := downRightGrapheme(i+2, j+2, wordSearch); g == "s" {
						fmt.Println("Found down right xmas at", i, j)
						xmasFound++
					}
				}
			}
		}
	}

	return xmasFound
}

func part2(wordSearch [][]string) int {
	xmasFound := 0

	for i, line := range wordSearch {
		for j, letter := range line {
			if letter != "a" {
				continue
			}

			if g := upLeftGrapheme(i, j, wordSearch); g == "m" {
				if g := upRightGrapheme(i, j, wordSearch); g == "m" {
					if g := downLeftGrapheme(i, j, wordSearch); g == "s" {
						if g := downRightGrapheme(i, j, wordSearch); g == "s" {
							fmt.Println("Found up left xmas at", i, j)
							xmasFound++
						}
					}
				}
			}

			if g := upLeftGrapheme(i, j, wordSearch); g == "s" {
				if g := upRightGrapheme(i, j, wordSearch); g == "s" {
					if g := downLeftGrapheme(i, j, wordSearch); g == "m" {
						if g := downRightGrapheme(i, j, wordSearch); g == "m" {
							fmt.Println("Found up left xmas at", i, j)
							xmasFound++
						}
					}
				}
			}

			if g := upLeftGrapheme(i, j, wordSearch); g == "m" {
				if g := upRightGrapheme(i, j, wordSearch); g == "s" {
					if g := downLeftGrapheme(i, j, wordSearch); g == "m" {
						if g := downRightGrapheme(i, j, wordSearch); g == "s" {
							fmt.Println("Found up left xmas at", i, j)
							xmasFound++
						}
					}
				}
			}

			if g := upLeftGrapheme(i, j, wordSearch); g == "s" {
				if g := upRightGrapheme(i, j, wordSearch); g == "m" {
					if g := downLeftGrapheme(i, j, wordSearch); g == "s" {
						if g := downRightGrapheme(i, j, wordSearch); g == "m" {
							fmt.Println("Found up left xmas at", i, j)
							xmasFound++
						}
					}
				}
			}
		}
	}

	return xmasFound
}
