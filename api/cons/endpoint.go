package cons

import (
	"strings"
)

type OnseReplaceString string

func (s OnseReplaceString) Replace(matchStr string) string {
	return strings.Replace(string(s), "{1}", matchStr, 1)
}

type DoubleReplaceString string

func (s DoubleReplaceString) Replace(firtstMatchStr, secondMatchStr string) string {
	str := strings.Replace(string(s), "{1}", firtstMatchStr, 1)
	return strings.Replace(str, "{2}", secondMatchStr, 1)
}

const (
	URICase  = "api/case"
	URIAlert = "api/alert"

	URIConnector       = "api/connector"
	URICortex          = URIConnector + "/cortex"
	URICortexAction    = URICortex + "/action"
	URICortexResponder = URICortex + "/responder"

	URIMisp = URIConnector + "/misp"

	URIAPIv0      = "api/v0"
	URIAPIv0Query = URIAPIv0 + "/query"
	URIAPIv0Case  = URIAPIv0 + "/case"

	URIAPIv1      = "api/v1"
	URIAPIv1Query = URIAPIv1 + "/query"
)

const (
	URICaseId = OnseReplaceString(URICase + "/{1}")
	URITaskId = OnseReplaceString(URICase + "/task/{1}")
	URILogId  = OnseReplaceString(URICase + "/task/log/{1}")

	URIAlertId = OnseReplaceString(URIAlert + "/{1}")

	URIAlertObservable = OnseReplaceString(URIAlert + "/artifact/{1}")
	URIObservableId    = OnseReplaceString(URIAlert + "/{1}/artifact")

	URICortexActionEntityId    = DoubleReplaceString(URICortexAction + "/{1}/{2}")
	URICortexResponderEntityId = DoubleReplaceString(URICortexResponder + "/{1}/{2}")

	URIMergeAlert = DoubleReplaceString(URIAlert + "/{1}/merge/{2}")
	URIMergeCase  = DoubleReplaceString(URIAPIv0 + "/{1}/_merge/{2}")

	URIMispExport = DoubleReplaceString(URIMisp + "/{1}/{2}")
)

//То что не реализовано ввиде констант ниже

/*
/api/case/{id}?force=1
/api/case/{id}/links
/api/case/{id}/task

/api/alert/{id}/createCase
/api/alert/{id}/markAsRead
/api/alert/{id}/markAsUnread

/api/v0/query?name

/api/v1/query?name=case-task-logs
/api/v1/query?name=alerts
/api/v1/query?name=alert-similar-cases
*/
