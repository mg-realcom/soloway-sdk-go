package solowaysdk

import "encoding/json"

type AccountInfo struct {
	Username string     `json:"username"`
	Client   ClientInfo `json:"client"`
	Agency   AgencyInfo `json:"agency"`
	User     UserInfo   `json:"user"`
}

type ClientInfo struct {
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	PartnerGuid interface{} `json:"partner_guid"`
	Guid        string      `json:"guid"`
}

type AgencyInfo struct {
	Name                        string      `json:"name"`
	AllowBmpAudit               int         `json:"allow_bmp_audit"`
	DmpAgencyGuid               string      `json:"dmp_agency_guid"`
	GlobalSharedModels          int         `json:"global_shared_models"`
	AllowAlerts                 int         `json:"allow_alerts"`
	SegmentsDumpsExist          int         `json:"segments_dumps_exist"`
	AllowAdraudPlacements       bool        `json:"allow_adraud_placements"`
	AllowClusterOrders          int         `json:"allow_cluster_orders"`
	AllowSegmentDump            int         `json:"allow_segment_dump"`
	AllowExtSegments            int         `json:"allow_ext_segments"`
	AdrUserId                   int         `json:"adr_user_id"`
	AllowAdraudAutoimport       bool        `json:"allow_adraud_autoimport"`
	AllowNativeCreatives        int         `json:"allow_native_creatives"`
	AllowAllAuditory            bool        `json:"allow_all_auditory"`
	AllowInappSites             int         `json:"allow_inapp_sites"`
	AllowExternalAnalytics      int         `json:"allow_external_analytics"`
	AllowForecastBs             bool        `json:"allow_forecast_bs"`
	AllowVpaid4Vast             bool        `json:"allow_vpaid4vast"`
	AllowCopyPlacements         int         `json:"allow_copy_placements"`
	AllowVastWrapperCreatives   int         `json:"allow_vast_wrapper_creatives"`
	AllowReports                int         `json:"allow_reports"`
	AllowSzAlias                int         `json:"allow_sz_alias"`
	AllowOrderSegmentDump       int         `json:"allow_order_segment_dump"`
	Timestamp                   int         `json:"timestamp"`
	DdEnableSkip                int         `json:"dd_enable_skip"`
	GlobalIndModels             int         `json:"global_ind_models"`
	AllowVastPlacements         int         `json:"allow_vast_placements"`
	AllowForecastIostream       bool        `json:"allow_forecast_iostream"`
	DdEnable                    int         `json:"dd_enable"`
	Email                       string      `json:"email"`
	DdAdrPools                  interface{} `json:"dd_adr_pools"`
	AllowAnalyzeTraffic         bool        `json:"allow_analyze_traffic"`
	AllowDmp                    int         `json:"allow_dmp"`
	AllowMicroAnalytics         bool        `json:"allow_micro_analytics"`
	LcUtmParams                 interface{} `json:"lc_utm_params"`
	AllowOwnExtSegments         int         `json:"allow_own_ext_segments"`
	AllowExtraTarget            bool        `json:"allow_extra_target"`
	ServiceType                 string      `json:"service_type"`
	AllowAuditPlacements        int         `json:"allow_audit_placements"`
	AllowAnalytics              int         `json:"allow_analytics"`
	Guid                        string      `json:"guid"`
	AllowSolowayAnalytics       int         `json:"allow_soloway_analytics"`
	UtmUrlUseDefault            int         `json:"utm_url_use_default"`
	RefundablePartnerCommission int         `json:"refundable_partner_commission"`
	GlobalAuditory              int         `json:"global_auditory"`
	AllowGbPlacements           bool        `json:"allow_gb_placements"`
	AllowForecaster             bool        `json:"allow_forecaster"`
	AllowInappPlacements        bool        `json:"allow_inapp_placements"`
	AllowTrackers               int         `json:"allow_trackers"`
	SingleClient                int         `json:"single_client"`
	AllowMicroSegments          bool        `json:"allow_micro_segments"`
	AllowVpaidCreatives         int         `json:"allow_vpaid_creatives"`
	AllowSzAutoimport           int         `json:"allow_sz_autoimport"`
	UtmUrlParams                interface{} `json:"utm_url_params"`
	LcUtmUseDefault             int         `json:"lc_utm_use_default"`
	AllowScreenshots            int         `json:"allow_screenshots"`
	AllowVpaidGbCreatives       int         `json:"allow_vpaid_gb_creatives"`
	PartnerGuid                 string      `json:"partner_guid"`
}

type UserInfo struct {
	Email       string            `json:"email"`
	Resources   PlacementResource `json:"resources"`
	Measures    []string          `json:"measures"`
	Login       string            `json:"login"`
	Guid        string            `json:"guid"`
	HideFinance int               `json:"hide_finance"`
	Dimensions  []string          `json:"dimensions"`
	Name        string            `json:"name"`
	Enable      int               `json:"enable"`
	Type        string            `json:"type"`
}

