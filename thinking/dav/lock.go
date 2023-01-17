package dav

type davLock struct {
	holders []*string
	shared  bool
	token   *string
}

func Lock(user, token *string) {

}

func Unlock() {

}
