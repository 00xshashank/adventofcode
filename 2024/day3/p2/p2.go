package main

import (
	"fmt"
	"os"
	"strconv"
)

type Parser struct {
	s          string
	idx        int
	stringLen  int
	mulEnabled bool
}

func (p *Parser) parseNumber() int {
	fmt.Printf("--- Parse Number ---\n")
	j := 0

	for j <= 4 {
		fmt.Printf("Checking idx: %d, value at idx: %c\n", p.idx+j, p.s[p.idx+j])
		if !(p.s[p.idx+j] >= '0' && p.s[p.idx+j] <= '9') {
			if p.s[p.idx+j] == ',' || p.s[p.idx+j] == ')' {
				break
			} else {
				return -1
			}
		}

		j++
	}

	if p.s[p.idx+j] == ',' || p.s[p.idx+j] == ')' {
		n, err := strconv.Atoi(p.s[p.idx : p.idx+j])
		if err != nil {
			fmt.Printf("=== ERROR WHILE CONVERTING %s TO NUMBER ===\n%s\n", p.s[j:p.idx], err)
		}
		return n
	} else {
		return -1
	}
}

func (p *Parser) parseMulInstruction() int {
	fmt.Printf("--- Match mul instruction ---\nNext few characters: %s\n", p.s[p.idx:p.idx+6])
	if p.idx < p.stringLen-5 && p.s[p.idx:p.idx+4] == "mul(" {
		p.idx += 4
		n1 := p.parseNumber()
		if n1 == -1 && p.s[p.idx] != ',' {
			return 0
		}
		p.idx++
		n2 := p.parseNumber()
		if n2 == -1 && p.s[p.idx] != ')' {
			return 0
		}
		p.idx++
		fmt.Printf("Result from multiplication: %d\n", n1*n2)
		return n1 * n2
	}

	fmt.Println("Absent")
	return 0
}

func (p *Parser) parseDoInstruction() {
	fmt.Printf("--- Parse do instruction ---\nValue: %s\n", p.s[p.idx:p.idx+4])
	if p.idx < p.stringLen-3 && p.s[p.idx:p.idx+4] == "do()" {
		fmt.Printf("Present\n")
		p.mulEnabled = true
	} else {
		fmt.Printf("Absent\n")
	}
}

func (p *Parser) parseDontInstruction() bool {
	fmt.Printf("---Parse dont instruction ---\nValue: %s\n", p.s[p.idx:p.idx+7])
	if p.idx < p.stringLen-6 && p.s[p.idx:p.idx+7] == "don't()" {
		fmt.Printf("Present\n")
		p.mulEnabled = false

		return true
	} else {
		fmt.Printf("Absent\n")

		return false
	}
}

func (p *Parser) matchNextInstruction() int {
	fmt.Printf("--- Match next Instruction ---\n")
	for p.idx < p.stringLen-1 {
		p.idx++
		fmt.Printf("idx: %d\n", p.idx)
		if p.s[p.idx] == 'm' && p.mulEnabled {
			fmt.Printf("Mul match at index: %d\n", p.idx)
			r := p.parseMulInstruction()
			return r
		}
		if p.s[p.idx] == 'd' {
			fmt.Printf("Do/don't instruction match\n")
			r := p.parseDontInstruction()
			if r {
				continue
			}
			p.parseDoInstruction()
		}
	}

	return 0
}

func (p *Parser) parse() int {
	result := 0
	for p.idx < p.stringLen-1 {
		result += p.matchNextInstruction()
	}

	return result
}

func main() {
	data, err := os.ReadFile("input-long.txt")
	if err != nil {
		fmt.Printf("=== ERROR WHILE READING FILE ===\n%s\n", err)
	}

	strdata := string(data)

	parser := Parser{
		s:          strdata,
		idx:        -1,
		stringLen:  len(strdata),
		mulEnabled: true,
	}

	result := parser.parse()
	fmt.Printf("Final result: %d\n", result)
}

// 101969991 --> Too low
