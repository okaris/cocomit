package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"cocomit/processor"
)

func main() {
	totalOnly := flag.Bool("total", false, "show only the total LoC and cost")
	flag.Parse()

	// ANSI color codes
	colorReset := "\033[0m"
	colorCyan := "\033[36m"
	colorYellow := "\033[33m"
	colorGreen := "\033[32m"
	colorRed := "\033[31m"
	colorMagenta := "\033[35m"

	fmt.Printf("%s#cocomit%s\n", colorCyan, colorReset)
	if !*totalOnly {
		fmt.Println("Assumptions:")
		fmt.Printf("%s- Model:        Single Developer Linear Estimate%s\n", colorYellow, colorReset)
		fmt.Printf("%s- Effort:       Effort = (SLOC / 50) hours → PM = hours / 152%s\n", colorYellow, colorReset)
		fmt.Printf("%s- EAF:          1.0 (nominal)%s\n", colorYellow, colorReset)
		fmt.Printf("%s- Hourly Wage:  $65.79 ($120K/year)%s\n", colorYellow, colorReset)
		fmt.Printf("%s- Overhead:     1.3 (benefits, infra, etc.)%s\n", colorYellow, colorReset)
		fmt.Printf("%s- Cost:         Effort × Hours × Hourly Wage × Overhead%s\n", colorYellow, colorReset)
		fmt.Println()
	}

	repo := "."

	pageSize := 50
	eaf := 1.0
	hourlyWage := int64(120000) / 160 / 12
	overhead := 1.3

	var totalLoc int64
	var totalCost float64

	for i := 0; ; i++ {
		skip := i * pageSize

		entries := gitLog(repo, skip, pageSize)
		if len(entries) == 0 {
			break
		}

		for _, entry := range entries {
			fields := strings.Split(entry, "|")
			if len(fields) != 4 {
				fmt.Printf("%sInvalid entry: %s%s\n", colorRed, entry, colorReset)
				continue
			}
			hash := fields[0]
			// author := fields[1]
			date := fields[2]
			locDelta, err := strconv.Atoi(fields[3])
			if err != nil {
				fmt.Printf("%sInvalid LOC delta: %s%s\n", colorRed, fields[3], colorReset)
				continue
			}
			if locDelta == 0 {
				continue
			}
			cost := processor.EstimateCost(processor.EstimateEffort(int64(locDelta), eaf), float64(hourlyWage), overhead)
			if !*totalOnly {
				fmt.Printf("%s%s%s | %s%s%s | %s$%.2f%s\n", colorGreen, hash[0:7], colorReset, colorCyan, date, colorReset, colorMagenta, cost, colorReset)
			}

			totalLoc += int64(locDelta)
			totalCost += cost
		}
	}

	if !*totalOnly {
		fmt.Println("\n" + strings.Repeat("─", 40))
	}
	fmt.Printf("%sTotal LoC:%s  %d\n", colorCyan, colorReset, totalLoc)
	fmt.Printf("%sTotal Cost:%s %s$%.2f%s\n", colorCyan, colorReset, colorMagenta, totalCost, colorReset)
}

func gitLog(repoPath string, skip, limit int) []string {
	cmd := exec.Command("git", "-C", repoPath,
		"-c", "core.pager=cat",
		"log",
		fmt.Sprintf("--skip=%d", skip),
		fmt.Sprintf("-n%d", limit),
		"--pretty=format:@@@%H|%an|%s",
		"--shortstat")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(out), "\n")
	var result []string
	var hash, author, date string

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "@@@") {
			parts := strings.SplitN(strings.TrimPrefix(line, "@@@"), "|", 3)
			if len(parts) != 3 {
				continue
			}
			hash, author, date = parts[0], parts[1], parts[2]

			if i+1 < len(lines) && strings.Contains(lines[i+1], "changed") {
				added, removed := parseLocChange(lines[i+1])
				loc := added + removed
				result = append(result, fmt.Sprintf("%s|%s|%s|%d", hash[0:7], author, date, loc))
				i++ // skip shortstat line
			}
		}
	}
	return result
}

func parseLocChange(line string) (int, int) {
	added := 0
	removed := 0
	fields := strings.Split(line, ",")
	for _, f := range fields {
		f = strings.TrimSpace(f)
		if strings.Contains(f, "insertion") {
			n, err := strconv.Atoi(strings.Fields(f)[0])
			if err == nil {
				added = n
			}
		} else if strings.Contains(f, "deletion") {
			n, err := strconv.Atoi(strings.Fields(f)[0])
			if err == nil {
				removed = n
			}
		}
	}
	return added, removed
}
