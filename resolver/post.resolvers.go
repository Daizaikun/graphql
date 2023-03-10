package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"example/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.Post, error) {

	// Create ID
	id := primitive.NewObjectID().Hex()

	// Create document
	post := &model.Post{
		ID:      id,
		Title:   input.Title,
		Content: input.Content,
		Author: &model.User{
			ID: input.AuthorID,
		},
	}

	//Insert operation
	_, err := r.Coll("post").InsertOne(ctx, post)

	// Check for errors
	if err != nil {
		return nil, err
	}

	return post, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.Post, error) {

	// Create filter
	filter := bson.M{"_id": input.ID}

	// Create update document
	updateDoc := bson.M{}
	if input.Title != nil {
		updateDoc["title"] = *input.Title
	}

	if input.Content != nil {
		updateDoc["Content"] = *input.Content
	}

	// Perform update operation
	err := r.Coll("post").UpdateId(ctx, input.ID, updateDoc)

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Get updated document

	var post model.Post
	err = r.Coll("post").Find(ctx, filter).One(&post)

	// Check for errors
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, input model.DeletePostInput) (*model.Post, error) {

	// Create delete filter
	filter := bson.M{"_id": input.ID}

	var post model.Post
	err := r.Coll("post").Find(ctx, filter).One(&post)

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Perform delete operation
	err = r.Coll("post").Remove(ctx, filter)

	// Check for errors
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, id string) (*model.Post, error) {


	// Create delete filter
	filter := bson.M{"_id": id}

	var post model.Post
	err := r.Coll("post").Find(ctx, filter).One(&post)

	// Check for errors
	if err != nil {
		return nil, err
	}

	Author, err := r.Query().GetUser(ctx, post.Author.ID)

	// Check for errors
	if err != nil {
		return nil, err
	}

	post.Author = Author

	return &post, nil
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*model.Post, error) {
	var Posts []*model.Post

	//Find operation
	err := r.Coll("post").Find(ctx, bson.M{}).All(&Posts)
	if err != nil {
		return nil, err
	}

	for _, Post := range Posts {
		user, err := r.Query().GetUser(ctx, Post.Author.ID)
		if err != nil {
			continue
		}
		Post.Author = user
	}

	return Posts, nil
}

func (r *queryResolver) GetPostsAutor(ctx context.Context, id string) ([]*model.Post, error) {
	var Posts []*model.Post

	//Find operation
	err := r.Coll("post").Find(ctx, bson.M{"author._id": id }).All(&Posts)

	if err != nil {
		return nil, err
	}

	return Posts, nil
}
