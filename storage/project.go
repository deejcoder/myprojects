package storage

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Project represents a Project
type Project struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	Status       string             `json:"status" bson:"status"`
	Tags         []string           `json:"tags" bson:"tags"`
	DateCreated  time.Time          `json:"date_created" bson:"date_created"`
	DateModified time.Time          `json:"date_modified" bson:"date_modified"`
	Summary      string             `json:"summary" bson:"summary"`
	ProjectLink  string             `json:"projectLink" bson:"projectLink"`
	Content      string             `json:"content" bson:"content"`
}

// GetProject finds a project by project ID, and returns it. If not found, nil.
func GetProject(col *mongo.Collection, id primitive.ObjectID) *Project {
	filter := bson.M{"_id": id}

	var project *Project
	if err := col.FindOne(context.TODO(), filter).Decode(project); err != nil {
		log.Error(err)
		return nil
	}
	return project
}

// GetProjects returns all projects
func GetProjects(col *mongo.Collection) []*Project {
	projects := make([]*Project, 0)

	cursor, err := col.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Error(err)
		return projects
	}

	defer cursor.Close(context.TODO())

	// fetch a project from mongo
	for cursor.Next(context.TODO()) {
		var project Project
		if err := cursor.Decode(&project); err != nil {
			log.Error(err)
			continue
		}

		projects = append(projects, &project)

	}

	if err := cursor.Err(); err != nil {
		log.Error(err)
	}
	return projects
}
