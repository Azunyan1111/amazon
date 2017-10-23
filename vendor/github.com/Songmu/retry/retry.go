package retry

import "time"
import "fmt"

// Retry calls the `fn` and if it returns the error, retry to call `fn` after `interval` duration.
// The `fn` is called up to `n` times.
func Retry(n int, interval int, fn func() error) (err error) {
	for n > 0 {
		n--
		err = fn()
		if err == nil || n <= 0 {
			break
		}
		fmt.Println("Retry")
		time.Sleep(time.Duration(interval * n))
	}
	return err
}
