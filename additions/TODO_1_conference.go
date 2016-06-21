func getProfileFromUser(r *http.Request) (*Profile, *datastore.Key, error) {
	//Return user Profile from datastore, creating new one if non-existent.
	//TODO
	//make sure user is authed
	c := endpoints.NewContext(r)
	user, err := endpoints.CurrentUser(c, []string{endpoints.EmailScope},
		[]string{WEB_CLIENT_ID, endpoints.APIExplorerClientID}, []string{WEB_CLIENT_ID, endpoints.APIExplorerClientID})
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, endpoints.UnauthorizedError
	}
	//get Profile from datastore
	userId := getUserId(user, "")
	appCtx := appengine.NewContext(r)
	key := datastore.NewKey(appCtx, "Profile", userId, 0, nil)
	var profile Profile
	//TODO
	if err != nil && err != datastore.ErrNoSuchEntity {
		return nil, nil, err
	}
	if err == datastore.ErrNoSuchEntity {
		profile = Profile{
			DisplayName: user.String(),
			MainEmail: user.Email,
			TeeShirtSize: TeeShirtSizeToStringEnum(NOT_SPECIFIED),
		}
		//TODO
	}
	return &profile, key, nil
}

func doProfile(r *http.Request, saveRequest *ProfileMiniForm) (*ProfileForm, error) {
	//Get user Profile and return to user, possibly updating it first.
	//get user Profile
	prof, key, err := getProfileFromUser(r)
	if err != nil {
		return nil, err
	}
	
	//if saveProfile(), process user-modifyable fields
	if saveRequest != nil {
		prof.TeeShirtSize = TeeShirtSizeToStringEnum(saveRequest.TeeShirtSize)
		prof.DisplayName = saveRequest.DisplayName
		appCtx := appengine.NewContext(r)
		//TODO
		if err != nil {
			return nil, err
		}
	}
	
	//return ProfileForm
	return copyProfileToForm(r, prof)
}

