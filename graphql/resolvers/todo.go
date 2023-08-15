package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/Da-max/todo-go/graphql/generated"
	"github.com/Da-max/todo-go/graphql/model"
	"github.com/Da-max/todo-go/utils/auth"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var (
		user *model.User = auth.ForContext(ctx)
		todo *model.Todo = &model.Todo{
			Text: input.Text,
			Done: false,
		}
	)

	todo.UserID = int(user.ID)
	todo.User = *user

	if todo.Text == "" {
		graphql.AddError(ctx, gqlerror.Errorf("An empty todo cannot be saved."))
		return nil, nil
	}

	if res := r.DB.Create(todo); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo "+todo.Text+" cannot be add."))
		return nil, nil
	}

	return todo, nil
}

// RemoveTodo is the resolver for the removeTodo field.
func (r *mutationResolver) RemoveTodo(ctx context.Context, todoID int) (int, error) {
	var (
		todo *model.Todo = &model.Todo{}
	)

	if res := r.DB.First(todo, todoID); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(todoID)+" is not found"))
	}

	if res := r.DB.Delete(todo); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo cannot be delete"))
	}

	return todoID, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.NewTodo, todoID int) (*model.Todo, error) {
	var (
		// user *model.User = &model.User{}
		todo *model.Todo = &model.Todo{}
	)

	// if res := r.DB.Find(user).Where("id = ?", input.UserID); res.Error != nil {
	// 	panic("The user with id " + fmt.Sprint(input.UserID) + " is not found")
	// }

	if res := r.DB.First(todo, todoID); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(todoID)+" is not found"))
	}

	todo.Text = input.Text
	// todo.UserID = input.UserID

	if res := r.DB.Save(todo); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(todoID)+" cannot be saved"))
	}

	return todo, nil
}

// MarkDoneTodo is the resolver for the markDoneTodo field.
func (r *mutationResolver) MarkDoneTodo(ctx context.Context, todoID int) (*model.Todo, error) {
	var todo *model.Todo = &model.Todo{}

	if res := r.DB.First(todo, todoID); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(todoID)+" cannot be found."))
	}

	todo.Done = true

	if res := r.DB.Save(todo); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(todo.ID)+" cannot be update"))
	}

	return todo, nil
}

// MarkUndoneTodo is the resolver for the markUndoneTodo field.
func (r *mutationResolver) MarkUndoneTodo(ctx context.Context, todoID int) (*model.Todo, error) {
	var todo *model.Todo = &model.Todo{}

	if res := r.DB.First(todo, todoID); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(todoID)+" cannot be found."))
	}

	todo.Done = false

	if res := r.DB.Save(todo); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(todo.ID)+" cannot be update"))
	}

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var (
		todos []*model.Todo
		// user  *model.User   = &model.User{}
	)

	// if res := r.DB.Where(&model.User{Username: ctx.Value("Username").(string)}).First(user); res.Error != nil || res.RowsAffected == 0 {
	// 	panic("The username cannot be found.")
	// }

	if result := r.DB. /*.Where(&model.Todo{UserID: int(user.ID)})*/ Order("done, updated_at desc").Find(&todos); result.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todos cannot be query."))
	}

	return todos, nil
}

// ID is the resolver for the id field.
func (r *todoResolver) ID(ctx context.Context, obj *model.Todo) (int, error) {
	if result := r.DB.First(&model.Todo{}, obj.ID); result.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The todo with id "+fmt.Sprint(obj.ID)+" cannot be found."))
	}

	return int(obj.ID), nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	var user *model.User = &model.User{}
	if res := r.DB.First(user, obj.UserID); res.Error != nil {
		graphql.AddError(ctx, gqlerror.Errorf("The user cannot be found."))
	}
	return user, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
