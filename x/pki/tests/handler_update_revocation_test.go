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

func TestHandler_UpdateRevocationPointForSameCertificateWithDifferentWhitespaces(t *testing.T) {
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
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
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

	dataURLNew := testconstants.DataURL + "/new"
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAACertWithNumericVidDifferentWhitespaces,
		Label:                "label",
		DataURL:              dataURLNew,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)

	revocationPointBySubjectKeyID, isFound = setup.Keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(setup.Ctx, testconstants.SubjectKeyIDWithoutColons)
	require.True(t, isFound)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].CrlSignerCertificate, updatePkiRevocationDistributionPoint.CrlSignerCertificate)
	require.Equal(t, revocationPointBySubjectKeyID.Points[0].DataURL, updatePkiRevocationDistributionPoint.DataURL)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_NegativeCases(t *testing.T) {
	accAddress := utils.GenerateAccAddress()
	vendorAcc := utils.GenerateAccAddress()

	cases := []struct {
		name              string
		accountVid        int32
		accountRole       dclauthtypes.AccountRole
		vendorAccVid      int32
		rootCertOptions   *utils.RootCertOptions
		addRevocation     *types.MsgAddPkiRevocationDistributionPoint
		updatedRevocation *types.MsgUpdatePkiRevocationDistributionPoint
		err               error
	}{
		{
			name:            "PAASenderNotVendor",
			accountVid:      testconstants.PAACertWithNumericVidVid,
			accountRole:     dclauthtypes.CertificationCenter,
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name:            "PAISenderNotVendor",
			accountVid:      testconstants.PAACertWithNumericVidVid,
			accountRole:     dclauthtypes.CertificationCenter,
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.PAICertWithNumericPidVidVid,
				CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			name:            "PAASenderVidNotEqualCertVid",
			accountVid:      testconstants.VendorID1,
			accountRole:     dclauthtypes.Vendor,
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrMessageVidNotEqualAccountVid,
		},
		{
			name:            "PAISenderVidNotEqualCertVid",
			accountVid:      testconstants.VendorID1,
			accountRole:     dclauthtypes.Vendor,
			vendorAccVid:    testconstants.PAICertWithPidVidVid,
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.PAICertWithPidVidVid),
			addRevocation:   createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               accAddress.String(),
				Vid:                  testconstants.PAICertWithPidVidVid,
				CrlSignerCertificate: testconstants.PAICertWithPidVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrMessageVidNotEqualAccountVid,
		},
		{
			name:            "PAIPidNotFound",
			vendorAccVid:    testconstants.PAICertWithPidVidVid,
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.PAICertWithPidVidVid),
			addRevocation:   createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.PAICertWithPidVidVid,
				CrlSignerCertificate: testconstants.PAICertWithVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrPidNotFound,
		},
		{
			name:         "DustributionPointNotFound",
			vendorAccVid: testconstants.PAACertWithNumericVidVid,
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrPkiRevocationDistributionPointDoesNotExists,
		},
		{
			name:            "PAANewCertificateNotPAA",
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  65521,
				CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrRootCertificateIsNotSelfSigned,
		},
		{
			name:            "PAANotOnLedger",
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: utils.CreatePAACertNoVidOptions(testconstants.PAACertWithNumericVidVid),
			addRevocation:   createAddRevocationMessageWithPAACertNoVid(vendorAcc.String(), testconstants.PAACertWithNumericVidVid),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrCertificateDoesNotExist,
		},
		{
			name:            "DataFieldsProvidedWhenRevocationType1",
			vendorAccVid:    testconstants.PAACertWithNumericVidVid,
			rootCertOptions: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:   createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             vendorAcc.String(),
				Vid:                testconstants.PAICertWithNumericPidVidVid,
				Label:              label,
				DataURL:            testconstants.DataURL + "/new",
				DataFileSize:       uint64(123),
				DataDigest:         testconstants.DataDigest + "new",
				DataDigestType:     1,
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:      0,
			},
			err: pkitypes.ErrDataFieldPresented,
		},
		{
			name:            "Invalid PAI Delegator certificate",
			accountVid:      testconstants.LeafCertWithVidVid,
			vendorAccVid:    testconstants.LeafCertWithVidVid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: utils.CreateRootWithVidOptions(),
			addRevocation:   createAddRevocationMessageWithLeafCertWithVid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.LeafCertWithVidVid,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				CrlSignerDelegator:   "invalid",
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrInvalidCertificate,
		},
		{
			name:            "CRL Signer Certificate is not chained back to Delegator PAI certificate",
			accountVid:      testconstants.LeafCertWithVidVid,
			vendorAccVid:    testconstants.LeafCertWithVidVid,
			accountRole:     dclauthtypes.Vendor,
			rootCertOptions: utils.CreateRootWithVidOptions(),
			addRevocation:   createAddRevocationMessageWithLeafCertWithVid(vendorAcc.String()),
			updatedRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.LeafCertWithVidVid,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				CrlSignerDelegator:   testconstants.IntermediateCertPem,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrCertNotChainedBack,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			setup.AddAccount(accAddress, []dclauthtypes.AccountRole{tc.accountRole}, tc.accountVid)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.vendorAccVid)

			if tc.rootCertOptions != nil {
				utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions)
			}

			if tc.addRevocation != nil {
				_, err := setup.Handler(setup.Ctx, tc.addRevocation)
				require.NoError(t, err)
			}

			_, err := setup.Handler(setup.Ctx, tc.updatedRevocation)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestHandler_UpdatePkiRevocationDistributionPoint_NotUniqueDataURLForIssuer(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// propose and approve root certificate
	rootCertOptions := utils.RootDaCertificateWithNumericVid(setup.Trustee1)
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint1 := createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String())
	addPkiRevocationDistributionPoint1.Label += "-1"
	addPkiRevocationDistributionPoint1.DataURL += "/1"
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint1)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint2 := createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String())
	addPkiRevocationDistributionPoint2.Label += "-2"
	addPkiRevocationDistributionPoint2.DataURL += "/2"
	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint2)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:             addPkiRevocationDistributionPoint1.Signer,
		Vid:                addPkiRevocationDistributionPoint1.Vid,
		Label:              addPkiRevocationDistributionPoint1.Label,
		DataURL:            addPkiRevocationDistributionPoint2.DataURL,
		IssuerSubjectKeyID: addPkiRevocationDistributionPoint1.IssuerSubjectKeyID,
		SchemaVersion:      addPkiRevocationDistributionPoint1.SchemaVersion,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrPkiRevocationDistributionPointAlreadyExists)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_DataURLNotUnique(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, 65522)

	baseVendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(baseVendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// propose and approve root certificate
	rootCertOptions := utils.CreatePAACertNoVidOptions(testconstants.Vid)
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint1 := createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String())
	addPkiRevocationDistributionPoint1.DataURL += "/1"
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint1)
	require.NoError(t, err)

	addPkiRevocationDistributionPoint2 := createAddRevocationMessageWithPAACertNoVid(baseVendorAcc.String(), testconstants.Vid)
	addPkiRevocationDistributionPoint2.DataURL += "/2"
	_, err = setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint2)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               baseVendorAcc.String(),
		Vid:                  testconstants.Vid,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                "label",
		DataURL:              addPkiRevocationDistributionPoint1.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		SchemaVersion:        0,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAI_NotChainedOnLedger(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// propose and approve root certificate
	rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	proposeRevokeRootCert := types.NewMsgProposeRevokeX509RootCert(setup.Trustee1.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, "", false, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeRevokeRootCert)
	require.NoError(t, err)

	approveRevokeRootCert := types.NewMsgApproveRevokeX509RootCert(
		setup.Trustee2.String(), testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID, "", testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveRevokeRootCert)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		SchemaVersion:        0,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCertNotChainedBack)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAI_VID_TO_PAI_NOVID(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAACertWithNumericVidVid)

	// add PAA for PAI_VID
	rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	// add PAA for PAI_NOVID
	rootCertOptions = utils.CreateTestRootCertOptions()
	rootCertOptions.Vid = testconstants.PAACertWithNumericVidVid
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	// add Revocation Point PAI_VID
	addPkiRevocationDistributionPoint := createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String())
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	// update Revocation Point to PAI_NOVID
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  testconstants.PAACertWithNumericVidVid,
		CrlSignerCertificate: testconstants.IntermediateCertPem,
		Label:                addPkiRevocationDistributionPoint.Label,
		IssuerSubjectKeyID:   addPkiRevocationDistributionPoint.IssuerSubjectKeyID,
		SchemaVersion:        0,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrCRLSignerCertificateVidNotEqualRevocationPointVid)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_NOVID_DifferentVID(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

	// add PAA NOVID 1 with VendorID1
	rootCertOptions := utils.CreatePAACertNoVidOptions(testconstants.VendorID1)
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	// add PAA NOVID 2 with VendorID2
	rootCertOptions = utils.CreateTestRootCertOptions()
	rootCertOptions.Vid = testconstants.VendorID2
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	// add Revocation Point PAA NOVID 1
	addPkiRevocationDistributionPoint := createAddRevocationMessageWithPAACertNoVid(vendorAcc.String(), testconstants.VendorID1)
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	// update to PAA NOVID 2 (with different vid)
	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  addPkiRevocationDistributionPoint.Vid,
		CrlSignerCertificate: testconstants.RootCertPem,
		Label:                addPkiRevocationDistributionPoint.Label,
		IssuerSubjectKeyID:   addPkiRevocationDistributionPoint.IssuerSubjectKeyID,
		SchemaVersion:        0,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.ErrorIs(t, err, pkitypes.ErrMessageVidNotEqualRootCertVid)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_VID(t *testing.T) {
	var err error
	vendorAcc := utils.GenerateAccAddress()
	addedRevocation := createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String())
	cases := []struct {
		name              string
		updatedRevocation types.MsgUpdatePkiRevocationDistributionPoint
		err               error
		schemaVersion     uint32
	}{
		{
			name: "Valid: PAAWithVid",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               addedRevocation.Signer,
				Vid:                  addedRevocation.Vid,
				CrlSignerCertificate: addedRevocation.CrlSignerCertificate,
				Label:                addedRevocation.Label,
				DataURL:              addedRevocation.DataURL,
				IssuerSubjectKeyID:   addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:        addedRevocation.SchemaVersion,
			},
			schemaVersion: uint32(0),
		},
		{
			name: "Valid: MinimalParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:      0,
			},
			schemaVersion: uint32(0),
		},
		{
			name: "Valid: AllParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocation.DataURL + "/new",
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:      0,
			},
			schemaVersion: uint32(0),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, addedRevocation.Vid)

			// propose and approve root certificate
			rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

			// add revocation
			if addedRevocation != nil {
				_, err = setup.Handler(setup.Ctx, addedRevocation)
				require.NoError(t, err)
			}
			_, err = setup.Handler(setup.Ctx, &tc.updatedRevocation)
			updatedPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, addedRevocation.Vid, addedRevocation.Label, addedRevocation.IssuerSubjectKeyID)

			require.NoError(t, err)
			require.True(t, isFound)
			require.Equal(t, updatedPoint.Vid, addedRevocation.Vid)
			require.Equal(t, updatedPoint.Pid, addedRevocation.Pid)
			require.Equal(t, updatedPoint.IsPAA, addedRevocation.IsPAA)
			require.Equal(t, updatedPoint.Label, addedRevocation.Label)
			require.Equal(t, updatedPoint.IssuerSubjectKeyID, addedRevocation.IssuerSubjectKeyID)
			require.Equal(t, updatedPoint.RevocationType, addedRevocation.RevocationType)
			require.Equal(t, updatedPoint.SchemaVersion, tc.schemaVersion)

			compareUpdatedStringFields(t, addedRevocation.DataURL, tc.updatedRevocation.DataURL, updatedPoint.DataURL)
			compareUpdatedStringFields(t, addedRevocation.DataDigest, tc.updatedRevocation.DataDigest, updatedPoint.DataDigest)
			compareUpdatedStringFields(t, addedRevocation.CrlSignerCertificate, tc.updatedRevocation.CrlSignerCertificate, updatedPoint.CrlSignerCertificate)
			compareUpdatedIntFields(t, int(addedRevocation.DataDigestType), int(tc.updatedRevocation.DataDigestType), int(updatedPoint.DataDigestType))
			compareUpdatedIntFields(t, int(addedRevocation.DataFileSize), int(tc.updatedRevocation.DataFileSize), int(updatedPoint.DataFileSize))
		})
	}
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAA_NOVID(t *testing.T) {
	var err error
	vendorAcc := utils.GenerateAccAddress()
	addedRevocation := createAddRevocationMessageWithPAACertNoVid(vendorAcc.String(), testconstants.VendorID1)
	cases := []struct {
		name              string
		updatedRevocation types.MsgUpdatePkiRevocationDistributionPoint
		err               error
	}{
		{
			name: "Valid: Same PAA",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               addedRevocation.Signer,
				Vid:                  addedRevocation.Vid,
				CrlSignerCertificate: addedRevocation.CrlSignerCertificate,
				Label:                addedRevocation.Label,
				DataURL:              addedRevocation.DataURL,
				IssuerSubjectKeyID:   addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:        0,
			},
		},
		{
			name: "Valid: MinimalParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:      0,
			},
		},
		{
			name: "Valid: AllParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocation.DataURL + "/new",
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:      0,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.VendorID1)

			// propose x509 root certificate by account Trustee1
			rootCertOptions := utils.CreatePAACertNoVidOptions(addedRevocation.Vid)
			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

			// add revocation
			if addedRevocation != nil {
				_, err = setup.Handler(setup.Ctx, addedRevocation)
				require.NoError(t, err)
			}
			_, err = setup.Handler(setup.Ctx, &tc.updatedRevocation)
			updatedPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, addedRevocation.Vid, addedRevocation.Label, addedRevocation.IssuerSubjectKeyID)

			require.NoError(t, err)
			require.True(t, isFound)
			require.Equal(t, updatedPoint.Vid, addedRevocation.Vid)
			require.Equal(t, updatedPoint.Pid, addedRevocation.Pid)
			require.Equal(t, updatedPoint.IsPAA, addedRevocation.IsPAA)
			require.Equal(t, updatedPoint.Label, addedRevocation.Label)
			require.Equal(t, updatedPoint.IssuerSubjectKeyID, addedRevocation.IssuerSubjectKeyID)
			require.Equal(t, updatedPoint.RevocationType, addedRevocation.RevocationType)

			compareUpdatedStringFields(t, addedRevocation.DataURL, tc.updatedRevocation.DataURL, updatedPoint.DataURL)
			compareUpdatedStringFields(t, addedRevocation.DataDigest, tc.updatedRevocation.DataDigest, updatedPoint.DataDigest)
			compareUpdatedStringFields(t, addedRevocation.CrlSignerCertificate, tc.updatedRevocation.CrlSignerCertificate, updatedPoint.CrlSignerCertificate)
			compareUpdatedIntFields(t, int(addedRevocation.DataDigestType), int(tc.updatedRevocation.DataDigestType), int(updatedPoint.DataDigestType))
			compareUpdatedIntFields(t, int(addedRevocation.DataFileSize), int(tc.updatedRevocation.DataFileSize), int(updatedPoint.DataFileSize))
		})
	}
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAI_VIDPID(t *testing.T) {
	var err error
	vendorAcc := utils.GenerateAccAddress()
	addedRevocation := createAddRevocationMessageWithPAICertWithNumericVidPid(vendorAcc.String())
	cases := []struct {
		name              string
		updatedRevocation types.MsgUpdatePkiRevocationDistributionPoint
		err               error
	}{
		{
			name: "Valid: Same PAI",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               addedRevocation.Signer,
				Vid:                  addedRevocation.Vid,
				CrlSignerCertificate: addedRevocation.CrlSignerCertificate,
				Label:                addedRevocation.Label,
				DataURL:              addedRevocation.DataURL,
				IssuerSubjectKeyID:   addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:        addedRevocation.SchemaVersion,
			},
		},
		{
			name: "Valid: MinimalParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:      addedRevocation.SchemaVersion,
			},
		},
		{
			name: "Valid: AllParams",
			updatedRevocation: types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:             addedRevocation.Signer,
				Vid:                addedRevocation.Vid,
				Label:              addedRevocation.Label,
				DataURL:            addedRevocation.DataURL + "/new",
				IssuerSubjectKeyID: addedRevocation.IssuerSubjectKeyID,
				SchemaVersion:      addedRevocation.SchemaVersion,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)
			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, addedRevocation.Vid)

			// propose and approve root certificate
			rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

			// add revocation
			if addedRevocation != nil {
				_, err = setup.Handler(setup.Ctx, addedRevocation)
				require.NoError(t, err)
			}
			_, err = setup.Handler(setup.Ctx, &tc.updatedRevocation)
			updatedPoint, isFound := setup.Keeper.GetPkiRevocationDistributionPoint(setup.Ctx, addedRevocation.Vid, addedRevocation.Label, addedRevocation.IssuerSubjectKeyID)

			require.NoError(t, err)
			require.True(t, isFound)
			require.Equal(t, updatedPoint.Vid, addedRevocation.Vid)
			require.Equal(t, updatedPoint.Pid, addedRevocation.Pid)
			require.Equal(t, updatedPoint.IsPAA, addedRevocation.IsPAA)
			require.Equal(t, updatedPoint.Label, addedRevocation.Label)
			require.Equal(t, updatedPoint.IssuerSubjectKeyID, addedRevocation.IssuerSubjectKeyID)
			require.Equal(t, updatedPoint.RevocationType, addedRevocation.RevocationType)

			compareUpdatedStringFields(t, addedRevocation.DataURL, tc.updatedRevocation.DataURL, updatedPoint.DataURL)
			compareUpdatedStringFields(t, addedRevocation.DataDigest, tc.updatedRevocation.DataDigest, updatedPoint.DataDigest)
			compareUpdatedStringFields(t, addedRevocation.CrlSignerCertificate, tc.updatedRevocation.CrlSignerCertificate, updatedPoint.CrlSignerCertificate)
			compareUpdatedIntFields(t, int(addedRevocation.DataDigestType), int(tc.updatedRevocation.DataDigestType), int(updatedPoint.DataDigestType))
			compareUpdatedIntFields(t, int(addedRevocation.DataFileSize), int(tc.updatedRevocation.DataFileSize), int(updatedPoint.DataFileSize))
		})
	}
}

