package hamming

import (
    "errors" 
    )

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("length is not same")
    }

    diff := 0
    for i, v := range a {
        if v != rune(b[i]) {
            diff++
        }
    }
    return diff, nil
}
