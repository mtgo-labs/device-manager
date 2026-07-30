package device

type clientCredentials struct {
	apiID     int32
	apiHash   string
	packageID string
}

var clientCredentialsByName = make(map[string]clientCredentials)

func registerClientCredentials(name string, apiID int32, apiHash, packageID string) {
	clientCredentialsByName[name] = clientCredentials{
		apiID:     apiID,
		apiHash:   apiHash,
		packageID: packageID,
	}
}

func withClientCredentials(name string, profile Profile) Profile {
	credentials, ok := clientCredentialsByName[name]
	if !ok {
		return profile
	}
	profile.APIID = credentials.apiID
	profile.APIHash = credentials.apiHash
	if credentials.packageID != "" {
		profile.PackageID = credentials.packageID
	}
	return profile
}
