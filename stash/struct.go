package stashClient

type Pr struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	State       string     `json:"state"`
	Open        bool       `json:"open"`
	Closed      bool       `json:"closed"`
	FromRef     Ref        `json:"fromRef"`
	ToRef       Ref        `json:"toRef"`
	Locked      bool       `json:"locked"`
	Reviewers   []Reviewer `json:"reviewers"`
}

type Ref struct {
	ID         string `json:"id"`
	Repository Repo   `json:"repository"`
}

type Repo struct {
	Slug    string      `json:"slug"`
	Name    interface{} `json:"name"`
	Project Proj        `json:"project"`
}

type Proj struct {
	Key string `json:"key"`
}

type Reviewer struct {
	User Usr `json:"user"`
}

type Usr struct {
	Name string `json:"name"`
}

type PrRes struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
	Open        bool   `json:"open"`
	Closed      bool   `json:"closed"`
	CreatedDate int    `json:"createdDate"`
	UpdatedDate int    `json:"updatedDate"`
	FromRef     struct {
		ID              string `json:"id"`
		DisplayID       string `json:"displayId"`
		LatestChangeset string `json:"latestChangeset"`
		LatestCommit    string `json:"latestCommit"`
		Repository      struct {
			Slug          string `json:"slug"`
			ID            int    `json:"id"`
			Name          string `json:"name"`
			ScmID         string `json:"scmId"`
			State         string `json:"state"`
			StatusMessage string `json:"statusMessage"`
			Forkable      bool   `json:"forkable"`
			Project       struct {
				Key         string `json:"key"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Public      bool   `json:"public"`
				Type        string `json:"type"`
				Link        struct {
					URL string `json:"url"`
					Rel string `json:"rel"`
				} `json:"link"`
				Links struct {
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"project"`
			Public   bool   `json:"public"`
			CloneURL string `json:"cloneUrl"`
			Link     struct {
				URL string `json:"url"`
				Rel string `json:"rel"`
			} `json:"link"`
			Links struct {
				Clone []struct {
					Href string `json:"href"`
					Name string `json:"name"`
				} `json:"clone"`
				Self []struct {
					URL string `json:"url"`
					Rel string `json:"rel"`
				} `json:"self"`
			} `json:"links"`
		} `json:"repository"`
	} `json:"fromRef"`
	ToRef struct {
		ID              string `json:"id"`
		DisplayID       string `json:"displayId"`
		LatestChangeset string `json:"latestChangeset"`
		LatestCommit    string `json:"latestCommit"`
		Repository      struct {
			Slug          string `json:"slug"`
			ID            int    `json:"id"`
			Name          string `json:"name"`
			ScmID         string `json:"scmId"`
			State         string `json:"state"`
			StatusMessage string `json:"statusMessage"`
			Forkable      bool   `json:"forkable"`
			Project       struct {
				Key         string `json:"key"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Public      bool   `json:"public"`
				Type        string `json:"type"`
				Link        struct {
					URL string `json:"url"`
					Rel string `json:"rel"`
				} `json:"link"`
				Links struct {
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"project"`
			Public   bool   `json:"public"`
			CloneURL string `json:"cloneUrl"`
			Link     struct {
				URL string `json:"url"`
				Rel string `json:"rel"`
			} `json:"link"`
			Links struct {
				Clone []struct {
					Href string `json:"href"`
					Name string `json:"name"`
				} `json:"clone"`
				Self []struct {
					URL string `json:"url"`
					Rel string `json:"rel"`
				} `json:"self"`
			} `json:"links"`
		} `json:"repository"`
	} `json:"toRef"`
	Locked bool `json:"locked"`
	Author struct {
		User struct {
			Name         string `json:"name"`
			EmailAddress string `json:"emailAddress"`
			ID           int    `json:"id"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			Slug         string `json:"slug"`
			Type         string `json:"type"`
		} `json:"user"`
		Role     string `json:"role"`
		Approved bool   `json:"approved"`
	} `json:"author"`
	Reviewers []struct {
		User struct {
			Name         string `json:"name"`
			EmailAddress string `json:"emailAddress"`
			ID           int    `json:"id"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			Slug         string `json:"slug"`
			Type         string `json:"type"`
		} `json:"user"`
		Role     string `json:"role"`
		Approved bool   `json:"approved"`
	} `json:"reviewers"`
	Participants []interface{} `json:"participants"`
	Link         struct {
		URL string `json:"url"`
		Rel string `json:"rel"`
	} `json:"link"`
	Links struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}
