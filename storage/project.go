package storage

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Project simply represents a project
type Project struct {
	ID           primitive.ObjectID `bson:"_id"`
	Title        string             `bson:"title"`
	Status       string             `bson:"status"`
	Tags         []string           `bson:"tags"`
	DateCreated  time.Time          `bson:"date_created"`
	DateModified time.Time          `bson:"date_modified"`
	Summary      string             `bson:"summary"`
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

// GetProjects returns all projects currently stored in the collection
func GetProjects(col *mongo.Collection) []*Project {
	projects := make([]*Project, 0)

	cursor, err := col.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Error(err)
		return projects
	}

	defer cursor.Close(context.TODO())

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

// AddProject adds a project to mongodb and returns the ID assigned to this new document
func AddProject(col *mongo.Collection, title string, status string, tags []string, summary string) string {

	result, err := col.InsertOne(context.TODO(), bson.M{
		"title":         title,
		"status":        status,
		"tags":          tags,
		"summary":       summary,
		"date_created":  time.Now(),
		"date_modified": time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}

	pid := result.InsertedID.(primitive.ObjectID).Hex()
	log.Infof("Created new project, %s (%s)", title, pid)
	return pid
}
