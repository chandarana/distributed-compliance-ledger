package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func TestHandler_RevocationPointsByIssuerSubjectKeyID(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAccount(65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.PAACertWithNumericVid,
		testconstants.Info,
		testconstants.Vid,
		testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(),
		testconstants.PAACertWithNumericVidSubject,
		testconstants.PAACertWithNumericVidSubjectKeyID,
		testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.False(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 0)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL + "/1",
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)

	addPkiRevocationDistributionPoint = types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label1",
		DataURL:              testconstants.DataURL + "/2",
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 2)

	dataURLNew := testconstants.DataURL + "/new"
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              dataURLNew,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 2)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, updatePkiRevocationDistributionPoint.CrlSignerCertificate)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataURL, updatePkiRevocationDistributionPoint.DataURL)

	deletePkiRevocationDistributionPoint := types.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAcc.String(),
		Vid:                65521,
		Label:              "label",
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &deletePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)
}

func TestHandler_AddRevocationPointForSameCertificateWithDifferentWhitespaces(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAccount(65521)

	// propose x509 root certificate by account Trustee1
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVid, testconstants.Info, testconstants.Vid, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.NoError(t, err)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint := types.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  8,
		CrlSignerCertificate: testconstants.PAACertWithNumericVidDifferentWhitespaces,
		Label:                "label",
		DataURL:              testconstants.DataURL + "/1",
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       1,
	}
	_, err = setup.Handler(setup.Ctx, &addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, len(revocationPointBySubjectKeyID.Points), 1)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, addPkiRevocationDistributionPoint.CrlSignerCertificate)
}

func TestHandler_AddPkiRevocationDistributionPoint_NegativeCases(t *testing.T) {
	accAddress := utils.GenerateAccAddress()

	cases := []struct {
		name            string
		accountVid      int32
		accountRole     dclauthtypes.AccountRole
		rootCertOptions *utils.RootCertOptions
		addRevocation   *types.MsgAddPkiRevocationDistributionPoint
		err             error
	}{
		{
			name:          "PAASenderNotVendor",
			accountVid:    testconstants.PAACertWithNumericVidVid,
			accountRole:   dclauthtypes.CertificationCenter,
			addRevocation: createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:           sdkerrors.ErrUnauthorized,
		},
		{
			name:          "PAISenderNotVendor",
			accountVid:    testconstants.PAICertWithNumericPidVidVid,
			accountRole:   dclauthtypes.CertificationCenter,
			addRevocation: createAddRevocationMessageWithPAICertWithNumericVidPid(accAddress.String()),
			err:           sdkerrors.ErrUnauthorized,
		},
		{
			name:          "PAACertEncodesVidSenderVidNotEqualVidField",
			accountVid:    testconstants.Vid,
			accountRole:   dclauthtypes.Vendor,
			addRevocation: createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:           pkitypes.ErrMessageVidNotEqualAccountVid,
		},
		{
			name:          "PAACertNotFound",
			accountVid:    testconstants.PAACertWithNumericVidVid,
			accountRole:   dclauthtypes.Vendor,
			addRevocation: createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:           pkitypes.ErrCertificateDoesNotExist,
		},
		{
			name:          "PAINotChainedBackToDCLCerts",
			accountVid:    testconstants.PAACertWithNumericVidVid,
			accountRole:   dclauthtypes.Vendor,
			addRevocation: createAddRevocationMessageWithPAICertWithNumericVidPid(accAddress.String()),
			err:           pkitypes.ErrCertNotChainedBack,
		},
		{
			name:        "InvalidCertificate",
			accountVid:  testconstants.PAACertWithNumericVidVid,
			accountRole: dclauthtypes.Vendor,
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				Pid:                  0,
				CrlSignerCertificate: "invalidpem",
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrInvalidCertificate,
		},
		{
			name:            "PAANotOnLedger",
			accountVid:      testconstants.PAACertWithNumericVidVid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: utils.CreateTestRootCertOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String()),
			err:             pkitypes.ErrCertificateDoesNotExist,
		},
		{
			name:            "PAANoVid_LedgerPAANoVid",
			accountVid:      testconstants.Vid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.VendorID1),
			addRevocation:   createAddRevocationMessageWithPAACertNoVid(accAddress.String(), testconstants.Vid),
			err:             pkitypes.ErrMessageVidNotEqualRootCertVid,
		},
		{
			name:        "PAANoVid_WrongVID",
			accountVid:  testconstants.Vid,
			accountRole: dclauthtypes.Vendor,
			rootCertOptions: &utils.RootCertOptions{
				PemCert:      testconstants.PAACertNoVid,
				Info:         testconstants.Info,
				Subject:      testconstants.PAACertNoVidSubject,
				SubjectKeyID: testconstants.PAACertNoVidSubjectKeyID,
				Vid:          testconstants.VendorID1,
			},
			addRevocation: createAddRevocationMessageWithPAACertNoVid(accAddress.String(), testconstants.Vid),
			err:           pkitypes.ErrMessageVidNotEqualRootCertVid,
		},
		{
			name:            "Invalid PAI Delegator certificate",
			accountVid:      testconstants.LeafCertWithVidVid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: utils.CreateRootWithVidOptions(),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.LeafCertWithVidVid,
				IsPAA:                false,
				Pid:                  0,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				CrlSignerDelegator:   "invalid",
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrInvalidCertificate,
		},
		{
			name:            "CRL Signer Certificate is not chained back to Delegator PAI certificate",
			accountVid:      testconstants.LeafCertWithVidVid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: utils.CreateRootWithVidOptions(),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.LeafCertWithVidVid,
				IsPAA:                false,
				Pid:                  0,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				CrlSignerDelegator:   testconstants.IntermediateCertPem,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateSubjectKeyIDWithoutColumns,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrCertNotChainedBack,
		},
		{
			name:            "Delegated CRL Signer Certificate is not chained back to root certificate on DCL",
			accountVid:      testconstants.LeafCertWithVidVid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: utils.CreateTestRootCertOptions(),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.LeafCertWithVidVid,
				IsPAA:                false,
				Pid:                  0,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				CrlSignerDelegator:   testconstants.IntermediateCertWithVid1,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrCertNotChainedBack,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)

			if tc.rootCertOptions != nil {
				utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions)
			}

			_, err := setup.Handler(setup.Ctx, tc.addRevocation)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_AddPkiRevocationDistributionPoint_PAAAlreadyExists(t *testing.T) {
	setup := utils.Setup(t)

	accAddress := setup.CreateVendorAccount(testconstants.PAACertWithNumericVidVid)

	// propose and approve x509 root certificate
	rootCert := utils.RootDaCertificateWithNumericVid(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCert)

	addPkiRevocationDistributionPoint := createAddRevocationMessageWithPAACertWithNumericVid(accAddress.String())

	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointAlreadyExists)
}

