package blclient

const (
	cmdHandShake         = 0x55
	cmdGetBootInfo       = 0x10
	cmdLoadBootHeader    = 0x11
	cmdLoadSegmentHeader = 0x17
	cmdLoadSegmentData   = 0x18
	cmdCheckImage        = 0x19
	cmdRunImage          = 0x1A
	cmdMemWrite          = 0x50
	cmdReadJedecid       = 0x36
	cmdFlashErase        = 0x30
	cmdFlashWrite        = 0x31
	cmdFlashWriteCHeck   = 0x3A
	cmdXipReadStart      = 0x60
	cmdFlashXipReadsha   = 0x3E
	cmdXipReadFinish     = 0x61
	cmdEfuseReadMacAddr  = 0x42
)
