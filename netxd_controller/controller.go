package controllers

import (
	"context"

	cus "github.com/20-VIGNESH-K/netxd-customer-proto"
	"github.com/20-VIGNESH-K/netxd-dal/netxd_dal_interfaces"
	"github.com/20-VIGNESH-K/netxd-dal/netxd_dal_models"
)

type RPCServer struct {
	cus.UnimplementedCustomerServiceServer
}

var (
	CustomerService netxd_dal_interfaces.INetxdCustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *cus.Customer) (*cus.CustomerResponse, error) {
	dbCustomer := &netxd_dal_models.NetxdCustomer{CustomerId: req.CustomerId, FirstName: req.FirstName, LastName: req.LastName, BankId: req.BankId, Balance: req.Balance, IsActive: req.IsActive}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &cus.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseCustomer, nil
	}
}
