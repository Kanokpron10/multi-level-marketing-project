package repository

import (
	"database/sql"
	"multi-level-marketing-project/models"
)

func GetMyPoint(db *sql.DB, memberID int) int {
	statement, err := db.Prepare(`SELECT SUM(point) FROM point_records WHERE member_id = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID)
	var myPoint int
	row.Scan(&myPoint)
	return myPoint
}

func GetMonthlyPoint(db *sql.DB, memberID int, month int, year int) int {
	statement, err := db.Prepare(`SELECT SUM(point) FROM point_records WHERE member_id = ? AND MONTH(create_date) = ? AND YEAR(create_date) = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID, month, year)
	var monthlyPoint int
	row.Scan(&monthlyPoint)
	return monthlyPoint
}

func CountTeamMember(db *sql.DB, memberID int) models.TeamMember {
	var teamMember models.TeamMember
	statement, err := db.Prepare(`SELECT COUNT(id) FROM members WHERE leader_id = ? AND level >= ?;`)
	if err != nil {
		return teamMember
	}
	rowCountHigherPearl := statement.QueryRow(memberID, 1)
	rowCountHigherEmmerald := statement.QueryRow(memberID, 4)
	rowCountHigherRuby := statement.QueryRow(memberID, 7)
	rowCountHigherPearl.Scan(&teamMember.HigherPearl)
	rowCountHigherEmmerald.Scan(&teamMember.HigherEmerald)
	rowCountHigherRuby.Scan(&teamMember.HigherRuby)
	return teamMember
}

func GetTeamPoint(db *sql.DB, memberID int) int {
	statement, err := db.Prepare(`SELECT SUM(point) FROM point_records INNER JOIN members ON members.id = point_records.member_id WHERE members.leader_id = ? OR members.id = ?`)
	if err != nil {
		return 0
	}
	row := statement.QueryRow(memberID, memberID)
	var teamPoint int
	row.Scan(&teamPoint)
	return teamPoint
}

func GetMemberData(db *sql.DB, memberID int) models.Member {
	var member models.Member
	statement, err := db.Prepare(`SELECT * FROM members WHERE id = ?`)
	if err != nil {
		return member
	}
	row := statement.QueryRow(memberID)
	err = row.Scan(
		&member.MemberID,
		&member.MemberName,
		&member.LeaderID,
		&member.Level,
	)
	if err != nil {
		panic(err.Error())
	}
	return member
}

func Promote(db *sql.DB, member models.Member) bool {
	statement, err := db.Prepare(`UPDATE members SET level = ? WHERE id = ?`)
	if err != nil {
		return false
	}
	_, err = statement.Exec(member.Level+1, member.MemberID)
	if err != nil {
		return false
	}
	return true
}

func Demote(db *sql.DB, member models.Member) bool {
	statement, err := db.Prepare(`UPDATE members SET level = ? WHERE id = ?`)
	if err != nil {
		return false
	}
	_, err = statement.Exec(member.Level-1, member.MemberID)
	if err != nil {
		return false
	}
	return true
}
