package main

// import (
// 	"log"
// 	"net/http"

// 	handler "microservice/services/orders/handler/orders"
// 	"microservice/services/orders/service"
// )

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	// router := http.NewServeMux()

	// orderService := service.NewOrderService()
	// orderHandler := handler.NewHttpOrdersHandler(orderService)
	// orderHandler.RegisterRouter(router)

	// log.Println("Starting server on", s.addr)
	// return http.ListenAndServe(s.addr, router)
	return nil
}
