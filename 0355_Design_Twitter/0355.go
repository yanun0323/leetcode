package main

import "container/heap"

// BUG: [0355] Timeout Passed

type Twitter struct {
	users map[int]*User
	time  int
}

func Constructor() Twitter {
	return Twitter{
		users: map[int]*User{},
		time:  0,
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	if _, exist := t.users[userId]; !exist {
		t.users[userId] = newUser()
	}
	t.users[userId].posts = append(t.users[userId].posts, newPost(tweetId, t.time))
	t.time++
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	if t.users[userId] == nil {
		return nil
	}
	follows := t.users[userId].follows

	h := &Heap{}
	heap.Init(h)
	for _, post := range t.users[userId].posts {
		heap.Push(h, post)
	}
	for follow := range follows {
		if t.users[follow] != nil {
			for _, post := range t.users[follow].posts {
				heap.Push(h, post)
			}
		}
	}
	result := []int{}
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(Post).ID)
		if len(result) == 10 {
			break
		}
	}
	return result
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if t.users[followerId] == nil {
		t.users[followerId] = newUser()
	}
	t.users[followerId].follows[followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if t.users[followerId].follows == nil || !t.users[followerId].follows[followeeId] {
		return
	}
	delete(t.users[followerId].follows, followeeId)
}

type User struct {
	posts   []Post
	follows map[int]bool
}

func newUser() *User {
	return &User{
		posts:   []Post{},
		follows: map[int]bool{},
	}
}

type Post struct {
	ID   int
	Time int
}

func newPost(id, time int) Post {
	return Post{ID: id, Time: time}
}

type Heap struct {
	posts []Post
}

func (h *Heap) Len() int {
	return len(h.posts)
}
func (h *Heap) Less(i, j int) bool {
	return h.posts[i].Time > h.posts[j].Time
}
func (h *Heap) Swap(i, j int) {
	h.posts[i], h.posts[j] = h.posts[j], h.posts[i]
}
func (h *Heap) Push(x interface{}) {
	h.posts = append(h.posts, x.(Post))
}
func (h *Heap) Pop() interface{} {
	x := h.posts[h.Len()-1]
	h.posts = h.posts[:h.Len()-1]
	return x
}
