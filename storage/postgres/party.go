package postgres

import (
	"database/sql"
	"fmt"
	pb "root/genprotos"

	"github.com/google/uuid"
)

type PartyStorage struct {
	db *sql.DB
}

func NewPartyStorage(db *sql.DB) *PartyStorage {
	return &PartyStorage{db: db}
}

func (p *PartyStorage) CreateParty(party *pb.Party) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO party (id, name, slogan, open_date, description)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, party.Name, party.Slogan, party.OpenDate, party.Description)
	return nil, err
}

func (p *PartyStorage) GetByIdParty(id *pb.ById) (*pb.Party, error) {
		query := `
			SELECT id, name, slogan, open_date, description
			FROM party
			WHERE id = $1
		`
		row := p.db.QueryRow(query, id.Id)

		var party pb.Party
		err := row.Scan(&party.Id,
			&party.Name,
			&party.Slogan,
			&party.OpenDate,
			&party.Description)
		if err != nil {
			return nil, err
		}

		return &party, nil
}

func (p *PartyStorage) GetAllParty(pa *pb.Party) (*pb.GetAllParty, error) {
	partys := &pb.GetAllParty{}
	var arr []interface{}
	count:=1
	query:=`select id, name, slogan, open_date, description from party where delated_at=0`
	if len(pa.Description)>0{
		query+=fmt.Sprintf(` and description=$%d`, count)
		arr = append(arr, pa.Description)
	}
	if len(pa.Name)>0{
		query+=fmt.Sprintf(` and name=$%d`, count)
		arr = append(arr, pa.Name)
	}
	if len(pa.Slogan)>0{
		query+=fmt.Sprintf(` and slogan=$%d`, count)
		arr = append(arr, pa.Slogan)
	}
	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}


	for row.Next() {
		party := &pb.Party{}
		err = row.Scan(&party.Id, &party.Name, &party.Slogan, &party.OpenDate, &party.Description)
		if err != nil {
			return nil, err
		}
		partys.Partys = append(partys.Partys, party)
	}
	return partys, nil
}

func (p *PartyStorage) UpdateParty(party *pb.Party) (*pb.Void, error) {
	query := `
		UPDATE party
		SET name = $2, open_date = $3, description = $4
		WHERE id = $1 
	`
	_, err := p.db.Exec(query, party.Id, party.Name, party.OpenDate, party.Description)
	return nil, err
}

func (p *PartyStorage) DeleteParty(id *pb.ById) (*pb.Void, error) {
	query := `
		delete from party where id = $1
	`
	_, err := p.db.Exec(query, id.Id)
	return nil, err
}



func (p *PartyStorage) GetAllPartyForCandidate(pa *pb.Party) (*pb.GetAllParty, error) {
	partys := &pb.GetAllParty{}
	row, err := p.db.Query(`select id, name, slogan, open_date, description from party
							where id =$1 `, )
	if err != nil {
		return nil, err
	}

	for row.Next() {
		party := &pb.Party{}
		err = row.Scan(&party.Id, &party.Name, &party.Slogan, &party.OpenDate, &party.Description)
		if err != nil {
			return nil, err
		}
		partys.Partys = append(partys.Partys, party)
	}
	return partys, nil
}