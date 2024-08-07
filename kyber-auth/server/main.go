package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime"

	"github.com/cloudflare/circl/kem"
	"github.com/cloudflare/circl/kem/kyber/kyber512"
	"gitlab.sliit.lk/r24-055/r24-055/kyber-auth"
	pb "gitlab.sliit.lk/r24-055/r24-055/kyber-auth/proto"
	"google.golang.org/grpc"
)

type keyExchangeServer struct {
	pb.UnimplementedKeyExchangeServiceServer
}

type registerServer struct {
	pb.UnimplementedRegisterServiceServer
}

var privateKey = utils.ImportPriKey("kyber_id")

func (s *keyExchangeServer) KeyExchange(ctx context.Context, in *pb.KeyExchangeRequest) (*pb.KeyExchangeResponse, error) {

	// Memory benchmarking
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", utils.BitsToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", utils.BitsToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", utils.BitsToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)

	encryptedAccessToken := Decapsulate(in.Kem, privateKey)
	return &pb.KeyExchangeResponse{EncryptedSharedSecret: encryptedAccessToken}, nil
}

func (s *registerServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Send to our vector DB and do the processing here and get a response

	return &pb.RegisterResponse{Success: true}, nil
}

func BytesToPublicKey(rawPublicKey []byte) (kem.PublicKey, error) {
	return kyber512.Scheme().UnmarshalBinaryPublicKey(rawPublicKey)
}

func Decapsulate(kemData []byte, privateKey kem.PrivateKey) []byte {
	sharedSecret, err := kyber512.Scheme().Decapsulate(privateKey, kemData)
	accessToken := []byte("SECURE_ACCESS_TOKEN")
	utils.CheckErr(err)

	return utils.CyclicXOR(sharedSecret, accessToken)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	utils.CheckErr(err)
	log.Printf("Server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterKeyExchangeServiceServer(s, &keyExchangeServer{})
	pb.RegisterRegisterServiceServer(s, &registerServer{})

	srvErr := s.Serve(lis)
	utils.CheckErr(srvErr)
}

