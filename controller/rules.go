package controller

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

)

func AddRule(rawRule string) {
	rule := parseRule(rawRule)

	Conf.Rule = append(Conf.Rule, rule)

	y := marshal(&Conf)
	writeToYaml(y)
}

func parseRule(rawRule string) (r string) {
	/*
	check wether the input rule is valid and return a formatted string.
	"rawRule" is basically made up by three parts: prefix, domain and adapter. Since the 
	prefix is of uppercase and seems a little bit long, this function could handle
	a shorthand or lowercase input of prefix such as "suffix".
	*/

	rule := strings.Split(rawRule, ",")
	if len(rule) != 3 {
		log.Fatal("invalid rule format")
	}

	// check whether the prefix is supported and reassign it with the right format.
	prefix := strings.ToLower(rule[0])
	switch prefix {
		case "domain", "ip-cidr", "geoip":
			prefix = strings.ToUpper(prefix)
		case "keyword", "suffix":
			prefix = "DOMAIN-" + strings.ToUpper(prefix)
		default:
			log.Fatalf("invalid rule prefix: %s", prefix)
	}

	r = strings.Join([]string{prefix, rule[1], rule[2]}, ",")

	return
}

func DeleteRule(rawRule string) {
	/*
	Two formats of rawRule are supported: a standard clash rule or just a 
	domain.
	This function will do a search first. If there is only one matched
	rule, it will deleted directly. But if there exist more than one rules
	contain the domain, they will be listed as candidates and wait for the user's
	choice.
	*/

	var result []int
	result = append(result, search(rawRule, Conf)...)

	switch len(result) {
	case 0:
		log.Fatal("Cannot delete a nonexistent rule...")
	case 1:
		i := result[0]
		Conf.Rule = append(Conf.Rule[:i], Conf.Rule[i+1: ]...)
	default:
		choices := chooseToDelete(Conf, result)
		for i, m := range choices {
			//todo: swap back to the original order.
			Conf.Rule[m] = Conf.Rule[len(Conf.Rule)-i-1]
		}
		Conf.Rule = Conf.Rule[:len(Conf.Rule)-len(choices)]
	}

	y := marshal(&Conf)
	writeToYaml(y)
}

func chooseToDelete(conf *yamlConfig, result []int) []int {
	fmt.Printf("Found %d rules, which do you want to delete?\n", len(result))
	for i, id := range result {
		fmt.Printf("%d. %s\n", i+1, conf.Rule[id])
	}
	fmt.Printf("%d. delete them all.\n", len(result)+1)

	// you can select several rules at one time.
	var choiceInput string
	var choice []int 
	fmt.Printf("=> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choiceInput = scanner.Text()

	// Since vscode doesn't support read STDIN form debug console, the assignment
	// below is just for test
	// choiceInput = "2"

	if choiceInput == "" {
		log.Fatal("No choice was made")
	}
	s := strings.Split(choiceInput, " ")
	for _, str := range s {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("input error: %s\n", err)
		}
		choice = append(choice, num)
	}

	/*note: "choice" is the list of serial numbers, don't mistake it for the 
	"result" indexes.*/
	/*let's check whether the choice is valid or contains the "all" option.*/
	var fResult []int
	for _, v := range choice {
		if v > len(result) + 1 || v < 1 {
			log.Fatalf("selection of %d goes beyond the options", v)
		}
		if v == len(result) + 1 {
			return result
		}
		fResult = append(fResult, result[v-1])
	}	
	return fResult
}

func search(rule string, conf *yamlConfig) (indexes []int) {
	/*
	return the indexes of the matched rules.
	*/

	for i, v := range conf.Rule {
		/*
		todo: use reg to search.
		use strings.Index may return extra rules. e.g. "xxx.com.cn" will be found 
		if you search "xxx.com" 
		*/
		in := strings.Index(v, rule)
		if in != -1 {
			indexes = append(indexes, i)
		}
	}	
	if len(indexes) == 0  {
		log.Fatal("rule not found")
	}

	return
}

func SearchDomain(domain string) {
	matched := search(domain, Conf)
	if len(matched) == 0 {
		log.Fatal("no such a domain")
	}


	var s string = ""
	if len(matched) > 1 {
		s = "s"
	}

	fmt.Printf("found %d rule%s\n", len(matched), s)
	for _, v := range matched {
		fmt.Println(Conf.Rule[v])
	}
}

