package society

import (
	"context"
	"log"

	"aimi-landing-back-go/domain/entity"
	g_ "aimi-landing-back-go/domain/repository/gorm"

	"github.com/jinzhu/gorm"

	"io"

	"errors"

	pb "aimi-landing-back-go/internal/protorender"
)

type Server struct {
	db *gorm.DB
	pb.UnimplementedSocietyServiceServer
}

func (s *Server) NewServer() {
	g := g_.Gorm{}
	s.db = g.NewDb()[0]

}

func (s *Server) SimpleRPC(ctx context.Context, in *pb.RequestSocieties) (*pb.ResponseSocieties, error) {
	log.Printf("Receive message body from client: %v", in)

	societies := make([]*entity.Society, 0)
	societies_ := make([]*pb.Society, 0)

	if in.Id > 0 {
		if s.db.Where("id = ?", in.Id).Find(&societies).Error != nil {
			return nil, errors.New("Societies not found")
		}
	} else {
		if s.db.Find(&societies).Error != nil {
			return nil, errors.New("Societies not found")
		}
	}

	for i := 0; i < len(societies); i++ {
		society := pb.Society{
			Id:         societies[i].ID,
			Abstract:   societies[i].Abstract,
			Country:    societies[i].Country,
			Department: societies[i].Department,
			Longitude:  societies[i].Longitude,
			Latitude:   societies[i].Latitude,
			Web:        societies[i].Web,
			Video:      societies[i].Video,
		}
		societies_ = append(societies_, &society)
	}

	return &pb.ResponseSocieties{Societies: societies_}, nil
}

func (s *Server) ServerSideStreamingRPC(in *pb.RequestSocieties, stream pb.SocietyService_ServerSideStreamingRPCServer) error {
	log.Printf("Receive message body from client: %v", in)

	societies := make([]*entity.Society, 0)
	// societies_ := make([]*Society, 0)

	if in.Id > 0 {
		if s.db.Where("id = ?", in.Id).Find(&societies).Error != nil {
			return errors.New("Societies not found")
		}
	} else {
		if s.db.Find(&societies).Error != nil {
			return errors.New("Societies not found")
		}
	}

	for _, s := range societies {
		stream.Send(&pb.Society{
			Id:         s.ID,
			Abstract:   s.Abstract,
			Country:    s.Country,
			Department: s.Department,
			Longitude:  s.Longitude,
			Latitude:   s.Latitude,
			Web:        s.Web,
			Video:      s.Video,
		})
	}
	return nil
}

func (s *Server) ClientSideStreamingRPC(stream pb.SocietyService_ClientSideStreamingRPCServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// log.Printf("Receive message body from client: %v", in.Id)
		log.Printf("Last Point: %v", in)

		societies := make([]*entity.Society, 0)
		societies_ := make([]*pb.Society, 0)

		if in.Id > 0 {
			if s.db.Where("id = ?", in.Id).Find(&societies).Error != nil {
				return errors.New("Societies not found")
			}
		} else {
			if s.db.Find(&societies).Error != nil {
				return errors.New("Societies not found")
			}
		}

		for i := 0; i < len(societies); i++ {
			society := pb.Society{
				Id:         societies[i].ID,
				Abstract:   societies[i].Abstract,
				Country:    societies[i].Country,
				Department: societies[i].Department,
				Longitude:  societies[i].Longitude,
				Latitude:   societies[i].Latitude,
				Web:        societies[i].Web,
				Video:      societies[i].Video,
			}
			societies_ = append(societies_, &society)
		}

		return stream.SendAndClose(&pb.ResponseSocieties{Societies: societies_})
	}
}

func (s *Server) BidirectionalStreamingRPC(stream pb.SocietyService_BidirectionalStreamingRPCServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// log.Printf("Receive message body from client: %v", in.Id)
		log.Printf("Last Point: %v", in)

		societies := make([]*entity.Society, 0)
		// societies_ := make([]*Society, 0)

		if in.Id > 0 {
			if s.db.Where("id = ?", in.Id).Find(&societies).Error != nil {
				return errors.New("Societies not found")
			}
		} else {
			if s.db.Find(&societies).Error != nil {
				return errors.New("Societies not found")
			}
		}

		for _, s := range societies {
			stream.Send(&pb.Society{
				Id:         s.ID,
				Abstract:   s.Abstract,
				Country:    s.Country,
				Department: s.Department,
				Longitude:  s.Longitude,
				Latitude:   s.Latitude,
				Web:        s.Web,
				Video:      s.Video,
			})
		}
		return nil
	}
}
