package DeployCli;

import "time"

type EmptyStruct struct {
	Name string `json:"name"`
}

type Buckets struct {
	Buckets []Bucket
}


type Frame struct {
	Created time.Time `json:"created"`
	ID string `json:"id"`
}


type Bucket struct {
	Storage int `json:"storage"`
	Transfer int `json:"transfer"`
	Status string `json:"status"`
	Pubkeys []string `json:"pubkeys"`
	User string `json:"user"`
	Name string `json:"name"`
	Created time.Time `json:"created"`
	ID string `json:"id"`
}


type FileStore struct {
	Frame string `json:"frame"`
	Mimetype string `json:"mimetype"`
	Filename string `json:"filename"`
}

type File struct {
	ID string `json:"id"`
	Bucket string `json:"bucket"`
	Mimetype string `json:"mimetype"`
	Filename string `json:"filename"`
	Size int `json:"size"`
}


type Pubkeys struct {
	keys []string `json:"pubkeys"`
}