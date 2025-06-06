package v3

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"

	"github.com/iotexproject/iotex-core/v2/action/protocol/abiutil"
	stakingComm "github.com/iotexproject/iotex-core/v2/action/protocol/staking/ethabi/common"
)

const _candidateByNameInterfaceABI = `[
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "candName",
				"type": "string"
			}
		],
		"name": "candidateByNameV3",
		"outputs": [
			{
				"components": [
					{
						"internalType": "address",
						"name": "ownerAddress",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "operatorAddress",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "rewardAddress",
						"type": "address"
					},
					{
						"internalType": "string",
						"name": "name",
						"type": "string"
					},
					{
						"internalType": "uint256",
						"name": "totalWeightedVotes",
						"type": "uint256"
					},
					{
						"internalType": "uint64",
						"name": "selfStakeBucketIdx",
						"type": "uint64"
					},
					{
						"internalType": "uint256",
						"name": "selfStakingTokens",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "id",
						"type": "address"
					}
				],
				"internalType": "struct IStaking.CandidateV3",
				"name": "",
				"type": "tuple"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

var _candidateByNameMethod abi.Method

func init() {
	_candidateByNameMethod = abiutil.MustLoadMethod(_candidateByNameInterfaceABI, "candidateByNameV3")
}

func newCandidateByNameStateContext(data []byte) (*stakingComm.CandidateByNameStateContext, error) {
	return stakingComm.NewCandidateByNameStateContext(data, &_candidateByNameMethod, iotexapi.ReadStakingDataMethod_CANDIDATE_BY_NAME)
}
