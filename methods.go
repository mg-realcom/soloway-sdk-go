package solowaysdk

const Host string = "https://dsp.soloway.ru"

type method string

const Login method = "/api/login"
const Whoami method = "/api/whoami"
const PlacementsStat method = "/api/placements_stat"
const PlacementStatByDay method = "/api/placements"
