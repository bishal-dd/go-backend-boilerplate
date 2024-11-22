package user

import (
	"github.com/bisal-dd/go-backend-boilerplate/graph/model"
	"github.com/bisal-dd/go-backend-boilerplate/helper/paginationUtil"
)


func convertEdges(edges []*paginationUtil.Edge[*model.User]) []*model.UserEdge {
    locationEdges := make([]*model.UserEdge, len(edges))
    for i, edge := range edges {
        locationEdges[i] = &model.UserEdge{
            Cursor: edge.Cursor,
            Node:   edge.Node,
        }
    }
    return locationEdges
}
