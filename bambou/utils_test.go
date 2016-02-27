//+build !test

package bambou

/*
   Fake Exposed
*/

var fakeIdentity = Identity{"fake", "fakes"}

type fakeObjectsList []*fakeObject
type fakeObject struct{ ExposedObject }

func (o *fakeObject) Save() *Error   { return nil }
func (o *fakeObject) Delete() *Error { return nil }
func (o *fakeObject) Fetch() *Error  { return nil }

/*
   Fake Rootable
*/
var fakeRootIdentity = Identity{"root", "root"}

type fakeRootObject struct {
	fakeObject
	Token string `json:"APIKey,omitempty"`
}

func (o *fakeRootObject) APIKey() string       { return o.Token }
func (o *fakeRootObject) SetAPIKey(key string) { o.Token = key }
