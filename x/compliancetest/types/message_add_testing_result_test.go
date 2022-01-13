package types

import (
	"testing"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgAddTestingResult_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgAddTestingResult
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddTestingResult{
				Signer:                "invalid_address",
				SoftwareVersionString: testconstants.SoftwareVersionString,
				TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
				TestResult:            testconstants.TestResult,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vid is 0",
			msg: MsgAddTestingResult{
				Signer:                sample.AccAddress(),
				Vid:                   0,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
				TestResult:            testconstants.TestResult,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 0",
			msg: MsgAddTestingResult{
				Signer:                sample.AccAddress(),
				Vid:                   -1,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
				TestResult:            testconstants.TestResult,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgAddTestingResult{
				Signer:                sample.AccAddress(),
				Vid:                   65536,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
				TestResult:            testconstants.TestResult,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "pid is 0",
			msg: MsgAddTestingResult{
				Signer:                sample.AccAddress(),
				Pid:                   0,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
				TestResult:            testconstants.TestResult,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid < 0",
			msg: MsgAddTestingResult{
				Signer:                sample.AccAddress(),
				Pid:                   -1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
				TestResult:            testconstants.TestResult,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid > 65535",
			msg: MsgAddTestingResult{
				Signer:                sample.AccAddress(),
				Pid:                   65536,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
				TestResult:            testconstants.TestResult,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgAddTestingResult
		err  error
	}{{
		name: "valid address",
		msg: MsgAddTestingResult{
			Signer:                sample.AccAddress(),
			SoftwareVersionString: testconstants.SoftwareVersionString,
			Pid:                   1,
			Vid:                   1,
			TestDate:              testconstants.CertificationDate.Format(time.RFC3339),
			TestResult:            testconstants.TestResult,
		},
	},
	}

	for _, tt := range negative_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positive_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}
}
