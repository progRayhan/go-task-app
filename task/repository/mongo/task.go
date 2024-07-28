package mongo

import (
	"context"

	"github.com/mchayapol/go-task-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId"`
	Completed bool               `bson:"completed"`
	Title     string             `bson:"title"`
}

type TaskRepository struct {
	db *mongo.Collection
}

func NewTaskRepository(db *mongo.Database, collection string) *TaskRepository {
	return &TaskRepository{
		db: db.Collection(collection),
	}
}

func (r TaskRepository) CreateTask(ctx context.Context, user *models.User, bm *models.Task) error {
	bm.UserID = user.ID

	model := toModel(bm)

	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	bm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r TaskRepository) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	uid, _ := primitive.ObjectIDFromHex(user.ID)
	cur, err := r.db.Find(ctx, bson.M{
		"userId": uid,
	})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Task, 0)

	for cur.Next(ctx) {
		user := new(Task)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toTasks(out), nil
}

func (r TaskRepository) DeleteTask(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	return err
}

func toModel(b *models.Task) *Task {
	uid, _ := primitive.ObjectIDFromHex(b.UserID)

	return &Task{
		UserID:    uid,
		Completed: b.Completed,
		Title:     b.Title,
	}
}

func toTask(b *Task) *models.Task {
	return &models.Task{
		ID:        b.ID.Hex(),
		UserID:    b.UserID.Hex(),
		Completed: b.Completed,
		Title:     b.Title,
	}
}

func toTasks(bs []*Task) []*models.Task {
	out := make([]*models.Task, len(bs))

	for i, b := range bs {
		out[i] = toTask(b)
	}

	return out
}
