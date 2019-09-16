/* A: attack、Stay
 * B: defend、Stay
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var health_a, health_b, damage_a, damage_b = 100, 50, 5, 5
	var attach_a, stay_a = 50, 50
	var action_a, action_b int
	var b2a_attach_defend, b2a_attach_stay, b2a_stay_defend, b2a_stay_stay = 50, 50, 50, 50

	for true {
		rand.Seed(time.Now().Unix())
		sum_a := attach_a + stay_a
		choose_a := rand.Intn(sum_a)

		if choose_a <= attach_a{
			action_a = 0
		} else {
			action_a = 1
		}

		rand.Seed(time.Now().Unix())
		if action_a == 0 {
			sum_b := b2a_attach_defend + b2a_attach_stay
			choose_b := rand.Intn(sum_b)

			if choose_b <= b2a_attach_defend {
				action_b = 0
			} else {
				action_b = 1
			}
		} else {
			sum_b := b2a_stay_defend + b2a_stay_stay
			choose_b := rand.Intn(sum_b)

			if choose_b <= b2a_stay_defend {
				action_b = 0
			} else {
				action_b = 1
			}
		}

		if action_a == 0 && action_b == 0 {
			health_a -= damage_b
			attach_a = int(0.8 * float64(action_a))
			b2a_attach_defend = int(1.2 * float64(b2a_attach_defend))
		} else if action_a == 0 && action_b == 1 {
			health_b -= damage_a
			attach_a = int(1.2 * float64(action_a))
			b2a_attach_stay = int(0.8 * float64(b2a_attach_stay))
		} else if action_a == 1 && action_b == 0 {
			health_b -= damage_a
			stay_a = int(1.2 * float64(stay_a))
			b2a_stay_defend = int(0.8 * float64(b2a_stay_defend))
		} else {
			health_a -= damage_b
			stay_a = int(0.8 * float64(stay_a))
			b2a_stay_stay = int(1.2 * float64(b2a_stay_stay))
		}

		fmt.Printf("health_a: %d, health_b: %d, action_a: %d, action_b: %d\n",
			health_a, health_b, action_a, action_b)

		if health_a <= 0 || health_b <= 0 {
			fmt.Println("Game Over!")
			break
		}
	}
}
