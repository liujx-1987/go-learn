package service

import context "context"

// ProdServiceServer is the server API for ProdService service.
// type ProdServiceServer interface {
	// 定义方法
	//GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error)
// }


type ProductService struct {

}

func (*ProductService) GetProductStock(c context.Context, req *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{
		ProdStock: req.ProdId,
	}, nil
}