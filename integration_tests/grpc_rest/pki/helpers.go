// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pki

import (
	"context"
	"fmt"
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	test_dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/grpc_rest/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func GetAllProposedDaX509RootCerts(suite *utils.TestSuite) (res []pkitypes.ProposedCertificate, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllProposedCertificateResponse
		err := suite.QueryREST("/dcl/pki/proposed-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificateAll(
			context.Background(),
			&pkitypes.QueryAllProposedCertificateRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	}

	return res, nil
}

func GetProposedDaX509RootCert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ProposedCertificate, error) {
	var res pkitypes.ProposedCertificate
	if suite.Rest {
		var resp pkitypes.QueryGetProposedCertificateResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/proposed-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificate(
			context.Background(),
			&pkitypes.QueryGetProposedCertificateRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificate()
	}

	return &res, nil
}

func GetAllDaX509Certs(suite *utils.TestSuite) (res []pkitypes.ApprovedCertificates, err error) {
	return getAllDaX509Certs(suite, "")
}

func GetAllDaX509certsBySubjectKeyID(suite *utils.TestSuite, subjectKeyID string) (res []pkitypes.ApprovedCertificates, err error) {
	return getAllDaX509Certs(suite, subjectKeyID)
}

func getAllDaX509Certs(suite *utils.TestSuite, subjectKeyID string) (res []pkitypes.ApprovedCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllApprovedCertificatesResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/pki/certificates?subjectKeyId=%s", subjectKeyID), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedCertificatesAll(
			context.Background(),
			&pkitypes.QueryAllApprovedCertificatesRequest{SubjectKeyId: subjectKeyID},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	}

	return res, nil
}

func GetDaX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ApprovedCertificates, error) {
	var res pkitypes.ApprovedCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetApprovedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedCertificates(
			context.Background(),
			&pkitypes.QueryGetApprovedCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificates()
	}

	return &res, nil
}

func GetAllRevokedDaX509Certs(suite *utils.TestSuite) (res []pkitypes.RevokedCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllRevokedCertificatesResponse
		err := suite.QueryREST("/dcl/pki/revoked-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedCertificatesAll(
			context.Background(),
			&pkitypes.QueryAllRevokedCertificatesRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	}

	return res, nil
}

func GetRevokedDaX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.RevokedCertificates, error) {
	var res pkitypes.RevokedCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetRevokedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/revoked-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedCertificates(
			context.Background(),
			&pkitypes.QueryGetRevokedCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRevokedCertificates()
	}

	return &res, nil
}

func GetAllProposedRevocationX509Certs(suite *utils.TestSuite) (res []pkitypes.ProposedCertificateRevocation, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllProposedCertificateRevocationResponse
		err := suite.QueryREST("/dcl/pki/proposed-revocation-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificateRevocationAll(
			context.Background(),
			&pkitypes.QueryAllProposedCertificateRevocationRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	}

	return res, nil
}

func GetProposedRevocationX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ProposedCertificateRevocation, error) {
	return getProposedRevocationX509Cert(suite, subject, subjectKeyID, "")
}

func GetProposedRevocationX509CertBySerialNumber(suite *utils.TestSuite, subject string, subjectKeyID string, serialNumber string) (*pkitypes.ProposedCertificateRevocation, error) {
	return getProposedRevocationX509Cert(suite, subject, subjectKeyID, serialNumber)
}

func getProposedRevocationX509Cert(suite *utils.TestSuite, subject string, subjectKeyID string, serialNumber string) (*pkitypes.ProposedCertificateRevocation, error) {
	var res pkitypes.ProposedCertificateRevocation
	if suite.Rest {
		var resp pkitypes.QueryGetProposedCertificateRevocationResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/proposed-revocation-certificates/%s/%s?serialNumber=%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID), url.QueryEscape(serialNumber),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ProposedCertificateRevocation(
			context.Background(),
			&pkitypes.QueryGetProposedCertificateRevocationRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
				SerialNumber: serialNumber,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetProposedCertificateRevocation()
	}

	return &res, nil
}

func GetAllDaRootX509Certs(suite *utils.TestSuite) (res pkitypes.ApprovedRootCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetApprovedRootCertificatesResponse
		err := suite.QueryREST("/dcl/pki/root-certificates", &resp)
		if err != nil {
			return pkitypes.ApprovedRootCertificates{}, err
		}
		res = resp.GetApprovedRootCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedRootCertificates(
			context.Background(),
			&pkitypes.QueryGetApprovedRootCertificatesRequest{},
		)
		if err != nil {
			return pkitypes.ApprovedRootCertificates{}, err
		}
		res = resp.GetApprovedRootCertificates()
	}

	return res, nil
}

func GetAllRevokedDaRootX509Certs(suite *utils.TestSuite) (res pkitypes.RevokedRootCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryGetRevokedRootCertificatesResponse
		err := suite.QueryREST("/dcl/pki/revoked-root-certificates", &resp)
		if err != nil {
			return pkitypes.RevokedRootCertificates{}, err
		}
		res = resp.GetRevokedRootCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RevokedRootCertificates(
			context.Background(),
			&pkitypes.QueryGetRevokedRootCertificatesRequest{},
		)
		if err != nil {
			return pkitypes.RevokedRootCertificates{}, err
		}
		res = resp.GetRevokedRootCertificates()
	}

	return res, nil
}

