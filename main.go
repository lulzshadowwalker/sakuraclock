package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type placeholder [5]string

func main() {
	zero := placeholder{
		"🌸🌸🌸",
		"🌸  🌸",
		"🌸  🌸",
		"🌸  🌸",
		"🌸🌸🌸",
	}

	one := placeholder{
		"🌸🌸 ",
		"  🌸 ",
		"  🌸 ",
		"  🌸 ",
		"🌸🌸🌸",
	}

	two := placeholder{
		"🌸🌸🌸",
		"    🌸",
		"🌸🌸🌸",
		"🌸   ",
		"🌸🌸🌸",
	}

	three := placeholder{
		"🌸🌸🌸",
		"    🌸",
		"🌸🌸🌸",
		"    🌸",
		"🌸🌸🌸",
	}

	four := placeholder{
		"🌸  🌸",
		"🌸  🌸",
		"🌸🌸🌸",
		"    🌸",
		"    🌸",
	}

	five := placeholder{
		"🌸🌸🌸",
		"🌸   ",
		"🌸🌸🌸",
		"    🌸",
		"🌸🌸🌸",
	}

	six := placeholder{
		"🌸🌸🌸",
		"🌸   ",
		"🌸🌸🌸",
		"🌸  🌸",
		"🌸🌸🌸",
	}

	seven := placeholder{
		"🌸🌸🌸",
		"    🌸",
		"    🌸",
		"    🌸",
		"    🌸",
	}

	eight := placeholder{
		"🌸🌸🌸",
		"🌸  🌸",
		"🌸🌸🌸",
		"🌸  🌸",
		"🌸🌸🌸",
	}

	nine := placeholder{
		"🌸🌸🌸",
		"🌸  🌸",
		"🌸🌸🌸",
		"    🌸",
		"🌸🌸🌸",
	}

	sep := placeholder{
		"",
		"🌺",
		"",
		"🌺",
		"",
	}

	digs := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hr, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digs[hr/10], digs[hr%10],
			sep,
			digs[min/10], digs[min%10],
			sep,
			digs[sec/10], digs[sec%10],
		}

		for line := range clock[0] {
			for i, dig := range clock {
				next := clock[i][line]
				if dig == sep && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "\t")

			}

			fmt.Println()

		}
		fmt.Println()

		time.Sleep(time.Second)
	}
}
