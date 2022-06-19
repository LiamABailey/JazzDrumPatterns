package patterngen

import ( 
	"fmt"
	"sort"
	"strings"
	"strconv"
)

const (
	beatIdentifier string = "b"
	eq string = "="
)

// split ?b0=ride,snare,kick,hihat&b1=... queries into 
// component strings
func splitBeatQuery(query map[string][]string) ([]string, error) {
	var unsupportedIds []string
	//for each segment, we load into a map[int]string for future ordering
	parsedBeats := make(map[int]string)
	for qkey, qval := range query {
		// try to parse id of beat 
		qkeyId := qkey[1:]
		beatId, atoierr := strconv.Atoi(qkeyId)
		if atoierr == nil {
			if len(query[qkey]) != 1 {
				return make([]string, 0), fmt.Errorf("Expect one value per param, saw %s for %s", qval, qkey)
			}
			parsedBeats[beatId] = qval[0]
		} else {
			unsupportedIds = append(unsupportedIds, qkeyId)
		}
	}
	// if unsupported ids, stop and exit
	if len(unsupportedIds) != 0 {
		return make([]string, 0), fmt.Errorf("Beat IDs not supported: %s", unsupportedIds)
	}
	// else, sort keys
	var mKeys []int
	for k := range parsedBeats {
		mKeys = append(mKeys, k)
	}
	sort.Ints(mKeys)
	// sort the beat values using the key order, from least to greatest
	var sortedBeats []string
	for sk := range mKeys {
		sortedBeats = append(sortedBeats, parsedBeats[sk])
	}	
	return sortedBeats, nil
}

// parse a string 
func splitConvertPatternLists(pat string, sep string) ([]int, error) {
	sepPat := strings.Split(pat, sep)
	iPat := make([]int, len(sepPat))
	var cerr error
	for i, v := range sepPat {
	  iPat[i], cerr = strconv.Atoi(v)
	  if cerr != nil {
		return nil, cerr
	  }
	}
	return iPat, nil
  }