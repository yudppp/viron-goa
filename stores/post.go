package stores

import (
	"fmt"
	"sort"

	"github.com/yudppp/viron-goa/app"
)

var postStore map[int]app.Post

// no atomic counter
var identifyCursor = 0

func init() {
	size := 10
	postStore = make(map[int]app.Post, size)
	for i := 0; i < size; i++ {
		id := identifyCursor + 1
		identifyCursor++
		post := app.Post{
			ID:       id,
			URLName:  fmt.Sprintf("post_%v", id),
			Title:    fmt.Sprintf("Title %v", id),
			Contents: "hello world",
			Status:   "draft",
		}
		postStore[id] = post
	}
}

type PostOptions struct {
	Status *string
}

func FindPosts(limit, offset int, opts PostOptions) ([]*app.Post, error) {
	list := filterAndSorted(opts)
	size := len(list)
	if size < offset {
		return nil, nil
	}
	if size < offset+limit {
		return list[offset:], nil
	}
	return list[offset : offset+limit], nil
}

func GetPosts(id int) (*app.Post, error) {
	post, ok := postStore[id]
	if !ok {
		return nil, nil
	}
	return &post, nil
}

func CountPosts(opts PostOptions) (int, error) {
	list := filterAndSorted(opts)
	return len(list), nil
}

func CreatePost(post app.Post) error {
	id := identifyCursor + 1
	identifyCursor++
	post.ID = id
	postStore[id] = post
	return nil
}

func UpdatePost(post app.Post) error {
	postStore[post.ID] = post
	return nil
}

func DeletePost(id int) error {
	delete(postStore, id)
	return nil
}

func filterAndSorted(opts PostOptions) []*app.Post {
	keys := make([]int, 0, len(postStore))
	for k := range postStore {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	list := make([]*app.Post, len(postStore))
	for i, key := range keys {
		post, ok := postStore[key]
		if ok {
			list[i] = &post
		}
	}
	// filter
	if opts.Status != nil {
		filterd := make([]*app.Post, 0, len(list))
		for _, post := range list {
			if post.Status == *opts.Status {
				filterd = append(filterd, post)
			}
		}
		list = filterd
	}
	return list
}
