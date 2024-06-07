package postgres

import (
	"database/sql"
	"fmt"
	"log"
	pb "root/genprotos"

	"github.com/google/uuid"
)

type PublicStorage struct {
	db *sql.DB
}

func NewPublicStorage(db *sql.DB) *PublicStorage {
	return &PublicStorage{db: db}
}

func (p *PublicStorage) CreatePublic(pub *pb.Public) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO public (id, first_name, last_name, birthday, gender, nation, party_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := p.db.Exec(query, id, pub.FirstName, pub.LastName, pub.Birthday, pub.Gender, pub.Nation, pub.Party.Id)
	return nil, err
}

func (p *PublicStorage) GetByIdPublic(id *pb.ById) (*pb.Public, error) {
	query := `
			SELECT p.id, p.first_name, p.last_name, p.birthday, p.gender, p.nation, 
			pa.name, pa.slogan, pa.open_date, pa.description
			FROM public p
			join party pa on p.party_id=pa.id
			WHERE p.id = $1
		`
	row := p.db.QueryRow(query, id.Id)

	var pub pb.Public
	var party pb.Party
	err := row.Scan(&pub.Id,
		&pub.FirstName,
		&pub.LastName,
		&pub.Birthday,
		&pub.Gender,
		&pub.Nation,
		&party.Name,
		&party.Slogan,
		&party.OpenDate,
		&party.Description)
	if err != nil {
		return nil, err
	}
	pub.Party = &party

	return &pub, nil
}

func (p *PublicStorage) GetAllPublic(_ *pb.Void) (*pb.GetAllPublic, error) {
	pubs := &pb.GetAllPublic{}
	row, err := p.db.Query(`SELECT p.id, p.first_name, p.last_name, p.birthday, p.gender, p.nation, 
		pa.name, pa.slogan, pa.open_date, pa.description
		FROM public p
		join party pa on p.party_id=pa.id
	`)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var party pb.Party
		pub := &pb.Public{}
		err = row.Scan(&pub.Id,
			&pub.FirstName,
			&pub.LastName,
			&pub.Birthday,
			&pub.Gender,
			&pub.Nation,
			&party.Name,
			&party.Slogan,
			&party.OpenDate,
			&party.Description)
		if err != nil {
			return nil, err
		}
		pub.Party = &party
		pubs.Publics = append(pubs.Publics, pub)
	}
	return pubs, nil
}

func (p *PublicStorage) UpdatePublic(pub *pb.Public) (*pb.Void, error) {
	query := `
		UPDATE public
		SET first_name = $2, last_name = $3, birthday = $4, gender = $5, nation = $6, party_id = $7
		WHERE id = $1 
	`
	_, err := p.db.Exec(query, pub.Id, pub.FirstName, pub.LastName, pub.Birthday, pub.Gender, pub.Nation, pub.Party.Id)
	return nil, err
}

func (p *PublicStorage) DeletePublic(id *pb.ById) (*pb.Void, error) {
	query := `
		delete from public where id = $1
	`
	_, err := p.db.Exec(query, id.Id)
	return nil, err
}

func (p *PublicStorage) CheakPublic(id *pb.ById) (*pb.Void, error) {
	query := ` select id from public
			   where EXTRACT(YEAR FROM AGE(birthday))>40 and nation='Uzbek' and id=$1
	`
	row, err := p.db.Query(query, id.Id)
	if err!=nil{
		log.Fatal("Error while CheakPublic: ", err.Error())
	}
	if !row.Next() {
		err = fmt.Errorf("yosh yoki Millat : %s", "togri kelmaydi")
		fmt.Println(err)
		return nil, err
	}
	return nil, nil
}
