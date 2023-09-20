package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"posts/types"
	"io/ioutil"
	"net/http"
)

var url = "https://jsonplaceholder.typicode.com/posts"
var resObj []types.Post

// Fetches all the posts from third party api
func FetchAllPosts(url string) ([]types.Post, error) {
	//http get method is used for fetching data from source
	//url := "https://jsonplaceholder.typicode.com/posts"
	fmt.Println("Making req to ->", url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	// required responseBody is stored in respData
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// an object of Post array will store the unmarshalled data
	//var resObj []types.Post
	err = json.Unmarshal(respData, &resObj)
	if err != nil {
		return nil, err
	}

	return resObj, nil
}

// Fetches all the comments for a post by taking post id as an argument
func FetchCommentForPostID(id int) ([]types.Comment, error) {
	//concatenating postId from argument to fetch comments
	//baseUrl := "https://jsonplaceholder.typicode.com/comments?postId=" // + strconv.Itoa(id)

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", id)

	if id > 100 {
		err := errors.New("Post id doesn't exist")
		return nil, err
	}
	fmt.Println("Making req to ->", url)
	//resC fetches all the comments using Get method
	resC, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	////error handling
	//ThrowError(err)

	defer resC.Body.Close()

	cData, err := ioutil.ReadAll(resC.Body)
	if err != nil {
		return nil, err
	}

	// an object of Comment array will store the unmarshalled data
	var cObj []types.Comment
	//Unmarshal data into cObj
	err = json.Unmarshal(cData, &cObj)
	if err != nil {
		return nil, err
	}

	return cObj, nil

}
