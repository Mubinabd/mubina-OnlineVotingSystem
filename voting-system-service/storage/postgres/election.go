package postgres

import (
	"database/sql"
	pb "project/genproto/public"

	"github.com/google/uuid"
)

type ElectionManager struct {
	conn *sql.DB
}

func NewElectionManager(conn *sql.DB) *ElectionManager {
	return &ElectionManager{conn: conn}
}

func (em *ElectionManager) Create(elec *pb.CreateElectionReq) error {
	id := uuid.NewString()
	query := `INSERT INTO election(id, name, date) VALUES($1, $2, $3)`
	_, err := em.conn.Exec(query, id, elec.Name, elec.Date)
	return err
}

func (em *ElectionManager) Get(elec *pb.ElectionReq) (*pb.ElectionReq, error) {
	res := &pb.ElectionReq{}
	query := `SELECT id, name, date FROM election WHERE id = $1`
	row := em.conn.QueryRow(query, elec.Id)
	err := row.Scan(
		&res.Id,
		&res.Name,
		&res.Date,
	)
	return res, err
}

func (em *ElectionManager) GetAll(filter *pb.Filter) (*pb.ElectionsGetAllResp, error) {
	query := `SELECT id, name, date FROM election`
	if filter.GetLimit() > 0 {
		query += " LIMIT $1"
	}
	if filter.GetOffset() > 0 {
		query += " OFFSET $2"
	}
	rows, err := em.conn.Query(query, filter.GetLimit(), filter.GetOffset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	electionList := &pb.ElectionsGetAllResp{}

	for rows.Next() {
		election := &pb.ElectionReq{}
		err := rows.Scan(
			&election.Id,
			&election.Name,
			&election.Date,
		)
		if err != nil {
			return nil, err
		}
		electionList.Elections = append(electionList.Elections, election)
	}
	return electionList, nil
}

func (em *ElectionManager) Update(elec *pb.ElectionUpdate) error {
	query := `UPDATE election SET name = $1, date = $2 WHERE id = $3`
	_, err := em.conn.Exec(query, elec.Name, elec.Date)
	return err
}

func (em *ElectionManager) Delete(elec *pb.GetByIdReq) error {
	_, err := em.conn.Exec("DELETE FROM election WHERE id = $1", elec.Id)
	if err != nil {
		return err
	}
	_, err = em.conn.Exec("UPDATE election SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1", elec.Id)
	return err
}
