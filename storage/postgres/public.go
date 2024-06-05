package postgres

import (
	"database/sql"
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
			SELECT id, first_name, last_name, birthday, gender, nation, party_id
			FROM public
			WHERE id = $1
		`
		row := p.db.QueryRow(query, id.Id)

		var pub pb.Public
		err := row.Scan(&pub.Id,
			&pub.FirstName,
			&pub.LastName,
			&pub.Birthday,
			&pub.Gender,
			&pub.Nation,
			&pub.Party.Id)
		if err != nil {
			return nil, err
		}

		return &pub, nil
}

func (p *PublicStorage) GetAllPublic(_ *pb.Void) (*pb.GetAllPublic, error) {
	pubs := &pb.GetAllPublic{}
	row, err := p.db.Query("select id, first_name, last_name, birthday, gender, nation, party_id from public")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		pub := &pb.Public{}
		err = row.Scan(&pub.Id, &pub.FirstName, &pub.LastName, &pub.Birthday, &pub.Gender, &pub.Nation, &pub.Party.Id)
		if err != nil {
			return nil, err
		}
		pubs.Publics = append(pubs.Publics, pub)
	}
	return pubs, nil
}

func (p *PublicStorage) UpdatePublic(pub *pb.Public) (*pb.Void, error) {
	query := `
		UPDATE public_
		SET first_name = $2, last_name = $3, birthday = $4, gender = $5, nation = $6, party_id = $6
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