func TestHandler_AddPkiRevocationDistributionPoint_PositiveCases(t *testing.T) {
	vendorAcc := utils.GenerateAccAddress()

	cases := []struct {
		name            string
		rootCertOptions *utils.RootCertOptions
		addRevocation   *types.MsgAddPkiRevocationDistributionPoint
		SchemaVersion   uint32
	}{
		{
			name:            "PAAWithVid",
			rootCertOptions: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			SchemaVersion:   0,
		},
		{
			name:            "PAIWithNumericVidPid",
			rootCertOptions: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String()),
			SchemaVersion:   0,
		},
		{
			name:            "PAIWithStringVidPid",
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.PAICertWithPidVidVid),
			addRevocation:   createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String()),
			SchemaVersion:   0,
		},
		{
			name:            "PAANoVid",
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.VendorID1),
			addRevocation:   createAddRevocationMessageWithPAACertNoVid(vendorAcc.String(), testconstants.VendorID1),
			SchemaVersion:   0,
		},
		{
			name:            "PAIWithVid",
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.PAICertWithVidVid),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.PAICertWithVidVid,
				IsPAA:                false,
				Pid:                  0,
				CrlSignerCertificate: testconstants.PAICertWithVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			SchemaVersion: testconstants.SchemaVersion,
		},
		{
			name:            "CrlSignerDelegatedByPAI",
			rootCertOptions: utils.CreateTestRootCertOptions(),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  65522,
				IsPAA:                false,
				Pid:                  0,
				CrlSignerCertificate: testconstants.LeafCertPem,
				CrlSignerDelegator:   testconstants.IntermediateCertPem,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.RootSubjectKeyIDWithoutColumns,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			SchemaVersion: testconstants.SchemaVersion,
		},
		{
			name:            "CrlSignerDelegatedByPAA",
			rootCertOptions: utils.CreateTestRootCertOptions(),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  65522,
				IsPAA:                true,
				Pid:                  0,
				CrlSignerCertificate: testconstants.IntermediateCertPem,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.RootSubjectKeyIDWithoutColumns,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			SchemaVersion: testconstants.SchemaVersion,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.addRevocation.Vid)

			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions)

			tc.addRevocation.SchemaVersion = tc.SchemaVersion
			_, err := setup.Handler(setup.Ctx, tc.addRevocation)
			require.NoError(t, err)

			revocationPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, tc.addRevocation.Vid, label, tc.addRevocation.IssuerSubjectKeyID)
			require.True(t, isFound)
			assertRevocationPointEqual(t, tc.addRevocation, &revocationPoint)

			revocationPointBySubjectKeyID, isFound := setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, tc.addRevocation.IssuerSubjectKeyID)
			require.True(t, isFound)
			assertRevocationPointEqual(t, tc.addRevocation, revocationPointBySubjectKeyID.Points[0])
		})
	}
}

func TestHandler_AddPkiRevocationDistributionPoint_DataURLNotUnique(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAccount(testconstants.PAICertWithPidVidVid)
	baseVendorAcc := setup.CreateVendorAccount(testconstants.Vid)

	// propose and approve root certificate
	rootCertOptions := utils.CreatePAACertNoVidOptions(testconstants.Vid)
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint = createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String())
	addPkiRevocationDistributionPoint.Label = "label-new"
	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointAlreadyExists)

	addPkiRevocationDistributionPoint = createAddRevocationMessageWithPAACertNoVid(baseVendorAcc.String(), testconstants.Vid)
	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)
}
