
const urlApi = {
  //home
  candidatesTop: `/api/stake/candidatesTop`,
  txsByDay: `/api/txsByDay`,
  blocksRecent: '/api/service/blocks/recent',
  txsRecent: '/api/service/txs/recent',
  navigation: '/api/service/home/navigation',
  //address
  address: '/api/service/txs/statistics?address={address}',
  addressTxTrans: '/api/service/tx/trans/{pageNumber}/{pageSize}?address={address}',
  addressTxStake: '/api/service/tx/stake/{pageNumber}/{pageSize}?address={address}',
  addressTxDeclaration: '/api/service/tx/declaration/{pageNumber}/{pageSize}?address={address}',
  addressTxGov: '/api/service/tx/gov/{pageNumber}/{pageSize}?address={address}',
  addressAccount: '/api/service/account/{address}',
  addressStakeCandidate: '/api/service/stake/candidate/{address}',
  addressStakeCandidateStatus: '/api/service/stake/candidate/{address}/status',
  addressTxByAddress : '/api/service/txsByAddress/{address}/1/30',
  addressStakeCandidateWeek: '/api/service/stake/candidate/{address}/power/week',
  addressStakeCandidateMonth: '/api/service/stake/candidate/{address}/power/month',
  addressStakeCandidateMonths: '/api/service/stake/candidate/{address}/power/months',
  addressStakeCandidateUptimeHour: '/api/service/stake/candidate/{address}/uptime/hour',
  addressStakeCandidateUptimeWeek: '/api/service/stake/candidate/{address}/uptime/week',
  addressStakeCandidateUptimeMonth: '/api/service/stake/candidate/{address}/uptime/month',
  //header
  headerTx: '/api/service/tx/{searchValue}',
  headerSearchAccount: '/api/service/account/{searchValue}',
  headerSearchCandidate: '/api/service/stake/candidate/{searchValue}',
  headerSearchValue: '/api/service/search/{searchValue}',
  headerConfig: '/api/service/config',
  //blockInfo
  blockInfoHeight: '/api/service/block/{height}',
  blockInfoValidatorSet: '/api/service/stake/validatorset?height={height}&page={pageNumber}&size={pageSize}',
  blockInfoTokenFlow: '/api/service/tx/token/flow?height={height}&page={pageNumber}&size={pageSize}',
  //blockList
  blockList: '/api/service/blocks/{pageNumber}/{pageSize}',
  //parameter
  govParams: '/api/service/gov/params',
  //proposal
  proposalDetail: '/api/service/gov/proposal/{proposalId}',
  //proposalList
  proposalList: '/api/service/gov/proposals?page={pageNumber}&size={pageSize}',
  //searchResult
  searchResult: '/api/service/search/{searchValue}',
  //txDetail
  txDetail: '/api/service/tx/{txHash}',
  //txList
  txListTransfer: '/api/service/tx/trans/{pageNumber}/{pageSize}',
  txListStake: '/api/service/tx/stake/{pageNumber}/{pageSize}',
  txListDeclaration: '/api/service/tx/declaration/{pageNumber}/{pageSize}',
  txListGov: '/api/service/tx/gov/{pageNumber}/{pageSize}',
  //validatorList
  validatorList: '/api/service/stake/validators?page={pageNumber}&size={pageSize}&type={validatorStatus}&origin=browser',
  validatorListHeaderImg: 'https://keybase.io/_/service/1.0/user/lookup.json?fields=pictures&key_suffix={keyBase}',
  //sysdate
  sysdate: '/api/service/sysdate',
};

export default urlApi
