package model

import (
	"database/sql/driver"
	"encoding/xml"
	"errors"
	"time"
)

type HattrickData struct {
	Players []*Player `xml:"Team>PlayerList>Player"`
	Player  *Player   `xml:"Player"`
	Error   *string   `xml:"Error"`
}

// NO SE OCUPA, PERO NO BORRAR
// Response es una copia temporal del tipo Response en helperHttp y se usa para el swagger
// @name Response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Player struct {
	ID                 uint       `json:"ID" gorm:"primaryKey;unique;autoIncrement" xml:"-"`
	PlayerID           uint       `json:"player_ID" xml:"PlayerID"`
	TeamID             uint       `json:"team_ID" xml:"-"`
	FirstName          string     `json:"first_name" xml:"FirstName"`
	NickName           *string    `json:"nickname" xml:"NickName"`
	LastName           string     `json:"last_name" xml:"LastName"`
	PlayerNumber       *uint8     `json:"player_number" xml:"PlayerNumber"`
	Age                uint16     `json:"age" xml:"Age"`
	AgeDays            uint8      `json:"age_days" xml:"AgeDays"`
	ArrivalDate        CustomTime `json:"arrival_date" xml:"ArrivalDate"`
	OwnerNotes         *string    `json:"owner_notes" xml:"OwnerNotes"`
	TSI                uint32     `json:"tsi" xml:"TSI"`
	PlayerForm         uint8      `json:"player_form" xml:"PlayerForm"`
	Statement          *string    `json:"statement" xml:"Statement"`
	Experience         uint8      `json:"experience" xml:"Experience"`
	Loyalty            uint8      `json:"loyalty" xml:"Loyalty"`
	MotherClubBonus    bool       `json:"mother_club_bonus" xml:"MotherClubBonus"`
	Leadership         uint8      `json:"leadership" xml:"Leadership"`
	Salary             uint32     `json:"salary" xml:"Salary"`
	Agreeability       uint8      `json:"agreeability" xml:"Agreeability"`
	Aggressiveness     uint8      `json:"aggressiveness" xml:"Aggressiveness"`
	Honesty            uint8      `json:"honesty" xml:"Honesty"`
	LeagueGoals        uint8      `json:"league_goals" xml:"LeagueGoals"`
	CupGoals           uint8      `json:"cup_goals" xml:"CupGoals"`
	FriendliesGoals    uint8      `json:"friendlies_goals" xml:"FriendliesGoals"`
	CareerGoals        uint32     `json:"career_goals" xml:"CareerGoals"`
	CareerHattricks    uint8      `json:"career_hattricks" xml:"CareerHattricks"`
	MatchesCurrentTeam uint16     `json:"matches_current_team" xml:"MatchesCurrentTeam"`
	GoalsCurrentTeam   uint32     `json:"goals_current_team" xml:"GoalsCurrentTeam"`
	Specialty          uint8      `json:"specialty" xml:"Specialty"`
	TransferListed     bool       `json:"transfer_listed" xml:"TransferListed"`
	NationalTeamID     *int       `json:"national_team_ID" xml:"NationalTeamID"`
	CountryID          uint       `json:"country_id" xml:"CountryID"`
	Caps               uint16     `json:"caps" xml:"Caps"`
	CapsU20            uint16     `json:"caps_u20" xml:"CapsU20"`
	Cards              uint8      `json:"cards" xml:"Cards"`
	InjuryLevel        int16      `json:"injury_level" xml:"InjuryLevel"`
	StaminaSkill       uint8      `json:"stamina_skill" xml:"StaminaSkill"`
	KeeperSkill        *uint8     `json:"keeper_skill" xml:"KeeperSkill"`
	PlaymakerSkill     *uint8     `json:"playmaker_skill" xml:"PlaymakerSkill"`
	ScorerSkill        *uint8     `json:"scorer_skill" xml:"ScorerSkill"`
	PassingSkill       *uint8     `json:"passing_skill" xml:"PassingSkill"`
	WingerSkill        *uint8     `json:"watching_skill" xml:"WingerSkill"`
	DefenderSkill      *uint8     `json:"defender_skill" xml:"DefenderSkill"`
	SetPiecesSkill     *uint8     `json:"set_pieces_skill" xml:"SetPiecesSkill"`
	PlayerCategoryId   uint8      `json:"player_category_ID" xml:"PlayerCategoryId"`
	TrainerData        struct {
		TrainerType       uint8 `json:"trainer_type" xml:"TrainerType"`
		TrainerSkillLevel uint8 `json:"trainer_skill_level" xml:"TrainerSkillLevel"`
	} `json:"trainer_data" gorm:"embedded" xml:"TrainerData"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create" xml:"-"`
	UpdatedAt time.Time `json:"update_at" xml:"-"`
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse("2006-01-02 15:04:05", v)
	if err != nil {
		return err
	}
	*c = CustomTime{parse}
	return nil
}

// Scan para leer el tiempo desde la base de datos
func (c *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*c = CustomTime{Time: time.Time{}}
		return nil
	}
	v, ok := value.(time.Time)
	if !ok {
		return errors.New("invalid type for CustomTime")
	}
	*c = CustomTime{v}
	return nil
}

// Value para escribir el tiempo en la base de datos
func (c CustomTime) Value() (driver.Value, error) {
	return c.Time, nil
}