func GetAllDaX509CertsBySubject(suite *utils.TestSuite, subject string) (*pkitypes.ApprovedCertificatesBySubject, error) {
	var res pkitypes.ApprovedCertificatesBySubject
	if suite.Rest {
		var resp pkitypes.QueryGetApprovedCertificatesBySubjectResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/certificates/%s",
				url.QueryEscape(subject),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificatesBySubject()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ApprovedCertificatesBySubject(
			context.Background(),
			&pkitypes.QueryGetApprovedCertificatesBySubjectRequest{
				Subject: subject,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetApprovedCertificatesBySubject()
	}

	return &res, nil
}

func GetAllChildX509Certs(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.ChildCertificates, error) {
	var res pkitypes.ChildCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetChildCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/child-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetChildCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.ChildCertificates(
			context.Background(),
			&pkitypes.QueryGetChildCertificatesRequest{
				Issuer:         subject,
				AuthorityKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetChildCertificates()
	}

	return &res, nil
}

func GetAllRejectedDaX509RootCerts(suite *utils.TestSuite) (res []pkitypes.RejectedCertificate, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllRejectedCertificatesResponse
		err := suite.QueryREST("dcl/pki/rejected-certificates", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RejectedCertificateAll(
			context.Background(),
			&pkitypes.QueryAllRejectedCertificatesRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	}

	return res, nil
}

func GetRejectedDaX509RootCert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.RejectedCertificate, error) {
	var res pkitypes.RejectedCertificate
	if suite.Rest {
		var resp pkitypes.QueryGetRejectedCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"dcl/pki/rejected-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.RejectedCertificate(
			context.Background(),
			&pkitypes.QueryGetRejectedCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetRejectedCertificate()
	}

	return &res, nil
}

func GetAllPkiRevocationDistributionPoints(suite *utils.TestSuite) (res []pkitypes.PkiRevocationDistributionPoint, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllPkiRevocationDistributionPointResponse
		err := suite.QueryREST("/dcl/pki/revocation-points", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetPkiRevocationDistributionPoint()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.PkiRevocationDistributionPointAll(
			context.Background(),
			&pkitypes.QueryAllPkiRevocationDistributionPointRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPkiRevocationDistributionPoint()
	}

	return res, nil
}

func GetPkiRevocationDistributionPointsBySubject(suite *utils.TestSuite, subjectKeyID string) (*pkitypes.PkiRevocationDistributionPointsByIssuerSubjectKeyID, error) {
	var res pkitypes.PkiRevocationDistributionPointsByIssuerSubjectKeyID
	if suite.Rest {
		var resp pkitypes.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/revocation-points/%s",
				url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.PkiRevocationDistributionPointsByIssuerSubjectKeyID(
			context.Background(),
			&pkitypes.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest{
				IssuerSubjectKeyID: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID()
	}

	return &res, nil
}

func GetPkiRevocationDistributionPoint(suite *utils.TestSuite, vendorID int32, subjectKeyID string, label string) (*pkitypes.PkiRevocationDistributionPoint, error) {
	var res pkitypes.PkiRevocationDistributionPoint
	if suite.Rest {
		var resp pkitypes.QueryGetPkiRevocationDistributionPointResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/revocation-points/%s/%v/%s",
				url.QueryEscape(subjectKeyID), vendorID, url.QueryEscape(label),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPkiRevocationDistributionPoint()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.PkiRevocationDistributionPoint(
			context.Background(),
			&pkitypes.QueryGetPkiRevocationDistributionPointRequest{
				Vid:                vendorID,
				Label:              label,
				IssuerSubjectKeyID: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPkiRevocationDistributionPoint()
	}

	return &res, nil
}

func GetAllCerts(suite *utils.TestSuite) (res []pkitypes.AllCertificates, err error) {
	return getAllCerts(suite, "")
}

func GetAllCertsBySubjectKeyId(suite *utils.TestSuite, subjectKeyID string) (res []pkitypes.AllCertificates, err error) {
	return getAllCerts(suite, subjectKeyID)
}

func getAllCerts(suite *utils.TestSuite, subjectKeyID string) (res []pkitypes.AllCertificates, err error) {
	if suite.Rest {
		var resp pkitypes.QueryAllCertificatesResponse
		err := suite.QueryREST(fmt.Sprintf("/dcl/pki/all-certificates?subjectKeyId=%s", subjectKeyID), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.CertificatesAll(
			context.Background(),
			&pkitypes.QueryAllCertificatesRequest{
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertificates()
	}

	return res, nil
}

func GetCert(suite *utils.TestSuite, subject string, subjectKeyID string) (*pkitypes.AllCertificates, error) {
	var res pkitypes.AllCertificates
	if suite.Rest {
		var resp pkitypes.QueryGetCertificatesResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/all-certificates/%s/%s",
				url.QueryEscape(subject), url.QueryEscape(subjectKeyID),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertificates()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.Certificates(
			context.Background(),
			&pkitypes.QueryGetCertificatesRequest{
				Subject:      subject,
				SubjectKeyId: subjectKeyID,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetCertificates()
	}

	return &res, nil
}

func GetAllCertsBySubject(suite *utils.TestSuite, subject string) (*pkitypes.AllCertificatesBySubject, error) {
	var res pkitypes.AllCertificatesBySubject
	if suite.Rest {
		var resp pkitypes.QueryGetAllCertificatesBySubjectResponse
		err := suite.QueryREST(
			fmt.Sprintf(
				"/dcl/pki/all-certificates/%s",
				url.QueryEscape(subject),
			),
			&resp,
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetAllCertificatesBySubject()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/pki service.
		pkiClient := pkitypes.NewQueryClient(grpcConn)
		resp, err := pkiClient.AllCertificatesBySubject(
			context.Background(),
			&pkitypes.QueryGetAllCertificatesBySubjectRequest{
				Subject: subject,
			},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetAllCertificatesBySubject()
	}

	return &res, nil
}

//nolint:funlen
func Demo(suite *utils.TestSuite) {
	// All requests return empty or 404 value
	proposedCertificates, _ := GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	_, err := GetProposedDaX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	certificates, _ := GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	_, err = GetDaX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	revokedCertificates, _ := GetAllRevokedDaX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedCertificates))

	_, err = GetRevokedDaX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	proposedRevocationCertificates, _ := GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	_, err = GetProposedRevocationX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	rootCertificates, _ := GetAllDaRootX509Certs(suite)
	require.Equal(suite.T, 0, len(rootCertificates.Certs))

	revokedRootCertificates, _ := GetAllRevokedDaRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	_, err = GetAllChildX509Certs(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllDaX509CertsBySubject(suite, testconstants.RootSubject)
	suite.AssertNotFound(err)

	allCertificates, _ := GetAllCerts(suite)
	require.Equal(suite.T, 0, len(allCertificates))

	allCertificatesBySkid, _ := GetAllCertsBySubjectKeyId(suite, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, 0, len(allCertificatesBySkid))

	_, err = GetAllCertsBySubject(suite, testconstants.RootSubject)
	suite.AssertNotFound(err)

	// Alice and Jack are predefined Trustees
	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	address, err := aliceKeyInfo.GetAddress()
	require.NoError(suite.T, err)
	aliceAccount, err := test_dclauth.GetAccount(suite, address)
	require.NoError(suite.T, err)

	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	address, err = jackKeyInfo.GetAddress()
	require.NoError(suite.T, err)
	jackAccount, err := test_dclauth.GetAccount(suite, address)
	require.NoError(suite.T, err)

	// Register new Vendor account
	vendorName := utils.RandString()
	vendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		testconstants.Vid,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// Register new Vendor Admin account
	vendorAdminName := utils.RandString()
	vendorAdminAccount := test_dclauth.CreateAccount(
		suite,
		vendorAdminName,
		dclauthtypes.AccountRoles{dclauthtypes.VendorAdmin},
		0,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAdminAccount)

	// Vendor (Not Trustee) propose Root certificate
	msgProposeAddX509RootCert := pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertPem,
		Signer: vendorAccount.Address,
		Vid:    testconstants.Vid,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, vendorName, vendorAccount)
	require.Error(suite.T, err)

	// Jack (Trustee) propose Root certificate
	msgProposeAddX509RootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertPem,
		Signer: jackAccount.Address,
		Vid:    testconstants.Vid,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.RootSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, proposedCertificates[0].SubjectAsText)

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Root certificate
	proposedCertificate, _ := GetProposedDaX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, jackAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.RootSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, proposedCertificates[0].SubjectAsText)

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Root certificate
	proposedCertificate, _ = GetProposedDaX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, jackAccount.Address, proposedCertificate.Owner)
	require.True(suite.T, proposedCertificate.HasApprovalFrom(jackAccount.Address))

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert := pkitypes.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgApproveAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Assign VID to Root certificate that already has VID
	msgAssignVid := pkitypes.MsgAssignVid{
		Signer:       vendorAdminAccount.Address,
		Subject:      testconstants.RootSubject,
		SubjectKeyId: testconstants.RootSubjectKeyID,
		Vid:          testconstants.Vid,
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAssignVid}, vendorAdminName, vendorAdminAccount)
	require.ErrorContains(suite.T, err, "vid is not empty")

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	certsBySubjectKeyID, _ := GetAllDaX509certsBySubjectKeyID(suite, testconstants.RootSubjectKeyID)
	for _, certs := range [][]pkitypes.ApprovedCertificates{certificates, certsBySubjectKeyID} {
		require.Equal(suite.T, 1, len(certs))
		require.Equal(suite.T, testconstants.RootSubjectKeyID, certs[0].SubjectKeyId)
		require.Equal(suite.T, testconstants.RootSubject, certs[0].Certs[0].Subject)
		require.Equal(suite.T, testconstants.RootSubjectAsText, certs[0].Certs[0].SubjectAsText)
	}

	// Request all approved Root certificates
	rootCertificates, _ = GetAllDaRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request approved Root certificate
	certificate, _ := GetDaX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.RootCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, jackAccount.Address, certificate.Certs[0].Owner)
	require.True(suite.T, certificate.Certs[0].IsRoot)

	// Try to add Intermediate certificate when sender is not Vendor account
	msgAddX509Cert := pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertPem,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, aliceName, aliceAccount)
	require.Error(suite.T, err)

	// Vandor add Intermediate certificate
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertPem,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 2, len(certificates))
	require.Equal(suite.T, testconstants.RootSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, certificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, certificates[1].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, certificates[1].Certs[0].SubjectAsText)

	// Request all approved Root certificates
	rootCertificates, _ = GetAllDaRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request Intermediate certificate
	certificate, _ = GetDaX509Cert(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(suite.T, testconstants.IntermediateSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, vendorAccount.Address, certificate.Certs[0].Owner)
	require.False(suite.T, certificate.Certs[0].IsRoot)

	// Vandor add Leaf certificate
	secondMsgAddX509Cert := pkitypes.MsgAddX509Cert{
		Cert:   testconstants.LeafCertPem,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 3, len(certificates))
	require.Equal(suite.T, testconstants.LeafSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, certificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.RootSubject, certificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, certificates[1].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, certificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, certificates[2].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, certificates[2].Certs[0].SubjectAsText)

	// Request all approved Root certificates
	rootCertificates, _ = GetAllDaRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request Leaf certificate
	certificate, _ = GetDaX509Cert(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(suite.T, testconstants.LeafSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.LeafCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, vendorAccount.Address, certificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.LeafSubject, certificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, certificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.False(suite.T, certificate.Certs[0].IsRoot)

	// Request all Subject certificates
	subjectCertificates, _ := GetAllDaX509CertsBySubject(suite, testconstants.LeafSubject)
	require.Equal(suite.T, testconstants.LeafSubject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	subjectCertificates, _ = GetAllDaX509CertsBySubject(suite, testconstants.IntermediateSubject)
	require.Equal(suite.T, testconstants.IntermediateSubject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	subjectCertificates, _ = GetAllDaX509CertsBySubject(suite, testconstants.RootSubject)
	require.Equal(suite.T, testconstants.RootSubject, subjectCertificates.Subject)
	require.Equal(suite.T, 1, len(subjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.RootSubjectKeyID, subjectCertificates.SubjectKeyIds[0])

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedDaX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedCertificates))

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedDaRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	// Get all child certificates
	childCertificates, _ := GetAllChildX509Certs(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootSubject, childCertificates.Issuer)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, childCertificates.AuthorityKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, childCertificates.CertIds[0].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, childCertificates.CertIds[0].SubjectKeyId)

	childCertificates, _ = GetAllChildX509Certs(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(suite.T, testconstants.IntermediateSubject, childCertificates.Issuer)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, childCertificates.AuthorityKeyId)
	require.Equal(suite.T, 1, len(childCertificates.CertIds))
	require.Equal(suite.T, testconstants.LeafSubject, childCertificates.CertIds[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, childCertificates.CertIds[0].SubjectKeyId)

	_, err = GetAllChildX509Certs(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	suite.AssertNotFound(err)

	// Request all certificates
	allCertificates, _ = GetAllCerts(suite)
	require.Equal(suite.T, 3, len(allCertificates))

	allCertificatesBySkid, _ = GetAllCertsBySubjectKeyId(suite, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, 1, len(allCertificatesBySkid))
	require.Equal(suite.T, testconstants.RootSubjectKeyID, allCertificatesBySkid[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubject, allCertificatesBySkid[0].Certs[0].Subject)

	// Request all Subject certificates
	allSubjectCertificates, _ := GetAllCertsBySubject(suite, testconstants.LeafSubject)
	require.Equal(suite.T, testconstants.LeafSubject, allSubjectCertificates.Subject)
	require.Equal(suite.T, 1, len(allSubjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, allSubjectCertificates.SubjectKeyIds[0])

	allSubjectCertificates, _ = GetAllCertsBySubject(suite, testconstants.IntermediateSubject)
	require.Equal(suite.T, testconstants.IntermediateSubject, allSubjectCertificates.Subject)
	require.Equal(suite.T, 1, len(allSubjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, allSubjectCertificates.SubjectKeyIds[0])

	allSubjectCertificates, _ = GetAllCertsBySubject(suite, testconstants.RootSubject)
	require.Equal(suite.T, testconstants.RootSubject, allSubjectCertificates.Subject)
	require.Equal(suite.T, 1, len(allSubjectCertificates.SubjectKeyIds))
	require.Equal(suite.T, testconstants.RootSubjectKeyID, allSubjectCertificates.SubjectKeyIds[0])

	// Try to Revoke Intermediate certificate when sender is not Vendor account
	msgRevokeX509Cert := pkitypes.MsgRevokeX509Cert{
		Subject:      testconstants.IntermediateSubject,
		SubjectKeyId: testconstants.IntermediateSubjectKeyID,
		RevokeChild:  true,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeX509Cert}, aliceName, aliceAccount)
	require.Error(suite.T, err)

	// Vendor revokes Intermediate certificate. With `RevokeChild` set to true, its child must also be revoked - Leaf certificate.
	msgRevokeX509Cert = pkitypes.MsgRevokeX509Cert{
		Subject:      testconstants.IntermediateSubject,
		SubjectKeyId: testconstants.IntermediateSubjectKeyID,
		RevokeChild:  true,
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedDaX509Certs(suite)
	require.Equal(suite.T, 2, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificates[1].Certs[0].SubjectAsText)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedDaRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	// Request revoked Intermediate certificate
	revokedCertificate, _ := GetRevokedDaX509Cert(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, vendorAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.False(suite.T, revokedCertificate.Certs[0].IsRoot)

	// Request revoked Leaf certificate
	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, vendorAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.False(suite.T, revokedCertificate.Certs[0].IsRoot)

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	certsBySubjectKeyID, _ = GetAllDaX509certsBySubjectKeyID(suite, testconstants.RootSubjectKeyID)
	for _, certs := range [][]pkitypes.ApprovedCertificates{certificates, certsBySubjectKeyID} {
		require.Equal(suite.T, 1, len(certs))
		require.Equal(suite.T, testconstants.RootSubjectKeyID, certs[0].SubjectKeyId)
		require.Equal(suite.T, testconstants.RootSubject, certs[0].Certs[0].Subject)
	}
	// Jack (Trustee) proposes to revoke Root certificate
	msgProposeRevokeX509RootCert := pkitypes.MsgProposeRevokeX509RootCert{
		Subject:      testconstants.RootSubject,
		SubjectKeyId: testconstants.RootSubjectKeyID,
		RevokeChild:  true,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeRevokeX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 1, len(proposedRevocationCertificates))
	require.Equal(suite.T, testconstants.RootSubject, proposedRevocationCertificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedRevocationCertificates[0].SubjectKeyId)

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedDaX509Certs(suite)
	require.Equal(suite.T, 2, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[1].SubjectKeyId)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedDaRootX509Certs(suite)
	require.Equal(suite.T, 0, len(revokedRootCertificates.Certs))

	// Request Root certificate proposed to revoke
	proposedCertificateRevocation, _ := GetProposedRevocationX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, testconstants.RootSubject, proposedCertificateRevocation.Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, proposedCertificateRevocation.SubjectKeyId)
	require.Equal(suite.T, true, proposedCertificateRevocation.RevokeChild)
	require.True(suite.T, proposedCertificateRevocation.HasRevocationFrom(jackAccount.Address))

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 1, len(certificates))
	require.Equal(suite.T, testconstants.RootSubject, certificates[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, certificates[0].SubjectKeyId)

	// Alice (Trustee) approves to revoke Root certificate
	msgApproveRevokeX509RootCert := pkitypes.MsgApproveRevokeX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveRevokeX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedDaX509Certs(suite)
	require.Equal(suite.T, 3, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[2].SubjectKeyId)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedDaRootX509Certs(suite)
	require.Equal(suite.T, 1, len(revokedRootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedRootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedRootCertificates.Certs[0].SubjectKeyId)

	// Request revoked Root certificate
	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, jackAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.RootSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.True(suite.T, revokedCertificate.Certs[0].IsRoot)

	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	_, err = GetProposedDaX509RootCert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetDaX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetDaX509Cert(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetDaX509Cert(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	suite.AssertNotFound(err)

	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	_, err = GetProposedRevocationX509Cert(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	rootCertificates, _ = GetAllDaRootX509Certs(suite)
	require.Equal(suite.T, 0, len(rootCertificates.Certs))

	_, err = GetAllChildX509Certs(suite, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllChildX509Certs(suite, testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllChildX509Certs(suite, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllDaX509CertsBySubject(suite, testconstants.RootSubject)
	suite.AssertNotFound(err)

	_, err = GetAllDaX509CertsBySubject(suite, testconstants.IntermediateSubject)
	suite.AssertNotFound(err)

	_, err = GetAllDaX509CertsBySubject(suite, testconstants.LeafSubject)
	suite.AssertNotFound(err)

	// CHECK GOOGLE ROOT CERTIFICATE WHICH INCLUDES VID
	_, err = GetRevokedDaX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetProposedRevocationX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)
	googlCertVid := int32(24582)
	// Alice (Trustee) propose Google Root certificate
	msgProposeAddX509GoogleRootcert := pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.GoogleCertPem,
		Signer: aliceAccount.Address,
		Vid:    googlCertVid,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509GoogleRootcert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificate
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.GoogleSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request all approved certificates by subjectKeyId
	certificates, _ = GetAllDaX509certsBySubjectKeyID(suite, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Google Root certificate
	proposedCertificate, _ = GetProposedDaX509RootCert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))

	// Jack (Trustee) rejects Root certificate after approval
	msgRejectAddX509RootCert := pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.GoogleSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request all approved certificates by subjectKeyId
	certificates, _ = GetAllDaX509certsBySubjectKeyID(suite, testconstants.RootSubjectKeyID)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Root certificate
	proposedCertificate, _ = GetProposedDaX509RootCert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.True(suite.T, proposedCertificate.HasApprovalFrom(aliceAccount.Address))

	// Jack (Trustee) approves Root certificate
	msgApproveAddX509RootCert := pkitypes.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	certsBySubjectKeyID, _ = GetAllDaX509certsBySubjectKeyID(suite, testconstants.GoogleSubjectKeyID)
	for _, certs := range [][]pkitypes.ApprovedCertificates{certificates, certsBySubjectKeyID} {
		require.Equal(suite.T, 1, len(certs))
		require.Equal(suite.T, testconstants.GoogleSubjectKeyID, certs[0].SubjectKeyId)
		require.Equal(suite.T, testconstants.GoogleSubject, certs[0].Certs[0].Subject)
	}
	// Request all approved Root certificates
	rootCertificates, _ = GetAllDaRootX509Certs(suite)
	require.Equal(suite.T, 1, len(rootCertificates.Certs))
	require.Equal(suite.T, testconstants.GoogleSubject, rootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, rootCertificates.Certs[0].SubjectKeyId)

	// Request approved Google Root certificate
	certificate, _ = GetDaX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleSubject, certificate.Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, certificate.SubjectKeyId)
	require.Equal(suite.T, 1, len(certificate.Certs))
	require.Equal(suite.T, testconstants.GoogleCertPem, certificate.Certs[0].PemCert)
	require.Equal(suite.T, aliceAccount.Address, certificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.GoogleSubjectAsText, certificate.Certs[0].SubjectAsText)
	require.True(suite.T, certificate.Certs[0].IsRoot)

	// Jack (Trustee) proposes to revoke Google Root certificate
	msgProposeRevokeX509RootCert = pkitypes.MsgProposeRevokeX509RootCert{
		Subject:      testconstants.GoogleSubject,
		SubjectKeyId: testconstants.GoogleSubjectKeyID,
		Signer:       jackAccount.Address,
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeRevokeX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 1, len(proposedRevocationCertificates))
	require.Equal(suite.T, testconstants.GoogleSubject, proposedRevocationCertificates[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedRevocationCertificates[0].SubjectKeyId)

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedDaX509Certs(suite)
	require.Equal(suite.T, 3, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, revokedCertificates[1].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[2].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificates[2].Certs[0].SubjectAsText)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedDaRootX509Certs(suite)
	require.Equal(suite.T, 1, len(revokedRootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedRootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedRootCertificates.Certs[0].SubjectKeyId)

	// Request Google Root certificate proposed to revoke
	proposedCertificateRevocation, _ = GetProposedRevocationX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, testconstants.GoogleSubject, proposedCertificateRevocation.Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, proposedCertificateRevocation.SubjectKeyId)
	require.True(suite.T, proposedCertificateRevocation.HasRevocationFrom(jackAccount.Address))

	// Request all approved certificates
	certificates, _ = GetAllDaX509Certs(suite)
	certsBySubjectKeyID, _ = GetAllDaX509certsBySubjectKeyID(suite, testconstants.GoogleSubjectKeyID)
	for _, certs := range [][]pkitypes.ApprovedCertificates{certificates, certsBySubjectKeyID} {
		require.Equal(suite.T, 1, len(certs))
		require.Equal(suite.T, testconstants.GoogleSubjectKeyID, certs[0].SubjectKeyId)
		require.Equal(suite.T, testconstants.GoogleSubject, certs[0].Certs[0].Subject)
		require.Equal(suite.T, testconstants.GoogleSubjectAsText, certs[0].Certs[0].SubjectAsText)
	}

	// Alice (Trustee) approves to revoke Google Root certificate
	msgApproveRevokeX509RootCert = pkitypes.MsgApproveRevokeX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveRevokeX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all Root certificates proposed to revoke
	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	// Request all revoked certificates
	revokedCertificates, _ = GetAllRevokedDaX509Certs(suite)
	require.Equal(suite.T, 4, len(revokedCertificates))
	require.Equal(suite.T, testconstants.LeafSubject, revokedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.LeafSubjectKeyID, revokedCertificates[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.LeafSubjectAsText, revokedCertificates[0].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.RootSubject, revokedCertificates[1].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedCertificates[1].SubjectKeyId)
	require.Equal(suite.T, testconstants.RootSubjectAsText, revokedCertificates[1].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.IntermediateSubject, revokedCertificates[2].Subject)
	require.Equal(suite.T, testconstants.IntermediateSubjectKeyID, revokedCertificates[2].SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateSubjectAsText, revokedCertificates[2].Certs[0].SubjectAsText)
	require.Equal(suite.T, testconstants.GoogleSubject, revokedCertificates[3].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedCertificates[3].SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleSubjectAsText, revokedCertificates[3].Certs[0].SubjectAsText)

	// Request all revoked Root certificates
	revokedRootCertificates, _ = GetAllRevokedDaRootX509Certs(suite)
	require.Equal(suite.T, 2, len(revokedRootCertificates.Certs))
	require.Equal(suite.T, testconstants.RootSubject, revokedRootCertificates.Certs[0].Subject)
	require.Equal(suite.T, testconstants.RootSubjectKeyID, revokedRootCertificates.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleSubject, revokedRootCertificates.Certs[1].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedRootCertificates.Certs[1].SubjectKeyId)

	// Request revoked Google Root certificate
	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.GoogleSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleSubject, revokedCertificate.Certs[0].Subject)
	require.Equal(suite.T, testconstants.GoogleSubjectKeyID, revokedCertificate.Certs[0].SubjectKeyId)
	require.Equal(suite.T, testconstants.GoogleCertPem, revokedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, aliceAccount.Address, revokedCertificate.Certs[0].Owner)
	require.Equal(suite.T, testconstants.GoogleSubjectAsText, revokedCertificate.Certs[0].SubjectAsText)
	require.True(suite.T, revokedCertificate.Certs[0].IsRoot)

	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	certificates, _ = GetAllDaX509certsBySubjectKeyID(suite, testconstants.GoogleSubjectKeyID)
	require.Equal(suite.T, 0, len(certificates))

	_, err = GetProposedDaX509RootCert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetDaX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	proposedRevocationCertificates, _ = GetAllProposedRevocationX509Certs(suite)
	require.Equal(suite.T, 0, len(proposedRevocationCertificates))

	_, err = GetProposedRevocationX509Cert(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	rootCertificates, _ = GetAllDaRootX509Certs(suite)
	require.Equal(suite.T, 0, len(rootCertificates.Certs))

	_, err = GetAllChildX509Certs(suite, testconstants.GoogleSubject, testconstants.GoogleSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetAllDaX509CertsBySubject(suite, testconstants.GoogleSubject)
	suite.AssertNotFound(err)

	// CHECK TEST ROOT CERTIFICATE FOR REJECTING
	_, err = GetRejectedDaX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetProposedDaX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	// Alice (Trustee) propose Test Root certificate
	msgProposeAddX509TestRootCert := pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.TestCertPem,
		Signer: aliceAccount.Address,
		Vid:    testconstants.TestCertPemVid,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509TestRootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request proposed Test Root certificate
	proposedCertificate, _ = GetProposedDaX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)

	// Alice (Trustee) rejects Test Root certificate
	msgRejectX509TestRootCert := pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectX509TestRootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	_, err = GetProposedDaX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetDaX509Cert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	_, err = GetRevokedDaX509Cert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	// Alice (Trustee) propose Test Root certificate
	msgProposeAddX509TestRootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.TestCertPem,
		Signer: aliceAccount.Address,
		Vid:    testconstants.TestCertPemVid,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509TestRootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.TestSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all rejected Root certificates
	rejectedCertificates, _ := GetAllRejectedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(rejectedCertificates))

	// Request all approved Root certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Test Root certificate
	proposedCertificate, _ = GetProposedDaX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))

	// Jack (Trustee) rejects Root certificate
	msgRejectAddX509RootCert = pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Jack (Trustee) rejects Root certificate for the second time
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRejectAddX509RootCert}, jackName, jackAccount)
	require.Error(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 1, len(proposedCertificates))
	require.Equal(suite.T, testconstants.TestSubject, proposedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, proposedCertificates[0].SubjectKeyId)

	// Request all rejected Root certificates
	rejectedCertificates, _ = GetAllRejectedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(rejectedCertificates))

	// Request all approved Root certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request proposed Test Root certificate
	proposedCertificate, _ = GetProposedDaX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestCertPem, proposedCertificate.PemCert)
	require.Equal(suite.T, aliceAccount.Address, proposedCertificate.Owner)
	require.Equal(suite.T, 1, len(proposedCertificate.Approvals))
	require.True(suite.T, proposedCertificate.HasRejectFrom(jackAccount.Address))

	// Alice (Trustee) rejects Root certificate
	secondMsgRejectAddX509RootCert := pkitypes.MsgRejectAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgRejectAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all proposed Root certificates
	proposedCertificates, _ = GetAllProposedDaX509RootCerts(suite)
	require.Equal(suite.T, 0, len(proposedCertificates))

	// Request all approved Root certificates
	certificates, _ = GetAllDaX509Certs(suite)
	require.Equal(suite.T, 0, len(certificates))

	// Request all rejected Root certificates
	rejectedCertificates, _ = GetAllRejectedDaX509RootCerts(suite)
	require.Equal(suite.T, 1, len(rejectedCertificates))
	require.Equal(suite.T, testconstants.TestSubject, rejectedCertificates[0].Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, rejectedCertificates[0].SubjectKeyId)

	// Request rejected Test Root certificate
	rejectedCertificate, _ := GetRejectedDaX509RootCert(suite, testconstants.TestSubject, testconstants.TestSubjectKeyID)
	require.Equal(suite.T, testconstants.TestSubject, rejectedCertificate.Subject)
	require.Equal(suite.T, testconstants.TestSubjectKeyID, rejectedCertificate.SubjectKeyId)
	require.Equal(suite.T, 1, len(rejectedCertificate.Certs))
	require.Equal(suite.T, testconstants.TestCertPem, rejectedCertificate.Certs[0].PemCert)
	require.Equal(suite.T, testconstants.TestSubjectAsText, rejectedCertificate.Certs[0].SubjectAsText)
	require.Equal(suite.T, aliceAccount.Address, rejectedCertificate.Certs[0].Owner)
	require.Equal(suite.T, jackAccount.Address, rejectedCertificate.Certs[0].Rejects[0].Address)
	require.Equal(suite.T, aliceAccount.Address, rejectedCertificate.Certs[0].Rejects[1].Address)

	// PKI Revocation Distribution Point tests

	// Create vendor account
	vendorName = utils.RandString()
	vendorAccount = test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		65521,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// Jack (Trustee) propose Root certificate
	msgProposeAddX509RootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.PAACertWithNumericVid,
		Signer: jackAccount.Address,
		Vid:    vendorAccount.VendorID,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	proposedCertificate, err = GetProposedDaX509RootCert(suite, testconstants.PAACertWithNumericVidSubject, testconstants.PAACertWithNumericVidSubjectKeyID)
	require.NoError(suite.T, err)

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert = pkitypes.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgApproveAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request all revocation distribution points
	revDistPoints, _ := GetAllPkiRevocationDistributionPoints(suite)
	require.Equal(suite.T, 0, len(revDistPoints))

	// Request revocation distribution points
	_, err = GetPkiRevocationDistributionPoint(suite, testconstants.Vid, testconstants.TestSubjectKeyID, "test")
	suite.AssertNotFound(err)

	// Request revocation distribution points by Issuer Subject Key ID
	_, err = GetPkiRevocationDistributionPointsBySubject(suite, testconstants.TestSubjectKeyID)
	suite.AssertNotFound(err)

	// Add revocation distribution point when sender is not Vendor account
	msgAddPkiRevDistPoints := pkitypes.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAccount.Address,
		Vid:                  vendorAccount.VendorID,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		IsPAA:                true,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		RevocationType:       1,
		SchemaVersion:        0,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddPkiRevDistPoints}, vendorAdminName, vendorAdminAccount)
	require.Error(suite.T, err)

	// Add revocation distribution point for PAA by Vendor
	msgAddPkiRevDistPoints = pkitypes.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAccount.Address,
		Vid:                  vendorAccount.VendorID,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		IsPAA:                true,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              testconstants.DataURL,
		RevocationType:       1,
		SchemaVersion:        0,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddPkiRevDistPoints}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	revDistPoints, _ = GetAllPkiRevocationDistributionPoints(suite)
	require.Equal(suite.T, 1, len(revDistPoints))
	require.Equal(suite.T, vendorAccount.VendorID, revDistPoints[0].Vid)
	require.Equal(suite.T, "label", revDistPoints[0].Label)
	require.Equal(suite.T, testconstants.SubjectKeyIDWithoutColons, revDistPoints[0].IssuerSubjectKeyID)

	revDistPoint, _ := GetPkiRevocationDistributionPoint(suite, vendorAccount.VendorID, testconstants.SubjectKeyIDWithoutColons, "label")
	require.Equal(suite.T, vendorAccount.VendorID, revDistPoint.Vid)
	require.Equal(suite.T, "label", revDistPoint.Label)
	require.Equal(suite.T, testconstants.SubjectKeyIDWithoutColons, revDistPoint.IssuerSubjectKeyID)

	revDistPointsBySubj, _ := GetPkiRevocationDistributionPointsBySubject(suite, testconstants.SubjectKeyIDWithoutColons)
	require.Equal(suite.T, vendorAccount.VendorID, revDistPointsBySubj.Points[0].Vid)
	require.Equal(suite.T, "label", revDistPointsBySubj.Points[0].Label)
	require.Equal(suite.T, testconstants.SubjectKeyIDWithoutColons, revDistPointsBySubj.IssuerSubjectKeyID)

	// Add revocation distribution point for PAI by Vendor
	venName65522 := utils.RandString()
	venAcc65522 := test_dclauth.CreateVendorAccount(
		suite,
		venName65522,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		65522,
		testconstants.ProductIDsEmpty,
		jackName,
		jackAccount,
		aliceName,
		aliceAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, venAcc65522)

	// Jack (Trustee) propose Root certificate
	msgProposeAddX509RootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.PAACertNoVid,
		Signer: jackAccount.Address,
		Vid:    venAcc65522.VendorID,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	proposedCertificate, err = GetProposedDaX509RootCert(suite, testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID)
	require.NoError(suite.T, err)

	// Alice (Trustee) approve Root certificate
	secondMsgApproveAddX509RootCert = pkitypes.MsgApproveAddX509RootCert{
		Subject:      proposedCertificate.Subject,
		SubjectKeyId: proposedCertificate.SubjectKeyId,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&secondMsgApproveAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	msgAddPkiRevDistPoints = pkitypes.MsgAddPkiRevocationDistributionPoint{
		Signer:               venAcc65522.Address,
		Vid:                  venAcc65522.VendorID,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		IsPAA:                false,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAICertWithVid,
		Label:                "label_PAI",
		DataURL:              testconstants.DataURL,
		RevocationType:       1,
		SchemaVersion:        0,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddPkiRevDistPoints}, venName65522, venAcc65522)
	require.NoError(suite.T, err)

	revDistPoints, _ = GetAllPkiRevocationDistributionPoints(suite)
	require.Equal(suite.T, 2, len(revDistPoints))

	revDistPoint, _ = GetPkiRevocationDistributionPoint(suite, venAcc65522.VendorID, testconstants.SubjectKeyIDWithoutColons, "label_PAI")
	require.Equal(suite.T, venAcc65522.VendorID, revDistPoint.Vid)
	require.Equal(suite.T, "label_PAI", revDistPoint.Label)
	require.Equal(suite.T, testconstants.SubjectKeyIDWithoutColons, revDistPoint.IssuerSubjectKeyID)

	revDistPointsBySubj, _ = GetPkiRevocationDistributionPointsBySubject(suite, testconstants.SubjectKeyIDWithoutColons)
	require.Equal(suite.T, 2, len(revDistPointsBySubj.Points))
	require.Equal(suite.T, vendorAccount.VendorID, revDistPointsBySubj.Points[0].Vid)
	require.Equal(suite.T, "label", revDistPointsBySubj.Points[0].Label)
	require.Equal(suite.T, venAcc65522.VendorID, revDistPointsBySubj.Points[1].Vid)
	require.Equal(suite.T, "label_PAI", revDistPointsBySubj.Points[1].Label)
	require.Equal(suite.T, testconstants.SubjectKeyIDWithoutColons, revDistPointsBySubj.IssuerSubjectKeyID)

	// Update revocation distribution point
	msgUpdPkiRevDistPoint := pkitypes.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAccount.Address,
		Vid:                  vendorAccount.VendorID,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                "label",
		DataURL:              "https://url2.data.dclmodel",
		SchemaVersion:        0,
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgUpdPkiRevDistPoint}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	revDistPoint, _ = GetPkiRevocationDistributionPoint(suite, vendorAccount.VendorID, testconstants.SubjectKeyIDWithoutColons, "label")
	require.Equal(suite.T, vendorAccount.VendorID, revDistPoint.Vid)
	require.Equal(suite.T, "https://url2.data.dclmodel", revDistPoint.DataURL)
	require.Equal(suite.T, testconstants.SubjectKeyIDWithoutColons, revDistPoint.IssuerSubjectKeyID)

	// Delete revocation distribution point
	msgDelPkiRevDistPoint := pkitypes.MsgDeletePkiRevocationDistributionPoint{
		Signer:             vendorAccount.Address,
		Vid:                vendorAccount.VendorID,
		IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
		Label:              "label",
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgDelPkiRevDistPoint}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	_, err = GetPkiRevocationDistributionPoint(suite, vendorAccount.VendorID, testconstants.SubjectKeyIDWithoutColons, "label")
	suite.AssertNotFound(err)

	revDistPointsBySubj, _ = GetPkiRevocationDistributionPointsBySubject(suite, testconstants.SubjectKeyIDWithoutColons)
	require.Equal(suite.T, 1, len(revDistPointsBySubj.Points))
	require.Equal(suite.T, venAcc65522.VendorID, revDistPointsBySubj.Points[0].Vid)
	require.Equal(suite.T, "label_PAI", revDistPointsBySubj.Points[0].Label)
	require.Equal(suite.T, testconstants.SubjectKeyIDWithoutColons, revDistPointsBySubj.IssuerSubjectKeyID)

	revDistPoints, _ = GetAllPkiRevocationDistributionPoints(suite)
	require.Equal(suite.T, 1, len(revDistPoints))

	// Add revocation distribution point for PAA by Vendor with certificate with different whitespaces
	label := "label-add-update"
	dataURL := testconstants.DataURL + "add-update"

	msgAddPkiRevDistPoints = pkitypes.MsgAddPkiRevocationDistributionPoint{
		Signer:               vendorAccount.Address,
		Vid:                  vendorAccount.VendorID,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		IsPAA:                true,
		CrlSignerCertificate: testconstants.PAACertWithNumericVidDifferentWhitespaces,
		Label:                label,
		DataURL:              dataURL,
		RevocationType:       1,
		SchemaVersion:        0,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddPkiRevDistPoints}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	revocationPointBySubjectKeyID, err := GetPkiRevocationDistributionPointsBySubject(suite, testconstants.SubjectKeyIDWithoutColons)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 2, len(revocationPointBySubjectKeyID.Points))
	require.Equal(suite.T, msgAddPkiRevDistPoints.CrlSignerCertificate, revocationPointBySubjectKeyID.Points[1].CrlSignerCertificate)

	// Update revocation distribution point
	msgUpdatePkiRevocationDistributionPoint := pkitypes.MsgUpdatePkiRevocationDistributionPoint{
		Signer:               vendorAccount.Address,
		Vid:                  vendorAccount.VendorID,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                label,
		DataURL:              dataURL + "/new",
		SchemaVersion:        0,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgUpdatePkiRevocationDistributionPoint}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	revocationPointBySubjectKeyID, err = GetPkiRevocationDistributionPointsBySubject(suite, testconstants.SubjectKeyIDWithoutColons)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 2, len(revocationPointBySubjectKeyID.Points))
	require.Equal(suite.T, msgUpdatePkiRevocationDistributionPoint.CrlSignerCertificate, revocationPointBySubjectKeyID.Points[1].CrlSignerCertificate)
	require.Equal(suite.T, msgUpdatePkiRevocationDistributionPoint.DataURL, revocationPointBySubjectKeyID.Points[1].DataURL)

	// Revoke certificates by serialNumber

	// Add root certificates
	msgProposeAddX509RootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertWithSameSubjectAndSKID1,
		Vid:    65521,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	msgApproveAddX509RootCert = pkitypes.MsgApproveAddX509RootCert{
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	msgProposeAddX509RootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertWithSameSubjectAndSKID2,
		Vid:    65521,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	msgApproveAddX509RootCert = pkitypes.MsgApproveAddX509RootCert{
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Add intermediate and leaf certificates
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateWithSameSubjectAndSKID1,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Try to add second intermediate certificate with same subject and SKID by another Vendor
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateWithSameSubjectAndSKID2,
		Signer: venAcc65522.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, venName65522, venAcc65522)
	require.Error(suite.T, err)

	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateWithSameSubjectAndSKID2,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.LeafCertWithSameSubjectAndSKID,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check approved certificate
	certs, _ := GetDaX509Cert(suite, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 2, len(certs.Certs))
	certs, _ = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 2, len(certs.Certs))
	certs, _ = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))

	// Revoke intermediate certificate with invalid serialNumber
	msgRevokeX509Cert = pkitypes.MsgRevokeX509Cert{
		Subject:      testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: "invalid",
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeX509Cert}, vendorName, vendorAccount)
	require.Error(suite.T, err)

	// Revoke intermediate certificate with serialNumber 1 only(child certs should not be removed)
	msgRevokeX509Cert.SerialNumber = testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRevokeX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Request revoked certificate with serialNumber 3
	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber, revokedCertificate.Certs[0].SerialNumber)

	// Check approved certificate
	certs, _ = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber, certs.Certs[0].SerialNumber)
	certs, _ = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber, certs.Certs[0].SerialNumber)

	// Revoke Root certificate with invalid serialNumber
	msgProposeRevokeX509RootCert = pkitypes.MsgProposeRevokeX509RootCert{
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: "invalid",
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeRevokeX509RootCert}, jackName, jackAccount)
	require.Error(suite.T, err)

	// Revoke Root certificate with serialNumber 1 only(child certs should not be removed)
	msgProposeRevokeX509RootCert.SerialNumber = testconstants.RootCertWithSameSubjectAndSKID1SerialNumber
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeRevokeX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	proposedCertificateRevocation, _ = GetProposedRevocationX509CertBySerialNumber(suite, msgProposeRevokeX509RootCert.Subject, msgProposeRevokeX509RootCert.SubjectKeyId, msgProposeRevokeX509RootCert.SerialNumber)
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKIDSubject, proposedCertificateRevocation.Subject)
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID, proposedCertificateRevocation.SubjectKeyId)
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKID1SerialNumber, proposedCertificateRevocation.SerialNumber)

	msgApproveRevokeX509RootCert = pkitypes.MsgApproveRevokeX509RootCert{
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: testconstants.RootCertWithSameSubjectAndSKID1SerialNumber,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveRevokeX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request revoked Root certificate with serialNumber 1
	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKIDSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKID1SerialNumber, revokedCertificate.Certs[0].SerialNumber)
	require.True(suite.T, revokedCertificate.Certs[0].IsRoot)

	// Check approved certificate
	certs, _ = GetDaX509Cert(suite, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKID2SerialNumber, certs.Certs[0].SerialNumber)

	certs, _ = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber, certs.Certs[0].SerialNumber)

	certs, _ = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber, certs.Certs[0].SerialNumber)

	// Remove x509 certificate with invalid serialNumber
	msgRemoveX509Cert := pkitypes.MsgRemoveX509Cert{
		Subject:      testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: "invalid",
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveX509Cert}, vendorName, vendorAccount)
	require.Error(suite.T, err)

	// Try to Remove x509 certificate by subject and subject key id when sender is not Vendor account
	msgRemoveX509Cert = pkitypes.MsgRemoveX509Cert{
		Subject:      testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: "invalid",
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveX509Cert}, aliceName, aliceAccount)
	require.Error(suite.T, err)

	// Remove revoked x509 certificate by subject and subject key id
	msgRemoveX509Cert = pkitypes.MsgRemoveX509Cert{
		Subject:      testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)
	// Check that two intermediate certificates removed
	_, err = GetRevokedDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)

	// Remove leaf x509 certificate by subject and subject key id
	msgRemoveX509Cert = pkitypes.MsgRemoveX509Cert{
		Subject:      testconstants.LeafCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID,
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)
	// Check that leaf certificate removed
	_, err = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)

	// Remove x509 by subject, subject key id and serial number
	// Add intermediate certificates
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateWithSameSubjectAndSKID1,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateWithSameSubjectAndSKID2,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.LeafCertWithSameSubjectAndSKID,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Remove x509 certificate by serial number
	msgRemoveX509Cert = pkitypes.MsgRemoveX509Cert{
		Subject:      testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber,
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check that leaf and x509 with different serial number is not removed
	certs, _ = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber, certs.Certs[0].SerialNumber)

	certs, _ = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber, certs.Certs[0].SerialNumber)

	// Remove revoked x509 certificate by subject and subject key id
	msgRemoveX509Cert = pkitypes.MsgRemoveX509Cert{
		Subject:      testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	_, err = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)
	certs, _ = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber, certs.Certs[0].SerialNumber)

	// Remove leaf x509 certificate by subject and subject key id
	msgRemoveX509Cert = pkitypes.MsgRemoveX509Cert{
		Subject:      testconstants.LeafCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID,
		Signer:       vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgRemoveX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	_, err = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)

	// Revoke Root certificate and its child certificates
	// Add intermediate and leaf certificates
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateWithSameSubjectAndSKID1,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.LeafCertWithSameSubjectAndSKID,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check that certs are added
	certs, _ = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber, certs.Certs[0].SerialNumber)

	certs, _ = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber, certs.Certs[0].SerialNumber)

	allCerts, _ := GetCert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(allCerts.Certs))
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber, allCerts.Certs[0].SerialNumber)

	// Revoke root cert and its child
	msgProposeRevokeX509RootCert = pkitypes.MsgProposeRevokeX509RootCert{
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: testconstants.RootCertWithSameSubjectAndSKID2SerialNumber,
		RevokeChild:  true,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeRevokeX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	msgApproveRevokeX509RootCert = pkitypes.MsgApproveRevokeX509RootCert{
		Subject:      testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyId: testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber: testconstants.RootCertWithSameSubjectAndSKID2SerialNumber,
		Signer:       aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveRevokeX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Request revoked certificates
	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 2, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKIDSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID, revokedCertificate.SubjectKeyId)
	require.Equal(suite.T, testconstants.RootCertWithSameSubjectAndSKID2SerialNumber, revokedCertificate.Certs[1].SerialNumber)
	require.True(suite.T, revokedCertificate.Certs[1].IsRoot)

	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID, revokedCertificate.SubjectKeyId)

	revokedCertificate, _ = GetRevokedDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	require.Equal(suite.T, 1, len(revokedCertificate.Certs))
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSubject, revokedCertificate.Subject)
	require.Equal(suite.T, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID, revokedCertificate.SubjectKeyId)

	// Check that all certs are removed from approved lists
	_, err = GetDaX509Cert(suite, testconstants.RootCertWithSameSubjectAndSKIDSubject, testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetDaX509Cert(suite, testconstants.IntermediateCertWithSameSubjectAndSKIDSubject, testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)
	_, err = GetDaX509Cert(suite, testconstants.LeafCertWithSameSubjectAndSKIDSubject, testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID)
	suite.AssertNotFound(err)

	// Add X509 certificates by Vendor Account

	// Check that if root cert is VID scoped and RootVID==CertVID==AccountVID then adding x509 should succeed
	// Add root certificate
	msgProposeAddX509RootCert = pkitypes.MsgProposeAddX509RootCert{
		Cert:   testconstants.RootCertWithVid,
		Vid:    testconstants.RootCertWithVidVid,
		Signer: aliceAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgProposeAddX509RootCert}, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	msgApproveAddX509RootCert = pkitypes.MsgApproveAddX509RootCert{
		Subject:      testconstants.RootCertWithVidSubject,
		SubjectKeyId: testconstants.RootCertWithVidSubjectKeyID,
		Signer:       jackAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgApproveAddX509RootCert}, jackName, jackAccount)
	require.NoError(suite.T, err)

	// Register new Vendor account
	vendorName = utils.RandString()
	vendorAccount = test_dclauth.CreateVendorAccount(
		suite,
		vendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		testconstants.RootCertWithVidVid,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, vendorAccount)

	// Add an intermediate certificate
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertWithVid1,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.NoError(suite.T, err)

	// Check approved certificates
	certs, _ = GetDaX509Cert(suite, testconstants.RootCertWithVidSubject, testconstants.RootCertWithVidSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	certs, _ = GetDaX509Cert(suite, testconstants.IntermediateCertWithVid1Subject, testconstants.IntermediateCertWithVid1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))

	// Check certificates (using global collection)
	allCerts, _ = GetCert(suite, testconstants.RootCertWithVidSubject, testconstants.RootCertWithVidSubjectKeyID)
	require.Equal(suite.T, 1, len(allCerts.Certs))
	allCerts, _ = GetCert(suite, testconstants.IntermediateCertWithVid1Subject, testconstants.IntermediateCertWithVid1SubjectKeyID)
	require.Equal(suite.T, 1, len(allCerts.Certs))

	// Check that if root cert is VID scoped and rootVID != CertVID then adding an intermediate cert should fail
	// Add an intermediate certificate
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.IntermediateCertWithVid2,
		Signer: vendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, vendorName, vendorAccount)
	require.Error(suite.T, err)

	// Check there is only one approved intermediate certificate
	certs, _ = GetDaX509Cert(suite, testconstants.IntermediateCertWithVid1Subject, testconstants.IntermediateCertWithVid1SubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, testconstants.IntermediateCertWithVid1SerialNumber, certs.Certs[0].SerialNumber)

	// Check that if root cert is non-VID scoped and CertVID != AccountVID then adding an intermediate cert should fail
	// Ensure that there is a non-VID root cert exists
	certs, _ = GetDaX509Cert(suite, testconstants.PAACertNoVidSubject, testconstants.PAACertNoVidSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))

	// Try to submit txn with another Vendor
	newVendorName := utils.RandString()
	newVendorAccount := test_dclauth.CreateVendorAccount(
		suite,
		newVendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		1234,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, newVendorAccount)

	// Add an intermediate certificate
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.PAICertWithNumericVid,
		Signer: newVendorAccount.Address,
	}

	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, newVendorName, newVendorAccount)
	require.Error(suite.T, err)

	// Check there is no an intermediate certificate
	_, err = GetDaX509Cert(suite, testconstants.PAICertWithNumericVidSubject, testconstants.PAICertWithNumericVidSubjectKeyID)
	suite.AssertNotFound(err)

	// Check that if root cert is non-VID scoped and CertVID==AccountVID then adding x509 should succeed
	// Create vendor with valid VID
	newVendorName = utils.RandString()
	newVendorAccount = test_dclauth.CreateVendorAccount(
		suite,
		newVendorName,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		testconstants.IntermediateCertWithVid2Vid,
		testconstants.ProductIDsEmpty,
		aliceName,
		aliceAccount,
		jackName,
		jackAccount,
		testconstants.Info,
	)
	require.NotNil(suite.T, newVendorAccount)

	// Add an intermediate certificate
	msgAddX509Cert = pkitypes.MsgAddX509Cert{
		Cert:   testconstants.PAICertWithNumericVid,
		Signer: newVendorAccount.Address,
	}
	_, err = suite.BuildAndBroadcastTx([]sdk.Msg{&msgAddX509Cert}, newVendorName, newVendorAccount)
	require.NoError(suite.T, err)
	// Check there is only one approved intermediate certificate
	certs, _ = GetDaX509Cert(suite, testconstants.PAICertWithNumericVidSubject, testconstants.PAICertWithNumericVidSubjectKeyID)
	require.Equal(suite.T, 1, len(certs.Certs))
	require.Equal(suite.T, int32(testconstants.PAICertWithNumericVidVid), certs.Certs[0].Vid)
}
