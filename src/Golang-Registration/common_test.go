package main

var tmpUserList []user

// existing code (not shown)
// .
// .
// .

func saveLists() {
	tmpUserList = userList
	tmpArticleList = articleList
}

func restoreLists() {
	userList = tmpUserList
	articleList = tmpArticleList
}