func compareUpdatedStringFields(t *testing.T, oldValue string, newValue string, updatedValue string) {
	t.Helper()
	if newValue == "" {
		require.Equal(t, oldValue, updatedValue)
	} else {
		require.Equal(t, newValue, updatedValue)
	}
}

func compareUpdatedIntFields(t *testing.T, oldValue int, newValue int, updatedValue int) {
	t.Helper()
	if newValue == 0 {
		require.Equal(t, oldValue, updatedValue)
	} else {
		require.Equal(t, newValue, updatedValue)
	}
}

func TestHandler_UpdatePkiRevocationDistributionPoint_PAIWithoutPid(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.PAICertWithPidVidVid)

	// propose x509 root certificate by account Trustee1
	rootCertOptions := utils.CreatePAACertNoVidOptions(testconstants.Vid)
	utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, rootCertOptions)

	addPkiRevocationDistributionPoint := createAddRevocationMessageWithPAICertWithVidPid(vendorAcc.String())
	addPkiRevocationDistributionPoint.Pid = 0
	_, err := setup.Handler(setup.Ctx, addPkiRevocationDistributionPoint)
	require.NoError(t, err)

	updatePkiRevocationDistributionPoint := types.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAcc.String(),
		Vid:                  addPkiRevocationDistributionPoint.Vid,
		CrlSignerCertificate: testconstants.PAICertWithVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		SchemaVersion:        0,
	}
	_, err = setup.Handler(setup.Ctx, &updatePkiRevocationDistributionPoint)
	require.NoError(t, err)
}

