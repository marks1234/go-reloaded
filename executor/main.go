package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"reload"
)

// simple error catcher
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Command struct {
	name   string
	amount int
}

func DefaultCommand(name string) Command {
	return Command{
		name:   name,
		amount: 1,
	}
}

// takes a (cap, 2) or anything of the sort and then places it into a struct with name and amount fields
func extract_Command(s string) (Command, bool) {
	command_list := map[string]int{
		"hex": 1,
		"bin": 1,
		"up":  1,
		"low": 1,
		"cap": 1,
	}

	store_name := ""
	store_number := ""

	for _, str := range s {
		if str < 97 || str > 122 {
			if str > 47 && str < 58 {
				store_number += string(str)
			}
			continue
		}
		store_name += string(str)
	}
	if command_list[store_name] == 1 {
		return_com := DefaultCommand(store_name)
		if store_number != "" {
			num := reload.Atoi(store_number)
			return_com.amount = num
		}
		return return_com, true
	}

	return DefaultCommand("none"), false
}

// splits the whole text and operates the commands it has
func commandFulfill(text string) string {
	var brackets int
	var start int
	var result []string
	text += " x"

	for i, r := range text {
		switch r {
		case '(':
			brackets++
		case ')':
			brackets--
		case ' ':
			if brackets == 0 {
				result = append(result, strings.TrimSpace(text[start:i]))
				start = i + 1
			}
		}
	}

	var was_com int
	for i, str := range result {
		com, isCom := extract_Command(str)
		if !isCom {
			if was_com != 0 {
				was_com--
			}
			continue
		}

		offset := i - was_com
		switch com.name {
		case "hex":
			result[i-1] = strconv.Itoa(reload.AtoiBase(result[i-1], "0123456789ABCDEF"))
		case "bin":
			result[i-1] = strconv.Itoa(reload.AtoiBase(result[i-1], "01"))
		case "up":
			for index := com.amount; index > 0; index-- {
				result[offset-index] = strings.ToUpper(result[offset-index])
			}
		case "low":
			for index := com.amount; index > 0; index-- {
				result[offset-index] = strings.ToLower(result[offset-index])
			}
		case "cap":
			for index := com.amount; index > 0; index-- {
				result[offset-index] = strings.Title(result[offset-index])
			}
		}

		was_com++
	}

	return strings.Join(result, " ")
}

func main() {
	files := os.Args[1:]

	// quick check on the amount of arguments
	if len(files) > 2 || len(files) == 0 {
		println("Not the right amount of arguments!")
		return
	}

	dat, err := os.ReadFile(files[0])
	str_dat := string(dat)
	check(err)

	file, err := os.Create(files[1])
	check(err)
	defer file.Close()

	// compiles the search for later commands
	re := regexp.MustCompile("\\((bin|hex|cap|up|low),? ?\\d?\\)")
	text_without_commands := commandFulfill(str_dat)
	// replaces all the commands in between () with a blank space
	text_without_commands = re.ReplaceAllString(text_without_commands, "")

	// removes the white spaces before the punctuations
	re = regexp.MustCompile(` +[,.:;!?]`)
	text_signs_correct := re.ReplaceAllStringFunc(string(text_without_commands), strings.TrimSpace)

	// adds space after specific punctuations
	re = regexp.MustCompile(`[,.:;!?][\w\d]`)
	text_signs_correct = re.ReplaceAllStringFunc(string(text_signs_correct), func(s string) string {
		s_arr := strings.Split(s, "")
		return s_arr[0] + " " + s_arr[1]
	})

	// removes white space between parantheses
	re = regexp.MustCompile(`'(.*?)'`)
	apostrophe_wrap := re.ReplaceAllStringFunc(text_signs_correct, func(s string) string {
		clean := strings.TrimSpace(s[1:])
		clean = strings.TrimSpace(clean[:len(clean)-1])
		return "'" + clean + "'"
	})

	// A to An | a to an
	re = regexp.MustCompile(`[aA] [aeiouh]`)
	a_to_an := re.ReplaceAllStringFunc(apostrophe_wrap, func(s string) string {
		s_arr := strings.Split(s, "")
		return s_arr[0] + "n" + strings.Join(s_arr[1:], "")
	})

	// removes white space
	space := regexp.MustCompile(` +`)
	s := space.ReplaceAllString(a_to_an, " ")
	file.WriteString(s)
}
