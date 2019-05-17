
const UrlApi = {
  //home
  candidatesTop: `/api/stake/candidatesTop`,
  txsByDay: `/api/txsByDay`,
  blocksRecent: '/service/blocks/recent',
  txsRecent: '/service/txs/recent',
  navigation: '/service/home/navigation',
  //address
  address: '/service/txs/statistics?address={address}',
  addressTxTrans: '/service/tx/trans/{pageNumber}/{pageSize}?address={address}',
  addressTxStake: '/service/tx/stake/{pageNumber}/{pageSize}?address={address}',
  addressTxDeclaration: '/service/tx/declaration/{pageNumber}/{pageSize}?address={address}',
  addressTxGov: '/service/tx/gov/{pageNumber}/{pageSize}?address={address}',
  addressAccount: '/service/account/{address}',
  addressStakeCandidate: '/service/stake/candidate/{address}',
  addressStakeCandidateStatus: '/service/stake/candidate/{address}/status',
  addressTxByAddress : '/service/txsByAddress/{address}/1/30',
  addressStakeCandidateWeek: '/service/stake/candidate/{address}/power/week',
  addressStakeCandidateMonth: '/service/stake/candidate/{address}/power/month',
  addressStakeCandidateMonths: '/service/stake/candidate/{address}/power/months',
  addressStakeCandidateUptimeHour: '/service/stake/candidate/{address}/uptime/hour',
  addressStakeCandidateUptimeWeek: '/service/stake/candidate/{address}/uptime/week',
  addressStakeCandidateUptimeMonth: '/service/stake/candidate/{address}/uptime/month',
  //header
  headerTx: '/service/tx/{searchValue}',
  headerSearchAccount: '/service/account/{searchValue}',
  headerSearchCandidate: '/service/stake/candidate/{searchValue}',
  headerSearchValue: '/service/search/{searchValue}',
  headerConfig: '/service/config',
  //blockInfo
  blockInfoHeight: '/service/block/{height}',
  blockInfoValidatorSet: '/service/stake/validatorset?height={height}&page={pageNumber}&size={pageSize}',
  blockInfoTokenFlow: '/service/tx/token/flow?height={height}&page={pageNumber}&size={pageSize}',
  //blockList
  blockList: '/service/blocks/{pageNumber}/{pageSize}',
  //parameter
  govParams: '/service/gov/params',
  //proposal
  proposalDetail: '/service/gov/proposal/{proposalId}',
  //proposalList
  proposalList: '/service/gov/proposals?page={pageNumber}&size={pageSize}',
  //searchResult
  searchResult: '/service/search/{searchValue}',
  //txDetail
  txDetail: '/service/tx/{txHash}',
  //txList
  txListTransfer: '/service/tx/trans/{pageNumber}/{pageSize}',
  txListStake: '/service/tx/stake/{pageNumber}/{pageSize}',
  txListDeclaration: '/service/tx/declaration/{pageNumber}/{pageSize}',
  txListGov: '/service/tx/gov/{pageNumber}/{pageSize}',
  //validatorList
  validatorList: '/service/stake/validators?page={pageNumber}&size={pageSize}&type={validatorStatus}&origin=browser',
  validatorListHeaderImg: 'https://keybase.io/_/service/1.0/user/lookup.json?fields=pictures&key_suffix={keyBase}',
  //sysdate
  sysdate: '/service/sysdate',
};

export default UrlApi
