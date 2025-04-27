package repository

import (
	"MedApp/internal/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *MedicineRepository) Create(ctx context.Context, medicine *model.Medicine) (string, error) {
	result, err := r.collection.InsertOne(ctx, medicine)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *MedicineRepository) GetByID(ctx context.Context, id string) (*model.Medicine, error) {
	var medicine model.Medicine
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&medicine)
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
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("medicine not found")
	}
	return nil
}
