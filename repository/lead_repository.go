package repository

import (
	"errors"
	"go-crud-grpc/pb"
	"sync"
)

type LeadRepository struct {
	mu    sync.Mutex
	leads map[int32]*pb.Lead
}

func NewLeadRepository() *LeadRepository {
	return &LeadRepository{
		leads: make(map[int32]*pb.Lead),
	}
}

func (r *LeadRepository) CreateLead(lead *pb.Lead) (*pb.Lead, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.leads[lead.Id]; exists {
		return nil, errors.New("lead already exists")
	}

	r.leads[lead.Id] = lead
	return lead, nil
}

func (r *LeadRepository) GetLead(id int32) (*pb.Lead, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	lead, exists := r.leads[id]
	if !exists {
		return nil, errors.New("lead not found")
	}
	return lead, nil
}

func (r *LeadRepository) GetAllLeads() []*pb.Lead {
	r.mu.Lock()
	defer r.mu.Unlock()

	leads := []*pb.Lead{}
	for _, lead := range r.leads {
		leads = append(leads, lead)
	}
	return leads
}

func (r *LeadRepository) UpdateLead(lead *pb.Lead) (*pb.Lead, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.leads[lead.Id]; !exists {
		return nil, errors.New("lead not found")
	}

	r.leads[lead.Id] = lead
	return lead, nil
}

func (r *LeadRepository) DeleteLead(id int32) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.leads[id]; !exists {
		return errors.New("lead not found")
	}

	delete(r.leads, id)
	return nil
}
