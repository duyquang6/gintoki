package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"gintoki/domain/dto"
	"gintoki/domain/interactor"
	"gintoki/utils/loggerV2"
	"net/http"
	"strconv"

	pb "gintoki/application/handler/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
)

type ProductInventoryHandler struct {
	Handler
	Interactor interactor.ProductInventoryInteractor
}

func (p ProductInventoryHandler) Get(w http.ResponseWriter, request *http.Request) {
	productIDStr, ok := mux.Vars(request)["product_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "product_id invalid")
		return
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		p.Handler.Error(w, err)
		return
	}
	product, err := p.Interactor.Get(productID)
	if err != nil {
		p.Handler.Error(w, err)
		return
	}
	res, _ := json.Marshal(&product)
	p.Data(w, http.StatusOK, res)
}

func (p ProductInventoryHandler) GetMultiProductInventory(ctx context.Context, req *pb.GetMultiProductInventoryRequest) (*pb.GetMultiProductInventoryResponse, error) {
	data, err := p.Interactor.GetMulti(ctx, req.ProductIDs)
	if err != nil {
		return &pb.GetMultiProductInventoryResponse{
			Meta: &pb.Meta{
				Code:    uint64(codes.InvalidArgument),
				Message: err.Error(),
			},
			Data: nil,
		}, nil
	}
	res := &pb.GetMultiProductInventoryResponse{
		Meta: &pb.Meta{
			Code: uint64(codes.OK),
		},
		Data: data,
	}
	return res, nil
}

func (p ProductInventoryHandler) UpdateLocalCacheFromQueue(ctx context.Context, messageValue []byte) error {
	data := dto.KafkaMessage{}
	err := json.Unmarshal(messageValue, &data)
	if err != nil {
		loggerV2.Errorf(ctx, "error when unmarshal: %v", err)
		return err
	}
	return p.Interactor.UpdateLocalCacheFromQueue(ctx, data.Payload.UpdateLocalCacheRequest)
}
