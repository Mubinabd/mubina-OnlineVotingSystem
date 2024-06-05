package postgres

import (
	"database/sql"
	pb "project/genproto/public"

	"github.com/google/uuid"
)

type PublicVoteManager struct{
	conn *sql.DB
}

func NewPublicVoteManager(conn *sql.DB) *PublicVoteManager{
	return &PublicVoteManager{conn: conn}
}

func (pvm *PublicVoteManager)Create(pubvote *pb.CreatePublicVoteReq)error{
	id := uuid.NewString()
	query := `INSERT INTO public_vote(id,election_id,public_id,candidate_id)VALUES($1,$2,$3,$4)`

	_,err := pvm.conn.Exec(query,id,pubvote.ElectionId,pubvote.PublicId,pubvote.CandidateId)
	return err
}

func (pvm *PublicVoteManager)Get(pubvote *pb.PublicVoteReq)(*pb.PublicVote,error){
	res := pb.PublicVote{}

	query := `SELECT id,election_id,public_id,candidate_id FROM public_vote WHERE id = $1`
	row := pvm.conn.QueryRow(query)
	err := row.Scan(
		&res.Id,
		&res.Public.Id,
		&res.Election.Id,
		&res.Candidate.Id,
	)
	return &res,err
}

func (pvm *PublicVoteManager)GetAll(filter *pb.Filter)(*pb.PublicVotesGetAllResp,error){
	query := `SELECT 
		p.id,
		p.election_id,
		p.public_id,
		p.candidate_id,
		e.name AS election_name,
		e.date AS election_date,
		pb.first_name,
		pb.last_name,
		pb.birthday,
		pb.gender,
		pb.nation,
		c.id
	FROM 
		public_vote p
	JOIN 
		election e ON p.election_id = e.id
	JOIN 
		public pb ON p.public_id = pb.id
	JOIN 
		candidate c ON p.candidate_id = c.id
	WHERE 
		p.id = $1
`
	if filter.GetLimit() > 0 {
		query += " LIMIT " + string(filter.GetLimit())
	}
	if filter.GetOffset() > 0 {
		query += " OFFSET " + string(filter.GetOffset())
	}
	rows,err := pvm.conn.Query(query)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	publicVotesList := &pb.PublicVotesGetAllResp{}

	for rows.Next() {
		var publicVotes pb.PublicVote
        err := rows.Scan(
            &publicVotes.Id,
            &publicVotes.Public.Id,
            &publicVotes.Public.FirstName,
            &publicVotes.Public.LastName,
            &publicVotes.Public.Birthday,
            &publicVotes.Public.Gender,
            &publicVotes.Public.Nation,
            &publicVotes.Election.Id,
            &publicVotes.Election.Name,
            &publicVotes.Election.Date,
            &publicVotes.Candidate.Id,
        )
        if err!= nil {
            return nil,err
        }
        publicVotesList.Publicvotes = append(publicVotesList.Publicvotes,&publicVotes)
	}
	return publicVotesList,nil
}

func (pvm *PublicVoteManager)Delete(pubvote *pb.GetByIdReq)error{
	_,err := pvm.conn.Exec("DELETE FROM public_votes WHERE id = $1",pubvote.Id)

	if err != nil{
		return err
	}

	_, err = pvm.conn.Exec("UPDATE public_votes SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1", pubvote.Id)
	return err
}