type PlacementResource struct {
	PlacementOs             string `json:"placement_os"`
	PlacementUniqLimits     string `json:"placement_uniq_limits"`
	PlacementLimits         string `json:"placement_limits"`
	PlacementWhitelists     string `json:"placement_whitelists"`
	PlacementTraffics       string `json:"placement_traffics"`
	PlacementOmp            string `json:"placement_omp"`
	PlacementSettings       string `json:"placement_settings"`
	PlacementDevice         string `json:"placement_device"`
	PlacementDatesWeekHours string `json:"placement_dates_week_hours"`
	PlacementGeo            string `json:"placement_geo"`
}

type PlacementsInfo struct {
	List []Placement `json:"list"`
}

type Placement struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Doc  struct {
		SrcPlacementGuid interface{} `json:"src_placement_guid"`
		Budget           int64       `json:"budget"`
		AgencyGuid       string      `json:"agency_guid"`
		TrafficType      string      `json:"traffic_type"`
		ExtraTargetGuid  interface{} `json:"extra_target_guid"`
		UniformDist      int         `json:"uniform_dist"`
		BotStopPercent   int         `json:"bot_stop_percent"`
		DefaultUrl       string      `json:"default_url"`
		Limits           struct {
			LimitClkDay   int `json:"limit_clk_day"`
			LimitClkTotal int `json:"limit_clk_total"`
		} `json:"limits"`
		OaAdrProfileId        int         `json:"oa_adr_profile_id"`
		BlockAnonymousTraffic int         `json:"block_anonymous_traffic"`
		RealStopDate          *string     `json:"real_stop_date"`
		NextExtraTargetGuid   interface{} `json:"next_extra_target_guid"`
		AutoRenewal           int         `json:"auto_renewal"`
		TargetGuid            string      `json:"target_guid"`
		RealStartDate         string      `json:"real_start_date"`
		StartTime             string      `json:"start_time"`
		Timestamp             int         `json:"timestamp"`
		ExtSegmentsUsed       int         `json:"ext_segments_used"`
		MaxClickLossRate      interface{} `json:"max_click_loss_rate"`
		StopTime              string      `json:"stop_time"`
		Name                  string      `json:"name"`
		Omp                   struct {
			LimitExpTotal int `json:"limit_exp_total"`
			LimitExpDay   int `json:"limit_exp_day"`
		} `json:"omp"`
		Enable            int         `json:"enable"`
		SiteGuid          string      `json:"site_guid"`
		BotStopStopWeight int         `json:"bot_stop_stop_weight"`
		Balance           int         `json:"balance"`
		Context           string      `json:"context"`
		Type              string      `json:"type"`
		Discount          int         `json:"discount"`
		Costs             int64       `json:"costs"`
		NextTargetGuid    interface{} `json:"next_target_guid"`
		DayBudget         int         `json:"day_budget"`
		DmpPlacementGuid  string      `json:"dmp_placement_guid"`
		Auditory          struct {
			Cst struct {
				Enable int           `json:"enable"`
				Cats   []interface{} `json:"cats"`
			} `json:"cst"`
			All json.Number `json:"all"`
			Aut struct {
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
			} `json:"aut"`
			CstInv struct {
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
			} `json:"cst_inv"`
			ExtInv struct {
				External struct {
				} `json:"external"`
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
			} `json:"ext_inv"`
			Socdem struct {
				GenderEnable int           `json:"gender_enable"`
				AgeEnable    int           `json:"age_enable"`
				AgeInvert    int           `json:"age_invert"`
				GenderInvert int           `json:"gender_invert"`
				Cats         []interface{} `json:"cats"`
				Enable       int           `json:"enable"`
			} `json:"socdem"`
			Ndr struct {
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
			} `json:"ndr"`
			Cls struct {
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
			} `json:"cls"`
			Ext struct {
				Enable   int `json:"enable"`
				External struct {
				} `json:"external"`
				Cats []interface{} `json:"cats"`
			} `json:"ext"`
			Sgm struct {
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
				Volume struct {
				} `json:"volume"`
			} `json:"sgm"`
			Ind struct {
				Enable int      `json:"enable"`
				Cats   []string `json:"cats"`
			} `json:"ind"`
			Smd struct {
				Volume struct {
				} `json:"volume"`
				Enable int           `json:"enable"`
				Cats   []interface{} `json:"cats"`
			} `json:"smd"`
			RtgInv struct {
				Enable int           `json:"enable"`
				Cats   []interface{} `json:"cats"`
			} `json:"rtg_inv"`
			Lal struct {
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
				Volume struct {
				} `json:"volume"`
			} `json:"lal"`
			Rtg struct {
				Enable int      `json:"enable"`
				Cats   []string `json:"cats"`
			} `json:"rtg"`
			Mcr struct {
				External struct {
				} `json:"external"`
				Cats   []interface{} `json:"cats"`
				Enable int           `json:"enable"`
			} `json:"mcr"`
			Csm struct {
				Enable   int `json:"enable"`
				External struct {
				} `json:"external"`
				Cats []interface{} `json:"cats"`
			} `json:"csm"`
			NdrInv struct {
				Enable int           `json:"enable"`
				Cats   []interface{} `json:"cats"`
			} `json:"ndr_inv"`
		} `json:"auditory"`
		IndModelGuid interface{} `json:"ind_model_guid"`
		Target       struct {
			Type string `json:"type"`
			Id   string `json:"id"`
			Doc  struct {
				Guid           string `json:"guid"`
				Timestamp      int    `json:"timestamp"`
				CheckpointGuid string `json:"checkpoint_guid"`
				CheckpointName string `json:"checkpoint_name"`
				RealPrice      int    `json:"real_price"`
				Price          int    `json:"price"`
			} `json:"doc"`
		} `json:"target"`
		HasVpaid4Vast       int         `json:"has_vpaid4vast"`
		AdStopCatId         int         `json:"ad_stop_cat_id"`
		TrafficAuditory     int         `json:"traffic_auditory"`
		IsGb                bool        `json:"is_gb"`
		SharedModelGuid     interface{} `json:"shared_model_guid"`
		MaxCpm              int         `json:"max_cpm"`
		MaxBounceRate       interface{} `json:"max_bounce_rate"`
		MinCtr              interface{} `json:"min_ctr"`
		NextSharedModelGuid interface{} `json:"next_shared_model_guid"`
		NextIndModelGuid    interface{} `json:"next_ind_model_guid"`
		BotStopStartWeight  int         `json:"bot_stop_start_weight"`
		MinCpm              int         `json:"min_cpm"`
		ClientGuid          string      `json:"client_guid"`
		AdrAdId             int         `json:"adr_ad_id"`
		DefaultPixels       struct {
			Dcm      []interface{} `json:"dcm"`
			Sizmek   []interface{} `json:"sizmek"`
			TmplHtml []interface{} `json:"tmpl_html"`
			Native   []interface{} `json:"native"`
			Tgb      []interface{} `json:"tgb"`
			Flash    []interface{} `json:"flash"`
		} `json:"default_pixels"`
		MinLeadRate   interface{} `json:"min_lead_rate"`
		Archive       int         `json:"archive"`
		Guid          string      `json:"guid"`
		NativeBanners int         `json:"native_banners"`
	} `json:"doc"`
}

