package orm

import (
	g_ "aimi-landing-back-go/domain/repository/gorm"

	"github.com/jinzhu/gorm"

	"aimi-landing-back-go/domain/entity"
)

// db, _ := gorm.NewDb()

type SocietyOrm struct {
	db *gorm.DB
}

func (s *SocietyOrm) NewsocietyOrm() {
	g := g_.Gorm{}
	s.db = g.NewDb()[0]

}

func (s *SocietyOrm) GetAll() ([]*entity.APISocietyListAll, error) {

	societies := make([]*entity.APISocietyListAll, 0)

	fields := `
		society.id, society.abstract, society.country, 
		society.department, society.longitude, society.latitude, 
		society.web, society.video, society_auth.name as person, society_auth.photo`

	rows, _ := s.db.Table("society_auth").Select(fields).Joins("left join society on society_auth.society_id = society.id").Rows()

	for rows.Next() {
		society := new(entity.APISocietyListAll)
		if err := rows.Scan(
			&society.ID, &society.Abstract, &society.Country,
			&society.Department, &society.Longitude, &society.Latitude,
			&society.Web, &society.Video, &society.Person,
			&society.Photo,
		); err != nil {
			panic(err)
		}
		societies = append(societies, society)
	}
	return societies, nil
}

// photo:,
// person:,
