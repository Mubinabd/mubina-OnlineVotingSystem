package service

type Service struct {
	CM  *CandidateService
	EM  *ElectionService
	PVM *PublicVoteService
}

func NewService(cm *CandidateService, em *ElectionService, pvm *PublicVoteService) *Service {
	return &Service{CM: cm, EM: em, PVM: pvm}
}
