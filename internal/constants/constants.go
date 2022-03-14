package constants

const (
	DEFAULT_PORT = "8080"

	// The paths that will be handled by each handler
	DEFAULT_PATH       = "/corona/"
	CASES_PATH         = "/corona/v1/cases/"
	POLICY_PATH        = "/corona/v1/policy/"
	STATUS_PATH        = "/corona/v1/status/"
	NOTIFICATIONS_PATH = "/corona/v1/notifications/"

	POLICY_API_URL = "https://covidtrackerapi.bsg.ox.ac.uk/api/"

	/*
		// Got from: https://stackoverflow.com/questions/41085409/country-code-validation-with-iso
		ALPHA_CODE_REGEX = "/^A(BW|FG|GO|IA|L[AB]|ND|R[EGM]|SM|T[A\nFG]|U[ST]|ZE)|B(DI|E[LNS]|FA|G[DR]|H[RS]|IH|L[MRZ]|MU|" +
			"OL|\nR[ABN]|TN|VT|WA)|C(A[FN]|CK|H[ELN]|IV|MR|O[DGKLM]|PV|RI|U\n[BW]|XR|Y[MP]|ZE)|D(EU|JI|MA|NK|OM|ZA)|E(CU|" +
			"GY|RI|S[HPT]|\nTH)|F(IN|JI|LK|R[AO]|SM)|G(AB|BR|EO|GY|HA|I[BN]|LP|MB|N[B\nQ]|R[CDL]|TM|U[FMY])|H(KG|MD|ND|RV|" +
			"TI|UN)|I(DN|MN|ND|OT|R\n[LNQ]|S[LR]|TA)|J(AM|EY|OR|PN)|K(AZ|EN|GZ|HM|IR|NA|OR|WT)\n|L(AO|B[NRY]|CA|IE|KA|SO|" +
			"TU|UX|VA)|M(A[CFR]|CO|D[AGV]|EX|\nHL|KD|L[IT]|MR|N[EGP]|OZ|RT|SR|TQ|US|WI|Y[ST])|N(AM|CL|ER\n|FK|GA|I[CU]|LD|" +
			"OR|PL|RU|ZL)|OMN|P(A[KN]|CN|ER|HL|LW|NG|O\nL|R[IKTY]|SE|YF)|QAT|R(EU|OU|US|WA)|S(AU|DN|EN|G[PS]|HN|J\nM|L[BEV]|" +
			"MR|OM|PM|RB|SD|TP|UR|V[KN]|W[EZ]|XM|Y[CR])|T(C[A\nD]|GO|HA|JK|K[LM]|LS|ON|TO|U[NRV]|WN|ZA)|U(GA|KR|MI|RY|SA\n|" +
			"ZB)|V(AT|CT|EN|GB|IR|NM|UT)|W(LF|SM)|YEM|Z(AF|MB|WE)$/ix"
	*/
	ALPHA_CODE_REGEX = "/^[a-z]|[A-Z]{3}$/"
	YEAR_REGEX       = "/^(2019|202\\d)$/"
	MONTH_REGEX      = "/^(0?[1-9]|1[012])$/"
	DAY_REGEX        = "/^(0?[1-9]|[12]\\d|3[01])$/"
)
