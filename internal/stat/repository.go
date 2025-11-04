package stat

import (
	"purpleschool-go/advanced/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat

	currentDate := datatypes.Date(time.Now())

	repo.DB.First(&stat, "link_id = ? and date = ?", linkId, currentDate)

	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkId: linkId,
			Date:   currentDate,
			Clicks: 1,
		})
	} else {
		stat.Clicks++
		repo.DB.Save(&stat)
	}
}

func (repo *StatRepository) GetAll(from, to time.Time, by string) []GetAllStatResponse {
	var stats []GetAllStatResponse
	var period string

	if by == GroupByMonth {
		period = "YYYY-MM"
	} else {
		period = "YYYY-MM-DD"
	}

	repo.Db.DB.
		Table("stats").
		Select("TO_CHAR(date, ?) AS period, sum(clicks)", period).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}
