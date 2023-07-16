package dp

import "sort"

// TODO solve with another approach f(skillmask) return min people needs with skillmask and use that.
func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	// return frecsol(req_skills, people)
	return frecsolbitsol(req_skills, people)
}

func frecsol(skills []string, peps [][]string) []int {
	minsize = 100
	req := make(map[string]bool)
	for _, v := range skills {
		req[v] = true
	}
	currskills := make(map[string]bool)
	frecminpep(req, peps, nil, 0, currskills)
	return minpeps
}

var (
	minsize int
	minpeps []int
)

func frecminpep(req map[string]bool, peps [][]string, selpeps []int, curridx int, currskill map[string]bool) {
	// check if currskill has all from req
	if compare(currskill, req) {
		if len(selpeps) < minsize {
			minsize = len(selpeps)
			minpeps = selpeps
		}
		return
	}

	// base case
	if curridx >= len(peps) {
		return
	}

	// choose curridx
	newskills := copymap(currskill)
	// add skills to map
	for _, skill := range peps[curridx] {
		newskills[skill] = true
	}

	newpeps := make([]int, len(selpeps))
	copy(newpeps, selpeps)
	// append to select peps
	newpeps = append(newpeps, curridx)
	frecminpep(req, peps, newpeps, curridx+1, newskills)
	// remove from select peps

	// not choose curridx
	frecminpep(req, peps, selpeps, curridx+1, currskill)
	return
}

func copymap(act map[string]bool) map[string]bool {
	copymp := make(map[string]bool)
	for k, v := range act {
		copymp[k] = v
	}
	return copymp
}

func compare(act, exp map[string]bool) bool {
	for k := range exp {
		if _, ok := act[k]; !ok {
			return false
		}
	}
	return true
}

func countSetBits(n uint64) int {
	c := 0
	for ; n > 0; c++ {
		n = n & (n - 1) // clear the least significant bit set
	}
	return c
}

func frecsolbitsol(skills []string, people [][]string) []int {
	sort.Strings(skills)
	skillidxmap := make(map[string]int)
	reqskill := 0
	for k, v := range skills {
		skillidxmap[v] = k
		reqskill = reqskill | 1<<k
	}

	// fmt.Printf("reqskils are %b\n", reqskill)

	peps := make([]int, len(people))

	for idx, skills := range people {
		pepskill := 0
		for _, skill := range skills {
			idx := skillidxmap[skill]
			pepskill = pepskill | 1<<idx
		}
		peps[idx] = pepskill
		// fmt.Printf("person %d skills are %b\n", idx, pepskill)
	}

	// fmt.Printf("minpeps bit: %d, %b\n", minpepsbit, minpepsbit)
	minsize = 100
	minpepsbit := frecbit(reqskill, peps, 0, 0, make(map[pair1125]uint64))

	// fmt.Printf("after run, selected: %b\n", minpepsbit)
	selected := make([]int, 0)
	for i := 0; i < 63; i++ {
		if minpepsbit&(1<<i) > 0 {
			selected = append(selected, i)
		}
	}
	return selected
}

var maxset uint64 = (1 << 63) - 1

type pair1125 struct {
	a int
	b int
}

func frecbit(reqskill int, peps []int, currskill, curridx int, dp map[pair1125]uint64) uint64 {
	// check if currskill has all from req
	if reqskill == currskill {
		return 0
	}
	// base case
	if curridx >= len(peps) {
		return maxset
	}
	currpair := pair1125{currskill, curridx}
	if val, ok := dp[currpair]; ok {
		return val
	}

	// select this person
	choose := frecbit(reqskill, peps, currskill|peps[curridx], curridx+1, dp)
	choose = choose | 1<<curridx

	// not choose curridx
	notchoose := frecbit(reqskill, peps, currskill, curridx+1, dp)

	// return min needed.
	chbits := countSetBits(choose)
	notchbits := countSetBits(notchoose)
	if chbits < notchbits {
		dp[currpair] = choose
	} else {
		dp[currpair] = notchoose
	}
	return dp[currpair]
}
