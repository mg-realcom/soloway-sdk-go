package solowaysdk

const Host string = "https://dsp.soloway.ru"

type method string

const (
	Login              method = "/api/login"
	Whoami             method = "/api/whoami"
	PlacementsStat     method = "/api/placements_stat"
	PlacementStatByDay method = "/api/placements"
)
