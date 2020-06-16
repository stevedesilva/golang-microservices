package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
     /*
        Test Marshalling take input interface and attempt to create a json string ( obj -> json )
    */
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang introductory respository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
    }
   
	bytes, err := json.Marshal(request)

	assert.Nil(t, err) // not error in marshalling
	assert.NotNil(t, bytes) // we have json string


    /*
        Test Unmarshalling
    */
	var target CreateRepoRequest // struct to populate
    err = json.Unmarshal(bytes, &target) // json string to target struct
    
    // compare to original request
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
    assert.EqualValues(t, target.Description, request.Description)
    assert.EqualValues(t, target.Homepage, request.Homepage)
    assert.True(t, target.Private)
    assert.True(t, target.HasProjects)
    assert.True(t, target.HasIssues)
    assert.True(t, target.HasWiki)
}

// TODO test response