func TestHandler_UpdatePkiRevocationDistributionPoint_CrlSignerCertificateField(t *testing.T) {
	vendorAcc := utils.GenerateAccAddress()

	cases := []struct {
		name             string
		rootCertOptions1 *utils.RootCertOptions
		rootCertOptions2 *utils.RootCertOptions
		addRevocation    *types.MsgAddPkiRevocationDistributionPoint
		updateRevocation *types.MsgUpdatePkiRevocationDistributionPoint
	}{
		{
			name:             "PAA_NOVID_TO_PAA_NOVID",
			rootCertOptions1: utils.CreatePAACertNoVidOptions(testconstants.Vid),
			rootCertOptions2: utils.CreateTestRootCertOptions(),
			addRevocation:    createAddRevocationMessageWithPAACertNoVid(vendorAcc.String(), testconstants.Vid),
			updateRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                label,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
		},
		{
			name:             "PAA_NOVID_TO_PAA_VID",
			rootCertOptions1: utils.CreatePAACertNoVidOptions(testconstants.PAACertWithNumericVidVid),
			rootCertOptions2: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:    createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			updateRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                label,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
		},
		{
			name:             "PAA_VID_TO_PAA_NOVID",
			rootCertOptions1: utils.CreatePAACertNoVidOptions(testconstants.PAACertWithNumericVidVid),
			rootCertOptions2: utils.CreatePAACertWithNumericVidOptions(),
			addRevocation:    createAddRevocationMessageWithPAACertWithNumericVid(vendorAcc.String()),
			updateRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertNoVid,
				Label:                label,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
		},
		{
			name:             "CrlSignerDelegatedByPAI",
			rootCertOptions1: utils.CreateTestRootCertOptions(),
			rootCertOptions2: utils.CreateRootWithVidOptions(),
			addRevocation:    createAddRevocationMessageWithLeafCertWithVid(vendorAcc.String()),
			updateRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				Vid:                  65521,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				CrlSignerDelegator:   testconstants.IntermediateCertWithVid1,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				SchemaVersion:        0,
			},
		},
		{
			name:             "CrlSignerDelegatedByPAA",
			rootCertOptions1: utils.CreateTestRootCertOptions(),
			rootCertOptions2: utils.CreateRootWithVidOptions(),
			addRevocation: &types.MsgAddPkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				IsPAA:                true,
				CrlSignerCertificate: testconstants.IntermediateCertPem,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.RootSubjectKeyIDWithoutColumns,
				RevocationType:       types.CRLRevocationType,
				SchemaVersion:        0,
			},
			updateRevocation: &types.MsgUpdatePkiRevocationDistributionPoint{
				Signer:               vendorAcc.String(),
				CrlSignerCertificate: testconstants.IntermediateCertWithoutVidPid,
				Label:                label,
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.RootSubjectKeyIDWithoutColumns,
				SchemaVersion:        0,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup := utils.Setup(t)

			setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, tc.addRevocation.Vid)

			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions1)
			utils.ProposeAndApproveRootCertificateByOptions(setup, setup.Trustee1, tc.rootCertOptions2)

			_, err := setup.Handler(setup.Ctx, tc.addRevocation)
			require.NoError(t, err)

			_, err = setup.Handler(setup.Ctx, tc.updateRevocation)
			require.NoError(t, err)
		})
	}
}
