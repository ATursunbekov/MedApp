package repository

import (
	"MedApp/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MedicineRepository struct {
	collection *mongo.Collection
}

func NewMedicineRepository(db *mongo.Database) *MedicineRepository {
	return &MedicineRepository{
		collection: db.Collection("medicines"),
	}
}

func (r *MedicineRepository) Create(ctx context.Context, medicine *model.Medicine) error {
	_, err := r.collection.InsertOne(ctx, medicine)
	return err
}

func (r *MedicineRepository) GetByID(ctx context.Context, id string) (*model.Medicine, error) {
	var medicine model.Medicine
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&medicine)
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *MedicineRepository) GetAll(ctx context.Context) ([]*model.Medicine, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var medicines []*model.Medicine
	if err := cursor.All(ctx, &medicines); err != nil {
		return nil, err
	}
	return medicines, nil
}

func (r *MedicineRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
