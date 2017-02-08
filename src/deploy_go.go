package DeployCli;

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"os"
	"crypto/sha256"
	"encoding/hex"
)

var (
	ApiEndpoint string
	ApiEmail string
	ApiPass string
)


func SetAPIInfo(name string, key string, endpoint string) string {
	ApiEmail = name
	ApiPass = key
	ApiEndpoint = endpoint
	return "Storj login creds set"
}


func GetUser() string {
	return ApiEmail
}



func GetFrames() ([]Frame, interface{}){
	outgoing, err := SendRequest("GET", "frames", string(""))
	var frames []Frame
	err = json.Unmarshal([]byte(outgoing), &frames)
	if err!=nil {
		panic(err)
	}
	return frames, err
}


func CreateFrame() (Frame, interface{}){
	outgoing, err := SendRequest("POST", "frames", string(""))
	var frames Frame
	err = json.Unmarshal([]byte(outgoing), &frames)
	if err!=nil {
		panic(err)
	}
	return frames, err
}


func DestroyFrame(id string) (bool, interface{}){
	statsCode, err := SendRequestStatusCode("DELETE", "frames/"+id, string(""))
	if err!=nil {
		panic(err)
	}
	if statsCode==204 {
		return true, err
	} else {
		return true, err
	}
}


func GetBuckets() ([]Bucket, interface{}){
	fmt.Print(ApiEmail,ApiPass)
	outgoing, err := SendRequest("GET", "buckets", string(""))
	fmt.Print(outgoing)
	var buckets []Bucket
	err = json.Unmarshal([]byte(outgoing), &buckets)
	if err!=nil {
		panic(err)
	}
	return buckets, err
}



func CreateBucket(name string, keys Pubkeys) (Bucket, interface{}){
	var thisReq = map[string]interface{}{"pubkeys": keys.keys, "name": name}
	jsoned, _ := json.Marshal(thisReq)
	outgoing, err := SendRequest("POST", "buckets", string(jsoned))
	var bucket Bucket
	err = json.Unmarshal([]byte(outgoing), &bucket)
	if err!=nil {
		panic(err)
	}
	return bucket, err
}



func GetBucket(id string) (Bucket, interface{}){
	outgoing, err := SendRequest("GET", "buckets/"+id, string(""))
	var bucket Bucket
	err = json.Unmarshal([]byte(outgoing), &bucket)
	if err!=nil {
		panic(err)
	}
	return bucket, err
}



func UploadFile(frame Frame, bucket Bucket, file string) (File, interface{}){
	extraParams := FileStore{Frame: frame.ID, Mimetype: "image/png", Filename: "cat.png"}
	output, err := UploadFileStore(extraParams, "buckets/"+bucket.ID+"/files", file)
	fmt.Print(output)
	var filein File
	//err = json.Unmarshal([]byte(output), &filein)
	//if err!=nil {
	//	panic(err)
	//}
	return filein, err
}



func DestroyBucket(bucket Bucket) (bool, interface{}){
	goTo := "buckets/"+bucket.ID
	delCode, err := SendRequestStatusCode("DELETE", goTo, string(""))
	if err!=nil {
		panic(err)
	}

	if delCode==204 {
		return true, err
	} else {
		return true, err
	}

}



func SendRequest(method string, url string, input string) (string, interface{}) {
	sendUrl := ApiEndpoint+url
	req, err := http.NewRequest(method, sendUrl, bytes.NewBuffer([]byte(input)))
	req.SetBasicAuth(ApiEmail,ApiPass)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	//fmt.Printf(string(body))

	return string(body), err
}



func SendRequestStatusCode(method string, url string, input string) (int, interface{}) {
	sendUrl := ApiEndpoint+url
	req, err := http.NewRequest(method, sendUrl, bytes.NewBuffer([]byte(input)))
	req.SetBasicAuth(ApiEmail,ApiPass)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	fmt.Printf(string(body))
	return resp.StatusCode, err
}



func GetNewVersion() string {
	url := "https://raw.githubusercontent.com/hunterlong/storj-go/master/version.txt"
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	//fmt.Printf(string(body))
	return string(body)

}


func UploadFileStore(fileStore FileStore, url string, file string) (string, error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	//fw, err := w.CreateFormFile("file", file)
	//if err != nil {
	//	return "", err
	//}
	//if _, err = io.Copy(fw, f); err != nil {
	//	return "", err
	//}


	var jsonStr = []byte(`{"frame":"`+fileStore.Frame+`", "mimetype":"`+fileStore.Mimetype+`", "filename":"`+fileStore.Filename+`"}`)
	b.Write(jsonStr)
	// Add the other fields
	//if fw, err = w.CreateFormField("file"); err != nil {
	//	return "", err
	//}
	//if _, err = fw.Write(jsonStr); err != nil {
	//	return "", err
	//}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()


	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	req.Header.Set("Content-Type", "application/json")
	fmt.Print(req)
	if err != nil {
		return "", err
	}
	// Don't forget to set the content type, this will contain the boundary.
	//req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	fmt.Print(res)
	if err != nil {
		return "", err
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}

	body, _ := ioutil.ReadAll(res.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	fmt.Printf(string(body))
	return string(body), err
}


func EncryptPassword(password string) string {
	aStringToHash := []byte(password)
	sha256Bytes := sha256.Sum256(aStringToHash)
	encrypted := hex.EncodeToString(sha256Bytes[:])
	return encrypted
}