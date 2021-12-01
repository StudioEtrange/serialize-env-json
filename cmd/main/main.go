package main

import (
	"encoding/json"
	"flag"
	"os"
	"regexp"
	"strings"
)

type envs struct {
	Envs []env `json:"envs"`
}

type env struct {
	Fullname string `json:"fullname"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

// ReplaceAllSubmatchFunc and ReplaceAllStringSubmatchFunc are variant of the same func for []byte and string
// They will replace all matching groups with the repl string
func ReplaceAllStringSubmatch(re *regexp.Regexp, src string, repl string, n int) string {
	return ReplaceAllStringSubmatchFunc(re, src, func(groups []string) []string {
		for i, _ := range groups {
			groups[i] = repl
		}
		return groups
	}, n)
}
func ReplaceAllSubmatch(re *regexp.Regexp, src []byte, repl []byte, n int) []byte {
	return ReplaceAllSubmatchFunc(re, src, func(groups [][]byte) [][]byte {
		for i, _ := range groups {
			groups[i] = repl
		}

		return groups
	}, n)
}

// ReplaceAllSubmatchFunc and ReplaceAllStringSubmatchFunc are variant of the same func for []byte and string
// They will replace all matching groups with the modified groups value by the callback func
// In the callback function groups index start at 0
//
// 	pattern := regexp.MustCompile(...)
// 	data = ReplaceAllSubmatchFunc(pattern, data, func(groups [][]byte) [][]byte {
// 		// mutate groups here
//		groups[0]=[]byte("REPLACED")
// 		return groups
// 	}, -1)
//
// 	data = ReplaceAllStringSubmatchFunc(pattern, data, func(groups []string) []string {
// 		// mutate groups here
//		groups[0]="REPLACED"
// 		return groups
// 	}, -1)
//
// If n >= 0, the function returns at most n matches/submatches; otherwise, it returns all of them.
// This snippet is MIT licensed. Please cite by leaving this comment in place. Find
// the original version is at:
//
//  https://gist.github.com/slimsag/14c66b88633bd52b7fa710349e4c6749
//
func ReplaceAllSubmatchFunc(re *regexp.Regexp, src []byte, repl func([][]byte) [][]byte, n int) []byte {
	var (
		result  = make([]byte, 0, len(src))
		matches = re.FindAllSubmatchIndex(src, n)
		last    = 0
	)
	for _, match := range matches {
		// Append bytes between our last match and this one (i.e. non-matched bytes).
		matchStart := match[0]
		matchEnd := match[1]
		result = append(result, src[last:matchStart]...)
		last = matchEnd

		// Determine the groups / submatch bytes and indices.
		groups := [][]byte{}
		groupIndices := [][2]int{}
		for i := 2; i < len(match); i += 2 {
			start := match[i]
			end := match[i+1]
			groups = append(groups, src[start:end])
			groupIndices = append(groupIndices, [2]int{start, end})
		}

		// Replace the groups as desired.
		groups = repl(groups)

		// Append match data.
		lastGroup := matchStart
		for i, newValue := range groups {
			// Append bytes between our last group match and this one (i.e. non-group-matched bytes)
			groupStart := groupIndices[i][0]
			groupEnd := groupIndices[i][1]
			result = append(result, src[lastGroup:groupStart]...)
			lastGroup = groupEnd

			// Append the new group value.
			result = append(result, newValue...)
		}
		result = append(result, src[lastGroup:matchEnd]...) // remaining
	}
	result = append(result, src[last:]...) // remaining
	return result
}
func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, src string, repl func([]string) []string, n int) string {
	return string(ReplaceAllSubmatchFunc(re, []byte(src), func(groups [][]byte) [][]byte {
		//https://stackoverflow.com/a/12829631
		groupsStr := make([]string, len(groups))
		for i, v := range groups {
			groupsStr[i] = string(v)
		}
		resultTmp := repl(groupsStr)
		result := make([][]byte, len(resultTmp))
		for i, v := range resultTmp {
			result[i] = []byte(v)
		}
		return result
	}, n))
}



func main() {
	filter := flag.String("filter", ".*", "Regexp to select env var")
	forceLower := flag.Bool("lower", false, "Force env var name to lower")
	forceUpper := flag.Bool("upper", false, "Force env var name to upper")
	removeFilter := flag.Bool("clean", false, "Remove from env var name all the matching regexp group")
	flag.Parse()

	var resultEnvs envs
	envVars := os.Environ()

	regx := regexp.MustCompile(*filter)

	for _, envVar := range envVars {
		var name string
		v := strings.SplitN(envVar, "=", 2)
		if regx.MatchString(v[0]) {
			if *removeFilter {
				name = ReplaceAllStringSubmatch(regx, v[0], "", -1)
			} else {
				name = v[0]
			}
			if *forceLower {
				name = strings.ToLower(name)
			} else {
				if *forceUpper {
					name = strings.ToUpper(name)
				}
			}
			resultEnvs.Envs = append(resultEnvs.Envs, env{
				Fullname: v[0],
				Name:     name,
				Value:    v[1],
			})
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(resultEnvs)
}
