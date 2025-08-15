package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("input must be a positive integer")
    }

	res, step := n, 0
   
    for res != 1 {
        if res % 2 == 0 {
            res = res / 2
        } else {
            res = res * 3 + 1
        }
        step++
    }
    return step, nil
}
