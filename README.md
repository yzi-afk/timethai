# timethai

Golang library for Thai [time.Time](https://pkg.go.dev/time) format

<!-- Add youtube embedded that auto on github -->
[![timethai](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT8o61O5zrmu3CIWa22zhvcyVU9gd30fecqMaOSFNw-&s)](https://www.youtube.com/watch?v=Uj4AWaR7gjU)

<!-- Add example -->
## Example

```go
package main

import (
  "fmt"
  "time"

  "github.com/yzi-afk/timethai"
)

func main() {
  now := time.Now()
  fmt.Println(timethai.Format(now, timethai.DateLayout)) // 18 พฤศจิกายน 2566
  fmt.Println(timethai.Format(now, timethai..ThaiYearLayout)) // 2566
  fmt.Println(timethai.Format(now, timethai.LongWeekDayLayout)) // อาทิตย์
  fmt.Println(timethai.Format(now, timethai.ShortWeekDayLayout)) // อา.
  fmt.Println(timethai.Format(now, timethai.LongMonthLayout)) // พฤศจิกายน
  fmt.Println(timethai.Format(now, timethai.ShortMonthLayout)) // พ.ย.
}
```
