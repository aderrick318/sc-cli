package lib

var productionRootUrl string = "https://support.pioneerrx.com"
var betaRootUrl string = "https://beta.pioneerrx.com"

func GetSupportCenterRootPath(isBeta bool) string {
	rootUrl := ""
	if isBeta {
		rootUrl = betaRootUrl
	} else {
		rootUrl = productionRootUrl
	}

	return rootUrl
}