package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	CandidateS *CandidateManager
	ElectionS *ElectionManager
	PublicVoteS *PublicVoteManager
	
}

func NewPostgresManager() (*Storage, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "1234", "localhost", 5432, "votingsystem")
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	cm := NewCandidateManager(db)
	pm := NewPublicVoteManager(db)
	em := NewElectionManager(db)
	
	return &Storage{Db: db, CandidateS: cm,ElectionS: em,PublicVoteS: pm}, err
	


}
