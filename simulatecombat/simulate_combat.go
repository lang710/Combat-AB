/*
 * 当前设计的是一个A-B对抗模型
 * 包含三种动作：包括Attach、Defend、Stay三种动作
 * 主要目的是使双方尽可能先消灭对手

 * 延伸：
 * 设计A-B石头剪子布对抗模型(同步模型)
 * 设计围棋对抗模型(异步模型)
 */


package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lang710/Combat-AB/simulatecombat/types"
)

func main() {
	robotA := types.RobotA{
		Health: 		80,
		Action:			types.Action{
			Frequency: 				2,
			Damage: 				1,
			Attach_Positive_Factor:	100,
			Attach_Negative_Factor: 100,
			Defend_Positive_Factor: 100,
			Defend_Negative_Factor: 100,
			Stay_Positive_Factor: 	100,
			Stay_Negative_Factor: 	100,
		},
		Reward_Factor:	5,
		Punish_Factor: 	-3,
	}
	robotB := types.RobotB{
		Health: 		100,
		Action: 		types.Action{
			Frequency:				1,
			Damage: 				2,
			Attach_Positive_Factor: 100,
			Attach_Negative_Factor: 100,
			Defend_Positive_Factor: 100,
			Defend_Negative_Factor: 100,
			Stay_Positive_Factor: 	100,
			Stay_Negative_Factor: 	100,
		},
		Reward_Factor:	2,
		Punish_Factor:	-5,
	}
	for true {
		if robotA.Health <= 0 || robotB.Health <= 0 {
			fmt.Println("Game Over!")
			break
		}

		rand.Seed(time.Now().Unix())

		for i:=0; i<robotA.Frequency; i++ {
			var robotA_action_type int
			robotA_action_fator_sum := robotA.Attach_Positive_Factor + robotA.Defend_Positive_Factor + robotA.Stay_Positive_Factor
			robotA_choose := rand.Intn(robotA_action_fator_sum)
			if robotA_choose <= robotA.Attach_Positive_Factor {
				robotA_action_type = 1
			} else if robotA_choose <= robotA.Attach_Positive_Factor + robotA.Defend_Positive_Factor {
				robotA_action_type = 2
			} else {
				robotA_action_type = 3
			}

			var robotB_action_type int
			robotB_action_fator_sum := robotB.Attach_Negative_Factor + robotB.Defend_Negative_Factor + robotB.Stay_Negative_Factor
			robotB_choose := rand.Intn(robotB_action_fator_sum)
			if robotB_choose <= robotB.Attach_Negative_Factor {
				robotB_action_type = 1
			} else if robotB_choose <= robotB.Attach_Negative_Factor + robotB.Defend_Negative_Factor {
				robotB_action_type = 2
			} else {
				robotB_action_type = 3
			}

			if robotA_action_type == 1 && robotB_action_type == 1 {
				robotB.Health -= robotA.Damage
				// robotB.Attach_Negative_Factor -=
			}
		}
	}
}
