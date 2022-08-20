package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

var carType = graphql.NewObject(graphql.ObjectConfig{
	Name: "car",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"speed": &graphql.Field{
			Type: graphql.Int,
		},
		"racing_number": &graphql.Field{
			Type: graphql.Int,
		},
		"country": &graphql.Field{
			Type: graphql.String,
		},
		"racing_type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"car": &graphql.Field{
			Type: carType,
			Args: graphql.FieldConfigArgument{
				"racing_number": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rnQuery, isOk := p.Args["racing_number"].(int)
				fmt.Printf("CAR : %v", rnQuery)
				if !isOk {
					return nil, nil
				}
				for _, car := range carList {
					if car.RacingNumber == rnQuery {
						return car, nil
					}
				}
				return nil, nil
			},
			Description: "Get Single Car",
		},
		"carList": &graphql.Field{
			Type:        graphql.NewList(carType),
			Description: "List of todos",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return carList, nil
			},
		},
	},
})

var carSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
