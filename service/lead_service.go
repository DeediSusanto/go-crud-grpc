package service

import (
	"context"
	"go-crud-grpc/internal/repository"
	"go-crud-grpc/pb"
)

type LeadService struct {
	pb.UnimplementedLeadServiceServer
	repo *repository.LeadRepository
}

func NewLeadService(repo *repository.LeadRepository) *LeadService {
	return &LeadService{repo: repo}
}

func (s *LeadService) CreateLead(ctx context.Context, req *pb.CreateLeadRequest) (*pb.LeadResponse, error) {
	lead := &pb.Lead{
		Id:      int32(len(s.repo.GetAllLeads()) + 1), // Auto-increment ID
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Company: req.Company,
		Source:  req.Source,
		Status:  req.Status,
		Notes:   req.Notes,
	}

	createdLead, err := s.repo.CreateLead(lead)
	if err != nil {
		return nil, err
	}

	return &pb.LeadResponse{Success: true, Message: "Lead created", Lead: createdLead}, nil
}

func (s *LeadService) GetLead(ctx context.Context, req *pb.GetLeadRequest) (*pb.LeadResponse, error) {
	lead, err := s.repo.GetLead(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.LeadResponse{Success: true, Message: "Lead found", Lead: lead}, nil
}

func (s *LeadService) GetAllLeads(ctx context.Context, req *pb.GetAllLeadsRequest) (*pb.LeadListResponse, error) {
	leads := s.repo.GetAllLeads()
	return &pb.LeadListResponse{Leads: leads}, nil
}

func (s *LeadService) UpdateLead(ctx context.Context, req *pb.UpdateLeadRequest) (*pb.LeadResponse, error) {
	lead := &pb.Lead{
		Id:      req.Id,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Company: req.Company,
		Source:  req.Source,
		Status:  req.Status,
		Notes:   req.Notes,
	}

	updatedLead, err := s.repo.UpdateLead(lead)
	if err != nil {
		return nil, err
	}

	return &pb.LeadResponse{Success: true, Message: "Lead updated", Lead: updatedLead}, nil
}

func (s *LeadService) DeleteLead(ctx context.Context, req *pb.DeleteLeadRequest) (*pb.LeadResponse, error) {
	err := s.repo.DeleteLead(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.LeadResponse{Success: true, Message: "Lead deleted"}, nil
}
