package serverapi

import (
	"context"
	"log"

	"github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	commonpb "github.com/airbloc/airbloc-go/proto/rpc/v1"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/pkg/errors"
)

type DAuthAPI struct {
	manager *dauth.Manager
}

func NewDAuthAPI(backend node.Backend) (node.API, error) {
	manager := dauth.NewManager(backend.Client())
	return &DAuthAPI{manager}, nil
}

func (api *DAuthAPI) Allow(ctx context.Context, req *pb.DAuthRequest) (*commonpb.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid collection ID: %s", req.GetCollectionId())
	}

	// err = api.manager.Allow(collectionId, req.GetPasswordSignature())
	log.Println(collectionId.String())
	return &commonpb.Result{}, err
}

func (api *DAuthAPI) Deny(ctx context.Context, req *pb.DAuthRequest) (*commonpb.Result, error) {
	collectionId, err := common.IDFromString(req.GetCollectionId())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid collection ID: %s", req.GetCollectionId())
	}

	log.Println(collectionId.String())
	// err = api.manager.Deny(collectionId, req.GetPasswordSignature())
	return &commonpb.Result{}, err
}

func (api *DAuthAPI) AttachToAPI(service *node.APIService) {
	pb.RegisterDAuthServer(service.GrpcServer, api)
}
