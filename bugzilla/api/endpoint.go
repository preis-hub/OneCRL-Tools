package api

// An Endpoint encodes what method, status return code, and
// REST resource that that Endpoint responds to.
//
// The pattern that implementors will see is that of declaring a struct
// that both implements this interface and can be JSON serialized to
// the expected input of that endpoint (if any).
//
// An example implementation may be...
//
//	type UpdateBug struct {
//		BugId      int `json:"-"` // This is not in the Bugzilla API, it is for building the resource
//		api.Update
//		api.Ok
//		Assignee   string `json:"assignee"`
//		Status     string `json:"status"`
//		Resolution string `json:"resolution"`
//		...
//	}
//
//	func (u *UpdateBug) Resource() string {
//		return fmt.Sprintf("/bug/%d", u.BugId)
//	}
//
type Endpoint interface {
	Methoder
	Expecter
	Resourcer
}
