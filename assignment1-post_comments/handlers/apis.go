package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/patrickmn/go-cache"
	"posts/cache"
	"posts/helpers"
	"posts/types"
	"net/http"
	"strconv"
	"sync"
)

type Interface interface {
	FetchCommentsForAllPosts(w http.ResponseWriter, r *http.Request)
	FetchCommentsForPostById(w http.ResponseWriter, r *http.Request)
}

var _ Interface = (*Server)(nil)
var url = "https://jsonplaceholder.typicode.com/posts"
var posts *[]types.Post

func (s *Server) FetchCommentsForAllPosts(w http.ResponseWriter, r *http.Request) {

	cacheData, ok := cache.Get("data")

	if ok {
		fmt.Println("Fetching from cache....")
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		enc.Encode(cacheData)

	} else {
		fmt.Println("expensive operation.. fetching from 3rd party api")
		posts, err := helpers.FetchAllPosts("https://jsonplaceholder.typicode.com/posts")
		// handle err
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		// initialise wait group
		var wg sync.WaitGroup
		// Loop in for each post
		for i := range posts {
			//Add 1 wait group
			wg.Add(1)
			go func(post *types.Post) {
				defer wg.Done()
				//Fetch Comments for each post by passing the post id
				//fmt.Println("Making req for comments ->", post.PostID)
				c, err := helpers.FetchCommentForPostID(post.PostID)

				if err != nil {
					fmt.Println(err.Error())
				}
				post.Comments = c

			}(&posts[i])
		}
		wg.Wait()
		w.Header().Set("Content-Type", "application/json")

		//Json Encoding
		enc := json.NewEncoder(w)
		//Prettify the content using Indentation
		enc.SetIndent("", "  ")
		//Marshaling and writing the result to the object
		enc.Encode(posts)

		cache.Set("data", &posts)

	}

}

func Caching() {
	fmt.Println("Caching operation.. fetching from 3rd party api")
	posts, err := helpers.FetchAllPosts("https://jsonplaceholder.typicode.com/posts")
	// handle err
	if err != nil {
		//fmt.Fprintln(w, err.Error())
		fmt.Println(err)
		return
	}
	// initialise wait group
	var wg sync.WaitGroup
	// Loop in for each post
	for i := range posts {
		//Add 1 wait group
		wg.Add(1)
		go func(post *types.Post) {
			defer wg.Done()
			//Fetch Comments for each post by passing the post id
			//fmt.Println("Making req for comments ->", post.PostID)
			c, err := helpers.FetchCommentForPostID(post.PostID)

			if err != nil {
				fmt.Println(err.Error())
			}
			post.Comments = c

		}(&posts[i])

		wg.Wait()
		//w.Header().Set("Content-Type", "application/json")

		//Json Encoding
		//enc := json.NewEncoder(w)
		//Prettify the content using Indentation
		//enc.SetIndent("", "  ")
		//Marshaling and writing the result to the object
		//enc.Encode(posts)

		cache.Set("data", &posts)
	}
}

// Fetches comments for a post by post id
func (s *Server) FetchCommentsForPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	postid, _ := strconv.Atoi(params)
	comments, err := helpers.FetchCommentForPostID(postid)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	//Content-Type header is set to json, to inform the client that JSON data is being sent.
	w.Header().Set("Content-Type", "application/json")

	//Json Encoding
	enc := json.NewEncoder(w)
	//Prettify the content using Indentation
	enc.SetIndent("", "  ")
	//Marshaling and writing the result to the object
	enc.Encode(comments)
}
