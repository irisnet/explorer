
const urlApi = {
  //home
  candidatesTop: `/api/stake/candidatesTop`,
  txsByDay: `/api/txsByDay`,
  blocksRecent: '/api/blocks/recent',
  txsRecent: '/api/txs/recent',
  navigation: '/api/home/navigation',
  //address
  address: '/api/txs/statistics?address={address}',
  addressTxTrans: '/api/tx/trans/{pageNumber}/{pageSize}?address={address}',
  addressTxStake: '/api/tx/stake/{pageNumber}/{pageSize}?address={address}',
  addressTxDeclaration: '/api/tx/declaration/{pageNumber}/{pageSize}?address={address}',
  addressTxGov: '/api/tx/gov/{pageNumber}/{pageSize}?address={address}',
  addressAccount: '/api/account/{address}',
  addressStakeCandidate: '/api/stake/candidate/{address}',
  addressStakeCandidateStatus: '/api/stake/candidate/{address}/status',
  addressTxByAddress : '/api/txsByAddress/{address}/{pageNumber}/{pageSize}',
  addressStakeCandidateWeek: '/api/stake/candidate/{address}/power/week',
  addressStakeCandidateMonth: '/api/stake/candidate/{address}/power/month',
  addressStakeCandidateMonths: '/api/stake/candidate/{address}/power/months',
  addressStakeCandidateUptimeHour: '/api/stake/candidate/{address}/uptime/hour',
  addressStakeCandidateUptimeWeek: '/api/stake/candidate/{address}/uptime/week',
  addressStakeCandidateUptimeMonth: '/api/stake/candidate/{address}/uptime/month',
  //validatorDetail
  validatorDelegations: '/api/stake/validators/{validatorAddr}/delegations?page={currentPage}&size={pageSize}',
  validatorUnbondingDelegations: '/api/stake/validators/{validatorAddr}/unbonding-delegations?page={currentPage}&size={pageSize}',
  validatorRedelegations: '/api/stake/validators/{validatorAddr}/redelegations?page={currentPage}&size={pageSize}',
  validatorVote: '/api/stake/validators/{validatorAddr}/vote?page={currentPage}&size={pageSize}',
  validatorDepositorTxs: '/api/stake/validators/{validatorAddr}/depositor_txs?page={currentPage}&size={pageSize}',
  //header
  headerTx: '/api/tx/{searchValue}',
  headerSearchAccount: '/api/account/{searchValue}',
  headerSearchCandidate: '/api/stake/candidate/{searchValue}',
  headerSearchValue: '/api/search/{searchValue}',
  headerConfig: '/api/config',
  //blockInfo
  blockInfoHeight: '/api/block/blockinfo/{height}',
  blockInfoGov: '/api/tx/gov/{currentPage}/{pageSize}?height={blockHeight}',
  blockInfoValidatorSet: '/api/block/validatorset/{blockHeight}?page={currentPage}&size={pageSize}',
  blockInfoTransfer: '/api/tx/trans/{currentPage}/{pageSize}?height={blockHeight}',
  blockInfoStake: '/api/tx/stake/{currentPage}/{pageSize}?height={blockHeight}',
  blockInfoDeclaration: '/api/tx/declaration/{currentPage}/{pageSize}?height={blockHeight}',
  //blockList
  blockList: '/api/blocks?page={pageNumber}&size={pageSize}',
  blockListLatestheight: '/api/block/latestheight',
  //parameter
  govParams: '/api/gov/params',
  //proposal
  proposalDetail: '/api/gov/proposals/{proposalId}',
  proposalDetailVoterTx:'/api/gov/proposals/{proposalId}/voter_txs?page={pageNumber}&size={perPageSize}',
  proposalDetailDepositorTx:'/api/gov/proposals/{proposalId}/depositor_txs?page={pageNumber}&size={perPageSize}',
  //proposalList
  proposalListVotingAndDeposit:'/api/gov/deposit_voting_proposals',
  proposalList: '/api/gov/proposals?page={pageNumber}&size={pageSize}',
  //searchResult
  searchResult: '/api/search/{searchValue}',
  //txDetail
  txDetail: '/api/tx/{txHash}',
  //txList
  txListTransfer: '/api/tx/trans/{pageNumber}/{pageSize}',
  txListStake: '/api/tx/stake/{pageNumber}/{pageSize}',
  txListDeclaration: '/api/tx/declaration/{pageNumber}/{pageSize}',
  txListGov: '/api/tx/gov/{pageNumber}/{pageSize}',
  //validatorList
  validatorList: '/api/stake/validators?page={pageNumber}&size={pageSize}&type={validatorStatus}&origin=browser',
  validatorListHeaderImg: 'https://keybase.io/_/api/1.0/user/lookup.json?fields=pictures&key_suffix={keyBase}',
  //sysdate
  sysdate: '/api/sysdate',
  //richList
  richListAccounts: '/api/accounts'
};

export default urlApi
