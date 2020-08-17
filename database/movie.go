package database

import (
	"context"
	"errors"
	"strconv"

	"movie_spider/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMovies returns all movies.
// start, end int, order, sort string
func (d *TenDatabase) GetMovies(paging *model.Paging) []*model.Movie {
	movies := []*model.Movie{}
	cursor, err := d.DB.Collection("movies").
		Find(context.Background(), paging.Condition,
			&options.FindOptions{
				Skip:  paging.Skip,
				Sort:  bson.D{bson.E{Key: paging.SortKey, Value: paging.SortVal}},
				Limit: paging.Limit,
			})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		movie := &model.Movie{}
		if err := cursor.Decode(movie); err != nil {
			return nil
		}
		movies = append(movies, movie)
	}

	return movies
}

// CreateMovie creates a movie.
func (d *TenDatabase) FindOneAndReplace(movie *model.Movie) *mongo.SingleResult {
	upsert := true
	result := d.DB.Collection("movies").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "name", Value: movie.Name}},
			movie,
			&options.FindOneAndReplaceOptions{
				Upsert: &upsert,
			},
		)
	if result != nil {
		return result
	}
	return nil
}

func (d *TenDatabase) CreateMovie(movie *model.Movie) error {
	if _, err := d.DB.Collection("movies").
		InsertOne(context.Background(), movie); err != nil {
		return err
	}
	return nil
}

// GetMovieByName returns the movie by the given name or nil.
func (d *TenDatabase) GetMovieByName(name string) *model.Movie {
	var movie *model.Movie
	err := d.DB.Collection("movies").
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(&movie)
	if err != nil {
		return nil
	}
	return movie
}

// GetMovieByIDs returns the movie by the given id or nil.
func (d *TenDatabase) GetMovieByID(id primitive.ObjectID) *model.Movie {
	var movie *model.Movie
	err := d.DB.Collection("movies").
		FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).
		Decode(&movie)
	if err != nil {
		return nil
	}
	return movie
}

// GetMovieByIDs returns the movie by the given id or nil.
func (d *TenDatabase) GetMovieByIDs(ids []primitive.ObjectID) []*model.Movie {
	var movies []*model.Movie
	cursor, err := d.DB.Collection("movies").
		Find(context.Background(), bson.D{{
			Key: "_id",
			Value: bson.D{{
				Key:   "$in",
				Value: ids,
			}},
		}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		movie := &model.Movie{}
		if err := cursor.Decode(movie); err != nil {
			return nil
		}
		movies = append(movies, movie)
	}

	return movies
}

// CountMovie returns the movie count
func (d *TenDatabase) CountMovie(condition interface{}) string {
	if condition == nil {
		condition = bson.D{{}}
	}
	total, err := d.DB.Collection("movies").CountDocuments(context.Background(), condition, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

// DeleteMovieByID deletes a movie by its id.
func (d *TenDatabase) DeleteMovieByID(id primitive.ObjectID) error {
	if d.CountPost(bson.D{{Key: "movieId", Value: id}}) == "0" {
		_, err := d.DB.Collection("movies").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
		return err
	}
	return errors.New("the current movie has posts published")
}
