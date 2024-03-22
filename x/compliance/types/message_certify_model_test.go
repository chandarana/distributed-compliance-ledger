package types

import (
	fmt "fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgCertifyModel_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgCertifyModel
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCertifyModel{
				Signer:                "invalid_address",
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vid is 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Vid:                   0,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Vid:                   -1,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Vid:                   65536,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "pid is 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   0,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid < 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   -1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid > 65535",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   65536,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "cd version number > 65535",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				CDVersionNumber:       65536,
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "certification date not set",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     "",
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "certification type not set",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "",
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "certification date is not RFC3339",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     "2020-01-01",
				CertificationType:     testconstants.TestResult,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: ErrInvalidTestDateFormat,
		},
		{
			name: "invalid certification type",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "invalid certification type",
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: ErrInvalidCertificationType,
		},
		{
			name: "software version string len > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: fmt.Sprintf("1.%063d", 0),
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "reason len > 102400 (100 KB)",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                tmrand.Str(102401),
				CDCertificateId:       testconstants.CDCertificateID,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "programTypeVersion > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				ProgramTypeVersion:    tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CDCertificateId > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "familyID > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				FamilyId:              tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "supportedClusters > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				SupportedClusters:     tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "compliancePlatformUsed > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				CompliantPlatformUsed: tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "compliancePlatformVersion > 64",
			msg: MsgCertifyModel{
				Signer:                   sample.AccAddress(),
				Pid:                      1,
				Vid:                      1,
				SoftwareVersionString:    testconstants.TestDate,
				CertificationDate:        testconstants.CertificationDate,
				CertificationType:        testconstants.CertificationType,
				CDVersionNumber:          uint32(testconstants.CdVersionNumber),
				Reason:                   testconstants.Reason,
				CDCertificateId:          testconstants.CDCertificateID,
				CompliantPlatformVersion: tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "OSVersion > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				OSVersion:             tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "certificationRoute > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				CertificationRoute:    tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "programType > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				ProgramType:           tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "transport > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				Transport:             tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "invalid parent/child value",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				ParentChild:           "invalid parent/child value",
			},
			err: ErrInvalidPFCCertificationRoute,
		},
		{
			name: "schemaVersion > 65535",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
				SchemaVersion:         65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgCertifyModel
	}{
		{
			name: "valid certification model msg",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "software version = 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "cd version number = 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       1,
				CDVersionNumber:       0,
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "certification type is zigbee",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "zigbee",
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "certification type is matter",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "matter",
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "minimal pid, vid values",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: "1",
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "max pid, vid values",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: "1",
				Pid:                   65535,
				Vid:                   65535,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "ProgramTypeVersion >= 0 && ProgramTypeVersion <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				ProgramTypeVersion:    testconstants.ProgramTypeVersion,
			},
		},
		{
			name: "CDCertificateId >= 0 && CDCertificateId <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
			},
		},
		{
			name: "FamilyId >= 0 && FamilyId <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				FamilyId:              testconstants.FamilyID,
			},
		},
		{
			name: "SupportedClusters >= 0 && SupportedClusters <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				SupportedClusters:     testconstants.SupportedClusters,
			},
		},
		{
			name: "CompliancePlatformUsed >= 0 && CompliancePlatformUsed <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				CompliantPlatformUsed: testconstants.CompliantPlatformUsed,
			},
		},
		{
			name: "CompliancePlatformVersion >= 0 && CompliancePlatformVersion <= 64",
			msg: MsgCertifyModel{
				Signer:                   sample.AccAddress(),
				Pid:                      1,
				Vid:                      1,
				SoftwareVersionString:    testconstants.TestDate,
				CertificationDate:        testconstants.CertificationDate,
				CertificationType:        testconstants.CertificationType,
				CDVersionNumber:          uint32(testconstants.CdVersionNumber),
				Reason:                   testconstants.Reason,
				CDCertificateId:          testconstants.CDCertificateID,
				CompliantPlatformVersion: testconstants.CompliantPlatformVersion,
			},
		},
		{
			name: "OSVersion >= 0 && OSVersion <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				OSVersion:             testconstants.OSVersion,
			},
		},
		{
			name: "CertificationRoute >= 0 && CertificationRoute <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				CertificationRoute:    testconstants.CertificationRoute,
			},
		},
		{
			name: "ProgramType >= 0 && ProgramType <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				ProgramType:           testconstants.ProgramType,
			},
		},
		{
			name: "Transport >= 0 && Transport <= 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
				CDCertificateId:       testconstants.CDCertificateID,
				Transport:             testconstants.Transport,
			},
		},
		{
			name: "parent/child value is parent",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
				ParentChild:           "parent",
			},
		},
		{
			name: "parent/child value is child",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				CDCertificateId:       testconstants.CDCertificateID,
				ParentChild:           "child",
			},
		},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}
}