func (pI *PlacementsInfo) ToMap() (placements map[string]Placement) {
	placements = map[string]Placement{}
	for _, place := range pI.List {
		placements[place.Id] = place
	}
	return placements
}

type ReqPlacementsStat struct {
	PlacementIds []string `json:"placement_ids"`
	StartDate    string   `json:"start_date"`
	StopDate     string   `json:"stop_date"`
	WithArchived int      `json:"with_archived"`
}

type PlacementsStatByDay struct {
	List []struct {
		Clicks      int    `json:"clicks"`
		Cost        int    `json:"cost"`
		PlacementId string `json:"placement_id"`
		Exposures   int    `json:"exposures"`
		Date        string `json:"date"`
	} `json:"list"`
	Total struct {
		ReachRising int `json:"reach_rising"`
		Checkpoints struct {
		} `json:"checkpoints"`
		Clicks    int `json:"clicks"`
		Reach     int `json:"reach"`
		Exposures int `json:"exposures"`
		Cost      int `json:"cost"`
	} `json:"total"`
	FormulaLog []struct {
		TargetGuid          string      `json:"target_guid"`
		SharedModelLeadRate interface{} `json:"shared_model_lead_rate"`
		PidLeadUsed         int         `json:"pid_lead_used"`
		AucLead             string      `json:"auc_lead"`
		IndModelGuid        string      `json:"ind_model_guid"`
		AucComplete         interface{} `json:"auc_complete"`
		PidClickUsed        int         `json:"pid_click_used"`
		SharedModelGuid     interface{} `json:"shared_model_guid"`
		IndModelLeadRate    interface{} `json:"ind_model_lead_rate"`
		IndModelAucLead     string      `json:"ind_model_auc_lead"`
		Timestamp           string      `json:"timestamp"`
		SharedModelAucLead  interface{} `json:"shared_model_auc_lead"`
		FormulaId           int         `json:"formula_id"`
		CheckpointGuid      string      `json:"checkpoint_guid"`
		ClickModelDefault   int         `json:"click_model_default"`
		AucClick            string      `json:"auc_click"`
	} `json:"formula_log"`
}
