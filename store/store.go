package store

import (
	"medium/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetSession(collection string, pk string) *mgo.Session {
	// Dial to database and Return mgo session
	var session *mgo.Session
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	ensureIndex(collection, pk, session)
	return session
}

func ensureIndex(collection string, pk string, s *mgo.Session) {
	// Ensure an index on the collection, why?
	session := s.Copy()
	defer session.Close()
	c := session.DB("medium").C(collection)
	index := mgo.Index{
		Key:        []string{pk},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func SavePost(post *models.Post) error {
	session := GetSession("Post", "ID")
	session = session.Copy()
	defer session.Close()
	c := session.DB("medium").C("Post")
	post.ID = bson.NewObjectId()
	err := c.Insert(&post)
	return err
}

func GetAllPosts() ([]models.Post, error) {
	session := GetSession("Post", "ID")
	session = session.Copy()
	defer session.Close()
	c := session.DB("medium").C("Post")
	var posts []models.Post
	err := c.Find(bson.M{}).All(&posts)
	return posts, err
}

func GetPost(postId string) (models.Post, error) {
	session := GetSession("Post", "ID")
	session = session.Copy()
	defer session.Close()
	c := session.DB("medium").C("Post")
	var post models.Post
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(postId)}).One(&post)
	return post, err
}
