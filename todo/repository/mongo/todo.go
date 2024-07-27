package mongo

import (
	"context"

	"github.com/mchayapol/go-todo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId"`
	Completed bool               `bson:"completed"`
	Title     string             `bson:"title"`
}

type TodoRepository struct {
	db *mongo.Collection
}

func NewTodoRepository(db *mongo.Database, collection string) *TodoRepository {
	return &TodoRepository{
		db: db.Collection(collection),
	}
}

func (r TodoRepository) CreateTodo(ctx context.Context, user *models.User, bm *models.Todo) error {
	bm.UserID = user.ID

	model := toModel(bm)

	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	bm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r TodoRepository) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	uid, _ := primitive.ObjectIDFromHex(user.ID)
	cur, err := r.db.Find(ctx, bson.M{
		"userId": uid,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Todo, 0)

	for cur.Next(ctx) {
		user := new(Todo)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toTodos(out), nil
}

func (r TodoRepository) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	return err
}

func toModel(b *models.Todo) *Todo {
	uid, _ := primitive.ObjectIDFromHex(b.UserID)

	return &Todo{
		UserID:    uid,
		Completed: b.Completed,
		Title:     b.Title,
	}
}

func toTodo(b *Todo) *models.Todo {
	return &models.Todo{
		ID:        b.ID.Hex(),
		UserID:    b.UserID.Hex(),
		Completed: b.Completed,
		Title:     b.Title,
	}
}

func toTodos(bs []*Todo) []*models.Todo {
	out := make([]*models.Todo, len(bs))

	for i, b := range bs {
		out[i] = toTodo(b)
	}

	return out
}
