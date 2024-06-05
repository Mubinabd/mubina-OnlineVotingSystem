package postgres

import (
	"database/sql"
	pb "project/genproto/public"

	"github.com/google/uuid"
)


type CandidateManager struct{
	conn *sql.DB
}

func NewCandidateManager(conn *sql.DB) *CandidateManager{
	return &CandidateManager{conn: conn}
}

func (cm *CandidateManager)Create(cand *pb.CreateCandidateReq) error{
	id := uuid.NewString()

	query := `INSERT INTO candidate(id,election_id,public_id)VALUES($1,$2,$3)`
	_,err := cm.conn.Exec(query,id,cand.ElectionId,cand.PublicId)
	return err
}

func (cm *CandidateManager)Get(cand *pb.GetCandidate)(*pb.CandidateRes,error){
	res := pb.CandidateRes{}
	query := `SELECT 
		c.id,
		c.election_id,
		c.public_id,
		e.name AS election_name,
		e.date AS election_date,
		p.first_name,
		p.last_name,
		p.birthday,
		p.gender,
		p.nation
	FROM 
		candidate c
	JOIN 
		election e ON c.election_id = e.id
	JOIN 
		public p ON c.public_id = p.id
	WHERE 
		c.id = $1
	`
	row := cm.conn.QueryRow(query)
	err := row.Scan(
		&res.Id,
		&res.Election.Id,
		&res.Election.Name,
		&res.Election.Date,
		&res.Public.Id,
		&res.Public.FirstName,
		&res.Public.LastName,
		&res.Public.Birthday,
		&res.Public.Gender,
		&res.Public.Nation,

	)
	return &res,err
}

func (cm *CandidateManager) GetAll(filter *pb.Filter) (*pb.CandidatiesGetAllResp, error) {
	query := `SELECT 
		c.id,
		c.election_id,
		c.public_id,
		e.name AS election_name,
		e.date AS election_date,
		p.first_name,
		p.last_name,
		p.birthday,
		p.gender,
		p.nation
	FROM 
		candidate c
	JOIN 
		election e ON c.election_id = e.id
	JOIN 
		public p ON c.public_id = p.id
`
	if filter.GetLimit() > 0 {
		query += " LIMIT " + string(filter.GetLimit())
	}
	if filter.GetOffset() > 0 {
		query += " OFFSET " + string(filter.GetOffset())
	}
	rows,err := cm.conn.Query(query)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	candidateList := &pb.CandidatiesGetAllResp{}

	for rows.Next() {
		var candidate pb.CandidateRes
		err := rows.Scan(
			&candidate.Id,
			&candidate.Public.Id,
			&candidate.Public.FirstName,
			&candidate.Public.LastName,
			&candidate.Public.Birthday,
			&candidate.Public.Gender,
			&candidate.Public.Nation,
			&candidate.Election.Id,
			&candidate.Election.Name,
			&candidate.Election.Date,
		)
		if err != nil{
			return nil,err
		}
		

		candidateList.Candidaties = append(candidateList.Candidaties,&candidate)
	}
	return candidateList,nil
}

func(cm *CandidateManager)Delete(cand *pb.GetByIdReq)error{
	_,err := cm.conn.Exec("DELETE FROM candidate WHERE id = $1",cand.Id)

	if err != nil{
		return err
	}

	_, err = cm.conn.Exec("UPDATE candidate SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1", cand.Id)
	return err
}