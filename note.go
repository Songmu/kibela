package kibela

import "fmt"

/*
   {
     "id": "QmxvZy8zNjY",
     "title": "APIテストpublic",
     "content": "コンテント!\nコンテント",
     "coediting": true,
     "folderName": "testtop/testsub1",
     "groups": [
       {
         "name": "Home",
         "id": "R3JvdXAvMQ"
       }
     ],
     "author": {
       "account": "Songmu"
     },
     "createdAt": "2019-06-23T16:54:09.447Z",
     "publishedAt": "2019-06-23T16:54:09.444Z",
     "contentUpdatedAt": "2019-06-23T16:54:09.445Z",
     "updatedAt": "2019-06-23T17:22:38.496Z"
   },
*/
type note struct {
	ID        `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CoEditing bool   `json:"coediting"`
	Folder    string `json:"folderName"`
	Groups    []struct {
		ID   `json:"id"`
		Name string `json:"name"`
	}
	Author struct {
		Account string `json:"account"`
	}
	UpdatedAt Time `json:"updatedAt"`
}

// .data/notes.totalCount
const totalCountQuery = `{
  notes() {
    totalCount
  }
}`

// .data.notes.nodes[]
func listNoteQuery(num int) string {
	return fmt.Sprintf(`{
  notes(first: %d) {
    nodes {
      id
      updatedAt
    }
  }
}`, num)
}

// .data.note
func getNoteQuery(id ID) string {
	return fmt.Sprintf(`{
  note(id: "%s") {
    title
    content
    coediting
    folderName
    groups {
      name
      id
    }
    author {
      account
    }
    updatedAt
  }
}`, string(id))
}
