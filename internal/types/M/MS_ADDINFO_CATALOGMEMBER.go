package win

//ref DWORD
//ref CRYPTCATSTORE
//ref CRYPTCATMEMBER

type MS_ADDINFO_CATALOGMEMBER struct {
	CbStruct DWORD
	PStore   *CRYPTCATSTORE
	PMember  *CRYPTCATMEMBER
}
