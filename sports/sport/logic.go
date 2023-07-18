package sport

import (
	"context"
	"log"
)

type Server struct {
	SportServiceServer
}

var teams = []*TeamResponse{
	{
		Id:    1,
		Name:  "Team A",
		Team:  "Team A",
		Sport: "Basketball",
	},
	{
		Id:    2,
		Name:  "Team B",
		Team:  "Team B",
		Sport: "Soccer",
	},
	{
		Id:    3,
		Name:  "Team C",
		Team:  "Team C",
		Sport: "Baseball",
	},
}


func (s *Server) GetAllTeams(ctx context.Context, in *AllTeamsRequest) (*TeamsResponse, error){
	log.Printf("this is a log log log , %s")
	return &TeamsResponse{Teams: teams}, nil
}

