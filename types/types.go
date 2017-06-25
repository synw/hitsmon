package types

type Db struct {
	Type  string
	Addr  string
	Host  string
	Port  int
	User  string
	Pwd   string
	Name  string
	Table string
}

type Conf struct {
	Db        *Db
	Frequency int
	Domain    string
	Separator string
	Dev       bool
	Verb      int
}

type Hit struct {
	Id              string
	Domain          string
	Path            string
	Method          string
	Ip              string
	UserAgent       string `gorm:"column:user_agent"`
	IsAuthenticated string
	IsStaff         string
	IsSuperuser     string
	User            string
	Referer         string
	View            string
	Module          string
	StatusCode      string
	ReasonPhrase    string
	RequestTime     string
	ContentLength   string
	NumQueries      string
	QueriesTime     string
}
