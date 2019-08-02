package model

type Member struct {
	MemberID     int    `json:"member_id"`
	MemberName   string `json:"member_name"`
	LeaderID     int    `json:"leader_id"`
	Level        int    `json:"user_referral"`
	MyPoint      int    `json:"my_point"`
	MonthlyPoint int    `json:"monthly_point"`
	TeamPoint    int    `json:"team_point"`
	TeamMember   TeamMember
}

type TeamMember struct {
	HigherPearl   int
	HigherEmerald int
	HigherRuby    int
}
