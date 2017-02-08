package DeployCli;

import (
	"testing"
	"os"
)


var newBucket Bucket
var newFrame Frame

func TestSetAPIInfo(t *testing.T) {
	ApiEmail = os.Getenv("STORJ_EMAIL")
	ApiPass = os.Getenv("STORJ_PASS")
	ApiEndpoint = "https://api.storj.io/"
	SetAPIInfo(ApiEmail,ApiPass, ApiEndpoint)
	t.Log("Storj API Keys have been set! \n")
}



func TestCreateFrame(t *testing.T) {
	newFrame, _ = CreateFrame()
	t.Log("New Frame Created: ",newFrame.ID, newFrame.Created)
}


func TestGetFrames(t *testing.T) {
	myFrames, _ := GetFrames()
	t.Log("I only have ",myFrames, "frames")
	t.Log("Getting Frames: ",myFrames)
}


func TestGetBuckets(t *testing.T) {
	myBuckets, err := GetBuckets()
	if err!=nil {
		t.Fail()
	}
	t.Log("I've got some buckets. Only", len(myBuckets))
	for _, v := range myBuckets {
		t.Log("Bucket:", v.Name, " | Status:", v.Status, " | Size: ",v.Storage,"\n")
	}
}


func TestCreateBucket(t *testing.T) {
	keys := Pubkeys{keys:[]string{"025d77a9bec982f1d66327d5e0840e071d5f7b9856877abb16be8159e35a361636"}}
	newBucket, _ = CreateBucket("devStorj", keys)
	if newBucket.Status!="Active" {
		t.Fail()
	}
	t.Log("New Bucket created: ",newBucket)
}



func TestGetBucket(t *testing.T) {
	thisBucket, _ := GetBucket(newBucket.ID)
	if thisBucket.Status!="Active" {
		t.Fail()
	}
	t.Log("Getting Bucket from ID: ",thisBucket.ID)
}



func TestUploadToBucket(t *testing.T) {
	toUpload := "cat.png"
	thisFile, _ := UploadFile(newFrame, newBucket, toUpload)
	if thisFile.Size < 0 {
		t.Fail()
	}
	t.Log("Successfully Uploaded File '",toUpload,"' into Bucket ",newBucket.Name)
}



func TestDestroyFrame(t *testing.T) {
	t.Skip()
	oldFrame, _ := DestroyFrame(newFrame.ID)
	t.Log("Goodbye Frame, you are now deleted: ",oldFrame)
}



func TestDestroyBucket(t *testing.T) {
	t.Skip()
	success, err := DestroyBucket(newBucket)
	if err!=nil {
		t.Fail()
	}
	if !success {
		t.Fail()
	}
	t.Log("Bucket was Deleted: ", newBucket.Name)
}