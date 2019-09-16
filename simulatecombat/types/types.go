package types

type Robot interface {
	Attach()
	Defend()
	Stay()
}

type Action struct {
	Frequency 					int
	Damage						int
	Attach_Positive_Factor		int
	Defend_Positive_Factor 		int
	Stay_Positive_Factor		int
	Attach_Negative_Factor		int
	Defend_Negative_Factor		int
	Stay_Negative_Factor		int

}

type RobotA struct {
	Health 			int
	Reward_Factor	float32
	Punish_Factor	float32
	Action
}

type RobotB struct {
	Health 			int
	Reward_Factor	float32
	Punish_Factor	float32
	Action
}